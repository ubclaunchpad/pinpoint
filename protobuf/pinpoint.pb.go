// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pinpoint.proto

package pinpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/ubclaunchpad/pinpoint/protobuf/models"
import request "github.com/ubclaunchpad/pinpoint/protobuf/request"
import response "github.com/ubclaunchpad/pinpoint/protobuf/response"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreClient interface {
	GetStatus(ctx context.Context, in *request.Status, opts ...grpc.CallOption) (*response.Status, error)
	Handshake(ctx context.Context, in *request.Empty, opts ...grpc.CallOption) (*response.Empty, error)
	// clubs
	CreateClub(ctx context.Context, in *request.CreateClub, opts ...grpc.CallOption) (*response.Message, error)
	CreatePeriod(ctx context.Context, in *request.CreatePeriod, opts ...grpc.CallOption) (*response.Message, error)
	CreateEvent(ctx context.Context, in *request.CreateEvent, opts ...grpc.CallOption) (*response.Message, error)
	// users
	CreateAccount(ctx context.Context, in *request.CreateAccount, opts ...grpc.CallOption) (*response.Message, error)
	Verify(ctx context.Context, in *request.Verify, opts ...grpc.CallOption) (*response.Message, error)
	Login(ctx context.Context, in *request.Login, opts ...grpc.CallOption) (*response.Message, error)
}

type coreClient struct {
	cc *grpc.ClientConn
}

