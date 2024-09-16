// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: account_associations.platform.proto

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
	GroupAccountAssociations_Create_FullMethodName = "/chainguard.platform.iam.GroupAccountAssociations/Create"
	GroupAccountAssociations_Update_FullMethodName = "/chainguard.platform.iam.GroupAccountAssociations/Update"
	GroupAccountAssociations_List_FullMethodName   = "/chainguard.platform.iam.GroupAccountAssociations/List"
	GroupAccountAssociations_Delete_FullMethodName = "/chainguard.platform.iam.GroupAccountAssociations/Delete"
	GroupAccountAssociations_Check_FullMethodName  = "/chainguard.platform.iam.GroupAccountAssociations/Check"
)

// GroupAccountAssociationsClient is the client API for GroupAccountAssociations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupAccountAssociationsClient interface {
	Create(ctx context.Context, in *AccountAssociations, opts ...grpc.CallOption) (*AccountAssociations, error)
	Update(ctx context.Context, in *AccountAssociations, opts ...grpc.CallOption) (*AccountAssociations, error)
	List(ctx context.Context, in *AccountAssociationsFilter, opts ...grpc.CallOption) (*AccountAssociationsList, error)
	Delete(ctx context.Context, in *DeleteAccountAssociationsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Check(ctx context.Context, in *AccountAssociationsCheckRequest, opts ...grpc.CallOption) (*AccountAssociationsStatus, error)
}

type groupAccountAssociationsClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupAccountAssociationsClient(cc grpc.ClientConnInterface) GroupAccountAssociationsClient {
	return &groupAccountAssociationsClient{cc}
}

func (c *groupAccountAssociationsClient) Create(ctx context.Context, in *AccountAssociations, opts ...grpc.CallOption) (*AccountAssociations, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountAssociations)
	err := c.cc.Invoke(ctx, GroupAccountAssociations_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAccountAssociationsClient) Update(ctx context.Context, in *AccountAssociations, opts ...grpc.CallOption) (*AccountAssociations, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountAssociations)
	err := c.cc.Invoke(ctx, GroupAccountAssociations_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAccountAssociationsClient) List(ctx context.Context, in *AccountAssociationsFilter, opts ...grpc.CallOption) (*AccountAssociationsList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountAssociationsList)
	err := c.cc.Invoke(ctx, GroupAccountAssociations_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAccountAssociationsClient) Delete(ctx context.Context, in *DeleteAccountAssociationsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupAccountAssociations_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAccountAssociationsClient) Check(ctx context.Context, in *AccountAssociationsCheckRequest, opts ...grpc.CallOption) (*AccountAssociationsStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountAssociationsStatus)
	err := c.cc.Invoke(ctx, GroupAccountAssociations_Check_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupAccountAssociationsServer is the server API for GroupAccountAssociations service.
// All implementations must embed UnimplementedGroupAccountAssociationsServer
// for forward compatibility.
type GroupAccountAssociationsServer interface {
	Create(context.Context, *AccountAssociations) (*AccountAssociations, error)
	Update(context.Context, *AccountAssociations) (*AccountAssociations, error)
	List(context.Context, *AccountAssociationsFilter) (*AccountAssociationsList, error)
	Delete(context.Context, *DeleteAccountAssociationsRequest) (*emptypb.Empty, error)
	Check(context.Context, *AccountAssociationsCheckRequest) (*AccountAssociationsStatus, error)
	mustEmbedUnimplementedGroupAccountAssociationsServer()
}

// UnimplementedGroupAccountAssociationsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGroupAccountAssociationsServer struct{}

func (UnimplementedGroupAccountAssociationsServer) Create(context.Context, *AccountAssociations) (*AccountAssociations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGroupAccountAssociationsServer) Update(context.Context, *AccountAssociations) (*AccountAssociations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGroupAccountAssociationsServer) List(context.Context, *AccountAssociationsFilter) (*AccountAssociationsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGroupAccountAssociationsServer) Delete(context.Context, *DeleteAccountAssociationsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGroupAccountAssociationsServer) Check(context.Context, *AccountAssociationsCheckRequest) (*AccountAssociationsStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedGroupAccountAssociationsServer) mustEmbedUnimplementedGroupAccountAssociationsServer() {
}
func (UnimplementedGroupAccountAssociationsServer) testEmbeddedByValue() {}

// UnsafeGroupAccountAssociationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupAccountAssociationsServer will
// result in compilation errors.
type UnsafeGroupAccountAssociationsServer interface {
	mustEmbedUnimplementedGroupAccountAssociationsServer()
}

func RegisterGroupAccountAssociationsServer(s grpc.ServiceRegistrar, srv GroupAccountAssociationsServer) {
	// If the following call pancis, it indicates UnimplementedGroupAccountAssociationsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GroupAccountAssociations_ServiceDesc, srv)
}

func _GroupAccountAssociations_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAssociations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAccountAssociationsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupAccountAssociations_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAccountAssociationsServer).Create(ctx, req.(*AccountAssociations))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAccountAssociations_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAssociations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAccountAssociationsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupAccountAssociations_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAccountAssociationsServer).Update(ctx, req.(*AccountAssociations))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAccountAssociations_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAssociationsFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAccountAssociationsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupAccountAssociations_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAccountAssociationsServer).List(ctx, req.(*AccountAssociationsFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAccountAssociations_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountAssociationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAccountAssociationsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupAccountAssociations_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAccountAssociationsServer).Delete(ctx, req.(*DeleteAccountAssociationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAccountAssociations_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAssociationsCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAccountAssociationsServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupAccountAssociations_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAccountAssociationsServer).Check(ctx, req.(*AccountAssociationsCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupAccountAssociations_ServiceDesc is the grpc.ServiceDesc for GroupAccountAssociations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupAccountAssociations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.iam.GroupAccountAssociations",
	HandlerType: (*GroupAccountAssociationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GroupAccountAssociations_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GroupAccountAssociations_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GroupAccountAssociations_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GroupAccountAssociations_Delete_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _GroupAccountAssociations_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_associations.platform.proto",
}
