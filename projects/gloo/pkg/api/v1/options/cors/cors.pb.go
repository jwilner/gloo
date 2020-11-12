// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/gloo/api/v1/options/cors/cors.proto

package cors

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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

// CorsPolicy defines Cross-Origin Resource Sharing for a virtual service.
type CorsPolicy struct {
	// Specifies the origins that will be allowed to make CORS requests.
	//
	// An origin is allowed if either allow_origin or allow_origin_regex match.
	AllowOrigin []string `protobuf:"bytes,1,rep,name=allow_origin,json=allowOrigin,proto3" json:"allow_origin,omitempty"`
	// Specifies regex patterns that match origins that will be allowed to make
	// CORS requests.
	//
	// An origin is allowed if either allow_origin or allow_origin_regex match.
	AllowOriginRegex []string `protobuf:"bytes,2,rep,name=allow_origin_regex,json=allowOriginRegex,proto3" json:"allow_origin_regex,omitempty"`
	// Specifies the content for the *access-control-allow-methods* header.
	AllowMethods []string `protobuf:"bytes,3,rep,name=allow_methods,json=allowMethods,proto3" json:"allow_methods,omitempty"`
	// Specifies the content for the *access-control-allow-headers* header.
	AllowHeaders []string `protobuf:"bytes,4,rep,name=allow_headers,json=allowHeaders,proto3" json:"allow_headers,omitempty"`
	// Specifies the content for the *access-control-expose-headers* header.
	ExposeHeaders []string `protobuf:"bytes,5,rep,name=expose_headers,json=exposeHeaders,proto3" json:"expose_headers,omitempty"`
	// Specifies the content for the *access-control-max-age* header.
	MaxAge string `protobuf:"bytes,6,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	// Specifies whether the resource allows credentials.
	AllowCredentials bool `protobuf:"varint,7,opt,name=allow_credentials,json=allowCredentials,proto3" json:"allow_credentials,omitempty"`
	// Optional, only applies to route-specific CORS Policies, defaults to false.
	// If set, the CORS Policy (specified on the virtual host) will be disabled for this route.
	DisableForRoute      bool     `protobuf:"varint,8,opt,name=disable_for_route,json=disableForRoute,proto3" json:"disable_for_route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CorsPolicy) Reset()         { *m = CorsPolicy{} }
func (m *CorsPolicy) String() string { return proto.CompactTextString(m) }
func (*CorsPolicy) ProtoMessage()    {}
func (*CorsPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc2768d8846e744c, []int{0}
}
func (m *CorsPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CorsPolicy.Unmarshal(m, b)
}
func (m *CorsPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CorsPolicy.Marshal(b, m, deterministic)
}
func (m *CorsPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CorsPolicy.Merge(m, src)
}
func (m *CorsPolicy) XXX_Size() int {
	return xxx_messageInfo_CorsPolicy.Size(m)
}
func (m *CorsPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_CorsPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_CorsPolicy proto.InternalMessageInfo

func (m *CorsPolicy) GetAllowOrigin() []string {
	if m != nil {
		return m.AllowOrigin
	}
	return nil
}

func (m *CorsPolicy) GetAllowOriginRegex() []string {
	if m != nil {
		return m.AllowOriginRegex
	}
	return nil
}

func (m *CorsPolicy) GetAllowMethods() []string {
	if m != nil {
		return m.AllowMethods
	}
	return nil
}

func (m *CorsPolicy) GetAllowHeaders() []string {
	if m != nil {
		return m.AllowHeaders
	}
	return nil
}

func (m *CorsPolicy) GetExposeHeaders() []string {
	if m != nil {
		return m.ExposeHeaders
	}
	return nil
}

func (m *CorsPolicy) GetMaxAge() string {
	if m != nil {
		return m.MaxAge
	}
	return ""
}

func (m *CorsPolicy) GetAllowCredentials() bool {
	if m != nil {
		return m.AllowCredentials
	}
	return false
}

func (m *CorsPolicy) GetDisableForRoute() bool {
	if m != nil {
		return m.DisableForRoute
	}
	return false
}

func init() {
	proto.RegisterType((*CorsPolicy)(nil), "cors.options.gloo.solo.io.CorsPolicy")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/gloo/api/v1/options/cors/cors.proto", fileDescriptor_cc2768d8846e744c)
}

var fileDescriptor_cc2768d8846e744c = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdf, 0xca, 0xd3, 0x30,
	0x18, 0xc6, 0xe9, 0x36, 0xf7, 0x27, 0x3a, 0x75, 0x41, 0xb0, 0xee, 0x60, 0x4c, 0x45, 0x18, 0xea,
	0x1a, 0xc4, 0x2b, 0x70, 0x03, 0x51, 0x50, 0x94, 0x1e, 0x7a, 0x52, 0xd2, 0xf6, 0x5d, 0x16, 0x4d,
	0xf7, 0x86, 0x24, 0x75, 0xf5, 0x8e, 0xbc, 0x04, 0x0f, 0xbd, 0x16, 0xef, 0xc1, 0x73, 0x69, 0xd2,
	0xb9, 0xc1, 0xf7, 0x9d, 0x7c, 0x27, 0x25, 0xfc, 0x9e, 0x5f, 0x9e, 0xb7, 0x90, 0x97, 0xbc, 0x17,
	0xd2, 0xed, 0xeb, 0x3c, 0x29, 0xb0, 0x62, 0x16, 0x15, 0xae, 0x25, 0x32, 0xa1, 0x10, 0xd7, 0x50,
	0x0a, 0x60, 0xda, 0xe0, 0x57, 0x28, 0x9c, 0xf5, 0x88, 0x71, 0x2d, 0xd9, 0xf7, 0x57, 0x0c, 0xb5,
	0x93, 0x78, 0xb0, 0xac, 0x40, 0x13, 0x3e, 0x89, 0x36, 0xe8, 0x90, 0x3e, 0xf2, 0xe7, 0x2e, 0x4d,
	0xda, 0x1b, 0x49, 0xdb, 0x98, 0x48, 0x9c, 0x3f, 0x10, 0x28, 0xd0, 0x5b, 0xac, 0x3d, 0x85, 0x0b,
	0xf3, 0x85, 0x40, 0x14, 0xca, 0x4f, 0x71, 0x98, 0xd7, 0x3b, 0x76, 0x34, 0x5c, 0x6b, 0x38, 0x15,
	0x5e, 0xcd, 0xcb, 0xda, 0xf0, 0xb6, 0xbd, 0xcb, 0x29, 0x34, 0x2e, 0x94, 0x42, 0xe3, 0x02, 0x7b,
	0xf2, 0xbb, 0x47, 0xc8, 0x16, 0x8d, 0xfd, 0x8c, 0x4a, 0x16, 0x3f, 0xe8, 0x63, 0x72, 0x87, 0x2b,
	0x85, 0xc7, 0x0c, 0x8d, 0x14, 0xf2, 0x10, 0x47, 0xcb, 0xfe, 0x6a, 0x92, 0xde, 0xf6, 0xec, 0x93,
	0x47, 0xf4, 0x25, 0xa1, 0x97, 0x4a, 0x66, 0x40, 0x40, 0x13, 0xf7, 0xbc, 0x78, 0xff, 0x42, 0x4c,
	0x5b, 0x4e, 0x9f, 0x92, 0x69, 0xb0, 0x2b, 0x70, 0x7b, 0x2c, 0x6d, 0xdc, 0xf7, 0x62, 0x98, 0xf2,
	0x31, 0xb0, 0xb3, 0xb4, 0x07, 0x5e, 0x82, 0xb1, 0xf1, 0xe0, 0x42, 0x7a, 0x17, 0x18, 0x7d, 0x46,
	0xee, 0x42, 0xa3, 0xd1, 0xc2, 0x7f, 0xeb, 0x96, 0xb7, 0xa6, 0x81, 0x9e, 0xb4, 0x87, 0x64, 0x54,
	0xf1, 0x26, 0xe3, 0x02, 0xe2, 0xe1, 0x32, 0x5a, 0x4d, 0xd2, 0x61, 0xc5, 0x9b, 0x37, 0x02, 0xe8,
	0x0b, 0x32, 0x0b, 0x43, 0x0a, 0x03, 0x25, 0x1c, 0x9c, 0xe4, 0xca, 0xc6, 0xa3, 0x65, 0xb4, 0x1a,
	0x77, 0xbf, 0xbd, 0x3d, 0x73, 0xfa, 0x9c, 0xcc, 0x4a, 0x69, 0x79, 0xae, 0x20, 0xdb, 0xa1, 0xc9,
	0x0c, 0xd6, 0x0e, 0xe2, 0xb1, 0x97, 0xef, 0x75, 0xc1, 0x5b, 0x34, 0x69, 0x8b, 0x37, 0x1f, 0x7e,
	0xfd, 0x1d, 0x44, 0x3f, 0xff, 0x2c, 0xa2, 0x2f, 0x9b, 0x1b, 0x2c, 0x87, 0xfe, 0x26, 0xae, 0x5b,
	0x90, 0x7c, 0xe8, 0xdf, 0xe5, 0xf5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x60, 0x6c, 0x75,
	0x69, 0x02, 0x00, 0x00,
}

func (this *CorsPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CorsPolicy)
	if !ok {
		that2, ok := that.(CorsPolicy)
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
	if len(this.AllowOrigin) != len(that1.AllowOrigin) {
		return false
	}
	for i := range this.AllowOrigin {
		if this.AllowOrigin[i] != that1.AllowOrigin[i] {
			return false
		}
	}
	if len(this.AllowOriginRegex) != len(that1.AllowOriginRegex) {
		return false
	}
	for i := range this.AllowOriginRegex {
		if this.AllowOriginRegex[i] != that1.AllowOriginRegex[i] {
			return false
		}
	}
	if len(this.AllowMethods) != len(that1.AllowMethods) {
		return false
	}
	for i := range this.AllowMethods {
		if this.AllowMethods[i] != that1.AllowMethods[i] {
			return false
		}
	}
	if len(this.AllowHeaders) != len(that1.AllowHeaders) {
		return false
	}
	for i := range this.AllowHeaders {
		if this.AllowHeaders[i] != that1.AllowHeaders[i] {
			return false
		}
	}
	if len(this.ExposeHeaders) != len(that1.ExposeHeaders) {
		return false
	}
	for i := range this.ExposeHeaders {
		if this.ExposeHeaders[i] != that1.ExposeHeaders[i] {
			return false
		}
	}
	if this.MaxAge != that1.MaxAge {
		return false
	}
	if this.AllowCredentials != that1.AllowCredentials {
		return false
	}
	if this.DisableForRoute != that1.DisableForRoute {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