func NewCoreClient(cc *grpc.ClientConn) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) GetStatus(ctx context.Context, in *request.Status, opts ...grpc.CallOption) (*response.Status, error) {
	out := new(response.Status)
	err := c.cc.Invoke(ctx, "/pinpoint.Core/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) Handshake(ctx context.Context, in *request.Empty, opts ...grpc.CallOption) (*response.Empty, error) {
	out := new(response.Empty)
	err := c.cc.Invoke(ctx, "/pinpoint.Core/Handshake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CreateClub(ctx context.Context, in *request.CreateClub, opts ...grpc.CallOption) (*response.Message, error) {
	out := new(response.Message)
	err := c.cc.Invoke(ctx, "/pinpoint.Core/CreateClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CreatePeriod(ctx context.Context, in *request.CreatePeriod, opts ...grpc.CallOption) (*response.Message, error) {
	out := new(response.Message)
	err := c.cc.Invoke(ctx, "/pinpoint.Core/CreatePeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CreateEvent(ctx context.Context, in *request.CreateEvent, opts ...grpc.CallOption) (*response.Message, error) {
	out := new(response.Message)
	err := c.cc.Invoke(ctx, "/pinpoint.Core/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CreateAccount(ctx context.Context, in *request.CreateAccount, opts ...grpc.CallOption) (*response.Message, error) {
	out := new(response.Message)
	err := c.cc.Invoke(ctx, "/pinpoint.Core/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) Verify(ctx context.Context, in *request.Verify, opts ...grpc.CallOption) (*response.Message, error) {
	out := new(response.Message)
	err := c.cc.Invoke(ctx, "/pinpoint.Core/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) Login(ctx context.Context, in *request.Login, opts ...grpc.CallOption) (*response.Message, error) {
	out := new(response.Message)
	err := c.cc.Invoke(ctx, "/pinpoint.Core/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServer is the server API for Core service.
type CoreServer interface {
	GetStatus(context.Context, *request.Status) (*response.Status, error)
	Handshake(context.Context, *request.Empty) (*response.Empty, error)
	// clubs
	CreateClub(context.Context, *request.CreateClub) (*response.Message, error)
	CreatePeriod(context.Context, *request.CreatePeriod) (*response.Message, error)
	CreateEvent(context.Context, *request.CreateEvent) (*response.Message, error)
	// users
	CreateAccount(context.Context, *request.CreateAccount) (*response.Message, error)
	Verify(context.Context, *request.Verify) (*response.Message, error)
	Login(context.Context, *request.Login) (*response.Message, error)
}

func RegisterCoreServer(s *grpc.Server, srv CoreServer) {
	s.RegisterService(&_Core_serviceDesc, srv)
}

func _Core_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Status)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinpoint.Core/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetStatus(ctx, req.(*request.Status))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinpoint.Core/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Handshake(ctx, req.(*request.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CreateClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateClub)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CreateClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinpoint.Core/CreateClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CreateClub(ctx, req.(*request.CreateClub))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CreatePeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreatePeriod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CreatePeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinpoint.Core/CreatePeriod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CreatePeriod(ctx, req.(*request.CreatePeriod))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinpoint.Core/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CreateEvent(ctx, req.(*request.CreateEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinpoint.Core/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CreateAccount(ctx, req.(*request.CreateAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Verify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinpoint.Core/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Verify(ctx, req.(*request.Verify))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Login)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinpoint.Core/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Login(ctx, req.(*request.Login))
	}
	return interceptor(ctx, in, info, handler)
}

var _Core_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pinpoint.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Core_GetStatus_Handler,
		},
		{
			MethodName: "Handshake",
			Handler:    _Core_Handshake_Handler,
		},
		{
			MethodName: "CreateClub",
			Handler:    _Core_CreateClub_Handler,
		},
		{
			MethodName: "CreatePeriod",
			Handler:    _Core_CreatePeriod_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Core_CreateEvent_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Core_CreateAccount_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Core_Verify_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Core_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pinpoint.proto",
}

func init() { proto.RegisterFile("pinpoint.proto", fileDescriptor_pinpoint_3d2b2f5d68a09e64) }

var fileDescriptor_pinpoint_3d2b2f5d68a09e64 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x03, 0x6a, 0xb1, 0xa3, 0xb6, 0x3a, 0xb5, 0x0a, 0x39, 0xe6, 0x2c, 0x89, 0x28, 0x88,
	0x88, 0x17, 0x09, 0x45, 0x0f, 0x0a, 0x82, 0xe0, 0x3d, 0x6d, 0xc6, 0x1a, 0x6c, 0x77, 0xd7, 0xdd,
	0x89, 0xd0, 0xd7, 0xf2, 0x09, 0x0b, 0x99, 0xfc, 0x29, 0x61, 0x4f, 0xb3, 0xf3, 0xed, 0xef, 0x9b,
	0xc3, 0x0c, 0x8c, 0x4c, 0xa1, 0x8c, 0x2e, 0x14, 0xc7, 0xc6, 0x6a, 0xd6, 0x78, 0xd8, 0xf4, 0xe1,
	0xd4, 0xd2, 0x6f, 0x49, 0x8e, 0x93, 0xba, 0x4a, 0x20, 0xbc, 0xb4, 0xe4, 0x8c, 0x56, 0x8e, 0x92,
	0xe6, 0x51, 0x7f, 0x4c, 0xd6, 0x3a, 0xa7, 0x95, 0x4b, 0xa4, 0x08, 0xbc, 0xf9, 0xdf, 0x83, 0xfd,
	0x54, 0x5b, 0xc2, 0x6b, 0x18, 0x3e, 0x13, 0x7f, 0x70, 0xc6, 0xa5, 0xc3, 0x71, 0xdc, 0xcc, 0x14,
	0x10, 0x9e, 0xc6, 0xed, 0x30, 0x21, 0x51, 0x80, 0x31, 0x0c, 0x5f, 0x32, 0x95, 0xbb, 0xef, 0xec,
	0x87, 0x70, 0xd4, 0x1a, 0xb3, 0xb5, 0xe1, 0x4d, 0x38, 0xee, 0x84, 0x0a, 0x44, 0x01, 0xde, 0x01,
	0xa4, 0x96, 0x32, 0xa6, 0x74, 0x55, 0xce, 0x71, 0xd2, 0x0a, 0x1d, 0x0c, 0xcf, 0x3a, 0xeb, 0x8d,
	0x9c, 0xcb, 0x96, 0x14, 0x05, 0xf8, 0x00, 0xc7, 0x12, 0x79, 0x27, 0x5b, 0xe8, 0x1c, 0xa7, 0x3d,
	0x53, 0xb0, 0xdf, 0xbd, 0x87, 0x23, 0x09, 0xcd, 0xfe, 0x48, 0x31, 0x9e, 0xf7, 0xd4, 0x8a, 0xfa,
	0xcd, 0x47, 0x38, 0x91, 0xcc, 0xd3, 0x62, 0xa1, 0x4b, 0xc5, 0x78, 0xd1, 0x73, 0x6b, 0xee, 0xb7,
	0x63, 0x18, 0x7c, 0x92, 0x2d, 0xbe, 0x36, 0x3b, 0xab, 0x14, 0xe0, 0xcf, 0x5f, 0xc1, 0xc1, 0xab,
	0x5e, 0x16, 0x6a, 0x67, 0x8f, 0x55, 0xef, 0x4d, 0xcf, 0x07, 0xd5, 0xed, 0x6e, 0xb7, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xfb, 0xb6, 0x51, 0x1c, 0x1c, 0x02, 0x00, 0x00,
}
