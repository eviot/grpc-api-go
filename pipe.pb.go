// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pipe.proto

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

type Pipe struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PipeDescriptorId     string          `protobuf:"bytes,2,opt,name=pipe_descriptor_id,json=pipeDescriptorId,proto3" json:"pipe_descriptor_id,omitempty"`
	X                    int32           `protobuf:"varint,7,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32           `protobuf:"varint,8,opt,name=y,proto3" json:"y,omitempty"`
	Params               map[string]*Any `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Inputs               []string        `protobuf:"bytes,5,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []string        `protobuf:"bytes,6,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Binds                []*PipeBind     `protobuf:"bytes,4,rep,name=binds,proto3" json:"binds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Pipe) Reset()         { *m = Pipe{} }
func (m *Pipe) String() string { return proto.CompactTextString(m) }
func (*Pipe) ProtoMessage()    {}
func (*Pipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{0}
}

func (m *Pipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipe.Unmarshal(m, b)
}
func (m *Pipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipe.Marshal(b, m, deterministic)
}
func (m *Pipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipe.Merge(m, src)
}
func (m *Pipe) XXX_Size() int {
	return xxx_messageInfo_Pipe.Size(m)
}
func (m *Pipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipe.DiscardUnknown(m)
}

var xxx_messageInfo_Pipe proto.InternalMessageInfo

func (m *Pipe) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Pipe) GetPipeDescriptorId() string {
	if m != nil {
		return m.PipeDescriptorId
	}
	return ""
}

func (m *Pipe) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Pipe) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Pipe) GetParams() map[string]*Any {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Pipe) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Pipe) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Pipe) GetBinds() []*PipeBind {
	if m != nil {
		return m.Binds
	}
	return nil
}

type PipeBind struct {
	FromOutputId         string   `protobuf:"bytes,1,opt,name=from_output_id,json=fromOutputId,proto3" json:"from_output_id,omitempty"`
	ToPipeId             string   `protobuf:"bytes,2,opt,name=to_pipe_id,json=toPipeId,proto3" json:"to_pipe_id,omitempty"`
	ToInputId            string   `protobuf:"bytes,3,opt,name=to_input_id,json=toInputId,proto3" json:"to_input_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PipeBind) Reset()         { *m = PipeBind{} }
func (m *PipeBind) String() string { return proto.CompactTextString(m) }
func (*PipeBind) ProtoMessage()    {}
func (*PipeBind) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{1}
}

func (m *PipeBind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipeBind.Unmarshal(m, b)
}
func (m *PipeBind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipeBind.Marshal(b, m, deterministic)
}
func (m *PipeBind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipeBind.Merge(m, src)
}
func (m *PipeBind) XXX_Size() int {
	return xxx_messageInfo_PipeBind.Size(m)
}
func (m *PipeBind) XXX_DiscardUnknown() {
	xxx_messageInfo_PipeBind.DiscardUnknown(m)
}

var xxx_messageInfo_PipeBind proto.InternalMessageInfo

func (m *PipeBind) GetFromOutputId() string {
	if m != nil {
		return m.FromOutputId
	}
	return ""
}

func (m *PipeBind) GetToPipeId() string {
	if m != nil {
		return m.ToPipeId
	}
	return ""
}

func (m *PipeBind) GetToInputId() string {
	if m != nil {
		return m.ToInputId
	}
	return ""
}

type PipeExt struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PipeDescriptorId     string          `protobuf:"bytes,2,opt,name=pipe_descriptor_id,json=pipeDescriptorId,proto3" json:"pipe_descriptor_id,omitempty"`
	Params               map[string]*Any `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Inputs               []string        `protobuf:"bytes,9,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []string        `protobuf:"bytes,10,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Binds                []*PipeBindExt  `protobuf:"bytes,5,rep,name=binds,proto3" json:"binds,omitempty"`
	PluginUri            string          `protobuf:"bytes,6,opt,name=plugin_uri,json=pluginUri,proto3" json:"plugin_uri,omitempty"`
	PluginId             string          `protobuf:"bytes,11,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	FlowId               string          `protobuf:"bytes,7,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	PipeDescriptorName   string          `protobuf:"bytes,8,opt,name=pipe_descriptor_name,json=pipeDescriptorName,proto3" json:"pipe_descriptor_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PipeExt) Reset()         { *m = PipeExt{} }
func (m *PipeExt) String() string { return proto.CompactTextString(m) }
func (*PipeExt) ProtoMessage()    {}
func (*PipeExt) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{2}
}

func (m *PipeExt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipeExt.Unmarshal(m, b)
}
func (m *PipeExt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipeExt.Marshal(b, m, deterministic)
}
func (m *PipeExt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipeExt.Merge(m, src)
}
func (m *PipeExt) XXX_Size() int {
	return xxx_messageInfo_PipeExt.Size(m)
}
func (m *PipeExt) XXX_DiscardUnknown() {
	xxx_messageInfo_PipeExt.DiscardUnknown(m)
}

