// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: role_binding.platform.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RoleBindings_Create_FullMethodName = "/chainguard.platform.iam.RoleBindings/Create"
	RoleBindings_Update_FullMethodName = "/chainguard.platform.iam.RoleBindings/Update"
	RoleBindings_List_FullMethodName   = "/chainguard.platform.iam.RoleBindings/List"
	RoleBindings_Delete_FullMethodName = "/chainguard.platform.iam.RoleBindings/Delete"
)

// RoleBindingsClient is the client API for RoleBindings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleBindingsClient interface {
	Create(ctx context.Context, in *CreateRoleBindingRequest, opts ...grpc.CallOption) (*RoleBinding, error)
	Update(ctx context.Context, in *RoleBinding, opts ...grpc.CallOption) (*RoleBinding, error)
	List(ctx context.Context, in *RoleBindingFilter, opts ...grpc.CallOption) (*RoleBindingList, error)
	Delete(ctx context.Context, in *DeleteRoleBindingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type roleBindingsClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleBindingsClient(cc grpc.ClientConnInterface) RoleBindingsClient {
	return &roleBindingsClient{cc}
}

func (c *roleBindingsClient) Create(ctx context.Context, in *CreateRoleBindingRequest, opts ...grpc.CallOption) (*RoleBinding, error) {
	out := new(RoleBinding)
	err := c.cc.Invoke(ctx, RoleBindings_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleBindingsClient) Update(ctx context.Context, in *RoleBinding, opts ...grpc.CallOption) (*RoleBinding, error) {
	out := new(RoleBinding)
	err := c.cc.Invoke(ctx, RoleBindings_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleBindingsClient) List(ctx context.Context, in *RoleBindingFilter, opts ...grpc.CallOption) (*RoleBindingList, error) {
	out := new(RoleBindingList)
	err := c.cc.Invoke(ctx, RoleBindings_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleBindingsClient) Delete(ctx context.Context, in *DeleteRoleBindingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RoleBindings_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleBindingsServer is the server API for RoleBindings service.
// All implementations must embed UnimplementedRoleBindingsServer
// for forward compatibility
type RoleBindingsServer interface {
	Create(context.Context, *CreateRoleBindingRequest) (*RoleBinding, error)
	Update(context.Context, *RoleBinding) (*RoleBinding, error)
	List(context.Context, *RoleBindingFilter) (*RoleBindingList, error)
	Delete(context.Context, *DeleteRoleBindingRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRoleBindingsServer()
}

// UnimplementedRoleBindingsServer must be embedded to have forward compatible implementations.
type UnimplementedRoleBindingsServer struct {
}

func (UnimplementedRoleBindingsServer) Create(context.Context, *CreateRoleBindingRequest) (*RoleBinding, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRoleBindingsServer) Update(context.Context, *RoleBinding) (*RoleBinding, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRoleBindingsServer) List(context.Context, *RoleBindingFilter) (*RoleBindingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRoleBindingsServer) Delete(context.Context, *DeleteRoleBindingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRoleBindingsServer) mustEmbedUnimplementedRoleBindingsServer() {}

// UnsafeRoleBindingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleBindingsServer will
// result in compilation errors.
type UnsafeRoleBindingsServer interface {
	mustEmbedUnimplementedRoleBindingsServer()
}

func RegisterRoleBindingsServer(s grpc.ServiceRegistrar, srv RoleBindingsServer) {
	s.RegisterService(&RoleBindings_ServiceDesc, srv)
}

func _RoleBindings_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleBindingsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleBindings_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleBindingsServer).Create(ctx, req.(*CreateRoleBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleBindings_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleBinding)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleBindingsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleBindings_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleBindingsServer).Update(ctx, req.(*RoleBinding))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleBindings_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleBindingFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleBindingsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleBindings_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleBindingsServer).List(ctx, req.(*RoleBindingFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleBindings_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleBindingsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleBindings_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleBindingsServer).Delete(ctx, req.(*DeleteRoleBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleBindings_ServiceDesc is the grpc.ServiceDesc for RoleBindings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleBindings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.iam.RoleBindings",
	HandlerType: (*RoleBindingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RoleBindings_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RoleBindings_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RoleBindings_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RoleBindings_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role_binding.platform.proto",
}
