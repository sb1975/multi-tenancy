/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"reflect"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	tenancy "github.com/kubernetes-sigs/multi-tenancy/incubator/hnc/api/v1alpha1"
	"github.com/kubernetes-sigs/multi-tenancy/incubator/hnc/pkg/forest"
)

// HierarchyReconciler is responsible for determining the forest structure from the Hierarchy CRs,
// as well as ensuring all objects in the forest are propagated correctly when the hierarchy
// changes. It can also set the status of the Hierarchy CRs, as well as (in rare cases) override
// part of its spec (i.e., if a parent namespace no longer exists).
type HierarchyReconciler struct {
	client.Client
	Log logr.Logger

	// Forest is the in-memory data structure that is shared with all other reconcilers.
	// HierarchyReconciler is responsible for keeping it up-to-date, but the other reconcilers
	// use it to determine how to propagate objects.
	Forest *forest.Forest

	// Types is a list of other reconcillers that HierarchyReconciler can call if the hierarchy
	// changes. This will force all objects to be re-propagated.
	//
	// This is probably wildly inefficient, and we can probably make better use of things like
	// owner references to make this better. But for a PoC, it works just fine.
	Types []NamespaceSyncer

	// Affected is a channel of event.GenericEvent (see "Watching Channels" in
	// https://book-v1.book.kubebuilder.io/beyond_basics/controller_watches.html) that is used to
	// enqueue additional namespaces that need updating.
	Affected chan event.GenericEvent
}

// NamespaceSyncer syncs various aspects of a namespace. The HierarchyReconciler both implements
// it (so it can be called by NamespaceSyncer) and uses it (to sync the objects in the
// namespace).
type NamespaceSyncer interface {
	SyncNamespace(context.Context, logr.Logger, string) error
}

// +kubebuilder:rbac:groups=hnc.x-k8s.io,resources=hierarchies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=hnc.x-k8s.io,resources=hierarchies/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;update;patch

// Reconcile simply calls SyncNamespace, which can also be called if a namespace is created or
// deleted.
func (r *HierarchyReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	ns := req.NamespacedName.Namespace
	log := r.Log.WithValues("ns", ns)
	return ctrl.Result{}, r.reconcile(ctx, log, ns)
}

func (r *HierarchyReconciler) reconcile(ctx context.Context, log logr.Logger, nm string) error {
	// Get the singleton and namespace
	inst, err := r.getSingleton(ctx, nm)
	if err != nil {
		log.Error(err, "Couldn't read singleton")
		return err
	}
	ns, err := r.getNamespace(ctx, nm)
	if err != nil {
		log.Error(err, "Couldn't read namespace")
		return err
	}

	// If either object exists but is being deleted, we won't update them when we're finished syncing
	// (we should sync our internal data structure anyway just in case something's changed).  I'm not
	// sure if this is the right thing to do but the kubebuilder boilerplate included this case
	// explicitly.
	update := true
	if !inst.GetDeletionTimestamp().IsZero() || !ns.GetDeletionTimestamp().IsZero() {
		log.Info("Singleton or namespace are being deleted; will not update")
		update = false
	}

	// Sync the tree.
	if err := r.updateTree(ctx, log, ns, inst, update); err != nil {
		return err
	}

	if update {
		// Update all the objects in this namespace. We have to do this at least *after* the tree is
		// updated, because if we don't, we could incorrectly think we've propagated the wrong objects
		// from our ancestors, or are propagating the wrong objects to our descendants.
		return r.updateObjects(ctx, log, nm)
	}

	return nil
}

// updateTree syncs the Hierarchy singleton with the in-memory forest (writing back to the apiserver
// if necessary and requested) and calls itself on any affected namespaces if the hierarchy has
// changed.
//
// TODO: store the conditions in the in-memory forest so that object propagation can be disabled if
// there's a problem on the namespace.
func (r *HierarchyReconciler) updateTree(ctx context.Context, log logr.Logger, nsInst *corev1.Namespace, inst *tenancy.Hierarchy, update bool) error {
	// Update the in-memory data structures
	orig := inst.DeepCopy()
	affected := r.syncWithForest(ctx, log, nsInst, inst)

	var updateErr error
	if update {
		updateErr = r.updateAPIServer(ctx, log, orig, inst)
	}

	// Enqueue any affected namespaces even if we failed to write back to the apiserver; we know
	// they may have changed somehow and the controller will ensure they eventually get reconciled.
	r.enqueueAffected(log, affected)

	return updateErr
}

func (r *HierarchyReconciler) updateAPIServer(ctx context.Context, log logr.Logger, orig, inst *tenancy.Hierarchy) error {
	if reflect.DeepEqual(orig, inst) {
		return nil
	}

	if inst.CreationTimestamp.IsZero() {
		log.Info("Creating singleton on apiserver")
		if err := r.Create(ctx, inst); err != nil {
			log.Error(err, "while creating on apiserver")
			return err
		}
	} else {
		log.Info("Updating singleton on apiserver")
		if err := r.Update(ctx, inst); err != nil {
			log.Error(err, "while updating apiserver")
			return err
		}
	}

	return nil
}

