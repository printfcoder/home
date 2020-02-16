// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/page/page.proto

package page

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Page struct {
	No                   int64    `protobuf:"varint,1,opt,name=no,proto3" json:"no,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Data                 *any.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Total                int64    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceb4d00551747c01, []int{0}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetNo() int64 {
	if m != nil {
		return m.No
	}
	return 0
}

func (m *Page) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Page) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Page) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Page)(nil), "page.Page")
}

func init() { proto.RegisterFile("common/page/page.proto", fileDescriptor_ceb4d00551747c01) }

var fileDescriptor_ceb4d00551747c01 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2f, 0x48, 0x4c, 0x4f, 0x05, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x2c, 0x20, 0xb6, 0x94, 0x64, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x2c, 0xa9, 0x34,
	0x4d, 0x3f, 0x31, 0xaf, 0x12, 0xa2, 0x40, 0x29, 0x8b, 0x8b, 0x25, 0x20, 0x31, 0x3d, 0x55, 0x88,
	0x8f, 0x8b, 0x29, 0x2f, 0x5f, 0x82, 0x51, 0x81, 0x51, 0x83, 0x39, 0x88, 0x29, 0x2f, 0x5f, 0x48,
	0x88, 0x8b, 0xa5, 0x38, 0xb3, 0x2a, 0x55, 0x82, 0x09, 0x2c, 0x02, 0x66, 0x0b, 0x69, 0x70, 0xb1,
	0xa4, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xe8, 0x41, 0x4c, 0xd5,
	0x83, 0x99, 0xaa, 0xe7, 0x98, 0x57, 0x19, 0x04, 0x56, 0x21, 0x24, 0xc2, 0xc5, 0x5a, 0x92, 0x5f,
	0x92, 0x98, 0x23, 0xc1, 0x02, 0xd6, 0x0e, 0xe1, 0x24, 0xb1, 0x81, 0x55, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x66, 0x8d, 0x5f, 0x2b, 0xad, 0x00, 0x00, 0x00,
}
