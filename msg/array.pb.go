// Code generated by protoc-gen-go. DO NOT EDIT.
// source: array.proto

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

type ArrayDataUserBag struct {
	Data                 []*DataUserBag `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ArrayDataUserBag) Reset()         { *m = ArrayDataUserBag{} }
func (m *ArrayDataUserBag) String() string { return proto.CompactTextString(m) }
func (*ArrayDataUserBag) ProtoMessage()    {}
func (*ArrayDataUserBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bd08b6c373ffa45, []int{0}
}

func (m *ArrayDataUserBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayDataUserBag.Unmarshal(m, b)
}
func (m *ArrayDataUserBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayDataUserBag.Marshal(b, m, deterministic)
}
func (m *ArrayDataUserBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayDataUserBag.Merge(m, src)
}
func (m *ArrayDataUserBag) XXX_Size() int {
	return xxx_messageInfo_ArrayDataUserBag.Size(m)
}
func (m *ArrayDataUserBag) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayDataUserBag.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayDataUserBag proto.InternalMessageInfo

func (m *ArrayDataUserBag) GetName() string {
	var buff bytes.Buffer
	buff.WriteString("ArrayDataUserBag")
	data := strconv.Quote(buff.String())
	return data
}

func (m *ArrayDataUserBag) GetData() []*DataUserBag {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ArrayDataUserBag)(nil), "msg.ArrayDataUserBag")
}

func init() { proto.RegisterFile("array.proto", fileDescriptor_6bd08b6c373ffa45) }

var fileDescriptor_6bd08b6c373ffa45 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0x2a, 0x4a,
	0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0xe2, 0x4a, 0x4a,
	0x2c, 0x4e, 0x85, 0x08, 0x28, 0x59, 0x70, 0x09, 0x38, 0x82, 0xe4, 0x5d, 0x12, 0x4b, 0x12, 0x43,
	0x8b, 0x53, 0x8b, 0x9c, 0x12, 0xd3, 0x85, 0x54, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x18,
	0x15, 0x98, 0x35, 0xb8, 0x8d, 0x04, 0xf4, 0x72, 0x8b, 0xd3, 0xf5, 0x90, 0xe4, 0x83, 0xc0, 0xb2,
	0x49, 0x6c, 0x60, 0x03, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x76, 0x62, 0x96, 0x60,
	0x00, 0x00, 0x00,
}