var xxx_messageInfo_PipeExt proto.InternalMessageInfo

func (m *PipeExt) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PipeExt) GetPipeDescriptorId() string {
	if m != nil {
		return m.PipeDescriptorId
	}
	return ""
}

func (m *PipeExt) GetParams() map[string]*Any {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *PipeExt) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PipeExt) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *PipeExt) GetBinds() []*PipeBindExt {
	if m != nil {
		return m.Binds
	}
	return nil
}

func (m *PipeExt) GetPluginUri() string {
	if m != nil {
		return m.PluginUri
	}
	return ""
}

func (m *PipeExt) GetPluginId() string {
	if m != nil {
		return m.PluginId
	}
	return ""
}

func (m *PipeExt) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *PipeExt) GetPipeDescriptorName() string {
	if m != nil {
		return m.PipeDescriptorName
	}
	return ""
}

type PipeBindExt struct {
	FromOutputId         string   `protobuf:"bytes,1,opt,name=from_output_id,json=fromOutputId,proto3" json:"from_output_id,omitempty"`
	ToPipeId             string   `protobuf:"bytes,2,opt,name=to_pipe_id,json=toPipeId,proto3" json:"to_pipe_id,omitempty"`
	ToPipe               *PipeExt `protobuf:"bytes,4,opt,name=to_pipe,json=toPipe,proto3" json:"to_pipe,omitempty"`
	ToInputId            string   `protobuf:"bytes,3,opt,name=to_input_id,json=toInputId,proto3" json:"to_input_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PipeBindExt) Reset()         { *m = PipeBindExt{} }
func (m *PipeBindExt) String() string { return proto.CompactTextString(m) }
func (*PipeBindExt) ProtoMessage()    {}
func (*PipeBindExt) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{3}
}

func (m *PipeBindExt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipeBindExt.Unmarshal(m, b)
}
func (m *PipeBindExt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipeBindExt.Marshal(b, m, deterministic)
}
func (m *PipeBindExt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipeBindExt.Merge(m, src)
}
func (m *PipeBindExt) XXX_Size() int {
	return xxx_messageInfo_PipeBindExt.Size(m)
}
func (m *PipeBindExt) XXX_DiscardUnknown() {
	xxx_messageInfo_PipeBindExt.DiscardUnknown(m)
}

var xxx_messageInfo_PipeBindExt proto.InternalMessageInfo

func (m *PipeBindExt) GetFromOutputId() string {
	if m != nil {
		return m.FromOutputId
	}
	return ""
}

func (m *PipeBindExt) GetToPipeId() string {
	if m != nil {
		return m.ToPipeId
	}
	return ""
}

func (m *PipeBindExt) GetToPipe() *PipeExt {
	if m != nil {
		return m.ToPipe
	}
	return nil
}

func (m *PipeBindExt) GetToInputId() string {
	if m != nil {
		return m.ToInputId
	}
	return ""
}

func init() {
	proto.RegisterType((*Pipe)(nil), "schema.Pipe")
	proto.RegisterMapType((map[string]*Any)(nil), "schema.Pipe.ParamsEntry")
	proto.RegisterType((*PipeBind)(nil), "schema.PipeBind")
	proto.RegisterType((*PipeExt)(nil), "schema.PipeExt")
	proto.RegisterMapType((map[string]*Any)(nil), "schema.PipeExt.ParamsEntry")
	proto.RegisterType((*PipeBindExt)(nil), "schema.PipeBindExt")
}

func init() { proto.RegisterFile("pipe.proto", fileDescriptor_d1979e1a5fdc07ed) }

var fileDescriptor_d1979e1a5fdc07ed = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0x49, 0x72, 0xd9, 0x5c, 0x26, 0xa5, 0x1e, 0xa3, 0xe8, 0xd2, 0xaa, 0x9c, 0x87, 0x48,
	0x04, 0x39, 0x4a, 0xfb, 0x22, 0xbe, 0x59, 0x3c, 0x21, 0x2f, 0x5a, 0x02, 0xbe, 0xf8, 0x12, 0xd2,
	0x6e, 0x5a, 0x17, 0x2f, 0xbb, 0x4b, 0xb2, 0xd1, 0xe4, 0xb3, 0xf8, 0x05, 0xfc, 0x00, 0x7e, 0x40,
	0xd9, 0xdd, 0xdc, 0x71, 0x3d, 0x90, 0x3e, 0xd8, 0xb7, 0xcc, 0xfc, 0x76, 0x67, 0x76, 0xfe, 0xff,
	0x09, 0x80, 0xe2, 0xaa, 0x5a, 0xaa, 0x46, 0x6a, 0x89, 0xa4, 0xbd, 0xfa, 0x56, 0xd5, 0xe5, 0x51,
	0x5c, 0x8a, 0xc1, 0xa5, 0x16, 0x7f, 0x7c, 0x98, 0x5c, 0x70, 0x55, 0xe1, 0x21, 0xf8, 0x9c, 0x51,
	0x6f, 0xee, 0xa5, 0x71, 0xee, 0x73, 0x86, 0x6f, 0x00, 0xcd, 0xcd, 0x82, 0x55, 0xed, 0x55, 0xc3,
	0x95, 0x96, 0x4d, 0xc1, 0x19, 0xf5, 0x2d, 0x9f, 0x19, 0xf2, 0x61, 0x0b, 0x32, 0x86, 0x07, 0xe0,
	0xf5, 0x34, 0x9a, 0x7b, 0x69, 0x98, 0x7b, 0xbd, 0x89, 0x06, 0x3a, 0x75, 0xd1, 0x80, 0x27, 0x40,
	0x54, 0xd9, 0x94, 0x75, 0x4b, 0x83, 0x79, 0x90, 0x26, 0xa7, 0x74, 0xe9, 0x9e, 0xb1, 0x34, 0x7d,
	0x97, 0x17, 0x16, 0xad, 0x84, 0x6e, 0x86, 0x7c, 0x3c, 0x87, 0x8f, 0x81, 0x70, 0xa1, 0x3a, 0xdd,
	0xd2, 0x70, 0x1e, 0xa4, 0x71, 0x3e, 0x46, 0x48, 0x21, 0x92, 0x9d, 0xb6, 0x80, 0x58, 0xb0, 0x09,
	0xf1, 0x15, 0x84, 0x97, 0x5c, 0xb0, 0x96, 0x4e, 0x6c, 0x8b, 0xd9, 0x6e, 0x8b, 0x73, 0x2e, 0x58,
	0xee, 0xf0, 0xd1, 0x47, 0x48, 0x76, 0x1a, 0xe2, 0x0c, 0x82, 0xef, 0xd5, 0x30, 0x4e, 0x6d, 0x3e,
	0xf1, 0x05, 0x84, 0x3f, 0xca, 0x75, 0x57, 0xd9, 0x49, 0x93, 0xd3, 0x64, 0x53, 0xe8, 0xbd, 0x18,
	0x72, 0x47, 0xde, 0xf9, 0x6f, 0xbd, 0x85, 0x80, 0xe9, 0xa6, 0x34, 0xbe, 0x84, 0xc3, 0xeb, 0x46,
	0xd6, 0x85, 0x7b, 0x4b, 0xb1, 0x55, 0xf1, 0xc0, 0x64, 0x3f, 0xdb, 0x64, 0xc6, 0xf0, 0x29, 0x80,
	0x96, 0x85, 0x95, 0x74, 0xab, 0xe3, 0x54, 0x4b, 0x53, 0x25, 0x63, 0xf8, 0x1c, 0x12, 0x2d, 0x0b,
	0x3b, 0xa6, 0xc1, 0x81, 0xc5, 0xb1, 0x96, 0x99, 0xb0, 0xb7, 0x17, 0xbf, 0x03, 0x88, 0xcc, 0xd1,
	0x55, 0xaf, 0xff, 0xd3, 0xa9, 0xb3, 0xad, 0x1b, 0x4e, 0xaa, 0xe3, 0x5d, 0xa9, 0x56, 0xbd, 0xbe,
	0xc3, 0x90, 0xf8, 0x5f, 0x86, 0xc0, 0x6d, 0x43, 0x5e, 0x6f, 0x0c, 0x09, 0x6d, 0x97, 0x87, 0xfb,
	0x86, 0xac, 0x7a, 0x3d, 0x7a, 0x82, 0xcf, 0x00, 0xd4, 0xba, 0xbb, 0xe1, 0xa2, 0xe8, 0x1a, 0x4e,
	0x89, 0x1b, 0xdd, 0x65, 0xbe, 0x34, 0x1c, 0x8f, 0x61, 0x0c, 0xcc, 0x54, 0x89, 0xd3, 0xcd, 0x25,
	0x32, 0x86, 0x4f, 0x20, 0xba, 0x5e, 0xcb, 0x9f, 0x06, 0x45, 0x16, 0x11, 0x13, 0x66, 0x0c, 0x4f,
	0xe0, 0xd1, 0xbe, 0x28, 0xa2, 0xac, 0x2b, 0xbb, 0x95, 0x71, 0x8e, 0xb7, 0x65, 0xf9, 0x54, 0xd6,
	0xd5, 0xbd, 0xad, 0xc6, 0x2f, 0x0f, 0x92, 0x9d, 0x29, 0xef, 0x65, 0x3d, 0x52, 0x88, 0x46, 0x4a,
	0x27, 0xb6, 0xf9, 0x83, 0x3d, 0xd7, 0x72, 0xe2, 0xce, 0xde, 0xb5, 0x48, 0xe7, 0xe4, 0xeb, 0xe4,
	0xa6, 0x54, 0xfc, 0x92, 0xd8, 0xdf, 0xff, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xa8,
	0x3f, 0xf8, 0x1f, 0x04, 0x00, 0x00,
}
