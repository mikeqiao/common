// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

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

type DataBaseConfig struct {
	UserCreatId          int64                  `protobuf:"varint,1,opt,name=userCreatId,proto3" json:"userCreatId,omitempty"`
	GroupCreatId         int64                  `protobuf:"varint,2,opt,name=groupCreatId,proto3" json:"groupCreatId,omitempty"`
	ItemCreatId          int64                  `protobuf:"varint,3,opt,name=ItemCreatId,proto3" json:"ItemCreatId,omitempty"`
	XXX_Update           map[string]interface{} `json:"-"`
	XXX_Update2          map[string]interface{} `json:"-"`
	XXX_Key              map[string]interface{} `json:"-"`
	XXX_Add              map[string]interface{} `json:"-"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DataBaseConfig) Reset()         { *m = DataBaseConfig{} }
func (m *DataBaseConfig) String() string { return proto.CompactTextString(m) }
func (*DataBaseConfig) ProtoMessage()    {}
func (*DataBaseConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *DataBaseConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBaseConfig.Unmarshal(m, b)
}
func (m *DataBaseConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBaseConfig.Marshal(b, m, deterministic)
}
func (m *DataBaseConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBaseConfig.Merge(m, src)
}
func (m *DataBaseConfig) XXX_Size() int {
	return xxx_messageInfo_DataBaseConfig.Size(m)
}
func (m *DataBaseConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBaseConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DataBaseConfig proto.InternalMessageInfo

func (m *DataBaseConfig) GetName() string {
	var buff bytes.Buffer
	buff.WriteString("DataBaseConfig")
	data := strconv.Quote(buff.String())
	return data
}
func (m *DataBaseConfig) DBDel() string {
	var buff bytes.Buffer
	buff.WriteString("DELETE FROM ")
	buff.WriteString("DataBaseConfig")
	buff.WriteString(" WHERE ")
	i := 0
	for k, v := range m.XXX_Update {
		if i > 0 {
			buff.WriteString(" AND ")
		}
		buff.WriteString(k)
		d := fmt.Sprint(v)
		buff.WriteString(" = ")
		buff.WriteString(d)
		i += 1
	}
	return buff.String()
}
func (m *DataBaseConfig) DBInsert() string {
	var buff bytes.Buffer
	buff.WriteString("INSERT INTO ")
	buff.WriteString("DataBaseConfig")
	buff.WriteString("( ")
	i := 0
	var key bytes.Buffer
	var value bytes.Buffer
	for k, v := range m.XXX_Update {
		if i > 0 {
			key.WriteString(", ")
			value.WriteString(", ")
		}
		key.WriteString(k)
		d := fmt.Sprint(v)
		value.WriteString(d)
		i += 1
	}
	buff.WriteString(key.String())
	buff.WriteString(" ) VALUES ( ")
	buff.WriteString(value.String())
	buff.WriteString(" )")
	return buff.String()
}
func (m *DataBaseConfig) DBUpdate() string {
	var buff bytes.Buffer
	buff.WriteString("UPDATE ")
	buff.WriteString("DataBaseConfig")
	buff.WriteString(" SET ")
	i := 0
	for k, v := range m.XXX_Update {
		if i > 0 {
			buff.WriteString(", ")
		}
		buff.WriteString(k)
		d := fmt.Sprint(v)
		buff.WriteString(" = ")
		buff.WriteString(d)
		i += 1
	}
	buff.WriteString(" WHERE ")
	i = 0
	for k, v := range m.XXX_Key {
		if i > 0 {
			buff.WriteString(" AND ")
		}
		buff.WriteString(k)
		d := fmt.Sprint(v)
		buff.WriteString(" = ")
		buff.WriteString(d)
		i += 1
	}
	return buff.String()
}
func (m *DataBaseConfig) DBSelect() (s string, ks []string) {
	var buff bytes.Buffer
	buff.WriteString("SELECT ")
	i := 0
	if len(m.XXX_Update) == 0 {
		ks = make([]string, 3)
		ks[0] = "UserCreatId"
		if i > 0 {
			buff.WriteString(", ")
		}
		buff.WriteString("UserCreatId")
		i += 1
		ks[1] = "GroupCreatId"
		if i > 0 {
			buff.WriteString(", ")
		}
		buff.WriteString("GroupCreatId")
		i += 1
		ks[2] = "ItemCreatId"
		if i > 0 {
			buff.WriteString(", ")
		}
		buff.WriteString("ItemCreatId")
		i += 1
	} else {
		ks = make([]string, len(m.XXX_Update))
		j := 0
		for k, _ := range m.XXX_Update {
			if i > 0 {
				buff.WriteString(", ")
			}
			buff.WriteString(k)
			ks[j] = k
			j += 1
			i += 1
		}
	}
	buff.WriteString(" FROM ")
	buff.WriteString("DataBaseConfig")
	i = 0
	for k, v := range m.XXX_Key {
		if i == 0 {
			buff.WriteString(" WHERE ")
		}
		if i > 0 {
			buff.WriteString(" AND ")
		}
		buff.WriteString(k)
		d := fmt.Sprint(v)
		buff.WriteString(" = ")
		buff.WriteString(d)
		i += 1
	}
	s = buff.String()
	return
}
func (m *DataBaseConfig) GetFieldID(ks []string) (p []interface{}) {
	if nil == m.XXX_Add {
		m.XXX_Add = make(map[string]interface{})
	}
	m.XXX_Add["UserCreatId"] = &m.UserCreatId
	m.XXX_Add["GroupCreatId"] = &m.GroupCreatId
	m.XXX_Add["ItemCreatId"] = &m.ItemCreatId
	p = make([]interface{}, len(ks))
	j := 0
	for _, v := range ks {
		if d, ok := m.XXX_Add[v]; ok && nil != d {
			p[j] = d
			j += 1
		}
	}
	return
}
func (m *DataBaseConfig) RedisSet(key interface{}) (table string, data []interface{}) {
	table = "DataBaseConfig" + "_" + fmt.Sprint(key)
	data = make([]interface{}, 2*len(m.XXX_Update2))
	j := 0
	for k, v := range m.XXX_Update2 {
		data[j] = k
		data[j+1] = v
		j += 2
	}
	return
}
func (m *DataBaseConfig) RedisGet(key interface{}) (table string, data []string) {
	table = "DataBaseConfig" + "_" + fmt.Sprint(key)
	if len(m.XXX_Update) == 0 {
		data = make([]string, 3)
		data[0] = "UserCreatId"
		data[1] = "GroupCreatId"
		data[2] = "ItemCreatId"
	} else {
		data = make([]string, len(m.XXX_Update))
		j := 0
		for k, _ := range m.XXX_Update {
			data[j] = k
			j += 1
		}
	}
	return
}
func (m *DataBaseConfig) RedisDel(key interface{}) (table string, data []string) {
	table = "DataBaseConfig" + "_" + fmt.Sprint(key)
	data = make([]string, len(m.XXX_Update))
	j := 0
	for k, _ := range m.XXX_Update {
		data[j] = k
		j += 1
	}
	return
}
func (m *DataBaseConfig) DoUpdate() (sql, table string, data []interface{}) {
	var buff bytes.Buffer
	buff.WriteString("UPDATE ")
	buff.WriteString("DataBaseConfig")
	buff.WriteString(" SET ")
	i := 0
	j := 0
	data = make([]interface{}, 2*len(m.XXX_Update))
	for k, v := range m.XXX_Update {
		if i > 0 {
			buff.WriteString(", ")
		}
		buff.WriteString(k)
		d := fmt.Sprint(v)
		buff.WriteString(" = ")
		buff.WriteString(d)
		i += 1
		data[j] = k
		if val, ok := m.XXX_Update2[k]; ok {
			data[j+1] = val
		} else {
			data[j+1] = nil
		}
		j += 2
		delete(m.XXX_Update, k)
		delete(m.XXX_Update2, k)
	}
	buff.WriteString(" WHERE ")
	i = 0
	for k, v := range m.XXX_Key {
		if i > 0 {
			buff.WriteString(" AND ")
		}
		buff.WriteString(k)
		d := fmt.Sprint(v)
		buff.WriteString(" = ")
		buff.WriteString(d)
		table = "DataBaseConfig" + "_" + d
		i += 1
	}
	sql = buff.String()
	return
}

func (m *DataBaseConfig) GetUserCreatId() int64 {
	if m != nil {
		return m.UserCreatId
	}
	return 0
}

func (m *DataBaseConfig) GetGroupCreatId() int64 {
	if m != nil {
		return m.GroupCreatId
	}
	return 0
}

func (m *DataBaseConfig) GetItemCreatId() int64 {
	if m != nil {
		return m.ItemCreatId
	}
	return 0
}

func (m *DataBaseConfig) SetUserCreatId(data int64) {
	if m != nil {
		m.UserCreatId = data
		if nil == m.XXX_Update2 {
			m.XXX_Update2 = make(map[string]interface{})
		}
		m.XXX_Update2["UserCreatId"] = data
		if nil == m.XXX_Update {
			m.XXX_Update = make(map[string]interface{})
		}
		m.XXX_Update["UserCreatId"] = data
	}
}

func (m *DataBaseConfig) SetKeyUserCreatId(data int64) {
	if m != nil {
		if nil == m.XXX_Key {
			m.XXX_Key = make(map[string]interface{})
		}
		m.XXX_Key["UserCreatId"] = data
	}
}

func (m *DataBaseConfig) SetGroupCreatId(data int64) {
	if m != nil {
		m.GroupCreatId = data
		if nil == m.XXX_Update2 {
			m.XXX_Update2 = make(map[string]interface{})
		}
		m.XXX_Update2["GroupCreatId"] = data
		if nil == m.XXX_Update {
			m.XXX_Update = make(map[string]interface{})
		}
		m.XXX_Update["GroupCreatId"] = data
	}
}

func (m *DataBaseConfig) SetKeyGroupCreatId(data int64) {
	if m != nil {
		if nil == m.XXX_Key {
			m.XXX_Key = make(map[string]interface{})
		}
		m.XXX_Key["GroupCreatId"] = data
	}
}

func (m *DataBaseConfig) SetItemCreatId(data int64) {
	if m != nil {
		m.ItemCreatId = data
		if nil == m.XXX_Update2 {
			m.XXX_Update2 = make(map[string]interface{})
		}
		m.XXX_Update2["ItemCreatId"] = data
		if nil == m.XXX_Update {
			m.XXX_Update = make(map[string]interface{})
		}
		m.XXX_Update["ItemCreatId"] = data
	}
}

func (m *DataBaseConfig) SetKeyItemCreatId(data int64) {
	if m != nil {
		if nil == m.XXX_Key {
			m.XXX_Key = make(map[string]interface{})
		}
		m.XXX_Key["ItemCreatId"] = data
	}
}

func init() {
	proto.RegisterType((*DataBaseConfig)(nil), "msg.DataBaseConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0xaa, 0xe0,
	0xe2, 0x73, 0x49, 0x2c, 0x49, 0x74, 0x4a, 0x2c, 0x4e, 0x75, 0x06, 0x4b, 0x0a, 0x29, 0x70, 0x71,
	0x97, 0x16, 0xa7, 0x16, 0x39, 0x17, 0xa5, 0x26, 0x96, 0x78, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0x30, 0x07, 0x21, 0x0b, 0x09, 0x29, 0x71, 0xf1, 0xa4, 0x17, 0xe5, 0x97, 0x16, 0xc0, 0x94, 0x30,
	0x81, 0x95, 0xa0, 0x88, 0x81, 0x4c, 0xf1, 0x2c, 0x49, 0xcd, 0x85, 0x29, 0x61, 0x86, 0x98, 0x82,
	0x24, 0x94, 0xc4, 0x06, 0x76, 0x85, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x52, 0xa8, 0xa9,
	0x95, 0x00, 0x00, 0x00,
}
