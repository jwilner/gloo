// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/gloo/api/v1/options/aws/ec2/aws_ec2.proto

package ec2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
	// The AWS Region where the desired EC2 instances exist
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Optional, if not set, Gloo will try to use the default AWS secret specified by environment variables.
	// If a secret is not provided, the environment must specify both the AWS access key and secret.
	// The environment variables used to indicate the AWS account can be:
	// - for the access key: "AWS_ACCESS_KEY_ID" or "AWS_ACCESS_KEY"
	// - for the secret: "AWS_SECRET_ACCESS_KEY" or "AWS_SECRET_KEY"
	// If set, a [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  ```
	// Gloo will create an EC2 API client with this credential. You may choose to use a credential with limited access
	// in conjunction with a list of Roles, specified by their Amazon Resource Number (ARN).
	SecretRef *core.ResourceRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// Optional, Amazon Resource Number (ARN) referring to IAM Role that should be assumed when the Upstream
	// queries for eligible EC2 instances.
	// If provided, Gloo will create an EC2 API client with the provided role. If not provided, Gloo will not assume
	// a role.
	RoleArn string `protobuf:"bytes,7,opt,name=role_arn,json=roleArn,proto3" json:"role_arn,omitempty"`
	// List of tag filters for selecting instances
	// An instance must match all the filters in order to be selected
	// Filter keys are not case-sensitive
	Filters []*TagFilter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	// If set, will use the EC2 public IP address. Defaults to the private IP address.
	PublicIp bool `protobuf:"varint,4,opt,name=public_ip,json=publicIp,proto3" json:"public_ip,omitempty"`
	// If set, will use this port on EC2 instances. Defaults to port 80.
	Port                 uint32   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5a14b365c5d168d, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() *core.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return nil
}

func (m *UpstreamSpec) GetRoleArn() string {
	if m != nil {
		return m.RoleArn
	}
	return ""
}

func (m *UpstreamSpec) GetFilters() []*TagFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *UpstreamSpec) GetPublicIp() bool {
	if m != nil {
		return m.PublicIp
	}
	return false
}

func (m *UpstreamSpec) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type TagFilter struct {
	// Types that are valid to be assigned to Spec:
	//	*TagFilter_Key
	//	*TagFilter_KvPair_
	Spec                 isTagFilter_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TagFilter) Reset()         { *m = TagFilter{} }
func (m *TagFilter) String() string { return proto.CompactTextString(m) }
func (*TagFilter) ProtoMessage()    {}
func (*TagFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5a14b365c5d168d, []int{1}
}
func (m *TagFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFilter.Unmarshal(m, b)
}
func (m *TagFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFilter.Marshal(b, m, deterministic)
}
func (m *TagFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter.Merge(m, src)
}
func (m *TagFilter) XXX_Size() int {
	return xxx_messageInfo_TagFilter.Size(m)
}
func (m *TagFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter proto.InternalMessageInfo

type isTagFilter_Spec interface {
	isTagFilter_Spec()
	Equal(interface{}) bool
}

type TagFilter_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof" json:"key,omitempty"`
}
type TagFilter_KvPair_ struct {
	KvPair *TagFilter_KvPair `protobuf:"bytes,2,opt,name=kv_pair,json=kvPair,proto3,oneof" json:"kv_pair,omitempty"`
}

func (*TagFilter_Key) isTagFilter_Spec()     {}
func (*TagFilter_KvPair_) isTagFilter_Spec() {}

func (m *TagFilter) GetSpec() isTagFilter_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TagFilter) GetKey() string {
	if x, ok := m.GetSpec().(*TagFilter_Key); ok {
		return x.Key
	}
	return ""
}

func (m *TagFilter) GetKvPair() *TagFilter_KvPair {
	if x, ok := m.GetSpec().(*TagFilter_KvPair_); ok {
		return x.KvPair
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TagFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TagFilter_Key)(nil),
		(*TagFilter_KvPair_)(nil),
	}
}

