// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enums/enum_example.proto

package enumpb

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

type DayOfTheWeek int32

const (
	DayOfTheWeek_UNKNOWN   DayOfTheWeek = 0
	DayOfTheWeek_MONDAY    DayOfTheWeek = 1
	DayOfTheWeek_TUESDAY   DayOfTheWeek = 2
	DayOfTheWeek_WEDNESDAY DayOfTheWeek = 3
	DayOfTheWeek_THURSDAY  DayOfTheWeek = 4
	DayOfTheWeek_FRIDAY    DayOfTheWeek = 5
	DayOfTheWeek_SATURDAY  DayOfTheWeek = 6
	DayOfTheWeek_SUNDAY    DayOfTheWeek = 7
)

var DayOfTheWeek_name = map[int32]string{
	0: "UNKNOWN",
	1: "MONDAY",
	2: "TUESDAY",
	3: "WEDNESDAY",
	4: "THURSDAY",
	5: "FRIDAY",
	6: "SATURDAY",
	7: "SUNDAY",
}

var DayOfTheWeek_value = map[string]int32{
	"UNKNOWN":   0,
	"MONDAY":    1,
	"TUESDAY":   2,
	"WEDNESDAY": 3,
	"THURSDAY":  4,
	"FRIDAY":    5,
	"SATURDAY":  6,
	"SUNDAY":    7,
}

func (x DayOfTheWeek) String() string {
	return proto.EnumName(DayOfTheWeek_name, int32(x))
}

func (DayOfTheWeek) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c2b4d8031bdbd27, []int{0}
}

type EnumMessage struct {
	Id                   int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DayOfTheWeek         DayOfTheWeek `protobuf:"varint,2,opt,name=day_of_the_week,json=dayOfTheWeek,proto3,enum=example.enumerations.DayOfTheWeek" json:"day_of_the_week,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EnumMessage) Reset()         { *m = EnumMessage{} }
func (m *EnumMessage) String() string { return proto.CompactTextString(m) }
func (*EnumMessage) ProtoMessage()    {}
func (*EnumMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c2b4d8031bdbd27, []int{0}
}

func (m *EnumMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumMessage.Unmarshal(m, b)
}
func (m *EnumMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumMessage.Marshal(b, m, deterministic)
}
func (m *EnumMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumMessage.Merge(m, src)
}
func (m *EnumMessage) XXX_Size() int {
	return xxx_messageInfo_EnumMessage.Size(m)
}
func (m *EnumMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EnumMessage proto.InternalMessageInfo

func (m *EnumMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EnumMessage) GetDayOfTheWeek() DayOfTheWeek {
	if m != nil {
		return m.DayOfTheWeek
	}
	return DayOfTheWeek_UNKNOWN
}

func init() {
	proto.RegisterEnum("example.enumerations.DayOfTheWeek", DayOfTheWeek_name, DayOfTheWeek_value)
	proto.RegisterType((*EnumMessage)(nil), "example.enumerations.EnumMessage")
}

func init() { proto.RegisterFile("enums/enum_example.proto", fileDescriptor_3c2b4d8031bdbd27) }

var fileDescriptor_3c2b4d8031bdbd27 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcd, 0x2b, 0xcd,
	0x2d, 0xd6, 0x07, 0x91, 0xf1, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x22, 0x30, 0x2e, 0x48, 0x2e, 0xb5, 0x28, 0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x58,
	0x29, 0x83, 0x8b, 0xdb, 0x35, 0xaf, 0x34, 0xd7, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x88,
	0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x88, 0x29, 0x33, 0x45, 0xc8,
	0x93, 0x8b, 0x3f, 0x25, 0xb1, 0x32, 0x3e, 0x3f, 0x2d, 0xbe, 0x24, 0x23, 0x35, 0xbe, 0x3c, 0x35,
	0x35, 0x5b, 0x82, 0x49, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x49, 0x0f, 0x9b, 0x71, 0x7a, 0x2e, 0x89,
	0x95, 0xfe, 0x69, 0x21, 0x19, 0xa9, 0xe1, 0xa9, 0xa9, 0xd9, 0x41, 0x3c, 0x29, 0x48, 0x3c, 0xad,
	0x72, 0x2e, 0x1e, 0x64, 0x59, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f,
	0x01, 0x06, 0x21, 0x2e, 0x2e, 0x36, 0x5f, 0x7f, 0x3f, 0x17, 0xc7, 0x48, 0x01, 0x46, 0x90, 0x44,
	0x48, 0xa8, 0x6b, 0x30, 0x88, 0xc3, 0x24, 0xc4, 0xcb, 0xc5, 0x19, 0xee, 0xea, 0xe2, 0x07, 0xe1,
	0x32, 0x0b, 0xf1, 0x70, 0x71, 0x84, 0x78, 0x84, 0x06, 0x81, 0x79, 0x2c, 0x20, 0x5d, 0x6e, 0x41,
	0x9e, 0x20, 0x36, 0x2b, 0x48, 0x26, 0xd8, 0x31, 0x24, 0x34, 0x08, 0xc4, 0x63, 0x03, 0xc9, 0x04,
	0x87, 0x82, 0xcd, 0x63, 0x77, 0xe2, 0x88, 0x62, 0x03, 0xb9, 0xb1, 0x20, 0x29, 0x89, 0x0d, 0x1c,
	0x12, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x33, 0xaf, 0xb6, 0x25, 0x01, 0x00, 0x00,
}