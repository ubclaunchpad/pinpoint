// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/response/response.proto

package response // import "github.com/ubclaunchpad/pinpoint/protobuf/response"

/*
package response defines response datatypes
*/

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_6ae7b399486ac24a, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
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
	return fileDescriptor_response_6ae7b399486ac24a, []int{1}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

type Message struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_6ae7b399486ac24a, []int{2}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "response.Empty")
	proto.RegisterType((*Status)(nil), "response.Status")
	proto.RegisterType((*Message)(nil), "response.Message")
}

func init() {
	proto.RegisterFile("protobuf/response/response.proto", fileDescriptor_response_6ae7b399486ac24a)
}

var fileDescriptor_response_6ae7b399486ac24a = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x85, 0x33, 0xf4,
	0xc0, 0x52, 0x42, 0x1c, 0x30, 0xbe, 0x12, 0x3b, 0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x12,
	0x07, 0x17, 0x5b, 0x70, 0x49, 0x62, 0x49, 0x69, 0xb1, 0x92, 0x32, 0x17, 0xbb, 0x6f, 0x6a, 0x71,
	0x71, 0x62, 0x7a, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x2e, 0x84, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0xe3, 0x3a, 0x99, 0x44, 0x19, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0x97, 0x26, 0x25, 0xe7, 0x24, 0x96, 0xe6, 0x25, 0x67, 0x14, 0x24, 0xa6, 0xe8, 0x17,
	0x64, 0xe6, 0x15, 0xe4, 0x67, 0xe6, 0x95, 0xe8, 0x63, 0x38, 0x23, 0x89, 0x0d, 0x2c, 0x64, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x6c, 0xaf, 0x12, 0xa2, 0x00, 0x00, 0x00,
}
