// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin.proto

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

type Plugin struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId               string    `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PluginUri            string    `protobuf:"bytes,5,opt,name=plugin_uri,json=pluginUri,proto3" json:"plugin_uri,omitempty"`
	Color                string    `protobuf:"bytes,8,opt,name=color,proto3" json:"color,omitempty"`
	CreatedAt            *UnixTime `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *UnixTime `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Plugin) Reset()         { *m = Plugin{} }
func (m *Plugin) String() string { return proto.CompactTextString(m) }
func (*Plugin) ProtoMessage()    {}
func (*Plugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{0}
}

func (m *Plugin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plugin.Unmarshal(m, b)
}
func (m *Plugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plugin.Marshal(b, m, deterministic)
}
func (m *Plugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plugin.Merge(m, src)
}
func (m *Plugin) XXX_Size() int {
	return xxx_messageInfo_Plugin.Size(m)
}
func (m *Plugin) XXX_DiscardUnknown() {
	xxx_messageInfo_Plugin.DiscardUnknown(m)
}

var xxx_messageInfo_Plugin proto.InternalMessageInfo

func (m *Plugin) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Plugin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plugin) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Plugin) GetPluginUri() string {
	if m != nil {
		return m.PluginUri
	}
	return ""
}

func (m *Plugin) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Plugin) GetCreatedAt() *UnixTime {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Plugin) GetUpdatedAt() *UnixTime {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Plugin)(nil), "schema.Plugin")
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor_22a625af4bc1cc87) }

var fileDescriptor_22a625af4bc1cc87 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0xc9, 0xb2, 0x4d, 0xdd, 0x51, 0x54, 0x06, 0xc1, 0x20, 0x08, 0xc5, 0x53, 0x4f, 0x2b,
	0xe8, 0x13, 0xd4, 0x9b, 0x37, 0x29, 0xf6, 0xe2, 0x25, 0xc4, 0x4d, 0xa8, 0x03, 0xcd, 0x07, 0xd9,
	0x04, 0xf6, 0x61, 0x7d, 0x18, 0xd9, 0x64, 0x3d, 0xf6, 0x36, 0xff, 0x8f, 0x1f, 0x33, 0x03, 0x57,
	0xe1, 0x94, 0x8f, 0xe4, 0xfa, 0x10, 0x7d, 0xf2, 0xc8, 0xc7, 0xe1, 0xc7, 0x58, 0xf5, 0x70, 0x93,
	0x1d, 0x4d, 0x32, 0x91, 0x35, 0x35, 0x78, 0xfa, 0x65, 0xc0, 0x3f, 0x4a, 0x13, 0xaf, 0xa1, 0x21,
	0x2d, 0xd8, 0x86, 0x6d, 0xbb, 0x7d, 0x43, 0x1a, 0x11, 0x5a, 0xa7, 0xac, 0x11, 0x4d, 0x71, 0xca,
	0x8c, 0xf7, 0xb0, 0xce, 0xa3, 0x89, 0x92, 0xb4, 0x68, 0x8b, 0xcd, 0x67, 0xf9, 0xae, 0xf1, 0x11,
	0xa0, 0x2e, 0x94, 0x39, 0x92, 0x58, 0x95, 0xac, 0xab, 0xce, 0x21, 0x12, 0xde, 0xc1, 0x6a, 0xf0,
	0x27, 0x1f, 0xc5, 0x45, 0x49, 0xaa, 0xc0, 0x67, 0x80, 0x21, 0x1a, 0x95, 0x8c, 0x96, 0x2a, 0x09,
	0xbe, 0x61, 0xdb, 0xcb, 0x97, 0xdb, 0xbe, 0x9e, 0xda, 0x1f, 0x1c, 0x4d, 0x9f, 0x64, 0xcd, 0xbe,
	0x5b, 0x3a, 0xbb, 0x34, 0x03, 0x39, 0xe8, 0x7f, 0x60, 0x7d, 0x0e, 0x58, 0x3a, 0xbb, 0xf4, 0xc6,
	0xbf, 0xda, 0xa3, 0x0a, 0xf4, 0xcd, 0xcb, 0xb7, 0xaf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3d,
	0x06, 0xc3, 0xcc, 0x16, 0x01, 0x00, 0x00,
}
