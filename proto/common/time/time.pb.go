// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/time/time.proto

package time

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TimeSpan struct {
	Start                int64    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeSpan) Reset()         { *m = TimeSpan{} }
func (m *TimeSpan) String() string { return proto.CompactTextString(m) }
func (*TimeSpan) ProtoMessage()    {}
func (*TimeSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b5dfee31eb07a3, []int{0}
}

func (m *TimeSpan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSpan.Unmarshal(m, b)
}
func (m *TimeSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSpan.Marshal(b, m, deterministic)
}
func (m *TimeSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSpan.Merge(m, src)
}
func (m *TimeSpan) XXX_Size() int {
	return xxx_messageInfo_TimeSpan.Size(m)
}
func (m *TimeSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSpan.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSpan proto.InternalMessageInfo

func (m *TimeSpan) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *TimeSpan) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*TimeSpan)(nil), "time.TimeSpan")
}

func init() { proto.RegisterFile("common/time/time.proto", fileDescriptor_21b5dfee31eb07a3) }

var fileDescriptor_21b5dfee31eb07a3 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d, 0x05, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x2c, 0x20, 0xb6, 0x92, 0x11, 0x17, 0x47, 0x48, 0x66, 0x6e, 0x6a, 0x70, 0x41, 0x62, 0x9e, 0x90,
	0x08, 0x17, 0x6b, 0x71, 0x49, 0x62, 0x51, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x84,
	0x23, 0x24, 0xc0, 0xc5, 0x9c, 0x9a, 0x97, 0x22, 0xc1, 0x04, 0x16, 0x03, 0x31, 0x93, 0xd8, 0xc0,
	0x06, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x89, 0x01, 0xe0, 0xe0, 0x5a, 0x00, 0x00, 0x00,
}