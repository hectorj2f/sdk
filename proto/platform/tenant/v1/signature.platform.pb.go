// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: signature.platform.proto

package v1

import (
	_ "chainguard.dev/sdk/proto/annotations"
	v1 "chainguard.dev/sdk/proto/platform/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id, The Signature UIDP at which this Signature resides.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name of the Signature.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// a short description of this Signature.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// last_seen tracks the timestamp at which this signature was last seen.
	LastSeen *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_seen,json=lastSeen,proto3" json:"last_seen,omitempty"`
	// predicate_type is an optional field that captures the type of signed
	// metadata this signature has signed.
	//   - For simple signatures (e.g. directly signed artifact), this will
	//     be empty to represent the "null claim".
	//   - For signed claims (e.g. attestations), this will hold the in-toto
	//     predicate type of the signed claim.
	PredicateType string `protobuf:"bytes,5,opt,name=predicate_type,json=predicateType,proto3" json:"predicate_type,omitempty"`
	// Types that are assignable to Kind:
	//
	//	*Signature_None
	//	*Signature_Keyless_
	//	*Signature_Key_
	Kind isSignature_Kind `protobuf_oneof:"kind"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signature_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_signature_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_signature_platform_proto_rawDescGZIP(), []int{0}
}

func (x *Signature) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Signature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Signature) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Signature) GetLastSeen() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSeen
	}
	return nil
}

func (x *Signature) GetPredicateType() string {
	if x != nil {
		return x.PredicateType
	}
	return ""
}

func (m *Signature) GetKind() isSignature_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Signature) GetNone() *emptypb.Empty {
	if x, ok := x.GetKind().(*Signature_None); ok {
		return x.None
	}
	return nil
}

func (x *Signature) GetKeyless() *Signature_Keyless {
	if x, ok := x.GetKind().(*Signature_Keyless_); ok {
		return x.Keyless
	}
	return nil
}

func (x *Signature) GetKey() *Signature_Key {
	if x, ok := x.GetKind().(*Signature_Key_); ok {
		return x.Key
	}
	return nil
}

type isSignature_Kind interface {
	isSignature_Kind()
}

type Signature_None struct {
	// none is the kind of signature that is attached when a policy
	// designates an image as "statically" trusted.
	None *emptypb.Empty `protobuf:"bytes,10,opt,name=none,proto3,oneof"`
}

type Signature_Keyless_ struct {
	// keyless is the kind of signature that is attached when a signature
	// was produced via a "keyless" flow.
	Keyless *Signature_Keyless `protobuf:"bytes,11,opt,name=keyless,proto3,oneof"`
}

type Signature_Key_ struct {
	// key is the kind of signature that is attached when a signature
	// was produced via a static key pair.
	Key *Signature_Key `protobuf:"bytes,12,opt,name=key,proto3,oneof"`
}

func (*Signature_None) isSignature_Kind() {}

func (*Signature_Keyless_) isSignature_Kind() {}

func (*Signature_Key_) isSignature_Kind() {}

type SignatureList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Signature `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SignatureList) Reset() {
	*x = SignatureList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signature_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureList) ProtoMessage() {}

func (x *SignatureList) ProtoReflect() protoreflect.Message {
	mi := &file_signature_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureList.ProtoReflect.Descriptor instead.
func (*SignatureList) Descriptor() ([]byte, []int) {
	return file_signature_platform_proto_rawDescGZIP(), []int{1}
}

func (x *SignatureList) GetItems() []*Signature {
	if x != nil {
		return x.Items
	}
	return nil
}

type SignatureFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uidp *v1.UIDPFilter `protobuf:"bytes,1,opt,name=uidp,proto3" json:"uidp,omitempty"`
	// active_since is the timestamp after which the records should
	// have last been observed in the returned context.
	ActiveSince   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=active_since,json=activeSince,proto3" json:"active_since,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	PredicateType *string                `protobuf:"bytes,4,opt,name=predicate_type,json=predicateType,proto3,oneof" json:"predicate_type,omitempty"`
}

func (x *SignatureFilter) Reset() {
	*x = SignatureFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signature_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureFilter) ProtoMessage() {}

