// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valide               bool     `protobuf:"varint,2,opt,name=valide,proto3" json:"valide,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValide() bool {
	if m != nil {
		return m.Valide
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*Error)(nil), "user.Error")
}

func init() {
	proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7)
}

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x6a, 0xe3, 0x40,
	0x10, 0xb4, 0xf5, 0xb2, 0xdd, 0xc2, 0x3e, 0x34, 0xbb, 0xcb, 0xe0, 0xc3, 0xa2, 0x95, 0x61, 0x49,
	0x08, 0x38, 0xe0, 0x9c, 0x73, 0x30, 0x21, 0xf8, 0xae, 0x3c, 0x0e, 0xb9, 0x29, 0x52, 0x43, 0x44,
	0x64, 0x8d, 0x3c, 0x33, 0x76, 0xc8, 0x87, 0xe6, 0x7f, 0xc2, 0xf4, 0xc8, 0xc1, 0x31, 0x81, 0xe4,
	0x22, 0x75, 0x55, 0x17, 0x4d, 0x55, 0x49, 0xf0, 0xbb, 0x55, 0xd2, 0xc8, 0xf3, 0xad, 0x26, 0xc5,
	0x8f, 0x39, 0x63, 0x0c, 0xec, 0x9c, 0xee, 0x20, 0xb8, 0xd3, 0xa4, 0x70, 0x02, 0x5e, 0x55, 0x8a,
	0x7e, 0xd2, 0x3f, 0x19, 0x65, 0x5e, 0x55, 0x22, 0x42, 0xd0, 0xe4, 0x6b, 0x12, 0x1e, 0x33, 0x3c,
	0xa3, 0x80, 0x41, 0x21, 0xd7, 0x6d, 0xde, 0xbc, 0x0a, 0x9f, 0xe9, 0x3d, 0xc4, 0x5f, 0x10, 0xd2,
	0x3a, 0xaf, 0x6a, 0x11, 0x30, 0xef, 0x00, 0x4e, 0x61, 0xd8, 0xe6, 0x5a, 0xbf, 0x48, 0x55, 0x8a,
	0x90, 0x17, 0x1f, 0x38, 0x1d, 0xc1, 0x20, 0xa3, 0xcd, 0x96, 0xb4, 0x49, 0x37, 0x30, 0xcc, 0x48,
	0xb7, 0xb2, 0xd1, 0x84, 0x7f, 0x81, 0x6d, 0xb1, 0x91, 0x78, 0x01, 0x73, 0xf6, 0x6b, 0x0d, 0x66,
	0xcc, 0x63, 0x02, 0xa1, 0x7d, 0x6b, 0xe1, 0x25, 0xfe, 0x91, 0xc0, 0x2d, 0x70, 0x06, 0x11, 0x29,
	0x25, 0x95, 0x16, 0x3e, 0x4b, 0x62, 0x27, 0xb9, 0xb6, 0x5c, 0xd6, 0xad, 0xd2, 0x07, 0x08, 0x6f,
	0xe5, 0x33, 0x35, 0xd6, 0xb8, 0xb1, 0x43, 0x97, 0xdc, 0x01, 0xfc, 0x03, 0xd1, 0x2e, 0xaf, 0xab,
	0xd2, 0xc5, 0x1f, 0x66, 0x1d, 0xfa, 0xd9, 0xed, 0x4b, 0x08, 0x99, 0xb0, 0x15, 0x16, 0xb2, 0x24,
	0x3e, 0x1d, 0x66, 0x3c, 0x63, 0x02, 0x71, 0x49, 0xba, 0x50, 0x55, 0x6b, 0x2a, 0xd9, 0x74, 0xed,
	0x1e, 0x52, 0x8b, 0xb7, 0x3e, 0xc4, 0x36, 0xcf, 0x0d, 0xa9, 0x5d, 0x55, 0x10, 0xfe, 0x87, 0xe8,
	0x4a, 0x51, 0x6e, 0x08, 0x0f, 0xc2, 0x4e, 0x27, 0x6e, 0xde, 0xf7, 0x96, 0xf6, 0x70, 0x06, 0xfe,
	0x8a, 0xcc, 0x37, 0xa2, 0x53, 0x88, 0x56, 0x64, 0x96, 0x75, 0x8d, 0xe3, 0xfd, 0x8e, 0xbf, 0xc1,
	0x17, 0xd2, 0x7f, 0x10, 0x2c, 0xb7, 0xe6, 0xe9, 0xd3, 0xc1, 0x2e, 0x2f, 0x57, 0x97, 0xf6, 0xf0,
	0x0c, 0xc6, 0xf7, 0xb6, 0x98, 0xdc, 0x90, 0x6b, 0xf3, 0x70, 0x7f, 0x24, 0x7e, 0x8c, 0xf8, 0xaf,
	0xbb, 0x78, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xfd, 0xe1, 0xd1, 0x8e, 0x02, 0x00, 0x00,
}
