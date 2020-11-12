// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/gloo/api/external/envoy/type/range.proto

package _type

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Specifies the int64 start and end of the range using half-open interval semantics [start,
// end).
type Int64Range struct {
	// start of the range (inclusive)
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Range) Reset()         { *m = Int64Range{} }
func (m *Int64Range) String() string { return proto.CompactTextString(m) }
func (*Int64Range) ProtoMessage()    {}
func (*Int64Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_57e989c7a8407978, []int{0}
}
func (m *Int64Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Range.Unmarshal(m, b)
}
func (m *Int64Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Range.Marshal(b, m, deterministic)
}
func (m *Int64Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Range.Merge(m, src)
}
func (m *Int64Range) XXX_Size() int {
	return xxx_messageInfo_Int64Range.Size(m)
}
func (m *Int64Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Range.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Range proto.InternalMessageInfo

func (m *Int64Range) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Int64Range) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

// Specifies the double start and end of the range using half-open interval semantics [start,
// end).
type DoubleRange struct {
	// start of the range (inclusive)
	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  float64  `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleRange) Reset()         { *m = DoubleRange{} }
func (m *DoubleRange) String() string { return proto.CompactTextString(m) }
func (*DoubleRange) ProtoMessage()    {}
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_57e989c7a8407978, []int{1}
}
func (m *DoubleRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleRange.Unmarshal(m, b)
}
func (m *DoubleRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleRange.Marshal(b, m, deterministic)
}
func (m *DoubleRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRange.Merge(m, src)
}
func (m *DoubleRange) XXX_Size() int {
	return xxx_messageInfo_DoubleRange.Size(m)
}
func (m *DoubleRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRange.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRange proto.InternalMessageInfo

func (m *DoubleRange) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DoubleRange) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*Int64Range)(nil), "envoy.type.Int64Range")
	proto.RegisterType((*DoubleRange)(nil), "envoy.type.DoubleRange")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/gloo/api/external/envoy/type/range.proto", fileDescriptor_57e989c7a8407978)
}

var fileDescriptor_57e989c7a8407978 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0x4e, 0x03, 0x31,
	0x10, 0x45, 0x65, 0x02, 0x14, 0x43, 0x83, 0x56, 0x29, 0x56, 0x29, 0x10, 0x4a, 0x45, 0x13, 0x4f,
	0x41, 0xe0, 0x00, 0x11, 0x42, 0xa2, 0x41, 0x51, 0x4a, 0xba, 0xdd, 0x64, 0x64, 0x16, 0x16, 0x8f,
	0xe5, 0x9d, 0x20, 0xef, 0x8d, 0x38, 0x02, 0xe7, 0xe1, 0x0e, 0xf4, 0xc8, 0xe3, 0x82, 0x26, 0x14,
	0x74, 0xdf, 0x4f, 0xfe, 0x5f, 0x7f, 0x3e, 0x3c, 0xba, 0x4e, 0x9e, 0xf7, 0xad, 0xdd, 0xf2, 0x1b,
	0x0e, 0xdc, 0xf3, 0xa2, 0x63, 0x74, 0x3d, 0xf3, 0x82, 0x76, 0x8e, 0x30, 0x44, 0x7e, 0xa1, 0xad,
	0x0c, 0x8a, 0xb0, 0x09, 0x1d, 0x52, 0x12, 0x8a, 0xbe, 0xe9, 0x91, 0xfc, 0x3b, 0x8f, 0x28, 0x63,
	0x20, 0x8c, 0x8d, 0x77, 0x64, 0x43, 0x64, 0xe1, 0x0a, 0x94, 0xdb, 0xcc, 0x67, 0x53, 0xc7, 0x8e,
	0x15, 0x63, 0x56, 0xe5, 0xc7, 0xac, 0xa2, 0x24, 0x05, 0x52, 0x92, 0xc2, 0xe6, 0x4b, 0x80, 0x07,
	0x2f, 0xb7, 0xcb, 0x4d, 0x4e, 0xaa, 0xa6, 0x70, 0x32, 0x48, 0x13, 0xa5, 0x36, 0x97, 0xe6, 0x6a,
	0xb2, 0x29, 0x8f, 0xea, 0x1c, 0x26, 0xe4, 0x77, 0xf5, 0x91, 0xb2, 0x2c, 0xe7, 0x37, 0x70, 0x76,
	0xc7, 0xfb, 0xb6, 0xa7, 0x03, 0x36, 0x73, 0xc0, 0x66, 0xd4, 0xb6, 0x4a, 0x9f, 0xdf, 0xc7, 0xe6,
	0xe3, 0xeb, 0xc2, 0x40, 0xdd, 0xb1, 0xd5, 0xbe, 0x21, 0x72, 0x1a, 0xed, 0x6f, 0xf5, 0x15, 0x68,
	0xe4, 0x3a, 0x97, 0x5b, 0x9b, 0xa7, 0xfb, 0x7f, 0x8c, 0x14, 0x5e, 0xdd, 0x5f, 0x43, 0xb5, 0xa7,
	0x7a, 0xed, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x58, 0xc4, 0xcd, 0x75, 0x01, 0x00,
	0x00,
}

func (this *Int64Range) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Int64Range)
	if !ok {
		that2, ok := that.(Int64Range)
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
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DoubleRange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DoubleRange)
	if !ok {
		that2, ok := that.(DoubleRange)
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
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