func (x *SignatureFilter) ProtoReflect() protoreflect.Message {
	mi := &file_signature_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureFilter.ProtoReflect.Descriptor instead.
func (*SignatureFilter) Descriptor() ([]byte, []int) {
	return file_signature_platform_proto_rawDescGZIP(), []int{2}
}

func (x *SignatureFilter) GetUidp() *v1.UIDPFilter {
	if x != nil {
		return x.Uidp
	}
	return nil
}

func (x *SignatureFilter) GetActiveSince() *timestamppb.Timestamp {
	if x != nil {
		return x.ActiveSince
	}
	return nil
}

func (x *SignatureFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignatureFilter) GetPredicateType() string {
	if x != nil && x.PredicateType != nil {
		return *x.PredicateType
	}
	return ""
}

type Signature_Keyless struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer  string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// TODO: Consider making this a oneof when other providers
	// start producing interesting claims.
	Github *Signature_Keyless_Github `protobuf:"bytes,3,opt,name=github,proto3" json:"github,omitempty"`
}

func (x *Signature_Keyless) Reset() {
	*x = Signature_Keyless{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signature_platform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature_Keyless) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature_Keyless) ProtoMessage() {}

func (x *Signature_Keyless) ProtoReflect() protoreflect.Message {
	mi := &file_signature_platform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature_Keyless.ProtoReflect.Descriptor instead.
func (*Signature_Keyless) Descriptor() ([]byte, []int) {
	return file_signature_platform_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Signature_Keyless) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Signature_Keyless) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Signature_Keyless) GetGithub() *Signature_Keyless_Github {
	if x != nil {
		return x.Github
	}
	return nil
}

type Signature_Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Signature_Key) Reset() {
	*x = Signature_Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signature_platform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature_Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature_Key) ProtoMessage() {}

func (x *Signature_Key) ProtoReflect() protoreflect.Message {
	mi := &file_signature_platform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature_Key.ProtoReflect.Descriptor instead.
func (*Signature_Key) Descriptor() ([]byte, []int) {
	return file_signature_platform_proto_rawDescGZIP(), []int{0, 1}
}

type Signature_Keyless_Github struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OID: 1.3.6.1.4.1.57264.1.2
	WorkflowTrigger string `protobuf:"bytes,1,opt,name=workflow_trigger,json=workflowTrigger,proto3" json:"workflow_trigger,omitempty"`
	// OID: 1.3.6.1.4.1.57264.1.3
	WorkflowSha string `protobuf:"bytes,2,opt,name=workflow_sha,json=workflowSha,proto3" json:"workflow_sha,omitempty"`
	// OID: 1.3.6.1.4.1.57264.1.4
	WorkflowName string `protobuf:"bytes,3,opt,name=workflow_name,json=workflowName,proto3" json:"workflow_name,omitempty"`
	// OID: 1.3.6.1.4.1.57264.1.5
	WorkflowRepo string `protobuf:"bytes,4,opt,name=workflow_repo,json=workflowRepo,proto3" json:"workflow_repo,omitempty"`
	// OID: 1.3.6.1.4.1.57264.1.6
	WorkflowRef string `protobuf:"bytes,5,opt,name=workflow_ref,json=workflowRef,proto3" json:"workflow_ref,omitempty"`
}

func (x *Signature_Keyless_Github) Reset() {
	*x = Signature_Keyless_Github{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signature_platform_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature_Keyless_Github) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature_Keyless_Github) ProtoMessage() {}

func (x *Signature_Keyless_Github) ProtoReflect() protoreflect.Message {
	mi := &file_signature_platform_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature_Keyless_Github.ProtoReflect.Descriptor instead.
func (*Signature_Keyless_Github) Descriptor() ([]byte, []int) {
	return file_signature_platform_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Signature_Keyless_Github) GetWorkflowTrigger() string {
	if x != nil {
		return x.WorkflowTrigger
	}
	return ""
}

func (x *Signature_Keyless_Github) GetWorkflowSha() string {
	if x != nil {
		return x.WorkflowSha
	}
	return ""
}

func (x *Signature_Keyless_Github) GetWorkflowName() string {
	if x != nil {
		return x.WorkflowName
	}
	return ""
}

func (x *Signature_Keyless_Github) GetWorkflowRepo() string {
	if x != nil {
		return x.WorkflowRepo
	}
	return ""
}

func (x *Signature_Keyless_Github) GetWorkflowRef() string {
	if x != nil {
		return x.WorkflowRef
	}
	return ""
}

var File_signature_platform_proto protoreflect.FileDescriptor

var file_signature_platform_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x69, 0x64, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xca, 0x05, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73,
	0x65, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x6f, 0x6e, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61,
	0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x4b, 0x65, 0x79,
	0x6c, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x12,
	0x3d, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x1a, 0xcf,
	0x02, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4c, 0x0a, 0x06,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x52, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x1a, 0xc3, 0x01, 0x0a, 0x06, 0x47,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x12, 0x29, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x68, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x53, 0x68, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x66,
	0x1a, 0x05, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22,
	0x4c, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xdf, 0x01,
	0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x3a, 0x0a, 0x04, 0x75, 0x69, 0x64, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x49, 0x44,
	0x50, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x04, 0x75, 0x69, 0x64, 0x70, 0x12, 0x3d, 0x0a,
	0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x32,
	0x9a, 0x01, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x8b,
	0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x8a, 0xaf,
	0xa8, 0xd2, 0x05, 0x08, 0x12, 0x06, 0x0a, 0x02, 0xef, 0x04, 0x10, 0x01, 0x42, 0x74, 0x0a, 0x25,
	0x64, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x1c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signature_platform_proto_rawDescOnce sync.Once
	file_signature_platform_proto_rawDescData = file_signature_platform_proto_rawDesc
)

