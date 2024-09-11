// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: role.platform.proto

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Roles_Create_FullMethodName = "/chainguard.platform.iam.Roles/Create"
	Roles_Update_FullMethodName = "/chainguard.platform.iam.Roles/Update"
	Roles_List_FullMethodName   = "/chainguard.platform.iam.Roles/List"
	Roles_Delete_FullMethodName = "/chainguard.platform.iam.Roles/Delete"
)

// RolesClient is the client API for Roles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolesClient interface {
	Create(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*Role, error)
	Update(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	List(ctx context.Context, in *RoleFilter, opts ...grpc.CallOption) (*RoleList, error)
	Delete(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type rolesClient struct {
	cc grpc.ClientConnInterface
}

func NewRolesClient(cc grpc.ClientConnInterface) RolesClient {
	return &rolesClient{cc}
}

func (c *rolesClient) Create(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Role)
	err := c.cc.Invoke(ctx, Roles_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Update(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Role)
	err := c.cc.Invoke(ctx, Roles_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) List(ctx context.Context, in *RoleFilter, opts ...grpc.CallOption) (*RoleList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleList)
	err := c.cc.Invoke(ctx, Roles_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Delete(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Roles_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServer is the server API for Roles service.
// All implementations must embed UnimplementedRolesServer
// for forward compatibility.
type RolesServer interface {
	Create(context.Context, *CreateRoleRequest) (*Role, error)
	Update(context.Context, *Role) (*Role, error)
	List(context.Context, *RoleFilter) (*RoleList, error)
	Delete(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRolesServer()
}

// UnimplementedRolesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRolesServer struct{}

func (UnimplementedRolesServer) Create(context.Context, *CreateRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRolesServer) Update(context.Context, *Role) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRolesServer) List(context.Context, *RoleFilter) (*RoleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRolesServer) Delete(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRolesServer) mustEmbedUnimplementedRolesServer() {}
func (UnimplementedRolesServer) testEmbeddedByValue()               {}

// UnsafeRolesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolesServer will
// result in compilation errors.
type UnsafeRolesServer interface {
	mustEmbedUnimplementedRolesServer()
}

func RegisterRolesServer(s grpc.ServiceRegistrar, srv RolesServer) {
	// If the following call pancis, it indicates UnimplementedRolesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Roles_ServiceDesc, srv)
}

func _Roles_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Create(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Update(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).List(ctx, req.(*RoleFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Delete(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Roles_ServiceDesc is the grpc.ServiceDesc for Roles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Roles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.iam.Roles",
	HandlerType: (*RolesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Roles_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Roles_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Roles_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Roles_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.platform.proto",
}
