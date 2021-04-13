// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: zitadel/features.proto

package features

import (
	object "github.com/caos/zitadel-go/pkg/client/zitadel/object"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FeaturesState int32

const (
	FeaturesState_FEATURES_STATE_ACTIVE          FeaturesState = 0
	FeaturesState_FEATURES_STATE_ACTION_REQUIRED FeaturesState = 1
	FeaturesState_FEATURES_STATE_CANCELED        FeaturesState = 2
	FeaturesState_FEATURES_STATE_GRANDFATHERED   FeaturesState = 3
)

// Enum value maps for FeaturesState.
var (
	FeaturesState_name = map[int32]string{
		0: "FEATURES_STATE_ACTIVE",
		1: "FEATURES_STATE_ACTION_REQUIRED",
		2: "FEATURES_STATE_CANCELED",
		3: "FEATURES_STATE_GRANDFATHERED",
	}
	FeaturesState_value = map[string]int32{
		"FEATURES_STATE_ACTIVE":          0,
		"FEATURES_STATE_ACTION_REQUIRED": 1,
		"FEATURES_STATE_CANCELED":        2,
		"FEATURES_STATE_GRANDFATHERED":   3,
	}
)

func (x FeaturesState) Enum() *FeaturesState {
	p := new(FeaturesState)
	*p = x
	return p
}

func (x FeaturesState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeaturesState) Descriptor() protoreflect.EnumDescriptor {
	return file_zitadel_features_proto_enumTypes[0].Descriptor()
}

func (FeaturesState) Type() protoreflect.EnumType {
	return &file_zitadel_features_proto_enumTypes[0]
}

func (x FeaturesState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeaturesState.Descriptor instead.
func (FeaturesState) EnumDescriptor() ([]byte, []int) {
	return file_zitadel_features_proto_rawDescGZIP(), []int{0}
}

type Features struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details                  *object.ObjectDetails `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	Tier                     *FeatureTier          `protobuf:"bytes,2,opt,name=tier,proto3" json:"tier,omitempty"`
	IsDefault                bool                  `protobuf:"varint,3,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	AuditLogRetention        *duration.Duration    `protobuf:"bytes,4,opt,name=audit_log_retention,json=auditLogRetention,proto3" json:"audit_log_retention,omitempty"`
	LoginPolicyUsernameLogin bool                  `protobuf:"varint,5,opt,name=login_policy_username_login,json=loginPolicyUsernameLogin,proto3" json:"login_policy_username_login,omitempty"`
	LoginPolicyRegistration  bool                  `protobuf:"varint,6,opt,name=login_policy_registration,json=loginPolicyRegistration,proto3" json:"login_policy_registration,omitempty"`
	LoginPolicyIdp           bool                  `protobuf:"varint,7,opt,name=login_policy_idp,json=loginPolicyIdp,proto3" json:"login_policy_idp,omitempty"`
	LoginPolicyFactors       bool                  `protobuf:"varint,8,opt,name=login_policy_factors,json=loginPolicyFactors,proto3" json:"login_policy_factors,omitempty"`
	LoginPolicyPasswordless  bool                  `protobuf:"varint,9,opt,name=login_policy_passwordless,json=loginPolicyPasswordless,proto3" json:"login_policy_passwordless,omitempty"`
	PasswordComplexityPolicy bool                  `protobuf:"varint,10,opt,name=password_complexity_policy,json=passwordComplexityPolicy,proto3" json:"password_complexity_policy,omitempty"`
	LabelPolicy              bool                  `protobuf:"varint,11,opt,name=label_policy,json=labelPolicy,proto3" json:"label_policy,omitempty"`
}

func (x *Features) Reset() {
	*x = Features{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_features_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Features) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Features) ProtoMessage() {}

func (x *Features) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_features_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Features.ProtoReflect.Descriptor instead.
func (*Features) Descriptor() ([]byte, []int) {
	return file_zitadel_features_proto_rawDescGZIP(), []int{0}
}

func (x *Features) GetDetails() *object.ObjectDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Features) GetTier() *FeatureTier {
	if x != nil {
		return x.Tier
	}
	return nil
}

func (x *Features) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

func (x *Features) GetAuditLogRetention() *duration.Duration {
	if x != nil {
		return x.AuditLogRetention
	}
	return nil
}

func (x *Features) GetLoginPolicyUsernameLogin() bool {
	if x != nil {
		return x.LoginPolicyUsernameLogin
	}
	return false
}

func (x *Features) GetLoginPolicyRegistration() bool {
	if x != nil {
		return x.LoginPolicyRegistration
	}
	return false
}

func (x *Features) GetLoginPolicyIdp() bool {
	if x != nil {
		return x.LoginPolicyIdp
	}
	return false
}

func (x *Features) GetLoginPolicyFactors() bool {
	if x != nil {
		return x.LoginPolicyFactors
	}
	return false
}