func file_signature_platform_proto_rawDescGZIP() []byte {
	file_signature_platform_proto_rawDescOnce.Do(func() {
		file_signature_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_signature_platform_proto_rawDescData)
	})
	return file_signature_platform_proto_rawDescData
}

var file_signature_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_signature_platform_proto_goTypes = []any{
	(*Signature)(nil),                // 0: chainguard.platform.tenant.Signature
	(*SignatureList)(nil),            // 1: chainguard.platform.tenant.SignatureList
	(*SignatureFilter)(nil),          // 2: chainguard.platform.tenant.SignatureFilter
	(*Signature_Keyless)(nil),        // 3: chainguard.platform.tenant.Signature.Keyless
	(*Signature_Key)(nil),            // 4: chainguard.platform.tenant.Signature.Key
	(*Signature_Keyless_Github)(nil), // 5: chainguard.platform.tenant.Signature.Keyless.Github
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 7: google.protobuf.Empty
	(*v1.UIDPFilter)(nil),            // 8: chainguard.platform.common.UIDPFilter
}
var file_signature_platform_proto_depIdxs = []int32{
	6, // 0: chainguard.platform.tenant.Signature.last_seen:type_name -> google.protobuf.Timestamp
	7, // 1: chainguard.platform.tenant.Signature.none:type_name -> google.protobuf.Empty
	3, // 2: chainguard.platform.tenant.Signature.keyless:type_name -> chainguard.platform.tenant.Signature.Keyless
	4, // 3: chainguard.platform.tenant.Signature.key:type_name -> chainguard.platform.tenant.Signature.Key
	0, // 4: chainguard.platform.tenant.SignatureList.items:type_name -> chainguard.platform.tenant.Signature
	8, // 5: chainguard.platform.tenant.SignatureFilter.uidp:type_name -> chainguard.platform.common.UIDPFilter
	6, // 6: chainguard.platform.tenant.SignatureFilter.active_since:type_name -> google.protobuf.Timestamp
	5, // 7: chainguard.platform.tenant.Signature.Keyless.github:type_name -> chainguard.platform.tenant.Signature.Keyless.Github
	2, // 8: chainguard.platform.tenant.Signatures.List:input_type -> chainguard.platform.tenant.SignatureFilter
	1, // 9: chainguard.platform.tenant.Signatures.List:output_type -> chainguard.platform.tenant.SignatureList
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_signature_platform_proto_init() }
func file_signature_platform_proto_init() {
	if File_signature_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signature_platform_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Signature); i {
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
		file_signature_platform_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SignatureList); i {
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
		file_signature_platform_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SignatureFilter); i {
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
		file_signature_platform_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Signature_Keyless); i {
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
		file_signature_platform_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Signature_Key); i {
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
		file_signature_platform_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Signature_Keyless_Github); i {
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
	file_signature_platform_proto_msgTypes[0].OneofWrappers = []any{
		(*Signature_None)(nil),
		(*Signature_Keyless_)(nil),
		(*Signature_Key_)(nil),
	}
	file_signature_platform_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_signature_platform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_signature_platform_proto_goTypes,
		DependencyIndexes: file_signature_platform_proto_depIdxs,
		MessageInfos:      file_signature_platform_proto_msgTypes,
	}.Build()
	File_signature_platform_proto = out.File
	file_signature_platform_proto_rawDesc = nil
	file_signature_platform_proto_goTypes = nil
	file_signature_platform_proto_depIdxs = nil
}
