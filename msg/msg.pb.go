// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package msg

import (
	"bytes"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	"strconv"
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

type CUserLoginRQ struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	AccountID            string   `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CUserLoginRQ) Reset()         { *m = CUserLoginRQ{} }
func (m *CUserLoginRQ) String() string { return proto.CompactTextString(m) }
func (*CUserLoginRQ) ProtoMessage()    {}
func (*CUserLoginRQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *CUserLoginRQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CUserLoginRQ.Unmarshal(m, b)
}
func (m *CUserLoginRQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CUserLoginRQ.Marshal(b, m, deterministic)
}
func (m *CUserLoginRQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CUserLoginRQ.Merge(m, src)
}
func (m *CUserLoginRQ) XXX_Size() int {
	return xxx_messageInfo_CUserLoginRQ.Size(m)
}
func (m *CUserLoginRQ) XXX_DiscardUnknown() {
	xxx_messageInfo_CUserLoginRQ.DiscardUnknown(m)
}

var xxx_messageInfo_CUserLoginRQ proto.InternalMessageInfo

func (m *CUserLoginRQ) GetName() string {
	var buff bytes.Buffer
	buff.WriteString("CUserLoginRQ")
	data := strconv.Quote(buff.String())
	return data
}

func (m *CUserLoginRQ) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CUserLoginRQ) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *CUserLoginRQ) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CUserLoginRS struct {
	Result               int32         `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	AccountID            string        `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Password             string        `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	User                 *DataUserBase `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CUserLoginRS) Reset()         { *m = CUserLoginRS{} }
func (m *CUserLoginRS) String() string { return proto.CompactTextString(m) }
func (*CUserLoginRS) ProtoMessage()    {}
func (*CUserLoginRS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *CUserLoginRS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CUserLoginRS.Unmarshal(m, b)
}
func (m *CUserLoginRS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CUserLoginRS.Marshal(b, m, deterministic)
}
func (m *CUserLoginRS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CUserLoginRS.Merge(m, src)
}
func (m *CUserLoginRS) XXX_Size() int {
	return xxx_messageInfo_CUserLoginRS.Size(m)
}
func (m *CUserLoginRS) XXX_DiscardUnknown() {
	xxx_messageInfo_CUserLoginRS.DiscardUnknown(m)
}

var xxx_messageInfo_CUserLoginRS proto.InternalMessageInfo

func (m *CUserLoginRS) GetName() string {
	var buff bytes.Buffer
	buff.WriteString("CUserLoginRS")
	data := strconv.Quote(buff.String())
	return data
}

func (m *CUserLoginRS) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *CUserLoginRS) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *CUserLoginRS) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CUserLoginRS) GetUser() *DataUserBase {
	if m != nil {
		return m.User
	}
	return nil
}

type CCreateAccountRQ struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	AccountID            string   `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CCreateAccountRQ) Reset()         { *m = CCreateAccountRQ{} }
func (m *CCreateAccountRQ) String() string { return proto.CompactTextString(m) }
func (*CCreateAccountRQ) ProtoMessage()    {}
func (*CCreateAccountRQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *CCreateAccountRQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCreateAccountRQ.Unmarshal(m, b)
}
func (m *CCreateAccountRQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCreateAccountRQ.Marshal(b, m, deterministic)
}
func (m *CCreateAccountRQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCreateAccountRQ.Merge(m, src)
}
func (m *CCreateAccountRQ) XXX_Size() int {
	return xxx_messageInfo_CCreateAccountRQ.Size(m)
}
func (m *CCreateAccountRQ) XXX_DiscardUnknown() {
	xxx_messageInfo_CCreateAccountRQ.DiscardUnknown(m)
}

var xxx_messageInfo_CCreateAccountRQ proto.InternalMessageInfo

func (m *CCreateAccountRQ) GetName() string {
	var buff bytes.Buffer
	buff.WriteString("CCreateAccountRQ")
	data := strconv.Quote(buff.String())
	return data
}

func (m *CCreateAccountRQ) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CCreateAccountRQ) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *CCreateAccountRQ) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CCreateAccountRS struct {
	Result               int32         `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	AccountID            string        `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Password             string        `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	User                 *DataUserBase `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CCreateAccountRS) Reset()         { *m = CCreateAccountRS{} }
func (m *CCreateAccountRS) String() string { return proto.CompactTextString(m) }
func (*CCreateAccountRS) ProtoMessage()    {}
func (*CCreateAccountRS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *CCreateAccountRS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCreateAccountRS.Unmarshal(m, b)
}
func (m *CCreateAccountRS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCreateAccountRS.Marshal(b, m, deterministic)
}
func (m *CCreateAccountRS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCreateAccountRS.Merge(m, src)
}
func (m *CCreateAccountRS) XXX_Size() int {
	return xxx_messageInfo_CCreateAccountRS.Size(m)
}
func (m *CCreateAccountRS) XXX_DiscardUnknown() {
	xxx_messageInfo_CCreateAccountRS.DiscardUnknown(m)
}

var xxx_messageInfo_CCreateAccountRS proto.InternalMessageInfo

func (m *CCreateAccountRS) GetName() string {
	var buff bytes.Buffer
	buff.WriteString("CCreateAccountRS")
	data := strconv.Quote(buff.String())
	return data
}

