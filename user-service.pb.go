// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user-service.proto

package gapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type GetUserReq struct {
	// Types that are valid to be assigned to Filter:
	//	*GetUserReq_ID
	//	*GetUserReq_Email
	//	*GetUserReq_Token
	Filter               isGetUserReq_Filter `protobuf_oneof:"filter"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetUserReq) Reset()         { *m = GetUserReq{} }
func (m *GetUserReq) String() string { return proto.CompactTextString(m) }
func (*GetUserReq) ProtoMessage()    {}
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{0}
}

func (m *GetUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserReq.Unmarshal(m, b)
}
func (m *GetUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserReq.Marshal(b, m, deterministic)
}
func (m *GetUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserReq.Merge(m, src)
}
func (m *GetUserReq) XXX_Size() int {
	return xxx_messageInfo_GetUserReq.Size(m)
}
func (m *GetUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserReq proto.InternalMessageInfo

type isGetUserReq_Filter interface {
	isGetUserReq_Filter()
}

type GetUserReq_ID struct {
	ID string `protobuf:"bytes,1,opt,name=ID,proto3,oneof"`
}

type GetUserReq_Email struct {
	Email string `protobuf:"bytes,2,opt,name=email,proto3,oneof"`
}

type GetUserReq_Token struct {
	Token string `protobuf:"bytes,3,opt,name=token,proto3,oneof"`
}

func (*GetUserReq_ID) isGetUserReq_Filter() {}

func (*GetUserReq_Email) isGetUserReq_Filter() {}

func (*GetUserReq_Token) isGetUserReq_Filter() {}

func (m *GetUserReq) GetFilter() isGetUserReq_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *GetUserReq) GetID() string {
	if x, ok := m.GetFilter().(*GetUserReq_ID); ok {
		return x.ID
	}
	return ""
}

func (m *GetUserReq) GetEmail() string {
	if x, ok := m.GetFilter().(*GetUserReq_Email); ok {
		return x.Email
	}
	return ""
}

func (m *GetUserReq) GetToken() string {
	if x, ok := m.GetFilter().(*GetUserReq_Token); ok {
		return x.Token
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetUserReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetUserReq_ID)(nil),
		(*GetUserReq_Email)(nil),
		(*GetUserReq_Token)(nil),
	}
}

type GetUserTokensReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserTokensReq) Reset()         { *m = GetUserTokensReq{} }
func (m *GetUserTokensReq) String() string { return proto.CompactTextString(m) }
func (*GetUserTokensReq) ProtoMessage()    {}
func (*GetUserTokensReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{1}
}

func (m *GetUserTokensReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserTokensReq.Unmarshal(m, b)
}
func (m *GetUserTokensReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserTokensReq.Marshal(b, m, deterministic)
}
func (m *GetUserTokensReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserTokensReq.Merge(m, src)
}
func (m *GetUserTokensReq) XXX_Size() int {
	return xxx_messageInfo_GetUserTokensReq.Size(m)
}
func (m *GetUserTokensReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserTokensReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserTokensReq proto.InternalMessageInfo

type GetUserTokensRes struct {
	Tokens               []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserTokensRes) Reset()         { *m = GetUserTokensRes{} }
func (m *GetUserTokensRes) String() string { return proto.CompactTextString(m) }
func (*GetUserTokensRes) ProtoMessage()    {}
func (*GetUserTokensRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{2}
}

func (m *GetUserTokensRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserTokensRes.Unmarshal(m, b)
}
func (m *GetUserTokensRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserTokensRes.Marshal(b, m, deterministic)
}
func (m *GetUserTokensRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserTokensRes.Merge(m, src)
}
func (m *GetUserTokensRes) XXX_Size() int {
	return xxx_messageInfo_GetUserTokensRes.Size(m)
}
func (m *GetUserTokensRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserTokensRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserTokensRes proto.InternalMessageInfo

func (m *GetUserTokensRes) GetTokens() []*Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type LoginReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Pass                 string   `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{3}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginReq) GetPass() string {
	if m != nil {
		return m.Pass
	}
	return ""
}

type CreateUserReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Pass                 string   `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReq) Reset()         { *m = CreateUserReq{} }
func (m *CreateUserReq) String() string { return proto.CompactTextString(m) }
func (*CreateUserReq) ProtoMessage()    {}
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{4}
}

func (m *CreateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReq.Unmarshal(m, b)
}
func (m *CreateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReq.Marshal(b, m, deterministic)
}
func (m *CreateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReq.Merge(m, src)
}
func (m *CreateUserReq) XXX_Size() int {
	return xxx_messageInfo_CreateUserReq.Size(m)
}
func (m *CreateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReq proto.InternalMessageInfo

func (m *CreateUserReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserReq) GetPass() string {
	if m != nil {
		return m.Pass
	}
	return ""
}

type SendResetPasswordLinkReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResetPasswordLinkReq) Reset()         { *m = SendResetPasswordLinkReq{} }
func (m *SendResetPasswordLinkReq) String() string { return proto.CompactTextString(m) }
func (*SendResetPasswordLinkReq) ProtoMessage()    {}
func (*SendResetPasswordLinkReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{5}
}

func (m *SendResetPasswordLinkReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResetPasswordLinkReq.Unmarshal(m, b)
}
func (m *SendResetPasswordLinkReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResetPasswordLinkReq.Marshal(b, m, deterministic)
}
func (m *SendResetPasswordLinkReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResetPasswordLinkReq.Merge(m, src)
}
func (m *SendResetPasswordLinkReq) XXX_Size() int {
	return xxx_messageInfo_SendResetPasswordLinkReq.Size(m)
}
func (m *SendResetPasswordLinkReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResetPasswordLinkReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendResetPasswordLinkReq proto.InternalMessageInfo

func (m *SendResetPasswordLinkReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type SendResetPasswordLinkRes struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResetPasswordLinkRes) Reset()         { *m = SendResetPasswordLinkRes{} }
func (m *SendResetPasswordLinkRes) String() string { return proto.CompactTextString(m) }
func (*SendResetPasswordLinkRes) ProtoMessage()    {}
func (*SendResetPasswordLinkRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{6}
}

func (m *SendResetPasswordLinkRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResetPasswordLinkRes.Unmarshal(m, b)
}
func (m *SendResetPasswordLinkRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResetPasswordLinkRes.Marshal(b, m, deterministic)
}
func (m *SendResetPasswordLinkRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResetPasswordLinkRes.Merge(m, src)
}
func (m *SendResetPasswordLinkRes) XXX_Size() int {
	return xxx_messageInfo_SendResetPasswordLinkRes.Size(m)
}
func (m *SendResetPasswordLinkRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResetPasswordLinkRes.DiscardUnknown(m)
}

var xxx_messageInfo_SendResetPasswordLinkRes proto.InternalMessageInfo

func (m *SendResetPasswordLinkRes) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type ResetPasswordReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	NewPass              string   `protobuf:"bytes,2,opt,name=newPass,proto3" json:"newPass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordReq) Reset()         { *m = ResetPasswordReq{} }
func (m *ResetPasswordReq) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordReq) ProtoMessage()    {}
func (*ResetPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{7}
}

func (m *ResetPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordReq.Unmarshal(m, b)
}
func (m *ResetPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordReq.Marshal(b, m, deterministic)
}
func (m *ResetPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordReq.Merge(m, src)
}
func (m *ResetPasswordReq) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordReq.Size(m)
}
func (m *ResetPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordReq proto.InternalMessageInfo

func (m *ResetPasswordReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ResetPasswordReq) GetNewPass() string {
	if m != nil {
		return m.NewPass
	}
	return ""
}

type ResetPasswordRes struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordRes) Reset()         { *m = ResetPasswordRes{} }
func (m *ResetPasswordRes) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordRes) ProtoMessage()    {}
func (*ResetPasswordRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{8}
}

