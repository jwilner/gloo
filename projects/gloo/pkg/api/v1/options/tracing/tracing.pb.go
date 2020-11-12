// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/gloo/api/v1/options/tracing/tracing.proto

package tracing

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

// Contains settings for configuring Envoy's tracing capabilities at the listener level.
// See here for additional information on Envoy's tracing capabilities: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing.html
// See here for additional information about configuring tracing with Gloo: https://gloo.solo.io/user_guides/setup_options/observability/#tracing
type ListenerTracingSettings struct {
	// Optional. If specified, Envoy will include the headers and header values for any matching request headers.
	RequestHeadersForTags []string `protobuf:"bytes,1,rep,name=request_headers_for_tags,json=requestHeadersForTags,proto3" json:"request_headers_for_tags,omitempty"`
	// Optional. If true, Envoy will include logs for streaming events. Default: false.
	Verbose bool `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
	// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
	// TracePercentages defines the limits for random, forced, and overall tracing percentages.
	TracePercentages     *TracePercentages `protobuf:"bytes,3,opt,name=trace_percentages,json=tracePercentages,proto3" json:"trace_percentages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListenerTracingSettings) Reset()         { *m = ListenerTracingSettings{} }
func (m *ListenerTracingSettings) String() string { return proto.CompactTextString(m) }
func (*ListenerTracingSettings) ProtoMessage()    {}
func (*ListenerTracingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f83a6ac2c2131a22, []int{0}
}
func (m *ListenerTracingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerTracingSettings.Unmarshal(m, b)
}
func (m *ListenerTracingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerTracingSettings.Marshal(b, m, deterministic)
}
func (m *ListenerTracingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerTracingSettings.Merge(m, src)
}
func (m *ListenerTracingSettings) XXX_Size() int {
	return xxx_messageInfo_ListenerTracingSettings.Size(m)
}
func (m *ListenerTracingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerTracingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerTracingSettings proto.InternalMessageInfo

func (m *ListenerTracingSettings) GetRequestHeadersForTags() []string {
	if m != nil {
		return m.RequestHeadersForTags
	}
	return nil
}

func (m *ListenerTracingSettings) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *ListenerTracingSettings) GetTracePercentages() *TracePercentages {
	if m != nil {
		return m.TracePercentages
	}
	return nil
}

// Contains settings for configuring Envoy's tracing capabilities at the route level.
// Note: must also specify ListenerTracingSettings for the associated listener.
// See here for additional information on Envoy's tracing capabilities: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing.html
// See here for additional information about configuring tracing with Gloo: https://gloo.solo.io/user_guides/setup_options/observability/#tracing
type RouteTracingSettings struct {
	// Optional. If set, will be used to identify the route that produced the trace.
	// Note that this value will be overridden if the "x-envoy-decorator-operation" header is passed.
	RouteDescriptor string `protobuf:"bytes,1,opt,name=route_descriptor,json=routeDescriptor,proto3" json:"route_descriptor,omitempty"`
	// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
	// TracePercentages defines the limits for random, forced, and overall tracing percentages.
	TracePercentages     *TracePercentages `protobuf:"bytes,2,opt,name=trace_percentages,json=tracePercentages,proto3" json:"trace_percentages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RouteTracingSettings) Reset()         { *m = RouteTracingSettings{} }
func (m *RouteTracingSettings) String() string { return proto.CompactTextString(m) }
func (*RouteTracingSettings) ProtoMessage()    {}
func (*RouteTracingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f83a6ac2c2131a22, []int{1}
}
func (m *RouteTracingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTracingSettings.Unmarshal(m, b)
}
func (m *RouteTracingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTracingSettings.Marshal(b, m, deterministic)
}
func (m *RouteTracingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTracingSettings.Merge(m, src)
}
func (m *RouteTracingSettings) XXX_Size() int {
	return xxx_messageInfo_RouteTracingSettings.Size(m)
}
func (m *RouteTracingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTracingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTracingSettings proto.InternalMessageInfo

func (m *RouteTracingSettings) GetRouteDescriptor() string {
	if m != nil {
		return m.RouteDescriptor
	}
	return ""
}

func (m *RouteTracingSettings) GetTracePercentages() *TracePercentages {
	if m != nil {
		return m.TracePercentages
	}
	return nil
}

// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
// TracePercentages defines the limits for random, forced, and overall tracing percentages.
type TracePercentages struct {
	// Percentage of requests that should produce traces when the `x-client-trace-id` header is provided.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	ClientSamplePercentage *types.FloatValue `protobuf:"bytes,1,opt,name=client_sample_percentage,json=clientSamplePercentage,proto3" json:"client_sample_percentage,omitempty"`
	// Percentage of requests that should produce traces by random sampling.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	RandomSamplePercentage *types.FloatValue `protobuf:"bytes,2,opt,name=random_sample_percentage,json=randomSamplePercentage,proto3" json:"random_sample_percentage,omitempty"`
	// Overall percentage of requests that should produce traces.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	OverallSamplePercentage *types.FloatValue `protobuf:"bytes,3,opt,name=overall_sample_percentage,json=overallSamplePercentage,proto3" json:"overall_sample_percentage,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}          `json:"-"`
	XXX_unrecognized        []byte            `json:"-"`
	XXX_sizecache           int32             `json:"-"`
}

func (m *TracePercentages) Reset()         { *m = TracePercentages{} }
func (m *TracePercentages) String() string { return proto.CompactTextString(m) }
func (*TracePercentages) ProtoMessage()    {}
func (*TracePercentages) Descriptor() ([]byte, []int) {
	return fileDescriptor_f83a6ac2c2131a22, []int{2}
}
func (m *TracePercentages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TracePercentages.Unmarshal(m, b)
}
func (m *TracePercentages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TracePercentages.Marshal(b, m, deterministic)
}
func (m *TracePercentages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracePercentages.Merge(m, src)
}
func (m *TracePercentages) XXX_Size() int {
	return xxx_messageInfo_TracePercentages.Size(m)
}
func (m *TracePercentages) XXX_DiscardUnknown() {
	xxx_messageInfo_TracePercentages.DiscardUnknown(m)
}

var xxx_messageInfo_TracePercentages proto.InternalMessageInfo

func (m *TracePercentages) GetClientSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.ClientSamplePercentage
	}
	return nil
}