func (m *CCreateAccountRS) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *CCreateAccountRS) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *CCreateAccountRS) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CCreateAccountRS) GetUser() *DataUserBase {
	if m != nil {
		return m.User
	}
	return nil
}

type CCreateUserRQ struct {
	User                 *DataUserBase `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CCreateUserRQ) Reset()         { *m = CCreateUserRQ{} }
func (m *CCreateUserRQ) String() string { return proto.CompactTextString(m) }
func (*CCreateUserRQ) ProtoMessage()    {}
func (*CCreateUserRQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *CCreateUserRQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCreateUserRQ.Unmarshal(m, b)
}
func (m *CCreateUserRQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCreateUserRQ.Marshal(b, m, deterministic)
}
func (m *CCreateUserRQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCreateUserRQ.Merge(m, src)
}
func (m *CCreateUserRQ) XXX_Size() int {
	return xxx_messageInfo_CCreateUserRQ.Size(m)
}
func (m *CCreateUserRQ) XXX_DiscardUnknown() {
	xxx_messageInfo_CCreateUserRQ.DiscardUnknown(m)
}

var xxx_messageInfo_CCreateUserRQ proto.InternalMessageInfo

func (m *CCreateUserRQ) GetName() string {
	var buff bytes.Buffer
	buff.WriteString("CCreateUserRQ")
	data := strconv.Quote(buff.String())
	return data
}

func (m *CCreateUserRQ) GetUser() *DataUserBase {
	if m != nil {
		return m.User
	}
	return nil
}

type CCreateUserRS struct {
	Result               int32         `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	User                 *DataUserBase `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CCreateUserRS) Reset()         { *m = CCreateUserRS{} }
func (m *CCreateUserRS) String() string { return proto.CompactTextString(m) }
func (*CCreateUserRS) ProtoMessage()    {}
func (*CCreateUserRS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}

func (m *CCreateUserRS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCreateUserRS.Unmarshal(m, b)
}
func (m *CCreateUserRS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCreateUserRS.Marshal(b, m, deterministic)
}
func (m *CCreateUserRS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCreateUserRS.Merge(m, src)
}
func (m *CCreateUserRS) XXX_Size() int {
	return xxx_messageInfo_CCreateUserRS.Size(m)
}
func (m *CCreateUserRS) XXX_DiscardUnknown() {
	xxx_messageInfo_CCreateUserRS.DiscardUnknown(m)
}

var xxx_messageInfo_CCreateUserRS proto.InternalMessageInfo

func (m *CCreateUserRS) GetName() string {
	var buff bytes.Buffer
	buff.WriteString("CCreateUserRS")
	data := strconv.Quote(buff.String())
	return data
}

func (m *CCreateUserRS) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *CCreateUserRS) GetUser() *DataUserBase {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*CUserLoginRQ)(nil), "msg.CUserLoginRQ")
	proto.RegisterType((*CUserLoginRS)(nil), "msg.CUserLoginRS")
	proto.RegisterType((*CCreateAccountRQ)(nil), "msg.CCreateAccountRQ")
	proto.RegisterType((*CCreateAccountRS)(nil), "msg.CCreateAccountRS")
	proto.RegisterType((*CCreateUserRQ)(nil), "msg.CCreateUserRQ")
	proto.RegisterType((*CCreateUserRS)(nil), "msg.CCreateUserRS")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0x85, 0x08, 0x28, 0xc5, 0x70, 0xf1, 0x38, 0x87, 0x16, 0xa7, 0x16, 0xf9, 0xe4, 0xa7, 0x67, 0xe6,
	0x05, 0x05, 0x0a, 0x09, 0x71, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0,
	0x06, 0x81, 0xd9, 0x42, 0x32, 0x5c, 0x9c, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x9e, 0x2e,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x21, 0x29, 0x2e, 0x8e, 0x82, 0xc4, 0xe2,
	0xe2, 0xf2, 0xfc, 0xa2, 0x14, 0x09, 0x66, 0xb0, 0x24, 0x9c, 0xaf, 0xd4, 0xce, 0x88, 0x62, 0x7c,
	0xb0, 0x90, 0x18, 0x17, 0x5b, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0xd4, 0x02, 0x28, 0x8f, 0x7c,
	0x2b, 0x84, 0x54, 0xb9, 0x58, 0x4a, 0x8b, 0x53, 0x8b, 0x24, 0x58, 0x14, 0x18, 0x35, 0xb8, 0x8d,
	0x04, 0xf5, 0x40, 0x7e, 0x75, 0x49, 0x2c, 0x49, 0x04, 0xd9, 0xea, 0x94, 0x58, 0x9c, 0x1a, 0x04,
	0x96, 0x56, 0x4a, 0xe0, 0x12, 0x70, 0x76, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x75, 0x84, 0x18, 0x4b,
	0x75, 0xbf, 0x76, 0x33, 0x62, 0x58, 0x31, 0x80, 0xfe, 0x35, 0xe3, 0xe2, 0x85, 0x3a, 0x06, 0x24,
	0x11, 0x14, 0x08, 0xd7, 0xc7, 0x88, 0x5f, 0x9f, 0x1f, 0xaa, 0x3e, 0xdc, 0x3e, 0x80, 0x99, 0xc7,
	0x84, 0xd7, 0xbc, 0x24, 0x36, 0x70, 0x32, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x99,
	0x3f, 0x94, 0x84, 0x02, 0x00, 0x00,
}