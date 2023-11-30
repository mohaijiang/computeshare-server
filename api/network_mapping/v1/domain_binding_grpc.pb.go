// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/network_mapping/v1/domain_binding.proto

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
	DomainBinding_CreateDomainBinding_FullMethodName = "/api.network_mapping.v1.DomainBinding/CreateDomainBinding"
	DomainBinding_UpdateDomainBinding_FullMethodName = "/api.network_mapping.v1.DomainBinding/UpdateDomainBinding"
	DomainBinding_DeleteDomainBinding_FullMethodName = "/api.network_mapping.v1.DomainBinding/DeleteDomainBinding"
	DomainBinding_GetDomainBinding_FullMethodName    = "/api.network_mapping.v1.DomainBinding/GetDomainBinding"
	DomainBinding_ListDomainBinding_FullMethodName   = "/api.network_mapping.v1.DomainBinding/ListDomainBinding"
	DomainBinding_NsLookup_FullMethodName            = "/api.network_mapping.v1.DomainBinding/NsLookup"
)

// DomainBindingClient is the client API for DomainBinding service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainBindingClient interface {
	CreateDomainBinding(ctx context.Context, in *CreateDomainBindingRequest, opts ...grpc.CallOption) (*CreateDomainBindingReply, error)
	UpdateDomainBinding(ctx context.Context, in *UpdateDomainBindingRequest, opts ...grpc.CallOption) (*UpdateDomainBindingReply, error)
	DeleteDomainBinding(ctx context.Context, in *DeleteDomainBindingRequest, opts ...grpc.CallOption) (*DeleteDomainBindingReply, error)
	GetDomainBinding(ctx context.Context, in *GetDomainBindingRequest, opts ...grpc.CallOption) (*GetDomainBindingReply, error)
	ListDomainBinding(ctx context.Context, in *ListDomainBindingRequest, opts ...grpc.CallOption) (*ListDomainBindingReply, error)
	NsLookup(ctx context.Context, in *NsLookupRequest, opts ...grpc.CallOption) (*NsLookupReply, error)
}

type domainBindingClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainBindingClient(cc grpc.ClientConnInterface) DomainBindingClient {
	return &domainBindingClient{cc}
}

