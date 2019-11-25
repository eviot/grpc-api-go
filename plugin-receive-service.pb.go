// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin-receive-service.proto

package gapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type InitPipeReq struct {
	Pipe                 *PipeExt `protobuf:"bytes,1,opt,name=pipe,proto3" json:"pipe,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitPipeReq) Reset()         { *m = InitPipeReq{} }
func (m *InitPipeReq) String() string { return proto.CompactTextString(m) }
func (*InitPipeReq) ProtoMessage()    {}
func (*InitPipeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f907c5fb48f6603, []int{0}
}

func (m *InitPipeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitPipeReq.Unmarshal(m, b)
}
func (m *InitPipeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitPipeReq.Marshal(b, m, deterministic)
}
func (m *InitPipeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitPipeReq.Merge(m, src)
}
func (m *InitPipeReq) XXX_Size() int {
	return xxx_messageInfo_InitPipeReq.Size(m)
}
func (m *InitPipeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InitPipeReq.DiscardUnknown(m)
}

var xxx_messageInfo_InitPipeReq proto.InternalMessageInfo

func (m *InitPipeReq) GetPipe() *PipeExt {
	if m != nil {
		return m.Pipe
	}
	return nil
}

type InitPipeRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitPipeRes) Reset()         { *m = InitPipeRes{} }
func (m *InitPipeRes) String() string { return proto.CompactTextString(m) }
func (*InitPipeRes) ProtoMessage()    {}
func (*InitPipeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f907c5fb48f6603, []int{1}
}

func (m *InitPipeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitPipeRes.Unmarshal(m, b)
}
func (m *InitPipeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitPipeRes.Marshal(b, m, deterministic)
}
func (m *InitPipeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitPipeRes.Merge(m, src)
}
func (m *InitPipeRes) XXX_Size() int {
	return xxx_messageInfo_InitPipeRes.Size(m)
}
func (m *InitPipeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_InitPipeRes.DiscardUnknown(m)
}

var xxx_messageInfo_InitPipeRes proto.InternalMessageInfo

type DestroyPipeReq struct {
	PipeID               string   `protobuf:"bytes,1,opt,name=pipeID,proto3" json:"pipeID,omitempty"`
	PipeDescriptorName   string   `protobuf:"bytes,2,opt,name=pipeDescriptorName,proto3" json:"pipeDescriptorName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyPipeReq) Reset()         { *m = DestroyPipeReq{} }
func (m *DestroyPipeReq) String() string { return proto.CompactTextString(m) }
func (*DestroyPipeReq) ProtoMessage()    {}
func (*DestroyPipeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f907c5fb48f6603, []int{2}
}

func (m *DestroyPipeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyPipeReq.Unmarshal(m, b)
}
func (m *DestroyPipeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyPipeReq.Marshal(b, m, deterministic)
}
func (m *DestroyPipeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyPipeReq.Merge(m, src)
}
func (m *DestroyPipeReq) XXX_Size() int {
	return xxx_messageInfo_DestroyPipeReq.Size(m)
}
func (m *DestroyPipeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyPipeReq.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyPipeReq proto.InternalMessageInfo

func (m *DestroyPipeReq) GetPipeID() string {
	if m != nil {
		return m.PipeID
	}
	return ""
}

func (m *DestroyPipeReq) GetPipeDescriptorName() string {
	if m != nil {
		return m.PipeDescriptorName
	}
	return ""
}

type DestroyPipeRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyPipeRes) Reset()         { *m = DestroyPipeRes{} }
func (m *DestroyPipeRes) String() string { return proto.CompactTextString(m) }
func (*DestroyPipeRes) ProtoMessage()    {}
func (*DestroyPipeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f907c5fb48f6603, []int{3}
}

func (m *DestroyPipeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyPipeRes.Unmarshal(m, b)
}
func (m *DestroyPipeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyPipeRes.Marshal(b, m, deterministic)
}
func (m *DestroyPipeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyPipeRes.Merge(m, src)
}
func (m *DestroyPipeRes) XXX_Size() int {
	return xxx_messageInfo_DestroyPipeRes.Size(m)
}
func (m *DestroyPipeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyPipeRes.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyPipeRes proto.InternalMessageInfo

type ReceiveMsgReq struct {
	Pipe                 *PipeExt          `protobuf:"bytes,1,opt,name=pipe,proto3" json:"pipe,omitempty"`
	InputID              string            `protobuf:"bytes,2,opt,name=inputID,proto3" json:"inputID,omitempty"`
	Message              map[string]*Value `protobuf:"bytes,3,rep,name=message,proto3" json:"message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReceiveMsgReq) Reset()         { *m = ReceiveMsgReq{} }
func (m *ReceiveMsgReq) String() string { return proto.CompactTextString(m) }
func (*ReceiveMsgReq) ProtoMessage()    {}
func (*ReceiveMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f907c5fb48f6603, []int{4}
}

func (m *ReceiveMsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveMsgReq.Unmarshal(m, b)
}
func (m *ReceiveMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveMsgReq.Marshal(b, m, deterministic)
}
func (m *ReceiveMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveMsgReq.Merge(m, src)
}
func (m *ReceiveMsgReq) XXX_Size() int {
	return xxx_messageInfo_ReceiveMsgReq.Size(m)
}
func (m *ReceiveMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveMsgReq proto.InternalMessageInfo

func (m *ReceiveMsgReq) GetPipe() *PipeExt {
	if m != nil {
		return m.Pipe
	}
	return nil
}

func (m *ReceiveMsgReq) GetInputID() string {
	if m != nil {
		return m.InputID
	}
	return ""
}

func (m *ReceiveMsgReq) GetMessage() map[string]*Value {
	if m != nil {
		return m.Message
	}
	return nil
}

type ReceiveMsgRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveMsgRes) Reset()         { *m = ReceiveMsgRes{} }
func (m *ReceiveMsgRes) String() string { return proto.CompactTextString(m) }
func (*ReceiveMsgRes) ProtoMessage()    {}
func (*ReceiveMsgRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f907c5fb48f6603, []int{5}
}

func (m *ReceiveMsgRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveMsgRes.Unmarshal(m, b)
}
func (m *ReceiveMsgRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveMsgRes.Marshal(b, m, deterministic)
}
func (m *ReceiveMsgRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveMsgRes.Merge(m, src)
}
func (m *ReceiveMsgRes) XXX_Size() int {
	return xxx_messageInfo_ReceiveMsgRes.Size(m)
}
func (m *ReceiveMsgRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveMsgRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveMsgRes proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InitPipeReq)(nil), "schema.InitPipeReq")
	proto.RegisterType((*InitPipeRes)(nil), "schema.InitPipeRes")
	proto.RegisterType((*DestroyPipeReq)(nil), "schema.DestroyPipeReq")
	proto.RegisterType((*DestroyPipeRes)(nil), "schema.DestroyPipeRes")
	proto.RegisterType((*ReceiveMsgReq)(nil), "schema.ReceiveMsgReq")
	proto.RegisterMapType((map[string]*Value)(nil), "schema.ReceiveMsgReq.MessageEntry")
	proto.RegisterType((*ReceiveMsgRes)(nil), "schema.ReceiveMsgRes")
}

func init() { proto.RegisterFile("plugin-receive-service.proto", fileDescriptor_6f907c5fb48f6603) }

var fileDescriptor_6f907c5fb48f6603 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x95, 0x81, 0x9a, 0x76, 0x5c, 0x0a, 0x9a, 0xb6, 0xc8, 0xb2, 0x7a, 0x40, 0xe6, 0xc2, 0x05,
	0x1f, 0xdc, 0x1e, 0x2a, 0xd4, 0x5e, 0x90, 0x39, 0xf8, 0x40, 0x84, 0x1c, 0x29, 0x8a, 0x72, 0x73,
	0xac, 0x91, 0xb3, 0x0a, 0xb6, 0x37, 0xbb, 0x0b, 0x0a, 0xbf, 0x21, 0x3f, 0x30, 0x7f, 0x27, 0xf2,
	0xc7, 0x2a, 0x26, 0xf2, 0x21, 0xb9, 0xed, 0xbc, 0xf7, 0xe6, 0xcd, 0xd7, 0xc2, 0x2f, 0xbe, 0x3f,
	0xa4, 0x2c, 0x5f, 0x0a, 0x4a, 0x88, 0x1d, 0x69, 0x29, 0x49, 0x1c, 0x59, 0x42, 0x1e, 0x17, 0x85,
	0x2a, 0xd0, 0x94, 0xc9, 0x1d, 0x65, 0xb1, 0x03, 0x9c, 0xf1, 0x06, 0x73, 0xac, 0x63, 0xbc, 0x3f,
	0x34, 0x81, 0xeb, 0x83, 0x15, 0xe6, 0x4c, 0xed, 0x18, 0xa7, 0x88, 0x1e, 0x70, 0x0e, 0x83, 0x52,
	0x69, 0x1b, 0x33, 0x63, 0x61, 0xf9, 0x63, 0xaf, 0x4e, 0xf7, 0x4a, 0x7a, 0xf3, 0xa8, 0xa2, 0x8a,
	0x74, 0x47, 0xed, 0x1c, 0xe9, 0x5e, 0xc3, 0xb7, 0x80, 0xa4, 0x12, 0xc5, 0x49, 0xbb, 0x4c, 0xc1,
	0x2c, 0x85, 0x61, 0x50, 0xf9, 0x7c, 0x89, 0x9a, 0x08, 0x3d, 0xc0, 0xf2, 0x15, 0x90, 0x4c, 0x04,
	0xe3, 0xaa, 0x10, 0x17, 0x71, 0x46, 0x76, 0xaf, 0xd2, 0x74, 0x30, 0xee, 0xe4, 0x8d, 0xb3, 0x74,
	0x9f, 0x0d, 0x18, 0x45, 0xf5, 0xa4, 0x5b, 0x99, 0xbe, 0xb7, 0x63, 0xb4, 0x61, 0xc8, 0x72, 0x7e,
	0x50, 0x61, 0xd0, 0x54, 0xd3, 0x21, 0xfe, 0x83, 0x61, 0x46, 0x52, 0xc6, 0x29, 0xd9, 0xfd, 0x59,
	0x7f, 0x61, 0xf9, 0xae, 0x76, 0x38, 0x2b, 0xe3, 0x6d, 0x6b, 0xd1, 0x26, 0x57, 0xe2, 0x14, 0xe9,
	0x14, 0x27, 0x84, 0xaf, 0x6d, 0x02, 0x27, 0xd0, 0xbf, 0xa7, 0x53, 0x33, 0x75, 0xf9, 0xc4, 0x39,
	0x7c, 0xaa, 0xd6, 0x5d, 0xd5, 0xb5, 0xfc, 0x91, 0x76, 0xbf, 0x2a, 0xc1, 0xa8, 0xe6, 0x56, 0xbd,
	0xbf, 0x86, 0x3b, 0x3e, 0x1f, 0x4c, 0xfa, 0x4f, 0x3d, 0xf8, 0xb1, 0xab, 0x6e, 0xdb, 0xe0, 0x97,
	0xf5, 0x65, 0xf1, 0x0f, 0x7c, 0xd6, 0xeb, 0xc7, 0xef, 0xda, 0xaf, 0x75, 0x44, 0xa7, 0x03, 0x94,
	0xf8, 0x1f, 0xac, 0xd6, 0x2e, 0x71, 0xaa, 0x35, 0xe7, 0xa7, 0x73, 0xba, 0x71, 0x89, 0x2b, 0x80,
	0xd7, 0xf6, 0xf0, 0x67, 0xe7, 0x92, 0x9c, 0x4e, 0x58, 0xe2, 0x1a, 0x26, 0x7a, 0x04, 0x25, 0x28,
	0xce, 0x3e, 0xec, 0xb0, 0x30, 0xd6, 0xe6, 0xcd, 0x20, 0x8d, 0x39, 0xbb, 0x35, 0xab, 0x6f, 0xfb,
	0xfb, 0x25, 0x00, 0x00, 0xff, 0xff, 0x72, 0xf9, 0xa0, 0x88, 0xf7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PluginReceiveServiceClient is the client API for PluginReceiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginReceiveServiceClient interface {
	InitPipe(ctx context.Context, in *InitPipeReq, opts ...grpc.CallOption) (*InitPipeRes, error)
	DestroyPipe(ctx context.Context, in *DestroyPipeReq, opts ...grpc.CallOption) (*DestroyPipeRes, error)
	ReceiveMsg(ctx context.Context, in *ReceiveMsgReq, opts ...grpc.CallOption) (*ReceiveMsgRes, error)
	ReceiveStreamMsg(ctx context.Context, opts ...grpc.CallOption) (PluginReceiveService_ReceiveStreamMsgClient, error)
}

type pluginReceiveServiceClient struct {
	cc *grpc.ClientConn
}

func NewPluginReceiveServiceClient(cc *grpc.ClientConn) PluginReceiveServiceClient {
	return &pluginReceiveServiceClient{cc}
}

func (c *pluginReceiveServiceClient) InitPipe(ctx context.Context, in *InitPipeReq, opts ...grpc.CallOption) (*InitPipeRes, error) {
	out := new(InitPipeRes)
	err := c.cc.Invoke(ctx, "/schema.PluginReceiveService/InitPipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginReceiveServiceClient) DestroyPipe(ctx context.Context, in *DestroyPipeReq, opts ...grpc.CallOption) (*DestroyPipeRes, error) {
	out := new(DestroyPipeRes)
	err := c.cc.Invoke(ctx, "/schema.PluginReceiveService/DestroyPipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginReceiveServiceClient) ReceiveMsg(ctx context.Context, in *ReceiveMsgReq, opts ...grpc.CallOption) (*ReceiveMsgRes, error) {
	out := new(ReceiveMsgRes)
	err := c.cc.Invoke(ctx, "/schema.PluginReceiveService/ReceiveMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginReceiveServiceClient) ReceiveStreamMsg(ctx context.Context, opts ...grpc.CallOption) (PluginReceiveService_ReceiveStreamMsgClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PluginReceiveService_serviceDesc.Streams[0], "/schema.PluginReceiveService/ReceiveStreamMsg", opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginReceiveServiceReceiveStreamMsgClient{stream}
	return x, nil
}

type PluginReceiveService_ReceiveStreamMsgClient interface {
	Send(*ReceiveMsgReq) error
	CloseAndRecv() (*ReceiveMsgRes, error)
	grpc.ClientStream
}

type pluginReceiveServiceReceiveStreamMsgClient struct {
	grpc.ClientStream
}

func (x *pluginReceiveServiceReceiveStreamMsgClient) Send(m *ReceiveMsgReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pluginReceiveServiceReceiveStreamMsgClient) CloseAndRecv() (*ReceiveMsgRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ReceiveMsgRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PluginReceiveServiceServer is the server API for PluginReceiveService service.
type PluginReceiveServiceServer interface {
	InitPipe(context.Context, *InitPipeReq) (*InitPipeRes, error)
	DestroyPipe(context.Context, *DestroyPipeReq) (*DestroyPipeRes, error)
	ReceiveMsg(context.Context, *ReceiveMsgReq) (*ReceiveMsgRes, error)
	ReceiveStreamMsg(PluginReceiveService_ReceiveStreamMsgServer) error
}

// UnimplementedPluginReceiveServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPluginReceiveServiceServer struct {
}

func (*UnimplementedPluginReceiveServiceServer) InitPipe(ctx context.Context, req *InitPipeReq) (*InitPipeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitPipe not implemented")
}
func (*UnimplementedPluginReceiveServiceServer) DestroyPipe(ctx context.Context, req *DestroyPipeReq) (*DestroyPipeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyPipe not implemented")
}
func (*UnimplementedPluginReceiveServiceServer) ReceiveMsg(ctx context.Context, req *ReceiveMsgReq) (*ReceiveMsgRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveMsg not implemented")
}
func (*UnimplementedPluginReceiveServiceServer) ReceiveStreamMsg(srv PluginReceiveService_ReceiveStreamMsgServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveStreamMsg not implemented")
}

func RegisterPluginReceiveServiceServer(s *grpc.Server, srv PluginReceiveServiceServer) {
	s.RegisterService(&_PluginReceiveService_serviceDesc, srv)
}

func _PluginReceiveService_InitPipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitPipeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginReceiveServiceServer).InitPipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.PluginReceiveService/InitPipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginReceiveServiceServer).InitPipe(ctx, req.(*InitPipeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginReceiveService_DestroyPipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyPipeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginReceiveServiceServer).DestroyPipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.PluginReceiveService/DestroyPipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginReceiveServiceServer).DestroyPipe(ctx, req.(*DestroyPipeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginReceiveService_ReceiveMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginReceiveServiceServer).ReceiveMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.PluginReceiveService/ReceiveMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginReceiveServiceServer).ReceiveMsg(ctx, req.(*ReceiveMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginReceiveService_ReceiveStreamMsg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PluginReceiveServiceServer).ReceiveStreamMsg(&pluginReceiveServiceReceiveStreamMsgServer{stream})
}

type PluginReceiveService_ReceiveStreamMsgServer interface {
	SendAndClose(*ReceiveMsgRes) error
	Recv() (*ReceiveMsgReq, error)
	grpc.ServerStream
}

type pluginReceiveServiceReceiveStreamMsgServer struct {
	grpc.ServerStream
}

func (x *pluginReceiveServiceReceiveStreamMsgServer) SendAndClose(m *ReceiveMsgRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pluginReceiveServiceReceiveStreamMsgServer) Recv() (*ReceiveMsgReq, error) {
	m := new(ReceiveMsgReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PluginReceiveService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schema.PluginReceiveService",
	HandlerType: (*PluginReceiveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitPipe",
			Handler:    _PluginReceiveService_InitPipe_Handler,
		},
		{
			MethodName: "DestroyPipe",
			Handler:    _PluginReceiveService_DestroyPipe_Handler,
		},
		{
			MethodName: "ReceiveMsg",
			Handler:    _PluginReceiveService_ReceiveMsg_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveStreamMsg",
			Handler:       _PluginReceiveService_ReceiveStreamMsg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "plugin-receive-service.proto",
}