type TagFilter_KvPair struct {
	// keys are not case-sensitive, as with AWS Condition Keys
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// values are case-sensitive
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagFilter_KvPair) Reset()         { *m = TagFilter_KvPair{} }
func (m *TagFilter_KvPair) String() string { return proto.CompactTextString(m) }
func (*TagFilter_KvPair) ProtoMessage()    {}
func (*TagFilter_KvPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5a14b365c5d168d, []int{1, 0}
}
func (m *TagFilter_KvPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFilter_KvPair.Unmarshal(m, b)
}
func (m *TagFilter_KvPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFilter_KvPair.Marshal(b, m, deterministic)
}
func (m *TagFilter_KvPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter_KvPair.Merge(m, src)
}
func (m *TagFilter_KvPair) XXX_Size() int {
	return xxx_messageInfo_TagFilter_KvPair.Size(m)
}
func (m *TagFilter_KvPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter_KvPair.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter_KvPair proto.InternalMessageInfo

func (m *TagFilter_KvPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TagFilter_KvPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "aws_ec2.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*TagFilter)(nil), "aws_ec2.options.gloo.solo.io.TagFilter")
	proto.RegisterType((*TagFilter_KvPair)(nil), "aws_ec2.options.gloo.solo.io.TagFilter.KvPair")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/gloo/api/v1/options/aws/ec2/aws_ec2.proto", fileDescriptor_c5a14b365c5d168d)
}

var fileDescriptor_c5a14b365c5d168d = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0xad, 0x49, 0xba, 0xc9, 0xba, 0x20, 0x21, 0xab, 0x42, 0x9b, 0x80, 0xd0, 0xaa, 0x17, 0xf6,
	0x52, 0x1b, 0xc2, 0x85, 0x6b, 0x2b, 0x81, 0x5a, 0x71, 0x28, 0x32, 0x70, 0xe1, 0xb2, 0xda, 0x98,
	0x59, 0x63, 0x76, 0x9b, 0xb1, 0x6c, 0x27, 0x84, 0xff, 0xe1, 0xc0, 0x27, 0xf0, 0x3d, 0xfc, 0x00,
	0x27, 0xee, 0x68, 0xed, 0x04, 0x71, 0x01, 0xd1, 0xd3, 0xce, 0x7b, 0xfb, 0x66, 0xde, 0x3c, 0x79,
	0xe8, 0x95, 0x36, 0xe1, 0xc3, 0x7a, 0xc9, 0x15, 0x5e, 0x0b, 0x8f, 0x3d, 0x9e, 0x1a, 0x14, 0xba,
	0x47, 0x3c, 0x85, 0xf7, 0x1a, 0x84, 0x75, 0xf8, 0x11, 0x54, 0xf0, 0x91, 0x12, 0x8d, 0x35, 0x62,
	0xf3, 0x44, 0xa0, 0x0d, 0x06, 0x57, 0x5e, 0x34, 0x9f, 0xbc, 0x00, 0xb5, 0x18, 0xbe, 0x35, 0xa8,
	0x05, 0xb7, 0x0e, 0x03, 0xb2, 0x07, 0x7b, 0xb8, 0x93, 0xf1, 0xa1, 0x95, 0x0f, 0xa3, 0xb9, 0xc1,
	0xf9, 0xb1, 0x46, 0x8d, 0x51, 0x28, 0x86, 0x2a, 0xf5, 0xcc, 0x19, 0x6c, 0x43, 0x22, 0x61, 0x1b,
	0x76, 0xdc, 0x2c, 0x6e, 0xd3, 0x99, 0xb0, 0xb7, 0x75, 0xd0, 0xa6, 0x5f, 0x27, 0x3f, 0x08, 0xbd,
	0xfd, 0xd6, 0xfa, 0xe0, 0xa0, 0xb9, 0x7e, 0x6d, 0x41, 0xb1, 0x7b, 0x34, 0x73, 0xa0, 0x0d, 0xae,
	0x0a, 0x52, 0x92, 0x2a, 0x97, 0x3b, 0xc4, 0x9e, 0x51, 0xea, 0x41, 0x39, 0x08, 0xb5, 0x83, 0xb6,
	0xb8, 0x55, 0x92, 0xea, 0x68, 0x31, 0xe3, 0x0a, 0x1d, 0xec, 0x17, 0xe2, 0x12, 0x3c, 0xae, 0x9d,
	0x02, 0x09, 0xad, 0xcc, 0x93, 0x58, 0x42, 0xcb, 0x66, 0x74, 0xea, 0xb0, 0x87, 0xba, 0x71, 0xab,
	0x62, 0x12, 0x67, 0x4e, 0x06, 0x7c, 0xe6, 0x56, 0xec, 0x8c, 0x4e, 0x5a, 0xd3, 0x07, 0x70, 0xbe,
	0x18, 0x95, 0xa3, 0xea, 0x68, 0xf1, 0x88, 0xff, 0x2b, 0x32, 0x7f, 0xd3, 0xe8, 0x17, 0x51, 0x2f,
	0xf7, 0x7d, 0xec, 0x3e, 0xcd, 0xed, 0x7a, 0xd9, 0x1b, 0x55, 0x1b, 0x5b, 0x8c, 0x4b, 0x52, 0x4d,
	0xe5, 0x34, 0x11, 0x97, 0x96, 0x31, 0x3a, 0xb6, 0xe8, 0x42, 0x71, 0x58, 0x92, 0xea, 0x8e, 0x8c,
	0xf5, 0xc9, 0x17, 0x42, 0xf3, 0xdf, 0x73, 0x18, 0xa3, 0xa3, 0x0e, 0x3e, 0xa7, 0xac, 0x17, 0x07,
	0x72, 0x00, 0xec, 0x92, 0x4e, 0xba, 0x4d, 0x6d, 0x1b, 0xe3, 0x76, 0x39, 0xf9, 0x7f, 0x6e, 0xc5,
	0x5f, 0x6e, 0x5e, 0x35, 0xc6, 0x5d, 0x1c, 0xc8, 0xac, 0x8b, 0xd5, 0xfc, 0x31, 0xcd, 0x12, 0xc7,
	0xee, 0xfe, 0x61, 0x94, 0x6c, 0x8e, 0xe9, 0xe1, 0xa6, 0xe9, 0xd7, 0x10, 0x4d, 0x72, 0x99, 0xc0,
	0x79, 0x46, 0xc7, 0xde, 0x82, 0x3a, 0xbf, 0xfa, 0xf6, 0x73, 0x4c, 0xbe, 0x7e, 0x7f, 0x48, 0xde,
	0x3d, 0xbf, 0xc1, 0x59, 0xd9, 0x4e, 0xff, 0xe5, 0xb4, 0x96, 0x59, 0x7c, 0xf0, 0xa7, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x8f, 0x58, 0x62, 0xb9, 0xa6, 0x02, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if !this.SecretRef.Equal(that1.SecretRef) {
		return false
	}
	if this.RoleArn != that1.RoleArn {
		return false
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if !this.Filters[i].Equal(that1.Filters[i]) {
			return false
		}
	}
	if this.PublicIp != that1.PublicIp {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TagFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter)
	if !ok {
		that2, ok := that.(TagFilter)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Spec == nil {
		if this.Spec != nil {
			return false
		}
	} else if this.Spec == nil {
		return false
	} else if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TagFilter_Key) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_Key)
	if !ok {
		that2, ok := that.(TagFilter_Key)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	return true
}
func (this *TagFilter_KvPair_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_KvPair_)
	if !ok {
		that2, ok := that.(TagFilter_KvPair_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.KvPair.Equal(that1.KvPair) {
		return false
	}
	return true
}
func (this *TagFilter_KvPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_KvPair)
	if !ok {
		that2, ok := that.(TagFilter_KvPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
