// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/compute/v1/sandbox.proto

package v1

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
	Sandbox_CreateSandbox_FullMethodName = "/api.server.compute.v1.Sandbox/CreateSandbox"
)

// SandboxClient is the client API for Sandbox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SandboxClient interface {
	CreateSandbox(ctx context.Context, in *CreateSandboxRequest, opts ...grpc.CallOption) (*CreateSandboxReply, error)
}

type sandboxClient struct {
	cc grpc.ClientConnInterface
}

func NewSandboxClient(cc grpc.ClientConnInterface) SandboxClient {
	return &sandboxClient{cc}
}

func (c *sandboxClient) CreateSandbox(ctx context.Context, in *CreateSandboxRequest, opts ...grpc.CallOption) (*CreateSandboxReply, error) {
	out := new(CreateSandboxReply)
	err := c.cc.Invoke(ctx, Sandbox_CreateSandbox_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SandboxServer is the server API for Sandbox service.
// All implementations must embed UnimplementedSandboxServer
// for forward compatibility
type SandboxServer interface {
	CreateSandbox(context.Context, *CreateSandboxRequest) (*CreateSandboxReply, error)
	mustEmbedUnimplementedSandboxServer()
}

// UnimplementedSandboxServer must be embedded to have forward compatible implementations.
type UnimplementedSandboxServer struct {
}

func (UnimplementedSandboxServer) CreateSandbox(context.Context, *CreateSandboxRequest) (*CreateSandboxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSandbox not implemented")
}
func (UnimplementedSandboxServer) mustEmbedUnimplementedSandboxServer() {}

// UnsafeSandboxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SandboxServer will
// result in compilation errors.
type UnsafeSandboxServer interface {
	mustEmbedUnimplementedSandboxServer()
}

func RegisterSandboxServer(s grpc.ServiceRegistrar, srv SandboxServer) {
	s.RegisterService(&Sandbox_ServiceDesc, srv)
}

func _Sandbox_CreateSandbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSandboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServer).CreateSandbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sandbox_CreateSandbox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServer).CreateSandbox(ctx, req.(*CreateSandboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sandbox_ServiceDesc is the grpc.ServiceDesc for Sandbox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sandbox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.server.compute.v1.Sandbox",
	HandlerType: (*SandboxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSandbox",
			Handler:    _Sandbox_CreateSandbox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/compute/v1/sandbox.proto",
}
