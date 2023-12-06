// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/network_mapping/v1/network_mapping.proto

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
	NetworkMapping_CreateNetworkMapping_FullMethodName = "/api.network_mapping.v1.NetworkMapping/CreateNetworkMapping"
	NetworkMapping_PageNetworkMapping_FullMethodName   = "/api.network_mapping.v1.NetworkMapping/PageNetworkMapping"
	NetworkMapping_GetNetworkMapping_FullMethodName    = "/api.network_mapping.v1.NetworkMapping/GetNetworkMapping"
	NetworkMapping_DeleteNetworkMapping_FullMethodName = "/api.network_mapping.v1.NetworkMapping/DeleteNetworkMapping"
	NetworkMapping_NextNetworkMapping_FullMethodName   = "/api.network_mapping.v1.NetworkMapping/NextNetworkMapping"
)

// NetworkMappingClient is the client API for NetworkMapping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkMappingClient interface {
	CreateNetworkMapping(ctx context.Context, in *CreateNetworkMappingRequest, opts ...grpc.CallOption) (*CreateNetworkMappingReply, error)
	PageNetworkMapping(ctx context.Context, in *PageNetworkMappingRequest, opts ...grpc.CallOption) (*PageNetworkMappingReply, error)
	GetNetworkMapping(ctx context.Context, in *GetNetworkMappingRequest, opts ...grpc.CallOption) (*GetNetworkMappingReply, error)
	DeleteNetworkMapping(ctx context.Context, in *DeleteNetworkMappingRequest, opts ...grpc.CallOption) (*DeleteNetworkMappingReply, error)
	NextNetworkMapping(ctx context.Context, in *NextNetworkMappingRequest, opts ...grpc.CallOption) (*NextNetworkMappingReply, error)
}

type networkMappingClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkMappingClient(cc grpc.ClientConnInterface) NetworkMappingClient {
	return &networkMappingClient{cc}
}