func (x *Features) GetLoginPolicyPasswordless() bool {
	if x != nil {
		return x.LoginPolicyPasswordless
	}
	return false
}

func (x *Features) GetPasswordComplexityPolicy() bool {
	if x != nil {
		return x.PasswordComplexityPolicy
	}
	return false
}

func (x *Features) GetLabelPolicy() bool {
	if x != nil {
		return x.LabelPolicy
	}
	return false
}

type FeatureTier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	State       FeaturesState `protobuf:"varint,3,opt,name=state,proto3,enum=zitadel.features.v1.FeaturesState" json:"state,omitempty"`
	StatusInfo  string        `protobuf:"bytes,4,opt,name=status_info,json=statusInfo,proto3" json:"status_info,omitempty"`
}

func (x *FeatureTier) Reset() {
	*x = FeatureTier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_features_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureTier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureTier) ProtoMessage() {}

func (x *FeatureTier) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_features_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureTier.ProtoReflect.Descriptor instead.
func (*FeatureTier) Descriptor() ([]byte, []int) {
	return file_zitadel_features_proto_rawDescGZIP(), []int{1}
}

func (x *FeatureTier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureTier) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FeatureTier) GetState() FeaturesState {
	if x != nil {
		return x.State
	}
	return FeaturesState_FEATURES_STATE_ACTIVE
}

func (x *FeatureTier) GetStatusInfo() string {
	if x != nil {
		return x.StatusInfo
	}
	return ""
}

var File_zitadel_features_proto protoreflect.FileDescriptor

var file_zitadel_features_proto_rawDesc = []byte{
	0x0a, 0x16, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65,
	0x6c, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x7a,
	0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x04, 0x0a, 0x08, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x54, 0x69, 0x65, 0x72, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x49, 0x0a, 0x13, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x11, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x1b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x3a, 0x0a, 0x19, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x69, 0x64, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x70, 0x12, 0x30, 0x0a, 0x14, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x66, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x3a, 0x0a,
	0x19, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x17, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x74, 0x79,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x74,
	0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x38, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x8d, 0x01, 0x0a, 0x0d,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x46, 0x45, 0x41, 0x54,
	0x55, 0x52, 0x45, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17,
	0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x46, 0x45, 0x41,
	0x54, 0x55, 0x52, 0x45, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x47, 0x52, 0x41, 0x4e,
	0x44, 0x46, 0x41, 0x54, 0x48, 0x45, 0x52, 0x45, 0x44, 0x10, 0x03, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6f, 0x73, 0x2f, 0x7a,
	0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zitadel_features_proto_rawDescOnce sync.Once
	file_zitadel_features_proto_rawDescData = file_zitadel_features_proto_rawDesc
)

func file_zitadel_features_proto_rawDescGZIP() []byte {
	file_zitadel_features_proto_rawDescOnce.Do(func() {
		file_zitadel_features_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_features_proto_rawDescData)
	})
	return file_zitadel_features_proto_rawDescData
}

var file_zitadel_features_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zitadel_features_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_zitadel_features_proto_goTypes = []interface{}{
	(FeaturesState)(0),           // 0: zitadel.features.v1.FeaturesState
	(*Features)(nil),             // 1: zitadel.features.v1.Features
	(*FeatureTier)(nil),          // 2: zitadel.features.v1.FeatureTier
	(*object.ObjectDetails)(nil), // 3: zitadel.v1.ObjectDetails
	(*duration.Duration)(nil),    // 4: google.protobuf.Duration
}
var file_zitadel_features_proto_depIdxs = []int32{
	3, // 0: zitadel.features.v1.Features.details:type_name -> zitadel.v1.ObjectDetails
	2, // 1: zitadel.features.v1.Features.tier:type_name -> zitadel.features.v1.FeatureTier
	4, // 2: zitadel.features.v1.Features.audit_log_retention:type_name -> google.protobuf.Duration
	0, // 3: zitadel.features.v1.FeatureTier.state:type_name -> zitadel.features.v1.FeaturesState
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_zitadel_features_proto_init() }
func file_zitadel_features_proto_init() {
	if File_zitadel_features_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zitadel_features_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Features); i {
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
		file_zitadel_features_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureTier); i {
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
			RawDescriptor: file_zitadel_features_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_features_proto_goTypes,
		DependencyIndexes: file_zitadel_features_proto_depIdxs,
		EnumInfos:         file_zitadel_features_proto_enumTypes,
		MessageInfos:      file_zitadel_features_proto_msgTypes,
	}.Build()
	File_zitadel_features_proto = out.File
	file_zitadel_features_proto_rawDesc = nil
	file_zitadel_features_proto_goTypes = nil
	file_zitadel_features_proto_depIdxs = nil
}
