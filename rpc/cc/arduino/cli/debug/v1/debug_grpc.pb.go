// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: cc/arduino/cli/debug/v1/debug.proto

package debug

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DebugService_Debug_FullMethodName          = "/cc.arduino.cli.debug.v1.DebugService/Debug"
	DebugService_GetDebugConfig_FullMethodName = "/cc.arduino.cli.debug.v1.DebugService/GetDebugConfig"
)

// DebugServiceClient is the client API for DebugService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DebugServiceClient interface {
	// Start a debug session and communicate with the debugger tool.
	Debug(ctx context.Context, opts ...grpc.CallOption) (DebugService_DebugClient, error)
	GetDebugConfig(ctx context.Context, in *GetDebugConfigRequest, opts ...grpc.CallOption) (*GetDebugConfigResponse, error)
}

type debugServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugServiceClient(cc grpc.ClientConnInterface) DebugServiceClient {
	return &debugServiceClient{cc}
}

func (c *debugServiceClient) Debug(ctx context.Context, opts ...grpc.CallOption) (DebugService_DebugClient, error) {
	stream, err := c.cc.NewStream(ctx, &DebugService_ServiceDesc.Streams[0], DebugService_Debug_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &debugServiceDebugClient{stream}
	return x, nil
}

type DebugService_DebugClient interface {
	Send(*DebugRequest) error
	Recv() (*DebugResponse, error)
	grpc.ClientStream
}

type debugServiceDebugClient struct {
	grpc.ClientStream
}

func (x *debugServiceDebugClient) Send(m *DebugRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *debugServiceDebugClient) Recv() (*DebugResponse, error) {
	m := new(DebugResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugServiceClient) GetDebugConfig(ctx context.Context, in *GetDebugConfigRequest, opts ...grpc.CallOption) (*GetDebugConfigResponse, error) {
	out := new(GetDebugConfigResponse)
	err := c.cc.Invoke(ctx, DebugService_GetDebugConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugServiceServer is the server API for DebugService service.
// All implementations must embed UnimplementedDebugServiceServer
// for forward compatibility
type DebugServiceServer interface {
	// Start a debug session and communicate with the debugger tool.
	Debug(DebugService_DebugServer) error
	GetDebugConfig(context.Context, *GetDebugConfigRequest) (*GetDebugConfigResponse, error)
	mustEmbedUnimplementedDebugServiceServer()
}

// UnimplementedDebugServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDebugServiceServer struct {
}

func (UnimplementedDebugServiceServer) Debug(DebugService_DebugServer) error {
	return status.Errorf(codes.Unimplemented, "method Debug not implemented")
}
func (UnimplementedDebugServiceServer) GetDebugConfig(context.Context, *GetDebugConfigRequest) (*GetDebugConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDebugConfig not implemented")
}
func (UnimplementedDebugServiceServer) mustEmbedUnimplementedDebugServiceServer() {}

// UnsafeDebugServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DebugServiceServer will
// result in compilation errors.
type UnsafeDebugServiceServer interface {
	mustEmbedUnimplementedDebugServiceServer()
}

func RegisterDebugServiceServer(s grpc.ServiceRegistrar, srv DebugServiceServer) {
	s.RegisterService(&DebugService_ServiceDesc, srv)
}

func _DebugService_Debug_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DebugServiceServer).Debug(&debugServiceDebugServer{stream})
}

type DebugService_DebugServer interface {
	Send(*DebugResponse) error
	Recv() (*DebugRequest, error)
	grpc.ServerStream
}

type debugServiceDebugServer struct {
	grpc.ServerStream
}

func (x *debugServiceDebugServer) Send(m *DebugResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *debugServiceDebugServer) Recv() (*DebugRequest, error) {
	m := new(DebugRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DebugService_GetDebugConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDebugConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServiceServer).GetDebugConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebugService_GetDebugConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServiceServer).GetDebugConfig(ctx, req.(*GetDebugConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DebugService_ServiceDesc is the grpc.ServiceDesc for DebugService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DebugService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cc.arduino.cli.debug.v1.DebugService",
	HandlerType: (*DebugServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDebugConfig",
			Handler:    _DebugService_GetDebugConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Debug",
			Handler:       _DebugService_Debug_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cc/arduino/cli/debug/v1/debug.proto",
}
