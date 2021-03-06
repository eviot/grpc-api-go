// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device_type.proto

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

type DeviceType struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FlowDescriptorId     string    `protobuf:"bytes,4,opt,name=flow_descriptor_id,json=flowDescriptorId,proto3" json:"flow_descriptor_id,omitempty"`
	UserId               string    `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt            *UnixTime `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeviceType) Reset()         { *m = DeviceType{} }
func (m *DeviceType) String() string { return proto.CompactTextString(m) }
func (*DeviceType) ProtoMessage()    {}
func (*DeviceType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a230965e42286ec, []int{0}
}

func (m *DeviceType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceType.Unmarshal(m, b)
}
func (m *DeviceType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceType.Marshal(b, m, deterministic)
}
func (m *DeviceType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceType.Merge(m, src)
}
func (m *DeviceType) XXX_Size() int {
	return xxx_messageInfo_DeviceType.Size(m)
}
func (m *DeviceType) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceType.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceType proto.InternalMessageInfo

func (m *DeviceType) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceType) GetFlowDescriptorId() string {
	if m != nil {
		return m.FlowDescriptorId
	}
	return ""
}

func (m *DeviceType) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DeviceType) GetCreatedAt() *UnixTime {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceType)(nil), "schema.DeviceType")
}

func init() { proto.RegisterFile("device_type.proto", fileDescriptor_1a230965e42286ec) }

var fileDescriptor_1a230965e42286ec = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xce, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x71, 0xb2, 0xae, 0x91, 0x8e, 0xa0, 0x75, 0x2e, 0x06, 0x4f, 0xc5, 0x53, 0x0f, 0xb2,
	0x82, 0x3e, 0x81, 0xd2, 0x4b, 0xaf, 0xa5, 0x5e, 0xbc, 0x84, 0x98, 0x19, 0x75, 0xc0, 0x6c, 0x42,
	0x36, 0xd5, 0xf6, 0x8d, 0x7c, 0x4c, 0xd9, 0xb8, 0xf4, 0x96, 0xcc, 0xff, 0x77, 0xf8, 0xe0, 0x8a,
	0xf8, 0x5b, 0x3c, 0xdb, 0x72, 0x48, 0xdc, 0xa5, 0x1c, 0x4b, 0x44, 0x3d, 0xf8, 0x4f, 0x0e, 0xee,
	0xe6, 0x72, 0xd7, 0xcb, 0xde, 0x16, 0x09, 0x53, 0xb8, 0xfd, 0x55, 0x00, 0xab, 0xca, 0xb7, 0x87,
	0xc4, 0x78, 0x01, 0x8d, 0x90, 0x51, 0x0b, 0xb5, 0x9c, 0x6d, 0x1a, 0x21, 0x44, 0x68, 0x7b, 0x17,
	0xd8, 0x34, 0xf5, 0x52, 0xdf, 0x78, 0x07, 0xf8, 0xfe, 0x15, 0x7f, 0x2c, 0xf1, 0xe0, 0xb3, 0xa4,
	0x12, 0xb3, 0x15, 0x32, 0x6d, 0x15, 0xf3, 0xb1, 0xac, 0x8e, 0x61, 0x4d, 0x78, 0x0d, 0x67, 0xbb,
	0x81, 0x2b, 0x39, 0xa9, 0x44, 0x8f, 0xdf, 0x35, 0xe1, 0x3d, 0x80, 0xcf, 0xec, 0x0a, 0x93, 0x75,
	0xc5, 0x9c, 0x2e, 0xd4, 0xf2, 0xfc, 0x61, 0xde, 0xfd, 0xef, 0xec, 0x5e, 0x7a, 0xd9, 0x6f, 0x25,
	0xf0, 0x66, 0x36, 0x99, 0xa7, 0xf2, 0xac, 0x5f, 0xdb, 0x0f, 0x97, 0xe4, 0x4d, 0xd7, 0xe5, 0x8f,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x4a, 0x0f, 0xa1, 0xe7, 0x00, 0x00, 0x00,
}
