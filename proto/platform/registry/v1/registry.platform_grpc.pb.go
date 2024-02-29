// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: registry.platform.proto

package v1

import (
	v1 "chainguard.dev/sdk/proto/platform/tenant/v1"
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
	Registry_CreateRepo_FullMethodName     = "/chainguard.platform.registry.Registry/CreateRepo"
	Registry_UpdateRepo_FullMethodName     = "/chainguard.platform.registry.Registry/UpdateRepo"
	Registry_ListRepos_FullMethodName      = "/chainguard.platform.registry.Registry/ListRepos"
	Registry_DeleteRepo_FullMethodName     = "/chainguard.platform.registry.Registry/DeleteRepo"
	Registry_CreateTag_FullMethodName      = "/chainguard.platform.registry.Registry/CreateTag"
	Registry_UpdateTag_FullMethodName      = "/chainguard.platform.registry.Registry/UpdateTag"
	Registry_DeleteTag_FullMethodName      = "/chainguard.platform.registry.Registry/DeleteTag"
	Registry_ListTags_FullMethodName       = "/chainguard.platform.registry.Registry/ListTags"
	Registry_ListTagHistory_FullMethodName = "/chainguard.platform.registry.Registry/ListTagHistory"
	Registry_DiffImage_FullMethodName      = "/chainguard.platform.registry.Registry/DiffImage"
	Registry_GetSbom_FullMethodName        = "/chainguard.platform.registry.Registry/GetSbom"
	Registry_GetVulnReport_FullMethodName  = "/chainguard.platform.registry.Registry/GetVulnReport"
)

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryClient interface {
	CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*Repo, error)
	UpdateRepo(ctx context.Context, in *Repo, opts ...grpc.CallOption) (*Repo, error)
	ListRepos(ctx context.Context, in *RepoFilter, opts ...grpc.CallOption) (*RepoList, error)
	DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error)
	UpdateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListTags(ctx context.Context, in *TagFilter, opts ...grpc.CallOption) (*TagList, error)
	ListTagHistory(ctx context.Context, in *TagHistoryFilter, opts ...grpc.CallOption) (*TagHistoryList, error)
	DiffImage(ctx context.Context, in *DiffImageRequest, opts ...grpc.CallOption) (*DiffImageResponse, error)
	GetSbom(ctx context.Context, in *SbomRequest, opts ...grpc.CallOption) (*v1.Sbom2, error)
	GetVulnReport(ctx context.Context, in *VulnReportRequest, opts ...grpc.CallOption) (*v1.VulnReport, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*Repo, error) {
	out := new(Repo)
	err := c.cc.Invoke(ctx, Registry_CreateRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) UpdateRepo(ctx context.Context, in *Repo, opts ...grpc.CallOption) (*Repo, error) {
	out := new(Repo)
	err := c.cc.Invoke(ctx, Registry_UpdateRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListRepos(ctx context.Context, in *RepoFilter, opts ...grpc.CallOption) (*RepoList, error) {
	out := new(RepoList)
	err := c.cc.Invoke(ctx, Registry_ListRepos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Registry_DeleteRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, Registry_CreateTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) UpdateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, Registry_UpdateTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Registry_DeleteTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListTags(ctx context.Context, in *TagFilter, opts ...grpc.CallOption) (*TagList, error) {
	out := new(TagList)
	err := c.cc.Invoke(ctx, Registry_ListTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListTagHistory(ctx context.Context, in *TagHistoryFilter, opts ...grpc.CallOption) (*TagHistoryList, error) {
	out := new(TagHistoryList)
	err := c.cc.Invoke(ctx, Registry_ListTagHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) DiffImage(ctx context.Context, in *DiffImageRequest, opts ...grpc.CallOption) (*DiffImageResponse, error) {
	out := new(DiffImageResponse)
	err := c.cc.Invoke(ctx, Registry_DiffImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetSbom(ctx context.Context, in *SbomRequest, opts ...grpc.CallOption) (*v1.Sbom2, error) {
	out := new(v1.Sbom2)
	err := c.cc.Invoke(ctx, Registry_GetSbom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetVulnReport(ctx context.Context, in *VulnReportRequest, opts ...grpc.CallOption) (*v1.VulnReport, error) {
	out := new(v1.VulnReport)
	err := c.cc.Invoke(ctx, Registry_GetVulnReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
// All implementations must embed UnimplementedRegistryServer
// for forward compatibility
type RegistryServer interface {
	CreateRepo(context.Context, *CreateRepoRequest) (*Repo, error)
	UpdateRepo(context.Context, *Repo) (*Repo, error)
	ListRepos(context.Context, *RepoFilter) (*RepoList, error)
	DeleteRepo(context.Context, *DeleteRepoRequest) (*emptypb.Empty, error)
	CreateTag(context.Context, *CreateTagRequest) (*Tag, error)
	UpdateTag(context.Context, *Tag) (*Tag, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*emptypb.Empty, error)
	ListTags(context.Context, *TagFilter) (*TagList, error)
	ListTagHistory(context.Context, *TagHistoryFilter) (*TagHistoryList, error)
	DiffImage(context.Context, *DiffImageRequest) (*DiffImageResponse, error)
	GetSbom(context.Context, *SbomRequest) (*v1.Sbom2, error)
	GetVulnReport(context.Context, *VulnReportRequest) (*v1.VulnReport, error)
	mustEmbedUnimplementedRegistryServer()
}

// UnimplementedRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (UnimplementedRegistryServer) CreateRepo(context.Context, *CreateRepoRequest) (*Repo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepo not implemented")
}
func (UnimplementedRegistryServer) UpdateRepo(context.Context, *Repo) (*Repo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRepo not implemented")
}
func (UnimplementedRegistryServer) ListRepos(context.Context, *RepoFilter) (*RepoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepos not implemented")
}
func (UnimplementedRegistryServer) DeleteRepo(context.Context, *DeleteRepoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepo not implemented")
}
func (UnimplementedRegistryServer) CreateTag(context.Context, *CreateTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedRegistryServer) UpdateTag(context.Context, *Tag) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedRegistryServer) DeleteTag(context.Context, *DeleteTagRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedRegistryServer) ListTags(context.Context, *TagFilter) (*TagList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTags not implemented")
}
func (UnimplementedRegistryServer) ListTagHistory(context.Context, *TagHistoryFilter) (*TagHistoryList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTagHistory not implemented")
}
func (UnimplementedRegistryServer) DiffImage(context.Context, *DiffImageRequest) (*DiffImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiffImage not implemented")
}
func (UnimplementedRegistryServer) GetSbom(context.Context, *SbomRequest) (*v1.Sbom2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSbom not implemented")
}
func (UnimplementedRegistryServer) GetVulnReport(context.Context, *VulnReportRequest) (*v1.VulnReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVulnReport not implemented")
}
func (UnimplementedRegistryServer) mustEmbedUnimplementedRegistryServer() {}

// UnsafeRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistryServer will
// result in compilation errors.
type UnsafeRegistryServer interface {
	mustEmbedUnimplementedRegistryServer()
}

func RegisterRegistryServer(s grpc.ServiceRegistrar, srv RegistryServer) {
	s.RegisterService(&Registry_ServiceDesc, srv)
}

func _Registry_CreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).CreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_CreateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).CreateRepo(ctx, req.(*CreateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_UpdateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).UpdateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_UpdateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).UpdateRepo(ctx, req.(*Repo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListRepos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListRepos(ctx, req.(*RepoFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_DeleteRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).DeleteRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_DeleteRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).DeleteRepo(ctx, req.(*DeleteRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_CreateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_UpdateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).UpdateTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_DeleteTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListTags(ctx, req.(*TagFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListTagHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagHistoryFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListTagHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListTagHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListTagHistory(ctx, req.(*TagHistoryFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_DiffImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).DiffImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_DiffImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).DiffImage(ctx, req.(*DiffImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetSbom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SbomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetSbom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetSbom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetSbom(ctx, req.(*SbomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetVulnReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VulnReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetVulnReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetVulnReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetVulnReport(ctx, req.(*VulnReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registry_ServiceDesc is the grpc.ServiceDesc for Registry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepo",
			Handler:    _Registry_CreateRepo_Handler,
		},
		{
			MethodName: "UpdateRepo",
			Handler:    _Registry_UpdateRepo_Handler,
		},
		{
			MethodName: "ListRepos",
			Handler:    _Registry_ListRepos_Handler,
		},
		{
			MethodName: "DeleteRepo",
			Handler:    _Registry_DeleteRepo_Handler,
		},
		{
			MethodName: "CreateTag",
			Handler:    _Registry_CreateTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _Registry_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _Registry_DeleteTag_Handler,
		},
		{
			MethodName: "ListTags",
			Handler:    _Registry_ListTags_Handler,
		},
		{
			MethodName: "ListTagHistory",
			Handler:    _Registry_ListTagHistory_Handler,
		},
		{
			MethodName: "DiffImage",
			Handler:    _Registry_DiffImage_Handler,
		},
		{
			MethodName: "GetSbom",
			Handler:    _Registry_GetSbom_Handler,
		},
		{
			MethodName: "GetVulnReport",
			Handler:    _Registry_GetVulnReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.platform.proto",
}
