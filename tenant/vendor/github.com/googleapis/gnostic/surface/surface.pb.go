// Code generated by protoc-gen-go.
// source: surface.proto
// DO NOT EDIT!

/*
Package surface_v1 is a generated protocol buffer package.

It is generated from these files:
	surface.proto

It has these top-level messages:
	Field
	Type
	Method
	Model
*/
package surface_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FieldKind int32

const (
	FieldKind_SCALAR    FieldKind = 0
	FieldKind_MAP       FieldKind = 1
	FieldKind_ARRAY     FieldKind = 2
	FieldKind_REFERENCE FieldKind = 3
	FieldKind_ANY       FieldKind = 4
)

var FieldKind_name = map[int32]string{
	0: "SCALAR",
	1: "MAP",
	2: "ARRAY",
	3: "REFERENCE",
	4: "ANY",
}
var FieldKind_value = map[string]int32{
	"SCALAR":    0,
	"MAP":       1,
	"ARRAY":     2,
	"REFERENCE": 3,
	"ANY":       4,
}

func (x FieldKind) String() string {
	return proto.EnumName(FieldKind_name, int32(x))
}
func (FieldKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TypeKind int32

const (
	TypeKind_STRUCT TypeKind = 0
	TypeKind_OBJECT TypeKind = 1
)

var TypeKind_name = map[int32]string{
	0: "STRUCT",
	1: "OBJECT",
}
var TypeKind_value = map[string]int32{
	"STRUCT": 0,
	"OBJECT": 1,
}

func (x TypeKind) String() string {
	return proto.EnumName(TypeKind_name, int32(x))
}
func (TypeKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Position int32

const (
	Position_BODY     Position = 0
	Position_HEADER   Position = 1
	Position_FORMDATA Position = 2
	Position_QUERY    Position = 3
	Position_PATH     Position = 4
)

var Position_name = map[int32]string{
	0: "BODY",
	1: "HEADER",
	2: "FORMDATA",
	3: "QUERY",
	4: "PATH",
}
var Position_value = map[string]int32{
	"BODY":     0,
	"HEADER":   1,
	"FORMDATA": 2,
	"QUERY":    3,
	"PATH":     4,
}

func (x Position) String() string {
	return proto.EnumName(Position_name, int32(x))
}
func (Position) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Field is a field in a definition and can be associated with
// a position in a request structure.
type Field struct {
	Name          string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type          string    `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Kind          FieldKind `protobuf:"varint,3,opt,name=kind,enum=surface.v1.FieldKind" json:"kind,omitempty"`
	Format        string    `protobuf:"bytes,4,opt,name=format" json:"format,omitempty"`
	Position      Position  `protobuf:"varint,5,opt,name=position,enum=surface.v1.Position" json:"position,omitempty"`
	NativeType    string    `protobuf:"bytes,6,opt,name=nativeType" json:"nativeType,omitempty"`
	FieldName     string    `protobuf:"bytes,7,opt,name=fieldName" json:"fieldName,omitempty"`
	ParameterName string    `protobuf:"bytes,8,opt,name=parameterName" json:"parameterName,omitempty"`
	Serialize     bool      `protobuf:"varint,9,opt,name=serialize" json:"serialize,omitempty"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Field) GetKind() FieldKind {
	if m != nil {
		return m.Kind
	}
	return FieldKind_SCALAR
}

func (m *Field) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *Field) GetPosition() Position {
	if m != nil {
		return m.Position
	}
	return Position_BODY
}

func (m *Field) GetNativeType() string {
	if m != nil {
		return m.NativeType
	}
	return ""
}

func (m *Field) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *Field) GetParameterName() string {
	if m != nil {
		return m.ParameterName
	}
	return ""
}

func (m *Field) GetSerialize() bool {
	if m != nil {
		return m.Serialize
	}
	return false
}

// Type typically corresponds to a definition, parameter, or response
// in an API and is represented by a type in generated code.
type Type struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Kind        TypeKind `protobuf:"varint,2,opt,name=kind,enum=surface.v1.TypeKind" json:"kind,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	ContentType string   `protobuf:"bytes,4,opt,name=contentType" json:"contentType,omitempty"`
	Fields      []*Field `protobuf:"bytes,5,rep,name=fields" json:"fields,omitempty"`
	TypeName    string   `protobuf:"bytes,6,opt,name=typeName" json:"typeName,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Type) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type) GetKind() TypeKind {
	if m != nil {
		return m.Kind
	}
	return TypeKind_STRUCT
}

func (m *Type) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Type) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Type) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetTypeName() string {
	if m != nil {
		return m.TypeName
	}
	return ""
}

