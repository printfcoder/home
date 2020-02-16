// Code generated by protoc-gen-go. DO NOT EDIT.
// source: finance/book/book.proto

package book

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/printfcoder/home/proto/common/response"
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

type Expense struct {
	Label                int32    `protobuf:"varint,1,opt,name=label,proto3" json:"label,omitempty"`
	Value                float32  `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	AccountType          string   `protobuf:"bytes,3,opt,name=accountType,proto3" json:"accountType,omitempty"`
	Members              []int32  `protobuf:"varint,4,rep,packed,name=members,proto3" json:"members,omitempty"`
	Time                 int64    `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expense) Reset()         { *m = Expense{} }
func (m *Expense) String() string { return proto.CompactTextString(m) }
func (*Expense) ProtoMessage()    {}
func (*Expense) Descriptor() ([]byte, []int) {
	return fileDescriptor_acfc4e7a0c92f58d, []int{0}
}

func (m *Expense) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expense.Unmarshal(m, b)
}
func (m *Expense) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expense.Marshal(b, m, deterministic)
}
func (m *Expense) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expense.Merge(m, src)
}
func (m *Expense) XXX_Size() int {
	return xxx_messageInfo_Expense.Size(m)
}
func (m *Expense) XXX_DiscardUnknown() {
	xxx_messageInfo_Expense.DiscardUnknown(m)
}

var xxx_messageInfo_Expense proto.InternalMessageInfo

func (m *Expense) GetLabel() int32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *Expense) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Expense) GetAccountType() string {
	if m != nil {
		return m.AccountType
	}
	return ""
}

func (m *Expense) GetMembers() []int32 {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Expense) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type NewExpenseRequest struct {
	Expense              *Expense `protobuf:"bytes,1,opt,name=expense,proto3" json:"expense,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewExpenseRequest) Reset()         { *m = NewExpenseRequest{} }
func (m *NewExpenseRequest) String() string { return proto.CompactTextString(m) }
func (*NewExpenseRequest) ProtoMessage()    {}
func (*NewExpenseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acfc4e7a0c92f58d, []int{1}
}

func (m *NewExpenseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewExpenseRequest.Unmarshal(m, b)
}
func (m *NewExpenseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewExpenseRequest.Marshal(b, m, deterministic)
}
func (m *NewExpenseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewExpenseRequest.Merge(m, src)
}
func (m *NewExpenseRequest) XXX_Size() int {
	return xxx_messageInfo_NewExpenseRequest.Size(m)
}
func (m *NewExpenseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewExpenseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewExpenseRequest proto.InternalMessageInfo

func (m *NewExpenseRequest) GetExpense() *Expense {
	if m != nil {
		return m.Expense
	}
	return nil
}

func init() {
	proto.RegisterType((*Expense)(nil), "Expense")
	proto.RegisterType((*NewExpenseRequest)(nil), "NewExpenseRequest")
}

func init() { proto.RegisterFile("finance/book/book.proto", fileDescriptor_acfc4e7a0c92f58d) }

var fileDescriptor_acfc4e7a0c92f58d = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0x3d, 0x4f, 0xc3, 0x30,
	0x14, 0xc4, 0x6d, 0x43, 0xe1, 0x75, 0xc2, 0x42, 0xc2, 0x74, 0xb2, 0x32, 0x65, 0xb2, 0xa5, 0x32,
	0x30, 0x21, 0xc1, 0xd0, 0x95, 0xc1, 0x82, 0x1f, 0x90, 0x38, 0xaf, 0x34, 0x22, 0xf6, 0x0b, 0xb6,
	0xc3, 0xc7, 0x0f, 0xe0, 0x7f, 0x23, 0x92, 0xa6, 0x42, 0x62, 0xb1, 0xee, 0xce, 0xa7, 0xd3, 0xbd,
	0x83, 0xab, 0x5d, 0xe3, 0x4b, 0x6f, 0x51, 0x57, 0x44, 0xaf, 0xc3, 0xa3, 0xba, 0x40, 0x89, 0xd6,
	0xf7, 0x2f, 0x4d, 0xda, 0xf7, 0x95, 0xb2, 0xe4, 0x74, 0x17, 0x1a, 0x9f, 0x76, 0x96, 0x6a, 0x0c,
	0x7a, 0x4f, 0x0e, 0xf5, 0x60, 0xd1, 0x96, 0x9c, 0x23, 0xaf, 0x03, 0xc6, 0x8e, 0x7c, 0xc4, 0x23,
	0x18, 0x13, 0xf2, 0x6f, 0x06, 0xcb, 0xed, 0x67, 0x87, 0x3e, 0x22, 0xbf, 0x84, 0xac, 0x2d, 0x2b,
	0x6c, 0x05, 0x93, 0xac, 0xc8, 0xcc, 0x48, 0x7e, 0xd5, 0xf7, 0xb2, 0xed, 0x51, 0xcc, 0x24, 0x2b,
	0x66, 0x66, 0x24, 0x5c, 0xc2, 0xaa, 0xb4, 0x96, 0x7a, 0x9f, 0x9e, 0xbe, 0x3a, 0x14, 0x73, 0xc9,
	0x8a, 0x73, 0xf3, 0x57, 0xe2, 0x02, 0x96, 0x0e, 0x5d, 0x85, 0x21, 0x8a, 0x85, 0x9c, 0x17, 0x99,
	0x99, 0x28, 0xe7, 0xb0, 0x48, 0x8d, 0x43, 0x91, 0x49, 0x56, 0xcc, 0xcd, 0x80, 0xf3, 0x5b, 0xb8,
	0x78, 0xc4, 0x8f, 0x43, 0x13, 0x83, 0x6f, 0x3d, 0xc6, 0xc4, 0x73, 0x58, 0xe2, 0xa8, 0x0c, 0x95,
	0x56, 0x9b, 0x33, 0x35, 0x39, 0xa6, 0x8f, 0xcd, 0x16, 0x16, 0xcf, 0x11, 0x03, 0xbf, 0x03, 0x78,
	0xa8, 0xeb, 0xe9, 0x14, 0xae, 0xfe, 0xa5, 0xad, 0xaf, 0xd5, 0xb8, 0x85, 0x3a, 0x4e, 0x60, 0x0e,
	0x20, 0x3f, 0xa9, 0x4e, 0x87, 0x39, 0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x55, 0x48, 0xe0,
	0x27, 0x6b, 0x01, 0x00, 0x00,
}