func (c *domainBindingClient) CreateDomainBinding(ctx context.Context, in *CreateDomainBindingRequest, opts ...grpc.CallOption) (*CreateDomainBindingReply, error) {
	out := new(CreateDomainBindingReply)
	err := c.cc.Invoke(ctx, DomainBinding_CreateDomainBinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainBindingClient) UpdateDomainBinding(ctx context.Context, in *UpdateDomainBindingRequest, opts ...grpc.CallOption) (*UpdateDomainBindingReply, error) {
	out := new(UpdateDomainBindingReply)
	err := c.cc.Invoke(ctx, DomainBinding_UpdateDomainBinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainBindingClient) DeleteDomainBinding(ctx context.Context, in *DeleteDomainBindingRequest, opts ...grpc.CallOption) (*DeleteDomainBindingReply, error) {
	out := new(DeleteDomainBindingReply)
	err := c.cc.Invoke(ctx, DomainBinding_DeleteDomainBinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainBindingClient) GetDomainBinding(ctx context.Context, in *GetDomainBindingRequest, opts ...grpc.CallOption) (*GetDomainBindingReply, error) {
	out := new(GetDomainBindingReply)
	err := c.cc.Invoke(ctx, DomainBinding_GetDomainBinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainBindingClient) ListDomainBinding(ctx context.Context, in *ListDomainBindingRequest, opts ...grpc.CallOption) (*ListDomainBindingReply, error) {
	out := new(ListDomainBindingReply)
	err := c.cc.Invoke(ctx, DomainBinding_ListDomainBinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainBindingClient) NsLookup(ctx context.Context, in *NsLookupRequest, opts ...grpc.CallOption) (*NsLookupReply, error) {
	out := new(NsLookupReply)
	err := c.cc.Invoke(ctx, DomainBinding_NsLookup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainBindingServer is the server API for DomainBinding service.
// All implementations must embed UnimplementedDomainBindingServer
// for forward compatibility
type DomainBindingServer interface {
	CreateDomainBinding(context.Context, *CreateDomainBindingRequest) (*CreateDomainBindingReply, error)
	UpdateDomainBinding(context.Context, *UpdateDomainBindingRequest) (*UpdateDomainBindingReply, error)
	DeleteDomainBinding(context.Context, *DeleteDomainBindingRequest) (*DeleteDomainBindingReply, error)
	GetDomainBinding(context.Context, *GetDomainBindingRequest) (*GetDomainBindingReply, error)
	ListDomainBinding(context.Context, *ListDomainBindingRequest) (*ListDomainBindingReply, error)
	NsLookup(context.Context, *NsLookupRequest) (*NsLookupReply, error)
	mustEmbedUnimplementedDomainBindingServer()
}

// UnimplementedDomainBindingServer must be embedded to have forward compatible implementations.
type UnimplementedDomainBindingServer struct {
}

func (UnimplementedDomainBindingServer) CreateDomainBinding(context.Context, *CreateDomainBindingRequest) (*CreateDomainBindingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomainBinding not implemented")
}
func (UnimplementedDomainBindingServer) UpdateDomainBinding(context.Context, *UpdateDomainBindingRequest) (*UpdateDomainBindingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomainBinding not implemented")
}
func (UnimplementedDomainBindingServer) DeleteDomainBinding(context.Context, *DeleteDomainBindingRequest) (*DeleteDomainBindingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomainBinding not implemented")
}
func (UnimplementedDomainBindingServer) GetDomainBinding(context.Context, *GetDomainBindingRequest) (*GetDomainBindingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainBinding not implemented")
}
func (UnimplementedDomainBindingServer) ListDomainBinding(context.Context, *ListDomainBindingRequest) (*ListDomainBindingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomainBinding not implemented")
}
func (UnimplementedDomainBindingServer) NsLookup(context.Context, *NsLookupRequest) (*NsLookupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NsLookup not implemented")
}
func (UnimplementedDomainBindingServer) mustEmbedUnimplementedDomainBindingServer() {}

// UnsafeDomainBindingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainBindingServer will
// result in compilation errors.
type UnsafeDomainBindingServer interface {
	mustEmbedUnimplementedDomainBindingServer()
}

func RegisterDomainBindingServer(s grpc.ServiceRegistrar, srv DomainBindingServer) {
	s.RegisterService(&DomainBinding_ServiceDesc, srv)
}

func _DomainBinding_CreateDomainBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainBindingServer).CreateDomainBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainBinding_CreateDomainBinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainBindingServer).CreateDomainBinding(ctx, req.(*CreateDomainBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainBinding_UpdateDomainBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainBindingServer).UpdateDomainBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainBinding_UpdateDomainBinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainBindingServer).UpdateDomainBinding(ctx, req.(*UpdateDomainBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainBinding_DeleteDomainBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainBindingServer).DeleteDomainBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainBinding_DeleteDomainBinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainBindingServer).DeleteDomainBinding(ctx, req.(*DeleteDomainBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainBinding_GetDomainBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainBindingServer).GetDomainBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainBinding_GetDomainBinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainBindingServer).GetDomainBinding(ctx, req.(*GetDomainBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainBinding_ListDomainBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainBindingServer).ListDomainBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainBinding_ListDomainBinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainBindingServer).ListDomainBinding(ctx, req.(*ListDomainBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainBinding_NsLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NsLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainBindingServer).NsLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainBinding_NsLookup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainBindingServer).NsLookup(ctx, req.(*NsLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DomainBinding_ServiceDesc is the grpc.ServiceDesc for DomainBinding service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DomainBinding_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.network_mapping.v1.DomainBinding",
	HandlerType: (*DomainBindingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDomainBinding",
			Handler:    _DomainBinding_CreateDomainBinding_Handler,
		},
		{
			MethodName: "UpdateDomainBinding",
			Handler:    _DomainBinding_UpdateDomainBinding_Handler,
		},
		{
			MethodName: "DeleteDomainBinding",
			Handler:    _DomainBinding_DeleteDomainBinding_Handler,
		},
		{
			MethodName: "GetDomainBinding",
			Handler:    _DomainBinding_GetDomainBinding_Handler,
		},
		{
			MethodName: "ListDomainBinding",
			Handler:    _DomainBinding_ListDomainBinding_Handler,
		},
		{
			MethodName: "NsLookup",
			Handler:    _DomainBinding_NsLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/network_mapping/v1/domain_binding.proto",
}