func (m *ResetPasswordRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordRes.Unmarshal(m, b)
}
func (m *ResetPasswordRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordRes.Marshal(b, m, deterministic)
}
func (m *ResetPasswordRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordRes.Merge(m, src)
}
func (m *ResetPasswordRes) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordRes.Size(m)
}
func (m *ResetPasswordRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordRes.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordRes proto.InternalMessageInfo

func (m *ResetPasswordRes) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type ChangeUserPasswordReq struct {
	OldPass              string   `protobuf:"bytes,1,opt,name=oldPass,proto3" json:"oldPass,omitempty"`
	NewPass              string   `protobuf:"bytes,2,opt,name=newPass,proto3" json:"newPass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeUserPasswordReq) Reset()         { *m = ChangeUserPasswordReq{} }
func (m *ChangeUserPasswordReq) String() string { return proto.CompactTextString(m) }
func (*ChangeUserPasswordReq) ProtoMessage()    {}
func (*ChangeUserPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{9}
}

func (m *ChangeUserPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeUserPasswordReq.Unmarshal(m, b)
}
func (m *ChangeUserPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeUserPasswordReq.Marshal(b, m, deterministic)
}
func (m *ChangeUserPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeUserPasswordReq.Merge(m, src)
}
func (m *ChangeUserPasswordReq) XXX_Size() int {
	return xxx_messageInfo_ChangeUserPasswordReq.Size(m)
}
func (m *ChangeUserPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeUserPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeUserPasswordReq proto.InternalMessageInfo

func (m *ChangeUserPasswordReq) GetOldPass() string {
	if m != nil {
		return m.OldPass
	}
	return ""
}

func (m *ChangeUserPasswordReq) GetNewPass() string {
	if m != nil {
		return m.NewPass
	}
	return ""
}

type ChangeUserPasswordRes struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeUserPasswordRes) Reset()         { *m = ChangeUserPasswordRes{} }
func (m *ChangeUserPasswordRes) String() string { return proto.CompactTextString(m) }
func (*ChangeUserPasswordRes) ProtoMessage()    {}
func (*ChangeUserPasswordRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{10}
}

func (m *ChangeUserPasswordRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeUserPasswordRes.Unmarshal(m, b)
}
func (m *ChangeUserPasswordRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeUserPasswordRes.Marshal(b, m, deterministic)
}
func (m *ChangeUserPasswordRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeUserPasswordRes.Merge(m, src)
}
func (m *ChangeUserPasswordRes) XXX_Size() int {
	return xxx_messageInfo_ChangeUserPasswordRes.Size(m)
}
func (m *ChangeUserPasswordRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeUserPasswordRes.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeUserPasswordRes proto.InternalMessageInfo

func (m *ChangeUserPasswordRes) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type UpdateUserReq struct {
	Email                *wrappers.StringValue  `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone                *wrappers.StringValue  `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Profile              *UpdateUserReq_Profile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UpdateUserReq) Reset()         { *m = UpdateUserReq{} }
func (m *UpdateUserReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserReq) ProtoMessage()    {}
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{11}
}

func (m *UpdateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserReq.Unmarshal(m, b)
}
func (m *UpdateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserReq.Merge(m, src)
}
func (m *UpdateUserReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserReq.Size(m)
}
func (m *UpdateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserReq proto.InternalMessageInfo

func (m *UpdateUserReq) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *UpdateUserReq) GetPhone() *wrappers.StringValue {
	if m != nil {
		return m.Phone
	}
	return nil
}

func (m *UpdateUserReq) GetProfile() *UpdateUserReq_Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type UpdateUserReq_Profile struct {
	FullName             *wrappers.StringValue `protobuf:"bytes,1,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Address              *wrappers.StringValue `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateUserReq_Profile) Reset()         { *m = UpdateUserReq_Profile{} }
func (m *UpdateUserReq_Profile) String() string { return proto.CompactTextString(m) }
func (*UpdateUserReq_Profile) ProtoMessage()    {}
func (*UpdateUserReq_Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{11, 0}
}

func (m *UpdateUserReq_Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserReq_Profile.Unmarshal(m, b)
}
func (m *UpdateUserReq_Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserReq_Profile.Marshal(b, m, deterministic)
}
func (m *UpdateUserReq_Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserReq_Profile.Merge(m, src)
}
func (m *UpdateUserReq_Profile) XXX_Size() int {
	return xxx_messageInfo_UpdateUserReq_Profile.Size(m)
}
func (m *UpdateUserReq_Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserReq_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserReq_Profile proto.InternalMessageInfo

func (m *UpdateUserReq_Profile) GetFullName() *wrappers.StringValue {
	if m != nil {
		return m.FullName
	}
	return nil
}

func (m *UpdateUserReq_Profile) GetAddress() *wrappers.StringValue {
	if m != nil {
		return m.Address
	}
	return nil
}

type UpdateUserRes struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRes) Reset()         { *m = UpdateUserRes{} }
func (m *UpdateUserRes) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRes) ProtoMessage()    {}
func (*UpdateUserRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{12}
}

func (m *UpdateUserRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRes.Unmarshal(m, b)
}
func (m *UpdateUserRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRes.Marshal(b, m, deterministic)
}
func (m *UpdateUserRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRes.Merge(m, src)
}
func (m *UpdateUserRes) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRes.Size(m)
}
func (m *UpdateUserRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRes proto.InternalMessageInfo

func (m *UpdateUserRes) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*GetUserReq)(nil), "schema.GetUserReq")
	proto.RegisterType((*GetUserTokensReq)(nil), "schema.GetUserTokensReq")
	proto.RegisterType((*GetUserTokensRes)(nil), "schema.GetUserTokensRes")
	proto.RegisterType((*LoginReq)(nil), "schema.LoginReq")
	proto.RegisterType((*CreateUserReq)(nil), "schema.CreateUserReq")
	proto.RegisterType((*SendResetPasswordLinkReq)(nil), "schema.SendResetPasswordLinkReq")
	proto.RegisterType((*SendResetPasswordLinkRes)(nil), "schema.SendResetPasswordLinkRes")
	proto.RegisterType((*ResetPasswordReq)(nil), "schema.ResetPasswordReq")
	proto.RegisterType((*ResetPasswordRes)(nil), "schema.ResetPasswordRes")
	proto.RegisterType((*ChangeUserPasswordReq)(nil), "schema.ChangeUserPasswordReq")
	proto.RegisterType((*ChangeUserPasswordRes)(nil), "schema.ChangeUserPasswordRes")
	proto.RegisterType((*UpdateUserReq)(nil), "schema.UpdateUserReq")
	proto.RegisterType((*UpdateUserReq_Profile)(nil), "schema.UpdateUserReq.Profile")
	proto.RegisterType((*UpdateUserRes)(nil), "schema.UpdateUserRes")
}

