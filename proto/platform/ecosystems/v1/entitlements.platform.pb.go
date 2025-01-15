// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: entitlements.platform.proto

package v1

import (
	_ "chainguard.dev/sdk/proto/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Ecosystem represents the language libraries an org can be entitled to.
type Ecosystem int32

const (
	Ecosystem_UNKNOWN Ecosystem = 0
	Ecosystem_JAVA    Ecosystem = 1
	Ecosystem_PYTHON  Ecosystem = 2
)

// Enum value maps for Ecosystem.
var (
	Ecosystem_name = map[int32]string{
		0: "UNKNOWN",
		1: "JAVA",
		2: "PYTHON",
	}
	Ecosystem_value = map[string]int32{
		"UNKNOWN": 0,
		"JAVA":    1,
		"PYTHON":  2,
	}
)

func (x Ecosystem) Enum() *Ecosystem {
	p := new(Ecosystem)
	*p = x
	return p
}

func (x Ecosystem) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ecosystem) Descriptor() protoreflect.EnumDescriptor {
	return file_entitlements_platform_proto_enumTypes[0].Descriptor()
}

func (Ecosystem) Type() protoreflect.EnumType {
	return &file_entitlements_platform_proto_enumTypes[0]
}

func (x Ecosystem) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ecosystem.Descriptor instead.
func (Ecosystem) EnumDescriptor() ([]byte, []int) {
	return file_entitlements_platform_proto_rawDescGZIP(), []int{0}
}

// Entitlement maps an org to an library ecosystem they are entitled to pull.
type Entitlement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the uidp of the entitlement, a child of a group.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ecosystem is the language ecosystem this entitlement grants access to.
	Ecosystem Ecosystem `protobuf:"varint,2,opt,name=ecosystem,proto3,enum=chainguard.platform.ecosystems.Ecosystem" json:"ecosystem,omitempty"`
}

func (x *Entitlement) Reset() {
	*x = Entitlement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entitlements_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entitlement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entitlement) ProtoMessage() {}

func (x *Entitlement) ProtoReflect() protoreflect.Message {
	mi := &file_entitlements_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entitlement.ProtoReflect.Descriptor instead.
func (*Entitlement) Descriptor() ([]byte, []int) {
	return file_entitlements_platform_proto_rawDescGZIP(), []int{0}
}

func (x *Entitlement) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Entitlement) GetEcosystem() Ecosystem {
	if x != nil {
		return x.Ecosystem
	}
	return Ecosystem_UNKNOWN
}

