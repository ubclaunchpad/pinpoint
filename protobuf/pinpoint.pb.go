// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pinpoint.proto

/*
Package pinpoint is a generated protocol buffer package.

It is generated from these files:
	pinpoint.proto

It has these top-level messages:
*/
package pinpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

// Client API for Core service

type CoreClient interface {
	GetStatus(ctx context.Context, in *request.Status, opts ...grpc.CallOption) (*response.Status, error)
	Handshake(ctx context.Context, in *request.Empty, opts ...grpc.CallOption) (*response.Empty, error)
}

type coreClient struct {
	cc *grpc.ClientConn
}

func NewCoreClient(cc *grpc.ClientConn) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) GetStatus(ctx context.Context, in *request.Status, opts ...grpc.CallOption) (*response.Status, error) {
	out := new(response.Status)
	err := grpc.Invoke(ctx, "/pinpoint.Core/GetStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) Handshake(ctx context.Context, in *request.Empty, opts ...grpc.CallOption) (*response.Empty, error) {
	out := new(response.Empty)
	err := grpc.Invoke(ctx, "/pinpoint.Core/Handshake", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Core service

type CoreServer interface {
	GetStatus(context.Context, *request.Status) (*response.Status, error)
	Handshake(context.Context, *request.Empty) (*response.Empty, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pinpoint.proto",
}

func init() { proto.RegisterFile("pinpoint.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xc8, 0xcc, 0x2b,
	0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x44,
	0x8b, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xf4, 0xa1, 0x34, 0x44, 0x81, 0x94, 0x78, 0x51, 0x6a,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x3e, 0x8c, 0x01, 0x91, 0x30, 0xca, 0xe0, 0x62, 0x71, 0xce,
	0x2f, 0x4a, 0x15, 0x32, 0xe0, 0xe2, 0x74, 0x4f, 0x2d, 0x09, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x16,
	0xe2, 0xd7, 0x83, 0xe9, 0x86, 0x08, 0x48, 0x09, 0xe8, 0xc1, 0xb5, 0x41, 0x44, 0x94, 0x18, 0x84,
	0xf4, 0xb8, 0x38, 0x3d, 0x12, 0xf3, 0x52, 0x8a, 0x33, 0x12, 0xb3, 0x53, 0x85, 0xf8, 0xe0, 0x3a,
	0x5c, 0x73, 0x0b, 0x4a, 0x2a, 0xa5, 0xf8, 0x11, 0x1a, 0xc0, 0x02, 0x4a, 0x0c, 0x49, 0x6c, 0x60,
	0x0b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xcb, 0x1d, 0xfb, 0xbc, 0x00, 0x00, 0x00,
}
