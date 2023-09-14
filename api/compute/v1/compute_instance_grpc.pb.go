// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/compute/v1/compute_instance.proto

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
	ComputeInstance_ListComputeSpec_FullMethodName             = "/api.compute.v1.ComputeInstance/ListComputeSpec"
	ComputeInstance_ListComputeImage_FullMethodName            = "/api.compute.v1.ComputeInstance/ListComputeImage"
	ComputeInstance_ListComputeInstanceDuration_FullMethodName = "/api.compute.v1.ComputeInstance/ListComputeInstanceDuration"
	ComputeInstance_Create_FullMethodName                      = "/api.compute.v1.ComputeInstance/Create"
	ComputeInstance_Delete_FullMethodName                      = "/api.compute.v1.ComputeInstance/Delete"
	ComputeInstance_Get_FullMethodName                         = "/api.compute.v1.ComputeInstance/Get"
	ComputeInstance_List_FullMethodName                        = "/api.compute.v1.ComputeInstance/List"
	ComputeInstance_StopInstance_FullMethodName                = "/api.compute.v1.ComputeInstance/StopInstance"
	ComputeInstance_StartInstance_FullMethodName               = "/api.compute.v1.ComputeInstance/StartInstance"
	ComputeInstance_SSHInstance_FullMethodName                 = "/api.compute.v1.ComputeInstance/SSHInstance"
)

// ComputeInstanceClient is the client API for ComputeInstance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComputeInstanceClient interface {
	// 查询规格
	ListComputeSpec(ctx context.Context, in *ListComputeSpecRequest, opts ...grpc.CallOption) (*ListComputeSpecReply, error)
	// 查询镜像
	ListComputeImage(ctx context.Context, in *ListComputeImageRequest, opts ...grpc.CallOption) (*ListComputeImageReply, error)
	// 查询到期时间
	ListComputeInstanceDuration(ctx context.Context, in *ListComputeDurationRequest, opts ...grpc.CallOption) (*ListComputeDurationReply, error)
	// 创建实例
	Create(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceReply, error)
	// 删除实例
	Delete(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceReply, error)
	// 获取实例详情
	Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceReply, error)
	// 实例列表
	List(ctx context.Context, in *ListInstanceRequest, opts ...grpc.CallOption) (*ListInstanceReply, error)
	// 停止实例
	StopInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*StopInstanceReply, error)
	// 启动实例
	StartInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*StartInstanceReply, error)
	// 连接ssh
	SSHInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*SSHInstanceReply, error)
}

type computeInstanceClient struct {
	cc grpc.ClientConnInterface
}

func NewComputeInstanceClient(cc grpc.ClientConnInterface) ComputeInstanceClient {
	return &computeInstanceClient{cc}
}

