// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/gloo/api/v1/enterprise/ratelimit.proto

package enterprise

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	math "math"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
//@solo-kit:xds-service=RateLimitDiscoveryService
//@solo-kit:resource.no_references
type RateLimitConfig struct {
	// @solo-kit:resource.name
	Domain               string                    `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*v1alpha1.Descriptor    `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	SetDescriptors       []*v1alpha1.SetDescriptor `protobuf:"bytes,3,rep,name=set_descriptors,json=setDescriptors,proto3" json:"set_descriptors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RateLimitConfig) Reset()         { *m = RateLimitConfig{} }
func (m *RateLimitConfig) String() string { return proto.CompactTextString(m) }
func (*RateLimitConfig) ProtoMessage()    {}
func (*RateLimitConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee1a2bf29d93fef4, []int{0}
}
func (m *RateLimitConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitConfig.Unmarshal(m, b)
}
func (m *RateLimitConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitConfig.Marshal(b, m, deterministic)
}
func (m *RateLimitConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitConfig.Merge(m, src)
}
func (m *RateLimitConfig) XXX_Size() int {
	return xxx_messageInfo_RateLimitConfig.Size(m)
}
func (m *RateLimitConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitConfig proto.InternalMessageInfo

func (m *RateLimitConfig) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitConfig) GetDescriptors() []*v1alpha1.Descriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimitConfig) GetSetDescriptors() []*v1alpha1.SetDescriptor {
	if m != nil {
		return m.SetDescriptors
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimitConfig)(nil), "glooe.solo.io.RateLimitConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/gloo/api/v1/enterprise/ratelimit.proto", fileDescriptor_ee1a2bf29d93fef4)
}

var fileDescriptor_ee1a2bf29d93fef4 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0x29, 0xaa, 0x84, 0x2b, 0xa8, 0xb4, 0x0a, 0xa8, 0x44, 0xa5, 0x94, 0xd0, 0x43, 0x84,
	0x14, 0x9b, 0xa4, 0xb7, 0x9e, 0x10, 0x8d, 0x38, 0xa0, 0x72, 0x49, 0x6e, 0x1c, 0x40, 0xee, 0xee,
	0xe0, 0x18, 0x76, 0x3d, 0xc6, 0x9e, 0xae, 0xda, 0x23, 0xbc, 0x02, 0x2f, 0xc1, 0x23, 0x70, 0xe1,
	0x29, 0x78, 0x03, 0xc4, 0x81, 0x17, 0xe0, 0x8e, 0xd6, 0xd9, 0x2e, 0xdb, 0xad, 0x8a, 0x84, 0xc4,
	0xcd, 0x9e, 0xf9, 0x7e, 0x3c, 0x9e, 0x19, 0x7e, 0xa4, 0x0d, 0x2d, 0x4f, 0x8e, 0x45, 0x8a, 0x85,
	0x0c, 0x98, 0xe3, 0xd8, 0xa0, 0xd4, 0x39, 0xe2, 0x18, 0x32, 0x0d, 0xd2, 0x79, 0x7c, 0x0b, 0x29,
	0x85, 0x18, 0x92, 0xca, 0x19, 0x59, 0x4e, 0x24, 0x58, 0x02, 0xef, 0xbc, 0x09, 0x20, 0xbd, 0x22,
	0xc8, 0x4d, 0x61, 0x48, 0x38, 0x8f, 0x84, 0xc9, 0xcd, 0x0a, 0x07, 0xa2, 0x12, 0x12, 0x06, 0x07,
	0xdb, 0x60, 0x4b, 0x3c, 0x5b, 0xf1, 0xa6, 0x32, 0x33, 0x21, 0xc5, 0x12, 0xfc, 0xd9, 0x0a, 0x3c,
	0xd8, 0xd6, 0x88, 0x3a, 0x87, 0x98, 0x56, 0xd6, 0x22, 0x29, 0x32, 0x68, 0x43, 0x9d, 0xdd, 0x8f,
	0xaf, 0x51, 0xce, 0x84, 0x08, 0xa8, 0x9c, 0xc6, 0xd1, 0x0a, 0xbc, 0x2c, 0x27, 0x2a, 0x77, 0x4b,
	0x35, 0xe9, 0xfa, 0x0f, 0xfa, 0x1a, 0x35, 0xc6, 0xa3, 0xac, 0x4e, 0x75, 0x34, 0x81, 0x53, 0x5a,
	0x05, 0xe1, 0xb4, 0x46, 0x0e, 0xbf, 0x32, 0xbe, 0x39, 0x57, 0x04, 0x47, 0x15, 0xfb, 0x10, 0xed,
	0x1b, 0xa3, 0x93, 0x3b, 0x7c, 0x3d, 0xc3, 0x42, 0x19, 0xbb, 0xc5, 0x76, 0xd9, 0xe8, 0xc6, 0xbc,
	0xbe, 0x25, 0x87, 0x7c, 0x23, 0x83, 0x90, 0x7a, 0xe3, 0x08, 0x7d, 0xd8, 0xea, 0xed, 0xae, 0x8d,
	0x36, 0xa6, 0x0f, 0xc4, 0x1f, 0x73, 0xe5, 0xcc, 0x79, 0xcd, 0x62, 0xd6, 0x20, 0xe7, 0x6d, 0x56,
	0xf2, 0x82, 0x6f, 0x06, 0xa0, 0xd7, 0x6d, 0xa1, 0xb5, 0x28, 0xb4, 0x77, 0x85, 0xd0, 0x02, 0xa8,
	0xa5, 0x75, 0x2b, 0xb4, 0xaf, 0x61, 0xfa, 0xb3, 0xc7, 0xef, 0x36, 0xef, 0x9f, 0x9d, 0xff, 0xec,
	0x02, 0x7c, 0x69, 0x52, 0x48, 0x5e, 0xf1, 0xdb, 0x0b, 0xf2, 0xa0, 0x8a, 0x6e, 0x89, 0x3b, 0x22,
	0xb6, 0x24, 0x1a, 0x95, 0x53, 0xd1, 0x10, 0xe7, 0xf0, 0xfe, 0x04, 0x02, 0x0d, 0xee, 0x5f, 0x99,
	0x0f, 0x0e, 0x6d, 0x80, 0xe1, 0xb5, 0x11, 0x7b, 0xcc, 0x12, 0xcd, 0xfb, 0x33, 0xc8, 0x49, 0x75,
	0xe5, 0x1f, 0x76, 0xe8, 0x15, 0xe6, 0x92, 0xc7, 0xde, 0xdf, 0x41, 0x17, 0x8c, 0x3e, 0x30, 0xde,
	0x7f, 0x06, 0x94, 0x2e, 0xff, 0x7b, 0x21, 0xa3, 0x8f, 0xdf, 0x7e, 0x7c, 0xea, 0x0d, 0x87, 0xf7,
	0x2e, 0x8c, 0xe6, 0x41, 0xd3, 0x8b, 0x34, 0xfa, 0x1c, 0xb0, 0x47, 0x4f, 0x9f, 0x7f, 0xf9, 0x75,
	0x9d, 0x7d, 0xfe, 0xbe, 0xc3, 0x5e, 0x3e, 0xf9, 0x87, 0x65, 0x71, 0xef, 0xf4, 0xe5, 0x85, 0x39,
	0x5e, 0x8f, 0xd3, 0xb7, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x8e, 0xc8, 0x16, 0x77, 0x03,
	0x00, 0x00,
}

func (this *RateLimitConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RateLimitConfig)
	if !ok {
		that2, ok := that.(RateLimitConfig)
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
	if this.Domain != that1.Domain {
		return false
	}
	if len(this.Descriptors) != len(that1.Descriptors) {
		return false
	}
	for i := range this.Descriptors {
		if !this.Descriptors[i].Equal(that1.Descriptors[i]) {
			return false
		}
	}
	if len(this.SetDescriptors) != len(that1.SetDescriptors) {
		return false
	}
	for i := range this.SetDescriptors {
		if !this.SetDescriptors[i].Equal(that1.SetDescriptors[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitDiscoveryServiceClient is the client API for RateLimitDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitDiscoveryServiceClient interface {
	StreamRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_StreamRateLimitConfigClient, error)
	DeltaRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_DeltaRateLimitConfigClient, error)
	FetchRateLimitConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type rateLimitDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitDiscoveryServiceClient(cc *grpc.ClientConn) RateLimitDiscoveryServiceClient {
	return &rateLimitDiscoveryServiceClient{cc}
}

func (c *rateLimitDiscoveryServiceClient) StreamRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_StreamRateLimitConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RateLimitDiscoveryService_serviceDesc.Streams[0], "/glooe.solo.io.RateLimitDiscoveryService/StreamRateLimitConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitDiscoveryServiceStreamRateLimitConfigClient{stream}
	return x, nil
}

type RateLimitDiscoveryService_StreamRateLimitConfigClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type rateLimitDiscoveryServiceStreamRateLimitConfigClient struct {
	grpc.ClientStream
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rateLimitDiscoveryServiceClient) DeltaRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_DeltaRateLimitConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RateLimitDiscoveryService_serviceDesc.Streams[1], "/glooe.solo.io.RateLimitDiscoveryService/DeltaRateLimitConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitDiscoveryServiceDeltaRateLimitConfigClient{stream}
	return x, nil
}

type RateLimitDiscoveryService_DeltaRateLimitConfigClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type rateLimitDiscoveryServiceDeltaRateLimitConfigClient struct {
	grpc.ClientStream
}

func (x *rateLimitDiscoveryServiceDeltaRateLimitConfigClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceDeltaRateLimitConfigClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rateLimitDiscoveryServiceClient) FetchRateLimitConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/glooe.solo.io.RateLimitDiscoveryService/FetchRateLimitConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitDiscoveryServiceServer is the server API for RateLimitDiscoveryService service.
type RateLimitDiscoveryServiceServer interface {
	StreamRateLimitConfig(RateLimitDiscoveryService_StreamRateLimitConfigServer) error
	DeltaRateLimitConfig(RateLimitDiscoveryService_DeltaRateLimitConfigServer) error
	FetchRateLimitConfig(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

// UnimplementedRateLimitDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRateLimitDiscoveryServiceServer struct {
}

func (*UnimplementedRateLimitDiscoveryServiceServer) StreamRateLimitConfig(srv RateLimitDiscoveryService_StreamRateLimitConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRateLimitConfig not implemented")
}
func (*UnimplementedRateLimitDiscoveryServiceServer) DeltaRateLimitConfig(srv RateLimitDiscoveryService_DeltaRateLimitConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaRateLimitConfig not implemented")
}
func (*UnimplementedRateLimitDiscoveryServiceServer) FetchRateLimitConfig(ctx context.Context, req *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRateLimitConfig not implemented")
}

func RegisterRateLimitDiscoveryServiceServer(s *grpc.Server, srv RateLimitDiscoveryServiceServer) {
	s.RegisterService(&_RateLimitDiscoveryService_serviceDesc, srv)
}

func _RateLimitDiscoveryService_StreamRateLimitConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitDiscoveryServiceServer).StreamRateLimitConfig(&rateLimitDiscoveryServiceStreamRateLimitConfigServer{stream})
}

type RateLimitDiscoveryService_StreamRateLimitConfigServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type rateLimitDiscoveryServiceStreamRateLimitConfigServer struct {
	grpc.ServerStream
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RateLimitDiscoveryService_DeltaRateLimitConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitDiscoveryServiceServer).DeltaRateLimitConfig(&rateLimitDiscoveryServiceDeltaRateLimitConfigServer{stream})
}

type RateLimitDiscoveryService_DeltaRateLimitConfigServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type rateLimitDiscoveryServiceDeltaRateLimitConfigServer struct {
	grpc.ServerStream
}

func (x *rateLimitDiscoveryServiceDeltaRateLimitConfigServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceDeltaRateLimitConfigServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RateLimitDiscoveryService_FetchRateLimitConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitDiscoveryServiceServer).FetchRateLimitConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooe.solo.io.RateLimitDiscoveryService/FetchRateLimitConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitDiscoveryServiceServer).FetchRateLimitConfig(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glooe.solo.io.RateLimitDiscoveryService",
	HandlerType: (*RateLimitDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRateLimitConfig",
			Handler:    _RateLimitDiscoveryService_FetchRateLimitConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRateLimitConfig",
			Handler:       _RateLimitDiscoveryService_StreamRateLimitConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRateLimitConfig",
			Handler:       _RateLimitDiscoveryService_DeltaRateLimitConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/solo-io/gloo-edge/projects/gloo/api/v1/enterprise/ratelimit.proto",
}