// Method is an operation of an API and typically has associated client and server code.
type Method struct {
	Operation          string `protobuf:"bytes,1,opt,name=operation" json:"operation,omitempty"`
	Path               string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Method             string `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	Description        string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Name               string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	HandlerName        string `protobuf:"bytes,6,opt,name=handlerName" json:"handlerName,omitempty"`
	ProcessorName      string `protobuf:"bytes,7,opt,name=processorName" json:"processorName,omitempty"`
	ClientName         string `protobuf:"bytes,8,opt,name=clientName" json:"clientName,omitempty"`
	ParametersTypeName string `protobuf:"bytes,9,opt,name=parametersTypeName" json:"parametersTypeName,omitempty"`
	ResponsesTypeName  string `protobuf:"bytes,10,opt,name=responsesTypeName" json:"responsesTypeName,omitempty"`
}

func (m *Method) Reset()                    { *m = Method{} }
func (m *Method) String() string            { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()               {}
func (*Method) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Method) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *Method) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Method) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Method) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Method) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Method) GetHandlerName() string {
	if m != nil {
		return m.HandlerName
	}
	return ""
}

func (m *Method) GetProcessorName() string {
	if m != nil {
		return m.ProcessorName
	}
	return ""
}

func (m *Method) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Method) GetParametersTypeName() string {
	if m != nil {
		return m.ParametersTypeName
	}
	return ""
}

func (m *Method) GetResponsesTypeName() string {
	if m != nil {
		return m.ResponsesTypeName
	}
	return ""
}

// Model represents an API for code generation.
type Model struct {
	Name    string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Types   []*Type   `protobuf:"bytes,2,rep,name=types" json:"types,omitempty"`
	Methods []*Method `protobuf:"bytes,3,rep,name=methods" json:"methods,omitempty"`
}

func (m *Model) Reset()                    { *m = Model{} }
func (m *Model) String() string            { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()               {}
func (*Model) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetTypes() []*Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Model) GetMethods() []*Method {
	if m != nil {
		return m.Methods
	}
	return nil
}

func init() {
	proto.RegisterType((*Field)(nil), "surface.v1.Field")
	proto.RegisterType((*Type)(nil), "surface.v1.Type")
	proto.RegisterType((*Method)(nil), "surface.v1.Method")
	proto.RegisterType((*Model)(nil), "surface.v1.Model")
	proto.RegisterEnum("surface.v1.FieldKind", FieldKind_name, FieldKind_value)
	proto.RegisterEnum("surface.v1.TypeKind", TypeKind_name, TypeKind_value)
	proto.RegisterEnum("surface.v1.Position", Position_name, Position_value)
}

func init() { proto.RegisterFile("surface.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x54, 0xed, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0x3e, 0x97, 0xdc, 0x31, 0xe4, 0x59, 0x80, 0x22, 0x84, 0x50, 0x54, 0x21, 0xd4, 0x55,
	0x53, 0x05, 0xe3, 0x09, 0xb2, 0x34, 0xd5, 0x04, 0xf4, 0x03, 0x93, 0xfd, 0xe8, 0xcf, 0xd0, 0xb8,
	0x6a, 0x44, 0x1b, 0x87, 0x38, 0x4c, 0x82, 0x07, 0xe2, 0x71, 0xe0, 0x95, 0x90, 0x6f, 0x92, 0x36,
	0x5b, 0xfb, 0xcf, 0x3e, 0xf7, 0xe4, 0xda, 0xe7, 0x9c, 0xeb, 0xc0, 0xb9, 0xfc, 0x59, 0xae, 0x92,
	0x25, 0x1f, 0x16, 0xa5, 0xa8, 0x04, 0x85, 0x76, 0x7b, 0xff, 0xbe, 0xf7, 0x47, 0x07, 0x6b, 0x9c,
	0xf1, 0x4d, 0x4a, 0x29, 0x98, 0x79, 0xb2, 0xe5, 0x9e, 0xe6, 0x6b, 0x7d, 0x97, 0xe1, 0x5a, 0x61,
	0xd5, 0xaf, 0x82, 0x7b, 0x7a, 0x8d, 0xa9, 0x35, 0xbd, 0x04, 0xf3, 0x7b, 0x96, 0xa7, 0x9e, 0xe1,
	0x6b, 0xfd, 0xa7, 0xd7, 0xcf, 0x87, 0xfb, 0x66, 0x43, 0x6c, 0xf4, 0x29, 0xcb, 0x53, 0x86, 0x14,
	0xfa, 0x02, 0xec, 0x95, 0x28, 0xb7, 0x49, 0xe5, 0x99, 0xd8, 0xa0, 0xd9, 0xd1, 0x77, 0xe0, 0x14,
	0x42, 0x66, 0x55, 0x26, 0x72, 0xcf, 0xc2, 0x36, 0xcf, 0xba, 0x6d, 0xe6, 0x4d, 0x8d, 0xed, 0x58,
	0xf4, 0x35, 0x40, 0x9e, 0x54, 0xd9, 0x3d, 0x8f, 0xd5, 0x75, 0x6c, 0xec, 0xd6, 0x41, 0xe8, 0x2b,
	0x70, 0x57, 0xea, 0xf0, 0xa9, 0x52, 0x70, 0x8a, 0xe5, 0x3d, 0x40, 0xdf, 0xc0, 0x79, 0x91, 0x94,
	0xc9, 0x96, 0x57, 0xbc, 0x44, 0x86, 0x83, 0x8c, 0x87, 0xa0, 0xea, 0x21, 0x79, 0x99, 0x25, 0x9b,
	0xec, 0x37, 0xf7, 0x5c, 0x5f, 0xeb, 0x3b, 0x6c, 0x0f, 0xf4, 0xfe, 0x69, 0x60, 0xe2, 0x51, 0xc7,
	0x7c, 0xea, 0x37, 0x9e, 0xe8, 0x87, 0x62, 0xd4, 0x37, 0x1d, 0x4b, 0x7c, 0x38, 0x4b, 0xb9, 0x5c,
	0x96, 0x59, 0x81, 0xea, 0x0d, 0x6c, 0xd2, 0x85, 0x14, 0x63, 0x29, 0xf2, 0x8a, 0xe7, 0x15, 0x6a,
	0xad, 0x9d, 0xeb, 0x42, 0xf4, 0x12, 0x6c, 0xd4, 0x26, 0x3d, 0xcb, 0x37, 0xfa, 0x67, 0xd7, 0x17,
	0x07, 0x19, 0xb0, 0x86, 0x40, 0x5f, 0x82, 0xa3, 0x42, 0x43, 0xd1, 0xb5, 0x6b, 0xbb, 0x7d, 0xef,
	0xaf, 0x0e, 0xf6, 0x84, 0x57, 0x6b, 0x91, 0x2a, 0xe9, 0xa2, 0xe0, 0x65, 0x82, 0x77, 0xaa, 0x85,
	0xed, 0x01, 0xa5, 0xb8, 0x48, 0xaa, 0x75, 0x3b, 0x05, 0x6a, 0xad, 0xa2, 0xdd, 0xe2, 0xb7, 0x8d,
	0x84, 0x66, 0xf7, 0x58, 0x9f, 0x79, 0xa8, 0xaf, 0xf5, 0xcf, 0xea, 0xf8, 0xe7, 0xc3, 0xd9, 0x3a,
	0xc9, 0xd3, 0x4d, 0x13, 0x4f, 0x7d, 0xd3, 0x2e, 0x84, 0x11, 0x96, 0x62, 0xc9, 0xa5, 0x14, 0x65,
	0x27, 0xe4, 0x87, 0xa0, 0x1a, 0x93, 0xe5, 0x26, 0xe3, 0x79, 0xd5, 0x49, 0xb9, 0x83, 0xd0, 0x21,
	0xd0, 0x5d, 0xe6, 0x32, 0x6e, 0x8d, 0x71, 0x91, 0x77, 0xa4, 0x42, 0xaf, 0xe0, 0xa2, 0xe4, 0xb2,
	0x10, 0xb9, 0xe4, 0x7b, 0x3a, 0x20, 0xfd, 0xb0, 0xd0, 0xfb, 0x01, 0xd6, 0x44, 0xa4, 0x7c, 0x73,
	0x74, 0x44, 0xde, 0x82, 0xa5, 0x9c, 0x97, 0x9e, 0x8e, 0x99, 0x91, 0xc7, 0x33, 0xc2, 0xea, 0x32,
	0xbd, 0x82, 0xd3, 0xda, 0x4a, 0xe9, 0x19, 0xc8, 0xa4, 0x5d, 0x66, 0x9d, 0x17, 0x6b, 0x29, 0x83,
	0x10, 0xdc, 0xdd, 0xa3, 0xa3, 0x00, 0xf6, 0xd7, 0x30, 0xf8, 0x1c, 0x30, 0x72, 0x42, 0x4f, 0xc1,
	0x98, 0x04, 0x73, 0xa2, 0x51, 0x17, 0xac, 0x80, 0xb1, 0x60, 0x41, 0x74, 0x7a, 0x0e, 0x2e, 0x8b,
	0xc6, 0x11, 0x8b, 0xa6, 0x61, 0x44, 0x0c, 0x45, 0x09, 0xa6, 0x0b, 0x62, 0x0e, 0x7a, 0xe0, 0xb4,
	0x53, 0x8a, 0x3d, 0x62, 0x76, 0x17, 0xc6, 0xe4, 0x44, 0xad, 0x67, 0x37, 0x1f, 0xa3, 0x30, 0x26,
	0xda, 0x20, 0x04, 0xa7, 0x7d, 0x96, 0xd4, 0x01, 0xf3, 0x66, 0x36, 0x5a, 0xd4, 0x8c, 0xdb, 0x28,
	0x18, 0x45, 0x8c, 0x68, 0xf4, 0x09, 0x38, 0xe3, 0x19, 0x9b, 0x8c, 0x82, 0x38, 0x20, 0xba, 0x3a,
	0xf6, 0xcb, 0x5d, 0xc4, 0x16, 0xc4, 0x50, 0xf4, 0x79, 0x10, 0xdf, 0x12, 0xf3, 0x9b, 0x8d, 0xff,
	0x9f, 0x0f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xcf, 0x90, 0x77, 0x90, 0x04, 0x00, 0x00,
}