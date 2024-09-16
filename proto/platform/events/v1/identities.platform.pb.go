// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: identities.platform.proto

package v1

import (
	_ "chainguard.dev/sdk/proto/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is unique identifier of this specific identity.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// subject of OIDC ID tokens issued for this identity. Matchs the `sub`
	// claim.
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// issuer of the OIDC ID tokens issued for this identity. Matches the `iss`
	// claim.
	Issuer string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// Optional JWKS formatted public keys for the issuer. If supplied
	// verification of ID tokens is attempted using these keys instead of the
	// normal OIDC discovery path. This enables e.g clusters behing NAT to
	// authenticate.
	IssuerKeys string `protobuf:"bytes,4,opt,name=issuer_keys,json=issuerKeys,proto3" json:"issuer_keys,omitempty"`
	// Expiration of identity / issuer keys. After this date /time the issuer
	// keys will not be trusted. Defaults / maximum of 30 days after creation
	// time.
	Expiration *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identities_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_identities_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_identities_platform_proto_rawDescGZIP(), []int{0}
}

func (x *Identity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Identity) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Identity) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Identity) GetIssuerKeys() string {
	if x != nil {
		return x.IssuerKeys
	}
	return ""
}

func (x *Identity) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

type IdentityMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnboardingQuestions *IdentityMetadata_OnboardingQuestions `protobuf:"bytes,1,opt,name=onboarding_questions,json=onboardingQuestions,proto3" json:"onboarding_questions,omitempty"`
	// Output only. This is the name of the user.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Whether the user has opted to receive updates from Chainguard.
	UpdatesOptIn bool `protobuf:"varint,3,opt,name=updatesOptIn,proto3" json:"updatesOptIn,omitempty"`
}

func (x *IdentityMetadata) Reset() {
	*x = IdentityMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identities_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityMetadata) ProtoMessage() {}

func (x *IdentityMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_identities_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityMetadata.ProtoReflect.Descriptor instead.
func (*IdentityMetadata) Descriptor() ([]byte, []int) {
	return file_identities_platform_proto_rawDescGZIP(), []int{1}
}

func (x *IdentityMetadata) GetOnboardingQuestions() *IdentityMetadata_OnboardingQuestions {
	if x != nil {
		return x.OnboardingQuestions
	}
	return nil
}

func (x *IdentityMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IdentityMetadata) GetUpdatesOptIn() bool {
	if x != nil {
		return x.UpdatesOptIn
	}
	return false
}

type IdentityMetadata_OnboardingQuestions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyName string   `protobuf:"bytes,1,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Providers   []string `protobuf:"bytes,2,rep,name=providers,proto3" json:"providers,omitempty"`
	Product     string   `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *IdentityMetadata_OnboardingQuestions) Reset() {
	*x = IdentityMetadata_OnboardingQuestions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identities_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityMetadata_OnboardingQuestions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityMetadata_OnboardingQuestions) ProtoMessage() {}

func (x *IdentityMetadata_OnboardingQuestions) ProtoReflect() protoreflect.Message {
	mi := &file_identities_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityMetadata_OnboardingQuestions.ProtoReflect.Descriptor instead.
func (*IdentityMetadata_OnboardingQuestions) Descriptor() ([]byte, []int) {
	return file_identities_platform_proto_rawDescGZIP(), []int{1, 0}
}

func (x *IdentityMetadata_OnboardingQuestions) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *IdentityMetadata_OnboardingQuestions) GetProviders() []string {
	if x != nil {
		return x.Providers
	}
	return nil
}

func (x *IdentityMetadata_OnboardingQuestions) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

var File_identities_platform_proto protoreflect.FileDescriptor

var file_identities_platform_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x08, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb1, 0x02, 0x0a, 0x10, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x73, 0x0a, 0x14, 0x6f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x13, 0x6f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x70, 0x74,
	0x49, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x4f, 0x70, 0x74, 0x49, 0x6e, 0x1a, 0x70, 0x0a, 0x13, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0xec, 0x02, 0x0a, 0x0a, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x81, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x8a, 0xaf, 0xa8,
	0xd2, 0x05, 0x08, 0x12, 0x06, 0x0a, 0x02, 0x85, 0x07, 0x10, 0x01, 0x12, 0xd9, 0x01, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x2c, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x3a, 0x01, 0x2a, 0x32, 0x1b, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x8a, 0xaf, 0xa8, 0xd2, 0x05, 0x04, 0x12, 0x02, 0x10, 0x01, 0xc2, 0xf0, 0x8e, 0xfc,
	0x0b, 0x35, 0x0a, 0x33, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61,
	0x72, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x75, 0x0a, 0x25, 0x64, 0x65, 0x76, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x1d, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x64, 0x65,
	0x76, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identities_platform_proto_rawDescOnce sync.Once
	file_identities_platform_proto_rawDescData = file_identities_platform_proto_rawDesc
)

func file_identities_platform_proto_rawDescGZIP() []byte {
	file_identities_platform_proto_rawDescOnce.Do(func() {
		file_identities_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_identities_platform_proto_rawDescData)
	})
	return file_identities_platform_proto_rawDescData
}

var file_identities_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_identities_platform_proto_goTypes = []any{
	(*Identity)(nil),                             // 0: chainguard.platform.events.Identity
	(*IdentityMetadata)(nil),                     // 1: chainguard.platform.events.IdentityMetadata
	(*IdentityMetadata_OnboardingQuestions)(nil), // 2: chainguard.platform.events.IdentityMetadata.OnboardingQuestions
	(*timestamppb.Timestamp)(nil),                // 3: google.protobuf.Timestamp
}
var file_identities_platform_proto_depIdxs = []int32{
	3, // 0: chainguard.platform.events.Identity.expiration:type_name -> google.protobuf.Timestamp
	2, // 1: chainguard.platform.events.IdentityMetadata.onboarding_questions:type_name -> chainguard.platform.events.IdentityMetadata.OnboardingQuestions
	0, // 2: chainguard.platform.events.Identities.Create:input_type -> chainguard.platform.events.Identity
	1, // 3: chainguard.platform.events.Identities.UpdateMetadata:input_type -> chainguard.platform.events.IdentityMetadata
	0, // 4: chainguard.platform.events.Identities.Create:output_type -> chainguard.platform.events.Identity
	1, // 5: chainguard.platform.events.Identities.UpdateMetadata:output_type -> chainguard.platform.events.IdentityMetadata
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_identities_platform_proto_init() }
func file_identities_platform_proto_init() {
	if File_identities_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identities_platform_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Identity); i {
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
		file_identities_platform_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*IdentityMetadata); i {
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
		file_identities_platform_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*IdentityMetadata_OnboardingQuestions); i {
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
			RawDescriptor: file_identities_platform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_identities_platform_proto_goTypes,
		DependencyIndexes: file_identities_platform_proto_depIdxs,
		MessageInfos:      file_identities_platform_proto_msgTypes,
	}.Build()
	File_identities_platform_proto = out.File
	file_identities_platform_proto_rawDesc = nil
	file_identities_platform_proto_goTypes = nil
	file_identities_platform_proto_depIdxs = nil
}