// updateObjects calls all type reconcillers in this namespace.
func (r *HierarchyReconciler) updateObjects(ctx context.Context, log logr.Logger, ns string) error {
	for _, tr := range r.Types {
		if err := tr.SyncNamespace(ctx, log, ns); err != nil {
			return err
		}
	}

	return nil
}

// syncWithForest synchronizes the in-memory forest with the (in-memory) Hierarchy instance. If any
// *other* namespaces have changed, it returns the list of affected namespaces.
func (r *HierarchyReconciler) syncWithForest(ctx context.Context, log logr.Logger, nsInst *corev1.Namespace, inst *tenancy.Hierarchy) []string {
	r.Forest.Lock()
	defer r.Forest.Unlock()
	ns := r.Forest.Get(inst.ObjectMeta.Namespace)

	// If the namespace has been deleted, just update that fact.
	if nsInst.Name == "" {
		if ns.UnsetExists() {
			log.Info("Removed namespace")
			return ns.RelativesNames()
		}
	}

	// Mark the namespace as existing. If this is the first time we're reconciling this namespace,
	// mark its relatives as being affected since they may have been waiting for this parent.
	affected := []string{}
	conds := []tenancy.Condition{}
	if ns.SetExists() {
		log.Info("Found new namespace")
		affected = append(affected, ns.RelativesNames()...)
	}

	// Sync this namespace with its current parent.
	var curParent *forest.Namespace
	if inst.Spec.Parent != "" {
		curParent = r.Forest.Get(inst.Spec.Parent)
		if !curParent.Exists() {
			log.Info("Missing parent", "parent", inst.Spec.Parent)
			conds = append(conds, tenancy.Condition{Msg: "missing parent"})
		}
	}

	// Update the in-memory hierarchy if it's changed
	oldParent := ns.Parent()
	if oldParent != curParent {
		log.Info("Parent has changed", "old", oldParent.Name(), "new", curParent.Name())
		if err := ns.SetParent(curParent); err != nil {
			log.Info("Couldn't set parent", "reason", err, "parent", inst.Spec.Parent)
			conds = append(conds, tenancy.Condition{Msg: err.Error()})
		} else {
			// Only call other parts of the hierarchy recursively if this one was successfully updated;
			// otherwise, if you get a cycle, this could get into an infinite loop.
			if oldParent != nil {
				affected = append(affected, oldParent.Name())
			}
			if curParent != nil {
				affected = append(affected, curParent.Name())
			}
		}
	}

	// Update all other changed fields. If they're empty, ensure they're nil so that they compare
	// properly.
	inst.Status.Children = ns.ChildNames()
	if len(conds) > 0 {
		inst.Status.Conditions = conds
	} else {
		inst.Status.Conditions = nil
	}

	return affected
}

// enqueueAffected enqueues all affected namespaces for later reconciliation.
func (r *HierarchyReconciler) enqueueAffected(log logr.Logger, affected []string) {
	for _, nm := range affected {
		log.Info("Enqueuing for reconcilation", "affected", nm)
		// The watch handler doesn't care about anything except the metadata.
		inst := &tenancy.Hierarchy{}
		inst.ObjectMeta.Name = tenancy.Singleton
		inst.ObjectMeta.Namespace = nm
		r.Affected <- event.GenericEvent{Meta: inst}
	}
}

// getSingleton returns the singleton if it exists, or creates an empty one if it doesn't.
func (r *HierarchyReconciler) getSingleton(ctx context.Context, nm string) (*tenancy.Hierarchy, error) {
	nnm := types.NamespacedName{Namespace: nm, Name: tenancy.Singleton}
	inst := &tenancy.Hierarchy{}
	if err := r.Get(ctx, nnm, inst); err != nil {
		if !errors.IsNotFound(err) {
			return nil, err
		}

		// It doesn't exist - initialize it to a sane initial value.
		inst.ObjectMeta.Name = tenancy.Singleton
		inst.ObjectMeta.Namespace = nm
	}

	return inst, nil
}

// getNamespace returns the namespace if it exists, or returns an invalid, blank, unnamed one if it
// doesn't. This allows it to be trivially identified as a namespace that doesn't exist, and also
// allows us to easily modify it if we want to create it.
func (r *HierarchyReconciler) getNamespace(ctx context.Context, nm string) (*corev1.Namespace, error) {
	ns := &corev1.Namespace{}
	nnm := types.NamespacedName{Name: nm}
	if err := r.Get(ctx, nnm, ns); err != nil {
		if !errors.IsNotFound(err) {
			return nil, err
		}
		return &corev1.Namespace{}, nil
	}
	return ns, nil
}

func (r *HierarchyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// Maps namespaces to their singletons
	nsMapFn := handler.ToRequestsFunc(
		func(a handler.MapObject) []reconcile.Request {
			return []reconcile.Request{
				{NamespacedName: types.NamespacedName{
					Name:      tenancy.Singleton,
					Namespace: a.Meta.GetName(),
				}},
			}
		})
	return ctrl.NewControllerManagedBy(mgr).
		For(&tenancy.Hierarchy{}).
		Watches(&source.Channel{Source: r.Affected}, &handler.EnqueueRequestForObject{}).
		Watches(&source.Kind{Type: &corev1.Namespace{}}, &handler.EnqueueRequestsFromMapFunc{ToRequests: nsMapFn}).
		Complete(r)
}