func (c *networkMappingClient) CreateNetworkMapping(ctx context.Context, in *CreateNetworkMappingRequest, opts ...grpc.CallOption) (*CreateNetworkMappingReply, error) {
	out := new(CreateNetworkMappingReply)
	err := c.cc.Invoke(ctx, NetworkMapping_CreateNetworkMapping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkMappingClient) PageNetworkMapping(ctx context.Context, in *PageNetworkMappingRequest, opts ...grpc.CallOption) (*PageNetworkMappingReply, error) {
	out := new(PageNetworkMappingReply)
	err := c.cc.Invoke(ctx, NetworkMapping_PageNetworkMapping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkMappingClient) GetNetworkMapping(ctx context.Context, in *GetNetworkMappingRequest, opts ...grpc.CallOption) (*GetNetworkMappingReply, error) {
	out := new(GetNetworkMappingReply)
	err := c.cc.Invoke(ctx, NetworkMapping_GetNetworkMapping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkMappingClient) DeleteNetworkMapping(ctx context.Context, in *DeleteNetworkMappingRequest, opts ...grpc.CallOption) (*DeleteNetworkMappingReply, error) {
	out := new(DeleteNetworkMappingReply)
	err := c.cc.Invoke(ctx, NetworkMapping_DeleteNetworkMapping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkMappingClient) NextNetworkMapping(ctx context.Context, in *NextNetworkMappingRequest, opts ...grpc.CallOption) (*NextNetworkMappingReply, error) {
	out := new(NextNetworkMappingReply)
	err := c.cc.Invoke(ctx, NetworkMapping_NextNetworkMapping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkMappingServer is the server API for NetworkMapping service.
// All implementations must embed UnimplementedNetworkMappingServer
// for forward compatibility
type NetworkMappingServer interface {
	CreateNetworkMapping(context.Context, *CreateNetworkMappingRequest) (*CreateNetworkMappingReply, error)
	PageNetworkMapping(context.Context, *PageNetworkMappingRequest) (*PageNetworkMappingReply, error)
	GetNetworkMapping(context.Context, *GetNetworkMappingRequest) (*GetNetworkMappingReply, error)
	DeleteNetworkMapping(context.Context, *DeleteNetworkMappingRequest) (*DeleteNetworkMappingReply, error)
	NextNetworkMapping(context.Context, *NextNetworkMappingRequest) (*NextNetworkMappingReply, error)
	mustEmbedUnimplementedNetworkMappingServer()
}

// UnimplementedNetworkMappingServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkMappingServer struct {
}

func (UnimplementedNetworkMappingServer) CreateNetworkMapping(context.Context, *CreateNetworkMappingRequest) (*CreateNetworkMappingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNetworkMapping not implemented")
}
func (UnimplementedNetworkMappingServer) PageNetworkMapping(context.Context, *PageNetworkMappingRequest) (*PageNetworkMappingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageNetworkMapping not implemented")
}
func (UnimplementedNetworkMappingServer) GetNetworkMapping(context.Context, *GetNetworkMappingRequest) (*GetNetworkMappingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkMapping not implemented")
}
func (UnimplementedNetworkMappingServer) DeleteNetworkMapping(context.Context, *DeleteNetworkMappingRequest) (*DeleteNetworkMappingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNetworkMapping not implemented")
}
func (UnimplementedNetworkMappingServer) NextNetworkMapping(context.Context, *NextNetworkMappingRequest) (*NextNetworkMappingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextNetworkMapping not implemented")
}
func (UnimplementedNetworkMappingServer) mustEmbedUnimplementedNetworkMappingServer() {}

// UnsafeNetworkMappingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkMappingServer will
// result in compilation errors.
type UnsafeNetworkMappingServer interface {
	mustEmbedUnimplementedNetworkMappingServer()
}

func RegisterNetworkMappingServer(s grpc.ServiceRegistrar, srv NetworkMappingServer) {
	s.RegisterService(&NetworkMapping_ServiceDesc, srv)
}

func _NetworkMapping_CreateNetworkMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkMappingServer).CreateNetworkMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkMapping_CreateNetworkMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkMappingServer).CreateNetworkMapping(ctx, req.(*CreateNetworkMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkMapping_PageNetworkMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageNetworkMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkMappingServer).PageNetworkMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkMapping_PageNetworkMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkMappingServer).PageNetworkMapping(ctx, req.(*PageNetworkMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkMapping_GetNetworkMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkMappingServer).GetNetworkMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkMapping_GetNetworkMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkMappingServer).GetNetworkMapping(ctx, req.(*GetNetworkMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkMapping_DeleteNetworkMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkMappingServer).DeleteNetworkMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkMapping_DeleteNetworkMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkMappingServer).DeleteNetworkMapping(ctx, req.(*DeleteNetworkMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkMapping_NextNetworkMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextNetworkMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkMappingServer).NextNetworkMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkMapping_NextNetworkMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkMappingServer).NextNetworkMapping(ctx, req.(*NextNetworkMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkMapping_ServiceDesc is the grpc.ServiceDesc for NetworkMapping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkMapping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.network_mapping.v1.NetworkMapping",
	HandlerType: (*NetworkMappingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNetworkMapping",
			Handler:    _NetworkMapping_CreateNetworkMapping_Handler,
		},
		{
			MethodName: "PageNetworkMapping",
			Handler:    _NetworkMapping_PageNetworkMapping_Handler,
		},
		{
			MethodName: "GetNetworkMapping",
			Handler:    _NetworkMapping_GetNetworkMapping_Handler,
		},
		{
			MethodName: "DeleteNetworkMapping",
			Handler:    _NetworkMapping_DeleteNetworkMapping_Handler,
		},
		{
			MethodName: "NextNetworkMapping",
			Handler:    _NetworkMapping_NextNetworkMapping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/network_mapping/v1/network_mapping.proto",
}
