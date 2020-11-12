// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/gloo/api/v1/circuit_breaker.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// CircuitBreakerConfig contains the options for customizing circuit breaking behavior.
// See the [envoy docs](https://www.envoyproxy.io/docs/envoy/v1.14.1/api-v2/api/v2/cluster/circuit_breaker.proto#envoy-api-msg-cluster-circuitbreakers)
// for the meaning of these values.
type CircuitBreakerConfig struct {
	MaxConnections       *types.UInt32Value `protobuf:"bytes,1,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	MaxPendingRequests   *types.UInt32Value `protobuf:"bytes,2,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	MaxRequests          *types.UInt32Value `protobuf:"bytes,3,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	MaxRetries           *types.UInt32Value `protobuf:"bytes,4,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CircuitBreakerConfig) Reset()         { *m = CircuitBreakerConfig{} }
func (m *CircuitBreakerConfig) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakerConfig) ProtoMessage()    {}
func (*CircuitBreakerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffebcf8da0fdbf17, []int{0}
}
func (m *CircuitBreakerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakerConfig.Unmarshal(m, b)
}
func (m *CircuitBreakerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakerConfig.Marshal(b, m, deterministic)
}
func (m *CircuitBreakerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakerConfig.Merge(m, src)
}
func (m *CircuitBreakerConfig) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakerConfig.Size(m)
}
func (m *CircuitBreakerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakerConfig proto.InternalMessageInfo

func (m *CircuitBreakerConfig) GetMaxConnections() *types.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakerConfig) GetMaxPendingRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakerConfig) GetMaxRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakerConfig) GetMaxRetries() *types.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakerConfig)(nil), "gloo.solo.io.CircuitBreakerConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/gloo/api/v1/circuit_breaker.proto", fileDescriptor_ffebcf8da0fdbf17)
}

var fileDescriptor_ffebcf8da0fdbf17 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0x73, 0x31,
	0x14, 0x86, 0xb9, 0xfd, 0xca, 0x37, 0xa4, 0x45, 0xe1, 0xd2, 0xa1, 0x14, 0x29, 0xe2, 0xe4, 0xd2,
	0x04, 0x5b, 0x1c, 0xa5, 0xd0, 0x22, 0xe2, 0x22, 0x52, 0xd0, 0xc1, 0xa5, 0xe4, 0xa6, 0xa7, 0x31,
	0xf6, 0xde, 0x9c, 0x98, 0xe4, 0x6a, 0x7e, 0x92, 0x8b, 0xbb, 0xbf, 0xc7, 0xff, 0xe0, 0x2e, 0x37,
	0xa9, 0x75, 0xad, 0xdb, 0xe1, 0x9c, 0xf7, 0x79, 0xde, 0xe1, 0x90, 0x2b, 0xa9, 0xfc, 0x63, 0x5d,
	0x50, 0x81, 0x15, 0x73, 0x58, 0xe2, 0x48, 0x21, 0x93, 0x25, 0xe2, 0x08, 0x56, 0x12, 0x98, 0xb1,
	0xf8, 0x04, 0xc2, 0xbb, 0xb8, 0x62, 0xdc, 0x28, 0xf6, 0x72, 0xc6, 0x84, 0xb2, 0xa2, 0x56, 0x7e,
	0x59, 0x58, 0xe0, 0x1b, 0xb0, 0xd4, 0x58, 0xf4, 0x98, 0x77, 0x9b, 0x08, 0x6d, 0x14, 0x54, 0xe1,
	0xa0, 0x27, 0x51, 0x62, 0x3c, 0xb0, 0x66, 0x4a, 0x99, 0xc1, 0x50, 0x22, 0xca, 0x32, 0x6a, 0x3d,
	0x16, 0xf5, 0x9a, 0xbd, 0x5a, 0x6e, 0x0c, 0x58, 0xb7, 0xbd, 0xe7, 0x10, 0x7c, 0x82, 0x20, 0xf8,
	0xb4, 0x3b, 0x79, 0x6f, 0x91, 0xde, 0x3c, 0x35, 0xce, 0x52, 0xe1, 0x1c, 0xf5, 0x5a, 0xc9, 0xfc,
	0x92, 0x1c, 0x56, 0x3c, 0x2c, 0x05, 0x6a, 0x0d, 0xc2, 0x2b, 0xd4, 0xae, 0x9f, 0x1d, 0x67, 0xa7,
	0x9d, 0xf1, 0x11, 0x4d, 0x35, 0xf4, 0xa7, 0x86, 0xde, 0x5d, 0x6b, 0x3f, 0x19, 0xdf, 0xf3, 0xb2,
	0x86, 0xc5, 0x41, 0xc5, 0xc3, 0xfc, 0x97, 0xc9, 0x6f, 0x48, 0xaf, 0xd1, 0x18, 0xd0, 0x2b, 0xa5,
	0xe5, 0xd2, 0xc2, 0x73, 0x0d, 0xce, 0xbb, 0x7e, 0x6b, 0x0f, 0x57, 0x5e, 0xf1, 0x70, 0x9b, 0xc0,
	0xc5, 0x96, 0xcb, 0xa7, 0xa4, 0xdb, 0xf8, 0x76, 0x9e, 0x7f, 0x7b, 0x78, 0x3a, 0x15, 0x0f, 0x3b,
	0xc1, 0x05, 0xe9, 0x24, 0x81, 0xb7, 0x0a, 0x5c, 0xbf, 0xbd, 0x07, 0x4f, 0x22, 0x1f, 0xf3, 0xb3,
	0xe9, 0xc7, 0x57, 0x3b, 0x7b, 0xfb, 0x1c, 0x66, 0x0f, 0xe7, 0x7f, 0x78, 0xad, 0xd9, 0xc8, 0xed,
	0x7b, 0x8b, 0xff, 0xb1, 0x62, 0xf2, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x28, 0x98, 0x9f, 0x7d, 0x1a,
	0x02, 0x00, 0x00,
}

func (this *CircuitBreakerConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CircuitBreakerConfig)
	if !ok {
		that2, ok := that.(CircuitBreakerConfig)
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
	if !this.MaxConnections.Equal(that1.MaxConnections) {
		return false
	}
	if !this.MaxPendingRequests.Equal(that1.MaxPendingRequests) {
		return false
	}
	if !this.MaxRequests.Equal(that1.MaxRequests) {
		return false
	}
	if !this.MaxRetries.Equal(that1.MaxRetries) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