type EntitlementList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Entitlement `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *EntitlementList) Reset() {
	*x = EntitlementList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entitlements_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntitlementList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntitlementList) ProtoMessage() {}

func (x *EntitlementList) ProtoReflect() protoreflect.Message {
	mi := &file_entitlements_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntitlementList.ProtoReflect.Descriptor instead.
func (*EntitlementList) Descriptor() ([]byte, []int) {
	return file_entitlements_platform_proto_rawDescGZIP(), []int{1}
}

func (x *EntitlementList) GetItems() []*Entitlement {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateEntitlementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// parent_id is the group to create the entitlement for.
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// ecosystem is the language ecosystem to entitle this group to.
	Ecosystem Ecosystem `protobuf:"varint,2,opt,name=ecosystem,proto3,enum=chainguard.platform.ecosystems.Ecosystem" json:"ecosystem,omitempty"`
}

func (x *CreateEntitlementRequest) Reset() {
	*x = CreateEntitlementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entitlements_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntitlementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntitlementRequest) ProtoMessage() {}

func (x *CreateEntitlementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entitlements_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntitlementRequest.ProtoReflect.Descriptor instead.
func (*CreateEntitlementRequest) Descriptor() ([]byte, []int) {
	return file_entitlements_platform_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEntitlementRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *CreateEntitlementRequest) GetEcosystem() Ecosystem {
	if x != nil {
		return x.Ecosystem
	}
	return Ecosystem_UNKNOWN
}

type EntitlementFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// parent_id is the group to list entitlements for. Required.
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// ecosystems are the language ecosystems to list entitlements for. If empty, all entitlements for the group will be listed.
	Ecosystems []Ecosystem `protobuf:"varint,2,rep,packed,name=ecosystems,proto3,enum=chainguard.platform.ecosystems.Ecosystem" json:"ecosystems,omitempty"`
}

func (x *EntitlementFilter) Reset() {
	*x = EntitlementFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entitlements_platform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntitlementFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntitlementFilter) ProtoMessage() {}

func (x *EntitlementFilter) ProtoReflect() protoreflect.Message {
	mi := &file_entitlements_platform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntitlementFilter.ProtoReflect.Descriptor instead.
func (*EntitlementFilter) Descriptor() ([]byte, []int) {
	return file_entitlements_platform_proto_rawDescGZIP(), []int{3}
}

func (x *EntitlementFilter) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *EntitlementFilter) GetEcosystems() []Ecosystem {
	if x != nil {
		return x.Ecosystems
	}
	return nil
}

type DeleteEntitlementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the exact uidp of the entitlement to delete.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEntitlementRequest) Reset() {
	*x = DeleteEntitlementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entitlements_platform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntitlementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntitlementRequest) ProtoMessage() {}

func (x *DeleteEntitlementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entitlements_platform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntitlementRequest.ProtoReflect.Descriptor instead.
func (*DeleteEntitlementRequest) Descriptor() ([]byte, []int) {
	return file_entitlements_platform_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteEntitlementRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_entitlements_platform_proto protoreflect.FileDescriptor

var file_entitlements_platform_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0b, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x90, 0xaf, 0xa8, 0xd2, 0x05, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x47, 0x0a, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x63, 0x6f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x45, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x54, 0x0a, 0x0f, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x88, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0x90, 0xaf, 0xa8, 0xd2, 0x05, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x47, 0x0a, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61,
	0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x63, 0x6f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x45, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x7b, 0x0a, 0x11, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x49, 0x0a,
	0x0a, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x73, 0x2e, 0x45, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x65, 0x63,
	0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x32, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0x90, 0xaf, 0xa8, 0xd2, 0x05, 0x01, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x2e, 0x0a, 0x09,
	0x45, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x41, 0x56, 0x41, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x50, 0x59, 0x54, 0x48, 0x4f, 0x4e, 0x10, 0x02, 0x32, 0x8c, 0x05, 0x0a,
	0x0c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x81, 0x02,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x38, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65,
	0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x8f, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x3a, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x22, 0x2a, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x2a, 0x7d, 0x8a,
	0xaf, 0xa8, 0xd2, 0x05, 0x06, 0x12, 0x04, 0x0a, 0x02, 0x88, 0x0e, 0xc2, 0xf0, 0x8e, 0xfc, 0x0b,
	0x40, 0x0a, 0x35, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x12, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x01, 0x12, 0x9b, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x2f, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x8a, 0xaf, 0xa8, 0xd2, 0x05, 0x06, 0x12, 0x04, 0x0a, 0x02, 0x89, 0x0e, 0x12,
	0xd9, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x38, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x7d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x25, 0x2a, 0x23, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x2a, 0x7d, 0x8a, 0xaf, 0xa8, 0xd2, 0x05, 0x06, 0x12,
	0x04, 0x0a, 0x02, 0x8a, 0x0e, 0xc2, 0xf0, 0x8e, 0xfc, 0x0b, 0x40, 0x0a, 0x35, 0x64, 0x65, 0x76,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2e,
	0x76, 0x31, 0x12, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x42, 0x6b, 0x0a, 0x29, 0x64,
	0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x63, 0x6f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x01, 0x5a, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entitlements_platform_proto_rawDescOnce sync.Once
	file_entitlements_platform_proto_rawDescData = file_entitlements_platform_proto_rawDesc
)

func file_entitlements_platform_proto_rawDescGZIP() []byte {
	file_entitlements_platform_proto_rawDescOnce.Do(func() {
		file_entitlements_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_entitlements_platform_proto_rawDescData)
	})
	return file_entitlements_platform_proto_rawDescData
}

var file_entitlements_platform_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_entitlements_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_entitlements_platform_proto_goTypes = []any{
	(Ecosystem)(0),                   // 0: chainguard.platform.ecosystems.Ecosystem
	(*Entitlement)(nil),              // 1: chainguard.platform.ecosystems.Entitlement
	(*EntitlementList)(nil),          // 2: chainguard.platform.ecosystems.EntitlementList
	(*CreateEntitlementRequest)(nil), // 3: chainguard.platform.ecosystems.CreateEntitlementRequest
	(*EntitlementFilter)(nil),        // 4: chainguard.platform.ecosystems.EntitlementFilter
	(*DeleteEntitlementRequest)(nil), // 5: chainguard.platform.ecosystems.DeleteEntitlementRequest
	(*emptypb.Empty)(nil),            // 6: google.protobuf.Empty
}
var file_entitlements_platform_proto_depIdxs = []int32{
	0, // 0: chainguard.platform.ecosystems.Entitlement.ecosystem:type_name -> chainguard.platform.ecosystems.Ecosystem
	1, // 1: chainguard.platform.ecosystems.EntitlementList.items:type_name -> chainguard.platform.ecosystems.Entitlement
	0, // 2: chainguard.platform.ecosystems.CreateEntitlementRequest.ecosystem:type_name -> chainguard.platform.ecosystems.Ecosystem
	0, // 3: chainguard.platform.ecosystems.EntitlementFilter.ecosystems:type_name -> chainguard.platform.ecosystems.Ecosystem
	3, // 4: chainguard.platform.ecosystems.Entitlements.Create:input_type -> chainguard.platform.ecosystems.CreateEntitlementRequest
	4, // 5: chainguard.platform.ecosystems.Entitlements.List:input_type -> chainguard.platform.ecosystems.EntitlementFilter
	5, // 6: chainguard.platform.ecosystems.Entitlements.Delete:input_type -> chainguard.platform.ecosystems.DeleteEntitlementRequest
	1, // 7: chainguard.platform.ecosystems.Entitlements.Create:output_type -> chainguard.platform.ecosystems.Entitlement
	2, // 8: chainguard.platform.ecosystems.Entitlements.List:output_type -> chainguard.platform.ecosystems.EntitlementList
	6, // 9: chainguard.platform.ecosystems.Entitlements.Delete:output_type -> google.protobuf.Empty
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_entitlements_platform_proto_init() }
func file_entitlements_platform_proto_init() {
	if File_entitlements_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entitlements_platform_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Entitlement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entitlements_platform_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EntitlementList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entitlements_platform_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateEntitlementRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entitlements_platform_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EntitlementFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entitlements_platform_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteEntitlementRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entitlements_platform_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_entitlements_platform_proto_goTypes,
		DependencyIndexes: file_entitlements_platform_proto_depIdxs,
		EnumInfos:         file_entitlements_platform_proto_enumTypes,
		MessageInfos:      file_entitlements_platform_proto_msgTypes,
	}.Build()
	File_entitlements_platform_proto = out.File
	file_entitlements_platform_proto_rawDesc = nil
	file_entitlements_platform_proto_goTypes = nil
	file_entitlements_platform_proto_depIdxs = nil
}
