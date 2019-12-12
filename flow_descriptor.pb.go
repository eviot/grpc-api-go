// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow_descriptor.proto

package gapi

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FlowDescriptor struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                    `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Variables            []*FlowVariableDescriptor `protobuf:"bytes,2,rep,name=variables,proto3" json:"variables,omitempty"`
	Pipes                []*Pipe                   `protobuf:"bytes,3,rep,name=pipes,proto3" json:"pipes,omitempty"`
	UserId               string                    `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt            *UnixTime                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *UnixTime                 `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *FlowDescriptor) Reset()         { *m = FlowDescriptor{} }
func (m *FlowDescriptor) String() string { return proto.CompactTextString(m) }
func (*FlowDescriptor) ProtoMessage()    {}
func (*FlowDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e0f8e9197cd596d, []int{0}
}

func (m *FlowDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowDescriptor.Unmarshal(m, b)
}
func (m *FlowDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowDescriptor.Marshal(b, m, deterministic)
}
func (m *FlowDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowDescriptor.Merge(m, src)
}
func (m *FlowDescriptor) XXX_Size() int {
	return xxx_messageInfo_FlowDescriptor.Size(m)
}
func (m *FlowDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_FlowDescriptor proto.InternalMessageInfo

func (m *FlowDescriptor) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FlowDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FlowDescriptor) GetVariables() []*FlowVariableDescriptor {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *FlowDescriptor) GetPipes() []*Pipe {
	if m != nil {
		return m.Pipes
	}
	return nil
}

func (m *FlowDescriptor) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *FlowDescriptor) GetCreatedAt() *UnixTime {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *FlowDescriptor) GetUpdatedAt() *UnixTime {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type FlowVariableDescriptor struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 VarType  `protobuf:"varint,2,opt,name=type,proto3,enum=schema.VarType" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DefaultValue         *Any     `protobuf:"bytes,4,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowVariableDescriptor) Reset()         { *m = FlowVariableDescriptor{} }
func (m *FlowVariableDescriptor) String() string { return proto.CompactTextString(m) }
func (*FlowVariableDescriptor) ProtoMessage()    {}
func (*FlowVariableDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e0f8e9197cd596d, []int{1}
}

func (m *FlowVariableDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowVariableDescriptor.Unmarshal(m, b)
}
func (m *FlowVariableDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowVariableDescriptor.Marshal(b, m, deterministic)
}
func (m *FlowVariableDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowVariableDescriptor.Merge(m, src)
}
func (m *FlowVariableDescriptor) XXX_Size() int {
	return xxx_messageInfo_FlowVariableDescriptor.Size(m)
}
func (m *FlowVariableDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowVariableDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_FlowVariableDescriptor proto.InternalMessageInfo

func (m *FlowVariableDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FlowVariableDescriptor) GetType() VarType {
	if m != nil {
		return m.Type
	}
	return VarType_INT
}

func (m *FlowVariableDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FlowVariableDescriptor) GetDefaultValue() *Any {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func init() {
	proto.RegisterType((*FlowDescriptor)(nil), "schema.FlowDescriptor")
	proto.RegisterType((*FlowVariableDescriptor)(nil), "schema.FlowVariableDescriptor")
}

func init() { proto.RegisterFile("flow_descriptor.proto", fileDescriptor_1e0f8e9197cd596d) }

var fileDescriptor_1e0f8e9197cd596d = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0x95, 0x34, 0x4d, 0x95, 0x9b, 0xfe, 0xd3, 0xbf, 0x2c, 0x01, 0x56, 0x07, 0x14, 0x95,
	0x25, 0x53, 0x41, 0x61, 0x65, 0x29, 0x42, 0x48, 0x6c, 0x28, 0x2a, 0x1d, 0x58, 0x22, 0xb7, 0x76,
	0xc1, 0x52, 0x3e, 0x2c, 0xc7, 0x69, 0x9b, 0xc7, 0xe0, 0x21, 0x78, 0x4f, 0x14, 0x27, 0x6e, 0x3b,
	0xc0, 0x76, 0xe3, 0x73, 0x7e, 0x27, 0xf7, 0x03, 0x2e, 0xb6, 0x59, 0xb9, 0x4f, 0x29, 0xab, 0x36,
	0x92, 0x0b, 0x55, 0xca, 0xb9, 0x90, 0xa5, 0x2a, 0x91, 0x5b, 0x6d, 0x3e, 0x59, 0x4e, 0xa6, 0x93,
	0xba, 0xe0, 0x87, 0x54, 0xf1, 0x9c, 0x75, 0xc2, 0x14, 0x04, 0x17, 0xa6, 0xf6, 0x48, 0xd1, 0xf4,
	0x65, 0xb0, 0x23, 0x32, 0x55, 0x8d, 0x91, 0x66, 0x5f, 0x36, 0x04, 0xcf, 0x59, 0xb9, 0x7f, 0x3a,
	0x06, 0xa3, 0x00, 0x6c, 0x4e, 0xb1, 0x15, 0x5a, 0x91, 0x97, 0xd8, 0x9c, 0x22, 0x04, 0x4e, 0x41,
	0x72, 0x86, 0x47, 0xfa, 0x45, 0xd7, 0xe8, 0x01, 0xbc, 0x1d, 0x91, 0x9c, 0xac, 0x33, 0x56, 0x61,
	0x3b, 0x1c, 0x44, 0x7e, 0x7c, 0x3d, 0xef, 0x5a, 0x99, 0xb7, 0x71, 0xab, 0x5e, 0x3c, 0xc5, 0x26,
	0x27, 0x00, 0xcd, 0x60, 0xd8, 0x76, 0x57, 0xe1, 0x81, 0x26, 0xc7, 0x86, 0x7c, 0xe5, 0x82, 0x25,
	0x9d, 0x84, 0xae, 0x60, 0x54, 0x57, 0x4c, 0xa6, 0x9c, 0x62, 0x47, 0xff, 0xd8, 0x6d, 0x3f, 0x5f,
	0x28, 0xba, 0x05, 0xd8, 0x48, 0x46, 0x14, 0xa3, 0x29, 0x51, 0x78, 0x18, 0x5a, 0x91, 0x1f, 0xff,
	0x37, 0x09, 0x6f, 0x05, 0x3f, 0x2c, 0x79, 0xce, 0x12, 0xaf, 0xf7, 0x2c, 0x54, 0x0b, 0xd4, 0x82,
	0x1a, 0xc0, 0xfd, 0x0b, 0xe8, 0x3d, 0x0b, 0x35, 0xfb, 0xb6, 0xe0, 0xf2, 0xf7, 0x21, 0x8e, 0xbb,
	0xb0, 0xce, 0x76, 0x71, 0x03, 0x4e, 0xbb, 0x50, 0x6c, 0x87, 0x56, 0x14, 0xc4, 0x13, 0x93, 0xbc,
	0x22, 0x72, 0xd9, 0x08, 0x96, 0x68, 0x11, 0x85, 0xe0, 0x9b, 0xdb, 0xf1, 0xb2, 0xc0, 0x03, 0xcd,
	0x9f, 0x3f, 0xa1, 0x3b, 0xf8, 0x47, 0xd9, 0x96, 0xd4, 0x99, 0x4a, 0x77, 0x24, 0xab, 0x99, 0x1e,
	0xdb, 0x8f, 0x7d, 0x93, 0xb7, 0x28, 0x9a, 0x64, 0xdc, 0x3b, 0x56, 0xad, 0xe1, 0xd1, 0x7d, 0x77,
	0x3e, 0x88, 0xe0, 0x6b, 0x57, 0x9f, 0xf2, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x3c, 0x55,
	0xec, 0x23, 0x02, 0x00, 0x00,
}