func init() { proto.RegisterFile("user-service.proto", fileDescriptor_2a3086c73a75cdba) }

var fileDescriptor_2a3086c73a75cdba = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5f, 0x6f, 0xd3, 0x40,
	0x0c, 0x67, 0xd9, 0xd6, 0x14, 0x97, 0xa2, 0xca, 0x62, 0x28, 0x8a, 0xf8, 0x53, 0x45, 0x42, 0x54,
	0x20, 0x32, 0x14, 0x10, 0x30, 0x1e, 0xdb, 0x49, 0x30, 0x31, 0xa1, 0x2a, 0x65, 0x08, 0xc1, 0x53,
	0xb6, 0xb8, 0x69, 0xd4, 0x34, 0x97, 0xdd, 0xa5, 0xf4, 0x81, 0x0f, 0xc1, 0x27, 0xe1, 0x3b, 0xa2,
	0x5c, 0x72, 0x4d, 0xd3, 0xb4, 0xeb, 0xde, 0xce, 0xf6, 0xcf, 0xf6, 0xcf, 0x3e, 0xff, 0x00, 0xe7,
	0x82, 0xf8, 0x2b, 0x41, 0xfc, 0x77, 0x78, 0x45, 0x76, 0xc2, 0x59, 0xca, 0xb0, 0x21, 0xae, 0x26,
	0x34, 0xf3, 0xcc, 0x56, 0xca, 0xa6, 0x14, 0xe7, 0x4e, 0x13, 0x32, 0x60, 0xf1, 0x7e, 0x12, 0x30,
	0x16, 0x44, 0x74, 0x2c, 0xad, 0xcb, 0xf9, 0xf8, 0x78, 0xc1, 0xbd, 0x24, 0x21, 0x2e, 0xf2, 0xb8,
	0xf5, 0x03, 0xe0, 0x13, 0xa5, 0x17, 0x82, 0xb8, 0x4b, 0xd7, 0xd8, 0x01, 0xed, 0xec, 0xd4, 0xd8,
	0xeb, 0xee, 0xf5, 0xee, 0x7e, 0xbe, 0xe3, 0x6a, 0x67, 0xa7, 0xf8, 0x10, 0x0e, 0x69, 0xe6, 0x85,
	0x91, 0xa1, 0x15, 0xce, 0xdc, 0xcc, 0xfc, 0xb2, 0xa5, 0xb1, 0xaf, 0xfc, 0xd2, 0xec, 0x37, 0xa1,
	0x31, 0x0e, 0xa3, 0x94, 0xb8, 0x85, 0xd0, 0x29, 0x2a, 0x7f, 0xcb, 0x22, 0xc2, 0xa5, 0x6b, 0xeb,
	0xa4, 0xe6, 0x13, 0xf8, 0x0c, 0x1a, 0x32, 0x55, 0x18, 0x7b, 0xdd, 0xfd, 0x5e, 0xcb, 0x69, 0xdb,
	0xf9, 0x4c, 0xb6, 0x84, 0xb8, 0x45, 0xd0, 0x7a, 0x0b, 0xcd, 0x73, 0x16, 0x84, 0x71, 0x46, 0xf3,
	0x81, 0x22, 0x25, 0x99, 0x2a, 0x4a, 0x08, 0x07, 0x89, 0x27, 0x44, 0xce, 0xd4, 0x95, 0x6f, 0xeb,
	0x04, 0xda, 0x03, 0x4e, 0x5e, 0x4a, 0x6a, 0xc2, 0xdb, 0xa7, 0xbe, 0x06, 0x63, 0x44, 0xb1, 0xef,
	0x92, 0xa0, 0x74, 0xe8, 0x09, 0xb1, 0x60, 0xdc, 0x3f, 0x0f, 0xe3, 0xe9, 0xd6, 0x2a, 0xd6, 0x8b,
	0xad, 0x19, 0x02, 0xef, 0x83, 0xc6, 0xa6, 0x12, 0xde, 0x74, 0x35, 0x36, 0xb5, 0xfa, 0xd0, 0xa9,
	0xe0, 0x8a, 0xaa, 0xf9, 0x4e, 0x8b, 0xaa, 0xd2, 0x40, 0x03, 0xf4, 0x98, 0x16, 0xc3, 0x92, 0x9e,
	0x32, 0x2d, 0xab, 0x56, 0xa3, 0xde, 0xe7, 0x0b, 0x1c, 0x0d, 0x26, 0x5e, 0x1c, 0xc8, 0x05, 0xac,
	0x36, 0x33, 0x40, 0x67, 0x91, 0x2f, 0xcb, 0xe6, 0xed, 0x94, 0x79, 0x43, 0xc3, 0xe7, 0x9b, 0x8b,
	0xd5, 0xbb, 0xfe, 0xd3, 0xa0, 0x7d, 0x91, 0xf8, 0x2b, 0x7b, 0x77, 0x56, 0x37, 0xd6, 0x72, 0x1e,
	0xd9, 0xf9, 0x5d, 0xda, 0xea, 0x2e, 0xed, 0x51, 0xca, 0xc3, 0x38, 0xf8, 0xee, 0x45, 0x73, 0x52,
	0xbf, 0xe2, 0xc0, 0x61, 0x32, 0x61, 0x31, 0x49, 0x1a, 0x3b, 0x73, 0x24, 0x14, 0xdf, 0x83, 0x9e,
	0x70, 0x36, 0x0e, 0x23, 0x92, 0x97, 0xd9, 0x72, 0x1e, 0xab, 0x73, 0xaa, 0xf0, 0xb1, 0x87, 0x39,
	0xc8, 0x55, 0x68, 0xf3, 0x0f, 0xe8, 0x85, 0x0f, 0x3f, 0x40, 0x73, 0x3c, 0x8f, 0xa2, 0xaf, 0xde,
	0x8c, 0x6e, 0x45, 0x77, 0x89, 0xc6, 0x77, 0xa0, 0x7b, 0xbe, 0xcf, 0xa9, 0x58, 0xdd, 0xae, 0x44,
	0x05, 0xb6, 0x9e, 0x56, 0xd7, 0x55, 0x5b, 0xa8, 0xf3, 0xf7, 0x00, 0x5a, 0x59, 0x6c, 0x94, 0xab,
	0x1f, 0x5f, 0x82, 0x5e, 0x08, 0x09, 0x51, 0x0d, 0x58, 0xea, 0xd8, 0xbc, 0xb7, 0x1c, 0x3a, 0x43,
	0x0c, 0xa0, 0x5d, 0x51, 0x1d, 0x1a, 0x6b, 0x29, 0x4b, 0x81, 0x9a, 0xdb, 0x22, 0x02, 0x7b, 0x70,
	0x28, 0xf5, 0x87, 0x1d, 0x05, 0x51, 0x72, 0x34, 0xab, 0x8a, 0x45, 0x07, 0xa0, 0xd4, 0x1c, 0x1e,
	0xa9, 0x60, 0x45, 0x87, 0xeb, 0x39, 0xbf, 0xe0, 0x68, 0xa3, 0x74, 0xb0, 0xab, 0x70, 0xdb, 0xb4,
	0x68, 0xee, 0x42, 0x88, 0x6c, 0xfe, 0x8a, 0xbf, 0x9c, 0x7f, 0x5d, 0x82, 0xe6, 0xb6, 0x88, 0x40,
	0x17, 0xb0, 0x7e, 0xfb, 0xb8, 0xbc, 0xae, 0x8d, 0x22, 0x33, 0x6f, 0x0c, 0x0b, 0xfc, 0x08, 0x50,
	0x7e, 0x7b, 0xb9, 0xa9, 0xca, 0xa5, 0x9a, 0x1b, 0xdd, 0xa2, 0xdf, 0xf8, 0x79, 0x10, 0x78, 0x49,
	0x78, 0xd9, 0x90, 0x97, 0xf5, 0xe6, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x28, 0x95, 0x35,
	0x1e, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*User, error)
	GetUserTokens(ctx context.Context, in *GetUserTokensReq, opts ...grpc.CallOption) (*GetUserTokensRes, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*Token, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*Token, error)
	SendResetPasswordLink(ctx context.Context, in *SendResetPasswordLinkReq, opts ...grpc.CallOption) (*SendResetPasswordLinkRes, error)
	ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordRes, error)
	ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordReq, opts ...grpc.CallOption) (*ChangeUserPasswordRes, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRes, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/schema.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserTokens(ctx context.Context, in *GetUserTokensReq, opts ...grpc.CallOption) (*GetUserTokensRes, error) {
	out := new(GetUserTokensRes)
	err := c.cc.Invoke(ctx, "/schema.UserService/GetUserTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/schema.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/schema.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendResetPasswordLink(ctx context.Context, in *SendResetPasswordLinkReq, opts ...grpc.CallOption) (*SendResetPasswordLinkRes, error) {
	out := new(SendResetPasswordLinkRes)
	err := c.cc.Invoke(ctx, "/schema.UserService/SendResetPasswordLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordRes, error) {
	out := new(ResetPasswordRes)
	err := c.cc.Invoke(ctx, "/schema.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordReq, opts ...grpc.CallOption) (*ChangeUserPasswordRes, error) {
	out := new(ChangeUserPasswordRes)
	err := c.cc.Invoke(ctx, "/schema.UserService/ChangeUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRes, error) {
	out := new(UpdateUserRes)
	err := c.cc.Invoke(ctx, "/schema.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUser(context.Context, *GetUserReq) (*User, error)
	GetUserTokens(context.Context, *GetUserTokensReq) (*GetUserTokensRes, error)
	Login(context.Context, *LoginReq) (*Token, error)
	CreateUser(context.Context, *CreateUserReq) (*Token, error)
	SendResetPasswordLink(context.Context, *SendResetPasswordLinkReq) (*SendResetPasswordLinkRes, error)
	ResetPassword(context.Context, *ResetPasswordReq) (*ResetPasswordRes, error)
	ChangeUserPassword(context.Context, *ChangeUserPasswordReq) (*ChangeUserPasswordRes, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserRes, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserReq) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) GetUserTokens(ctx context.Context, req *GetUserTokensReq) (*GetUserTokensRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTokens not implemented")
}
func (*UnimplementedUserServiceServer) Login(ctx context.Context, req *LoginReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) SendResetPasswordLink(ctx context.Context, req *SendResetPasswordLinkReq) (*SendResetPasswordLinkRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendResetPasswordLink not implemented")
}
func (*UnimplementedUserServiceServer) ResetPassword(ctx context.Context, req *ResetPasswordReq) (*ResetPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (*UnimplementedUserServiceServer) ChangeUserPassword(ctx context.Context, req *ChangeUserPasswordReq) (*ChangeUserPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPassword not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(ctx context.Context, req *UpdateUserReq) (*UpdateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTokensReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.UserService/GetUserTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserTokens(ctx, req.(*GetUserTokensReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendResetPasswordLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendResetPasswordLinkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendResetPasswordLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.UserService/SendResetPasswordLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendResetPasswordLink(ctx, req.(*SendResetPasswordLinkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.UserService/ChangeUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeUserPassword(ctx, req.(*ChangeUserPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schema.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUserTokens",
			Handler:    _UserService_GetUserTokens_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "SendResetPasswordLink",
			Handler:    _UserService_SendResetPasswordLink_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "ChangeUserPassword",
			Handler:    _UserService_ChangeUserPassword_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user-service.proto",
}
