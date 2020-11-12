// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/ingress/api/v1/service.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

//
//A simple wrapper for a Kubernetes Service Object.
type KubeService struct {
	// a raw byte representation of the kubernetes service this resource wraps
	KubeServiceSpec *types.Any `protobuf:"bytes,1,opt,name=kube_service_spec,json=kubeServiceSpec,proto3" json:"kube_service_spec,omitempty"`
	// a raw byte representation of the service status of the kubernetes service object
	KubeServiceStatus *types.Any `protobuf:"bytes,2,opt,name=kube_service_status,json=kubeServiceStatus,proto3" json:"kube_service_status,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *KubeService) Reset()         { *m = KubeService{} }
func (m *KubeService) String() string { return proto.CompactTextString(m) }
func (*KubeService) ProtoMessage()    {}
func (*KubeService) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41331169e99fee, []int{0}
}
func (m *KubeService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubeService.Unmarshal(m, b)
}
func (m *KubeService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubeService.Marshal(b, m, deterministic)
}
func (m *KubeService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubeService.Merge(m, src)
}
func (m *KubeService) XXX_Size() int {
	return xxx_messageInfo_KubeService.Size(m)
}
func (m *KubeService) XXX_DiscardUnknown() {
	xxx_messageInfo_KubeService.DiscardUnknown(m)
}

var xxx_messageInfo_KubeService proto.InternalMessageInfo

func (m *KubeService) GetKubeServiceSpec() *types.Any {
	if m != nil {
		return m.KubeServiceSpec
	}
	return nil
}

func (m *KubeService) GetKubeServiceStatus() *types.Any {
	if m != nil {
		return m.KubeServiceStatus
	}
	return nil
}

func (m *KubeService) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*KubeService)(nil), "ingress.solo.io.KubeService")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/ingress/api/v1/service.proto", fileDescriptor_4e41331169e99fee)
}

var fileDescriptor_4e41331169e99fee = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x49, 0x15, 0x41, 0xe5, 0x0e, 0x55, 0x43, 0x55, 0x95, 0x0a, 0x15, 0xc4, 0xc4, 0x52,
	0x5b, 0xc0, 0x52, 0x31, 0x20, 0xa8, 0xd8, 0x10, 0x4b, 0xbb, 0xb1, 0x54, 0x8e, 0x7b, 0x18, 0x93,
	0x36, 0x67, 0xc5, 0x4e, 0xa4, 0xae, 0x7d, 0x1a, 0x1e, 0x85, 0xa7, 0x60, 0x60, 0x61, 0xee, 0x1b,
	0xa0, 0x3a, 0x0e, 0xff, 0x24, 0x90, 0xd8, 0x12, 0x7d, 0x77, 0xbf, 0xfb, 0xee, 0x4c, 0x46, 0x52,
	0xd9, 0x87, 0x3c, 0xa6, 0x02, 0x17, 0xcc, 0xe0, 0x1c, 0x07, 0x0a, 0x99, 0x9c, 0x23, 0x0e, 0x60,
	0x26, 0x81, 0xe9, 0x0c, 0x1f, 0x41, 0x58, 0xc3, 0x54, 0x2a, 0x33, 0x30, 0x86, 0x71, 0xad, 0x58,
	0x71, 0xc2, 0x0c, 0x64, 0x85, 0x12, 0x40, 0x75, 0x86, 0x16, 0xa3, 0xa6, 0xa7, 0x74, 0x13, 0x40,
	0x15, 0xf6, 0xda, 0x12, 0x25, 0x3a, 0xc6, 0x36, 0x5f, 0x65, 0x59, 0x6f, 0x4f, 0x22, 0xca, 0xb9,
	0x0b, 0xb5, 0x18, 0xe7, 0xf7, 0x8c, 0xa7, 0x4b, 0x8f, 0xfa, 0x6e, 0x74, 0xa2, 0x6c, 0x35, 0x60,
	0x01, 0x96, 0xcf, 0xb8, 0xe5, 0x9e, 0xef, 0xff, 0xe4, 0xc6, 0x72, 0x9b, 0x9b, 0xdf, 0xba, 0xab,
	0xff, 0x92, 0x1f, 0xbd, 0x05, 0xa4, 0x71, 0x93, 0xc7, 0x30, 0x29, 0xad, 0xa3, 0x4b, 0xd2, 0x4a,
	0xf2, 0x18, 0xa6, 0x7e, 0x8b, 0xa9, 0xd1, 0x20, 0xba, 0xc1, 0x61, 0x70, 0xdc, 0x38, 0x6d, 0xd3,
	0x52, 0x92, 0x56, 0x92, 0xf4, 0x2a, 0x5d, 0x8e, 0x9b, 0xc9, 0x67, 0xf7, 0x44, 0x83, 0x88, 0xae,
	0xc9, 0xee, 0xf7, 0x04, 0xa7, 0xd3, 0xad, 0xfd, 0x91, 0xd1, 0xfa, 0x9a, 0xe1, 0xca, 0xa3, 0x21,
	0xa9, 0x57, 0x7b, 0x76, 0x77, 0x5c, 0x6b, 0x87, 0x0a, 0xcc, 0xa0, 0xba, 0x23, 0xbd, 0xf5, 0x74,
	0x14, 0x3e, 0xbf, 0x1c, 0x6c, 0x8d, 0x3f, 0xaa, 0xcf, 0x3b, 0xab, 0x75, 0x18, 0x92, 0x9a, 0x29,
	0x56, 0xeb, 0x90, 0x44, 0x75, 0xaf, 0x61, 0x46, 0x17, 0x4f, 0xaf, 0xfd, 0xe0, 0x6e, 0xf8, 0xbf,
	0x37, 0xd5, 0x89, 0xf4, 0x87, 0x8b, 0xb7, 0x9d, 0xf2, 0xd9, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x28, 0xb3, 0x33, 0x03, 0x16, 0x02, 0x00, 0x00,
}

func (this *KubeService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubeService)
	if !ok {
		that2, ok := that.(KubeService)
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
	if !this.KubeServiceSpec.Equal(that1.KubeServiceSpec) {
		return false
	}
	if !this.KubeServiceStatus.Equal(that1.KubeServiceStatus) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