func (c *computeInstanceClient) ListComputeSpec(ctx context.Context, in *ListComputeSpecRequest, opts ...grpc.CallOption) (*ListComputeSpecReply, error) {
	out := new(ListComputeSpecReply)
	err := c.cc.Invoke(ctx, ComputeInstance_ListComputeSpec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeInstanceClient) ListComputeImage(ctx context.Context, in *ListComputeImageRequest, opts ...grpc.CallOption) (*ListComputeImageReply, error) {
	out := new(ListComputeImageReply)
	err := c.cc.Invoke(ctx, ComputeInstance_ListComputeImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeInstanceClient) ListComputeInstanceDuration(ctx context.Context, in *ListComputeDurationRequest, opts ...grpc.CallOption) (*ListComputeDurationReply, error) {
	out := new(ListComputeDurationReply)
	err := c.cc.Invoke(ctx, ComputeInstance_ListComputeInstanceDuration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeInstanceClient) Create(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceReply, error) {
	out := new(CreateInstanceReply)
	err := c.cc.Invoke(ctx, ComputeInstance_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeInstanceClient) Delete(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceReply, error) {
	out := new(DeleteInstanceReply)
	err := c.cc.Invoke(ctx, ComputeInstance_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeInstanceClient) Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceReply, error) {
	out := new(GetInstanceReply)
	err := c.cc.Invoke(ctx, ComputeInstance_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeInstanceClient) List(ctx context.Context, in *ListInstanceRequest, opts ...grpc.CallOption) (*ListInstanceReply, error) {
	out := new(ListInstanceReply)
	err := c.cc.Invoke(ctx, ComputeInstance_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeInstanceClient) StopInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*StopInstanceReply, error) {
	out := new(StopInstanceReply)
	err := c.cc.Invoke(ctx, ComputeInstance_StopInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeInstanceClient) StartInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*StartInstanceReply, error) {
	out := new(StartInstanceReply)
	err := c.cc.Invoke(ctx, ComputeInstance_StartInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeInstanceClient) SSHInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*SSHInstanceReply, error) {
	out := new(SSHInstanceReply)
	err := c.cc.Invoke(ctx, ComputeInstance_SSHInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeInstanceServer is the server API for ComputeInstance service.
// All implementations must embed UnimplementedComputeInstanceServer
// for forward compatibility
type ComputeInstanceServer interface {
	// 查询规格
	ListComputeSpec(context.Context, *ListComputeSpecRequest) (*ListComputeSpecReply, error)
	// 查询镜像
	ListComputeImage(context.Context, *ListComputeImageRequest) (*ListComputeImageReply, error)
	// 查询到期时间
	ListComputeInstanceDuration(context.Context, *ListComputeDurationRequest) (*ListComputeDurationReply, error)
	// 创建实例
	Create(context.Context, *CreateInstanceRequest) (*CreateInstanceReply, error)
	// 删除实例
	Delete(context.Context, *DeleteInstanceRequest) (*DeleteInstanceReply, error)
	// 获取实例详情
	Get(context.Context, *GetInstanceRequest) (*GetInstanceReply, error)
	// 实例列表
	List(context.Context, *ListInstanceRequest) (*ListInstanceReply, error)
	// 停止实例
	StopInstance(context.Context, *GetInstanceRequest) (*StopInstanceReply, error)
	// 启动实例
	StartInstance(context.Context, *GetInstanceRequest) (*StartInstanceReply, error)
	// 连接ssh
	SSHInstance(context.Context, *GetInstanceRequest) (*SSHInstanceReply, error)
	mustEmbedUnimplementedComputeInstanceServer()
}

// UnimplementedComputeInstanceServer must be embedded to have forward compatible implementations.
type UnimplementedComputeInstanceServer struct {
}

func (UnimplementedComputeInstanceServer) ListComputeSpec(context.Context, *ListComputeSpecRequest) (*ListComputeSpecReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComputeSpec not implemented")
}
func (UnimplementedComputeInstanceServer) ListComputeImage(context.Context, *ListComputeImageRequest) (*ListComputeImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComputeImage not implemented")
}
func (UnimplementedComputeInstanceServer) ListComputeInstanceDuration(context.Context, *ListComputeDurationRequest) (*ListComputeDurationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComputeInstanceDuration not implemented")
}
func (UnimplementedComputeInstanceServer) Create(context.Context, *CreateInstanceRequest) (*CreateInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedComputeInstanceServer) Delete(context.Context, *DeleteInstanceRequest) (*DeleteInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedComputeInstanceServer) Get(context.Context, *GetInstanceRequest) (*GetInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedComputeInstanceServer) List(context.Context, *ListInstanceRequest) (*ListInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedComputeInstanceServer) StopInstance(context.Context, *GetInstanceRequest) (*StopInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopInstance not implemented")
}
func (UnimplementedComputeInstanceServer) StartInstance(context.Context, *GetInstanceRequest) (*StartInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartInstance not implemented")
}
func (UnimplementedComputeInstanceServer) SSHInstance(context.Context, *GetInstanceRequest) (*SSHInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SSHInstance not implemented")
}
func (UnimplementedComputeInstanceServer) mustEmbedUnimplementedComputeInstanceServer() {}

// UnsafeComputeInstanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComputeInstanceServer will
// result in compilation errors.
type UnsafeComputeInstanceServer interface {
	mustEmbedUnimplementedComputeInstanceServer()
}

func RegisterComputeInstanceServer(s grpc.ServiceRegistrar, srv ComputeInstanceServer) {
	s.RegisterService(&ComputeInstance_ServiceDesc, srv)
}

func _ComputeInstance_ListComputeSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListComputeSpecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeInstanceServer).ListComputeSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputeInstance_ListComputeSpec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeInstanceServer).ListComputeSpec(ctx, req.(*ListComputeSpecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeInstance_ListComputeImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListComputeImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeInstanceServer).ListComputeImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputeInstance_ListComputeImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeInstanceServer).ListComputeImage(ctx, req.(*ListComputeImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeInstance_ListComputeInstanceDuration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListComputeDurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeInstanceServer).ListComputeInstanceDuration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputeInstance_ListComputeInstanceDuration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeInstanceServer).ListComputeInstanceDuration(ctx, req.(*ListComputeDurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeInstance_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeInstanceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputeInstance_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeInstanceServer).Create(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeInstance_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeInstanceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputeInstance_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeInstanceServer).Delete(ctx, req.(*DeleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeInstance_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeInstanceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputeInstance_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeInstanceServer).Get(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeInstance_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeInstanceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputeInstance_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeInstanceServer).List(ctx, req.(*ListInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeInstance_StopInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeInstanceServer).StopInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputeInstance_StopInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeInstanceServer).StopInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeInstance_StartInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeInstanceServer).StartInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputeInstance_StartInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeInstanceServer).StartInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeInstance_SSHInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeInstanceServer).SSHInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComputeInstance_SSHInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeInstanceServer).SSHInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComputeInstance_ServiceDesc is the grpc.ServiceDesc for ComputeInstance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComputeInstance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.compute.v1.ComputeInstance",
	HandlerType: (*ComputeInstanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListComputeSpec",
			Handler:    _ComputeInstance_ListComputeSpec_Handler,
		},
		{
			MethodName: "ListComputeImage",
			Handler:    _ComputeInstance_ListComputeImage_Handler,
		},
		{
			MethodName: "ListComputeInstanceDuration",
			Handler:    _ComputeInstance_ListComputeInstanceDuration_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ComputeInstance_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ComputeInstance_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ComputeInstance_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ComputeInstance_List_Handler,
		},
		{
			MethodName: "StopInstance",
			Handler:    _ComputeInstance_StopInstance_Handler,
		},
		{
			MethodName: "StartInstance",
			Handler:    _ComputeInstance_StartInstance_Handler,
		},
		{
			MethodName: "SSHInstance",
			Handler:    _ComputeInstance_SSHInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/compute/v1/compute_instance.proto",
}
