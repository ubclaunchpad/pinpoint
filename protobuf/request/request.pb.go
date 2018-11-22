// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request/request.proto

// package request defines request datatypes

package request

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_30d73066fed3fbb2, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Status struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_30d73066fed3fbb2, []int{1}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

type CreateAccount struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccount) Reset()         { *m = CreateAccount{} }
func (m *CreateAccount) String() string { return proto.CompactTextString(m) }
func (*CreateAccount) ProtoMessage()    {}
func (*CreateAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_30d73066fed3fbb2, []int{2}
}

func (m *CreateAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccount.Unmarshal(m, b)
}
func (m *CreateAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccount.Marshal(b, m, deterministic)
}
func (m *CreateAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccount.Merge(m, src)
}
func (m *CreateAccount) XXX_Size() int {
	return xxx_messageInfo_CreateAccount.Size(m)
}
func (m *CreateAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccount.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccount proto.InternalMessageInfo

func (m *CreateAccount) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAccount) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Verify struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Verify) Reset()         { *m = Verify{} }
func (m *Verify) String() string { return proto.CompactTextString(m) }
func (*Verify) ProtoMessage()    {}
func (*Verify) Descriptor() ([]byte, []int) {
	return fileDescriptor_30d73066fed3fbb2, []int{3}
}

func (m *Verify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Verify.Unmarshal(m, b)
}
func (m *Verify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Verify.Marshal(b, m, deterministic)
}
func (m *Verify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Verify.Merge(m, src)
}
func (m *Verify) XXX_Size() int {
	return xxx_messageInfo_Verify.Size(m)
}
func (m *Verify) XXX_DiscardUnknown() {
	xxx_messageInfo_Verify.DiscardUnknown(m)
}

var xxx_messageInfo_Verify proto.InternalMessageInfo

func (m *Verify) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type Login struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}
func (*Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_30d73066fed3fbb2, []int{4}
}

func (m *Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login.Unmarshal(m, b)
}
func (m *Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login.Marshal(b, m, deterministic)
}
func (m *Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login.Merge(m, src)
}
func (m *Login) XXX_Size() int {
	return xxx_messageInfo_Login.Size(m)
}
func (m *Login) XXX_DiscardUnknown() {
	xxx_messageInfo_Login.DiscardUnknown(m)
}

var xxx_messageInfo_Login proto.InternalMessageInfo

func (m *Login) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Login) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "request.Empty")
	proto.RegisterType((*Status)(nil), "request.Status")
	proto.RegisterType((*CreateAccount)(nil), "request.CreateAccount")
	proto.RegisterType((*Verify)(nil), "request.Verify")
	proto.RegisterType((*Login)(nil), "request.Login")
}

func init() { proto.RegisterFile("request/request.proto", fileDescriptor_30d73066fed3fbb2) }

var fileDescriptor_30d73066fed3fbb2 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbf, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0xb9, 0xd3, 0xdd, 0x3b, 0x07, 0x6c, 0x82, 0xc2, 0x22, 0x16, 0x92, 0xca, 0xea, 0x82,
	0x5c, 0x65, 0xa9, 0x62, 0x67, 0xa5, 0x68, 0x61, 0x37, 0xc9, 0x66, 0x37, 0x81, 0xcd, 0x0f, 0x93,
	0x09, 0xb2, 0xff, 0xbd, 0x18, 0x57, 0xd1, 0xc2, 0x2a, 0xef, 0x7d, 0x90, 0x6f, 0x98, 0x81, 0xd3,
	0xa4, 0xdf, 0x8a, 0xce, 0x24, 0x96, 0x77, 0x17, 0x53, 0xa0, 0xc0, 0x36, 0x4b, 0xe5, 0x1b, 0x68,
	0xee, 0x5d, 0xa4, 0x99, 0x6f, 0xa1, 0x7d, 0x22, 0xa4, 0x92, 0xf9, 0x33, 0x1c, 0xdf, 0x25, 0x8d,
	0xa4, 0x6f, 0x94, 0x0a, 0xc5, 0x13, 0x3b, 0x81, 0x46, 0x3b, 0xb4, 0x53, 0xb7, 0xba, 0x58, 0x5d,
	0x1e, 0x3d, 0x7e, 0x15, 0xc6, 0xe0, 0xd0, 0xa3, 0xd3, 0xdd, 0xba, 0xc2, 0x9a, 0xd9, 0x19, 0x6c,
	0x23, 0xe6, 0xfc, 0x1e, 0x52, 0xdf, 0x1d, 0x54, 0xfe, 0xd3, 0xf9, 0x39, 0xb4, 0x2f, 0x3a, 0xd9,
	0x61, 0xfe, 0xfc, 0x69, 0x30, 0x9b, 0x45, 0x57, 0x33, 0xbf, 0x86, 0xe6, 0x21, 0x8c, 0xd6, 0xff,
	0x33, 0xec, 0xb7, 0x78, 0xfd, 0x57, 0x7c, 0xbb, 0x7f, 0xbd, 0x1a, 0x2d, 0x99, 0x22, 0x77, 0x2a,
	0x38, 0x51, 0xa4, 0x9a, 0xb0, 0x78, 0x65, 0x22, 0xf6, 0x22, 0x5a, 0x1f, 0x83, 0xf5, 0x24, 0xea,
	0xd6, 0xb2, 0x0c, 0xdf, 0x67, 0x90, 0x6d, 0x25, 0xfb, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45,
	0xa9, 0xa2, 0xfc, 0x20, 0x01, 0x00, 0x00,
}
