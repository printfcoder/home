// Code generated by protoc-gen-go. DO NOT EDIT.
// source: finance/book/book.proto

package book

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	page "github.com/printfcoder/home/proto/common/page"
	_ "github.com/printfcoder/home/proto/common/response"
	time "github.com/printfcoder/home/proto/common/time"
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
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                int32    `protobuf:"varint,2,opt,name=label,proto3" json:"label,omitempty"`
	Value                float32  `protobuf:"fixed32,3,opt,name=value,proto3" json:"value,omitempty"`
	AccountType          string   `protobuf:"bytes,4,opt,name=accountType,proto3" json:"accountType,omitempty"`
	Members              []int32  `protobuf:"varint,5,rep,packed,name=members,proto3" json:"members,omitempty"`
	Time                 int64    `protobuf:"varint,6,opt,name=time,proto3" json:"time,omitempty"`
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

func (m *Expense) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

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

type RemoveExpenseRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveExpenseRequest) Reset()         { *m = RemoveExpenseRequest{} }
func (m *RemoveExpenseRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveExpenseRequest) ProtoMessage()    {}
func (*RemoveExpenseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acfc4e7a0c92f58d, []int{2}
}

func (m *RemoveExpenseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveExpenseRequest.Unmarshal(m, b)
}
func (m *RemoveExpenseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveExpenseRequest.Marshal(b, m, deterministic)
}
func (m *RemoveExpenseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveExpenseRequest.Merge(m, src)
}
func (m *RemoveExpenseRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveExpenseRequest.Size(m)
}
func (m *RemoveExpenseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveExpenseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveExpenseRequest proto.InternalMessageInfo

func (m *RemoveExpenseRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateExpenseRequest struct {
	Expense              *Expense `protobuf:"bytes,1,opt,name=expense,proto3" json:"expense,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateExpenseRequest) Reset()         { *m = UpdateExpenseRequest{} }
func (m *UpdateExpenseRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateExpenseRequest) ProtoMessage()    {}
func (*UpdateExpenseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acfc4e7a0c92f58d, []int{3}
}

func (m *UpdateExpenseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateExpenseRequest.Unmarshal(m, b)
}
func (m *UpdateExpenseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateExpenseRequest.Marshal(b, m, deterministic)
}
func (m *UpdateExpenseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateExpenseRequest.Merge(m, src)
}
func (m *UpdateExpenseRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateExpenseRequest.Size(m)
}
func (m *UpdateExpenseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateExpenseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateExpenseRequest proto.InternalMessageInfo

func (m *UpdateExpenseRequest) GetExpense() *Expense {
	if m != nil {
		return m.Expense
	}
	return nil
}

type ListExpenseRequest struct {
	Page                 *page.Page     `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	TimeSpan             *time.TimeSpan `protobuf:"bytes,2,opt,name=timeSpan,proto3" json:"timeSpan,omitempty"`
	Labels               []int32        `protobuf:"varint,3,rep,packed,name=labels,proto3" json:"labels,omitempty"`
	AccountTypes         []string       `protobuf:"bytes,4,rep,name=accountTypes,proto3" json:"accountTypes,omitempty"`
	Members              []int32        `protobuf:"varint,5,rep,packed,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListExpenseRequest) Reset()         { *m = ListExpenseRequest{} }
func (m *ListExpenseRequest) String() string { return proto.CompactTextString(m) }
func (*ListExpenseRequest) ProtoMessage()    {}
func (*ListExpenseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acfc4e7a0c92f58d, []int{4}
}

func (m *ListExpenseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListExpenseRequest.Unmarshal(m, b)
}
func (m *ListExpenseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListExpenseRequest.Marshal(b, m, deterministic)
}
func (m *ListExpenseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListExpenseRequest.Merge(m, src)
}
func (m *ListExpenseRequest) XXX_Size() int {
	return xxx_messageInfo_ListExpenseRequest.Size(m)
}
func (m *ListExpenseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListExpenseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListExpenseRequest proto.InternalMessageInfo

func (m *ListExpenseRequest) GetPage() *page.Page {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *ListExpenseRequest) GetTimeSpan() *time.TimeSpan {
	if m != nil {
		return m.TimeSpan
	}
	return nil
}

func (m *ListExpenseRequest) GetLabels() []int32 {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ListExpenseRequest) GetAccountTypes() []string {
	if m != nil {
		return m.AccountTypes
	}
	return nil
}

func (m *ListExpenseRequest) GetMembers() []int32 {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*Expense)(nil), "Expense")
	proto.RegisterType((*NewExpenseRequest)(nil), "NewExpenseRequest")
	proto.RegisterType((*RemoveExpenseRequest)(nil), "RemoveExpenseRequest")
	proto.RegisterType((*UpdateExpenseRequest)(nil), "UpdateExpenseRequest")
	proto.RegisterType((*ListExpenseRequest)(nil), "ListExpenseRequest")
}

func init() { proto.RegisterFile("finance/book/book.proto", fileDescriptor_acfc4e7a0c92f58d) }

var fileDescriptor_acfc4e7a0c92f58d = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x97, 0x24, 0xed, 0xa4, 0x20, 0x75, 0x48, 0x60, 0xe9, 0x93, 0x65, 0x09, 0xe4, 0xa7,
	0xb5, 0x08, 0x0f, 0x20, 0x24, 0xa4, 0x02, 0xe2, 0x0d, 0x21, 0xb4, 0x94, 0x0f, 0xf0, 0x65, 0x9a,
	0x5a, 0x8d, 0x77, 0x17, 0xef, 0xa6, 0xc0, 0x67, 0xf0, 0xc0, 0x1f, 0xf1, 0x61, 0xc8, 0x6b, 0x3b,
	0x24, 0x6d, 0x29, 0x97, 0x97, 0xd1, 0x5c, 0x8f, 0x67, 0xce, 0x59, 0xc3, 0xfd, 0xd3, 0x4a, 0x66,
	0xb2, 0xa0, 0x34, 0x57, 0xea, 0xdc, 0x19, 0xae, 0x1b, 0x65, 0xd5, 0xd1, 0xf1, 0xb2, 0xb2, 0x67,
	0xeb, 0x9c, 0x17, 0xaa, 0x4e, 0x75, 0x53, 0x49, 0x7b, 0x5a, 0xa8, 0x92, 0x9a, 0xf4, 0x4c, 0xd5,
	0x94, 0xba, 0x96, 0xb4, 0x50, 0x75, 0xad, 0x64, 0xda, 0x90, 0xd1, 0x4a, 0x1a, 0xda, 0x38, 0x3d,
	0xc2, 0xb3, 0xbf, 0x46, 0xb0, 0x55, 0x4d, 0xce, 0xfc, 0xf3, 0xa4, 0xce, 0x96, 0xe4, 0x4c, 0x37,
	0x19, 0x7f, 0xf7, 0x60, 0xf2, 0xe6, 0x8b, 0x26, 0x69, 0x08, 0xef, 0x80, 0x5f, 0x95, 0xcc, 0x8b,
	0xbc, 0x64, 0x24, 0xfc, 0xaa, 0xc4, 0x19, 0x8c, 0x56, 0x59, 0x4e, 0x2b, 0xe6, 0xbb, 0x54, 0x17,
	0xb4, 0xd9, 0x8b, 0x6c, 0xb5, 0x26, 0x16, 0x44, 0x5e, 0xe2, 0x8b, 0x2e, 0xc0, 0x08, 0xa6, 0x59,
	0x51, 0xa8, 0xb5, 0xb4, 0x27, 0x5f, 0x35, 0xb1, 0x30, 0xf2, 0x92, 0x7d, 0xb1, 0x9d, 0x42, 0x06,
	0x93, 0x9a, 0xea, 0x9c, 0x1a, 0xc3, 0x46, 0x51, 0x90, 0x8c, 0xc4, 0x10, 0x22, 0x42, 0xd8, 0xde,
	0xc2, 0xc6, 0x91, 0x97, 0x04, 0xc2, 0xf9, 0xf1, 0x53, 0x38, 0x7c, 0x47, 0x9f, 0xfb, 0xcd, 0x04,
	0x7d, 0x5a, 0x93, 0xb1, 0x18, 0xc3, 0x84, 0xba, 0x8c, 0xdb, 0x72, 0xba, 0xd8, 0xe3, 0x43, 0xc7,
	0x50, 0x88, 0x1f, 0xc1, 0x4c, 0x50, 0xad, 0x2e, 0xe8, 0xd2, 0xec, 0xaf, 0xe3, 0x82, 0xf6, 0xb8,
	0xf8, 0x39, 0xcc, 0x3e, 0xea, 0x32, 0xb3, 0xf4, 0x1f, 0xdf, 0xf8, 0xe1, 0x01, 0xbe, 0xad, 0x8c,
	0xbd, 0x34, 0xfa, 0x10, 0xc2, 0x96, 0xd9, 0x7e, 0xee, 0x90, 0x77, 0x94, 0x73, 0xc7, 0xf6, 0xfb,
	0x6c, 0x49, 0xc2, 0x95, 0xf1, 0x31, 0xec, 0xb5, 0x27, 0x7e, 0xd0, 0x99, 0x74, 0xcc, 0x4e, 0x17,
	0xf3, 0xa1, 0xd5, 0x49, 0x7a, 0xd2, 0x17, 0xc5, 0xa6, 0x0d, 0xef, 0xc1, 0xd8, 0x91, 0x6f, 0x58,
	0xe0, 0xa8, 0xeb, 0x23, 0x8c, 0xe1, 0x60, 0x8b, 0x62, 0xc3, 0xc2, 0x28, 0x48, 0xf6, 0xc5, 0x4e,
	0xee, 0xf7, 0xbc, 0x2f, 0xbe, 0xf9, 0x10, 0xbe, 0x52, 0xea, 0x1c, 0x5f, 0x00, 0xbc, 0x2c, 0xcb,
	0xe1, 0x19, 0x20, 0xbf, 0xc2, 0xfc, 0xd1, 0x83, 0x61, 0xc3, 0xcd, 0x93, 0x15, 0xbd, 0x13, 0xdf,
	0xc2, 0xd7, 0x70, 0x7b, 0x87, 0x72, 0x9c, 0xf3, 0xeb, 0x24, 0xf8, 0x23, 0xc8, 0x8e, 0x1e, 0x38,
	0xe7, 0xd7, 0xe9, 0x73, 0x33, 0xc8, 0x31, 0x1c, 0x6c, 0xe9, 0x62, 0xf0, 0x2e, 0xbf, 0x2a, 0xd3,
	0x8d, 0x08, 0xf9, 0xd8, 0xfd, 0x16, 0x4f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xb1, 0x58,
	0xc9, 0xe7, 0x03, 0x00, 0x00,
}