func (m *TracePercentages) GetRandomSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.RandomSamplePercentage
	}
	return nil
}

func (m *TracePercentages) GetOverallSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.OverallSamplePercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ListenerTracingSettings)(nil), "tracing.options.gloo.solo.io.ListenerTracingSettings")
	proto.RegisterType((*RouteTracingSettings)(nil), "tracing.options.gloo.solo.io.RouteTracingSettings")
	proto.RegisterType((*TracePercentages)(nil), "tracing.options.gloo.solo.io.TracePercentages")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/gloo/api/v1/options/tracing/tracing.proto", fileDescriptor_f83a6ac2c2131a22)
}

var fileDescriptor_f83a6ac2c2131a22 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xe5, 0x5d, 0x04, 0xd4, 0x3d, 0xb0, 0x44, 0x85, 0x86, 0x82, 0xaa, 0x55, 0x4f, 0xcb,
	0xa1, 0xb6, 0x28, 0x07, 0xee, 0x08, 0x2a, 0x0e, 0x48, 0x45, 0x69, 0x01, 0x09, 0x0e, 0x91, 0x37,
	0x99, 0xba, 0x06, 0x6f, 0xc6, 0x8c, 0x27, 0x4b, 0x5f, 0x85, 0x0b, 0x67, 0x1e, 0x81, 0x57, 0xe0,
	0x35, 0x78, 0x07, 0xee, 0x28, 0xc9, 0x2e, 0x45, 0x61, 0x97, 0x3f, 0x12, 0xa7, 0xc4, 0xdf, 0xcc,
	0xf7, 0x9b, 0xcf, 0xb6, 0x2c, 0x8f, 0xac, 0xe3, 0xb3, 0x7a, 0xaa, 0x0a, 0x9c, 0xe9, 0x88, 0x1e,
	0xf7, 0x1d, 0x6a, 0xeb, 0x11, 0xf7, 0xa1, 0xb4, 0xa0, 0x03, 0xe1, 0x1b, 0x28, 0x38, 0xb6, 0x92,
	0x36, 0xc1, 0xe9, 0xf9, 0x3d, 0x8d, 0x81, 0x1d, 0x56, 0x51, 0x33, 0x99, 0xc2, 0x55, 0x76, 0xf9,
	0x55, 0x81, 0x90, 0x31, 0xb9, 0xb3, 0x5c, 0x2e, 0xda, 0x54, 0x63, 0x55, 0x0d, 0x5a, 0x39, 0xdc,
	0xd9, 0xb2, 0x68, 0xb1, 0x6d, 0xd4, 0xcd, 0x5f, 0xe7, 0xd9, 0xd9, 0xb5, 0x88, 0xd6, 0xb7, 0xe3,
	0x18, 0xa7, 0xf5, 0xa9, 0x7e, 0x4f, 0x26, 0x04, 0xa0, 0xb8, 0xae, 0x5e, 0xd6, 0x64, 0x1a, 0xfa,
	0xa2, 0x9e, 0xc0, 0x39, 0x77, 0x50, 0x38, 0xe7, 0x4e, 0xdb, 0xfb, 0x22, 0xe4, 0xf6, 0x53, 0x17,
	0x19, 0x2a, 0xa0, 0x93, 0x2e, 0xd2, 0x31, 0x30, 0xbb, 0xca, 0xc6, 0xe4, 0x81, 0x4c, 0x09, 0xde,
	0xd5, 0x10, 0x39, 0x3f, 0x03, 0x53, 0x02, 0xc5, 0xfc, 0x14, 0x29, 0x67, 0x63, 0x63, 0x2a, 0xc6,
	0xc3, 0xc9, 0x46, 0x76, 0x63, 0x51, 0x7f, 0xd2, 0x95, 0x0f, 0x91, 0x4e, 0x8c, 0x8d, 0x49, 0x2a,
	0xaf, 0xcc, 0x81, 0xa6, 0x18, 0x21, 0x1d, 0x8c, 0xc5, 0xe4, 0x6a, 0xb6, 0x5c, 0x26, 0xaf, 0xe5,
	0xf5, 0x66, 0xe3, 0x90, 0x07, 0xa0, 0x02, 0x2a, 0x36, 0x16, 0x62, 0x3a, 0x1c, 0x8b, 0xc9, 0xe6,
	0x81, 0x52, 0xbf, 0x3b, 0x12, 0xd5, 0x84, 0x83, 0x67, 0x17, 0xae, 0x6c, 0xc4, 0x3d, 0x65, 0xef,
	0xa3, 0x90, 0x5b, 0x19, 0xd6, 0x0c, 0xfd, 0x8d, 0xdc, 0x95, 0x23, 0x6a, 0xf4, 0xbc, 0x84, 0x58,
	0x90, 0x0b, 0x8c, 0x94, 0x8a, 0xb1, 0x98, 0x6c, 0x64, 0xd7, 0x5a, 0xfd, 0xd1, 0x0f, 0x79, 0x75,
	0xc0, 0xc1, 0x7f, 0x0a, 0xf8, 0x61, 0x20, 0x47, 0xfd, 0xb6, 0xe4, 0xb9, 0x4c, 0x0b, 0xef, 0xa0,
	0xe2, 0x3c, 0x9a, 0x59, 0xf0, 0x3f, 0x4f, 0x6e, 0x43, 0x6e, 0x1e, 0xdc, 0x56, 0xdd, 0xc5, 0xaa,
	0xe5, 0xc5, 0xaa, 0x43, 0x8f, 0x86, 0x5f, 0x18, 0x5f, 0x43, 0x76, 0xb3, 0x33, 0x1f, 0xb7, 0xde,
	0x0b, 0x6e, 0x83, 0x25, 0x53, 0x95, 0x38, 0x5b, 0x81, 0x1d, 0xfc, 0x05, 0xb6, 0x33, 0xff, 0x82,
	0x7d, 0x29, 0x6f, 0xe1, 0x1c, 0xc8, 0x78, 0xbf, 0x82, 0x3b, 0xfc, 0x33, 0x77, 0x7b, 0xe1, 0xee,
	0x83, 0x1f, 0x1e, 0x7d, 0xfe, 0x76, 0x49, 0x7c, 0xfa, 0xba, 0x2b, 0x5e, 0x3d, 0xfe, 0x87, 0xb7,
	0x16, 0xde, 0xda, 0x35, 0xef, 0x6d, 0x7a, 0xb9, 0x1d, 0x7f, 0xff, 0x7b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x5c, 0x47, 0x52, 0x27, 0xbb, 0x03, 0x00, 0x00,
}

func (this *ListenerTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerTracingSettings)
	if !ok {
		that2, ok := that.(ListenerTracingSettings)
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
	if len(this.RequestHeadersForTags) != len(that1.RequestHeadersForTags) {
		return false
	}
	for i := range this.RequestHeadersForTags {
		if this.RequestHeadersForTags[i] != that1.RequestHeadersForTags[i] {
			return false
		}
	}
	if this.Verbose != that1.Verbose {
		return false
	}
	if !this.TracePercentages.Equal(that1.TracePercentages) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTracingSettings)
	if !ok {
		that2, ok := that.(RouteTracingSettings)
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
	if this.RouteDescriptor != that1.RouteDescriptor {
		return false
	}
	if !this.TracePercentages.Equal(that1.TracePercentages) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TracePercentages) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TracePercentages)
	if !ok {
		that2, ok := that.(TracePercentages)
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
	if !this.ClientSamplePercentage.Equal(that1.ClientSamplePercentage) {
		return false
	}
	if !this.RandomSamplePercentage.Equal(that1.RandomSamplePercentage) {
		return false
	}
	if !this.OverallSamplePercentage.Equal(that1.OverallSamplePercentage) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
