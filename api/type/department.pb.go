// Code generated by protoc-gen-go. DO NOT EDIT.
// source: department.proto

package type_pb // import "github.com/ProgrammingLab/prolab-accounts/api/type"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Department struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShortName            string   `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Department) Reset()         { *m = Department{} }
func (m *Department) String() string { return proto.CompactTextString(m) }
func (*Department) ProtoMessage()    {}
func (*Department) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_0011adb89030f85b, []int{0}
}
func (m *Department) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Department.Unmarshal(m, b)
}
func (m *Department) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Department.Marshal(b, m, deterministic)
}
func (dst *Department) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Department.Merge(dst, src)
}
func (m *Department) XXX_Size() int {
	return xxx_messageInfo_Department.Size(m)
}
func (m *Department) XXX_DiscardUnknown() {
	xxx_messageInfo_Department.DiscardUnknown(m)
}

var xxx_messageInfo_Department proto.InternalMessageInfo

func (m *Department) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Department) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Department) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func init() {
	proto.RegisterType((*Department)(nil), "programming_lab.prolab_accounts.type.Department")
}

func init() { proto.RegisterFile("department.proto", fileDescriptor_department_0011adb89030f85b) }

var fileDescriptor_department_0011adb89030f85b = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x49, 0x2d, 0x48,
	0x2c, 0x2a, 0xc9, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x29, 0x28,
	0xca, 0x4f, 0x2f, 0x4a, 0xcc, 0xcd, 0xcd, 0xcc, 0x4b, 0x8f, 0xcf, 0x49, 0x4c, 0x02, 0x09, 0xe7,
	0x24, 0x26, 0xc5, 0x27, 0x26, 0x27, 0xe7, 0x97, 0xe6, 0x95, 0x14, 0xeb, 0x95, 0x54, 0x16, 0xa4,
	0x2a, 0xf9, 0x73, 0x71, 0xb9, 0xc0, 0x75, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0xf0, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0xb2, 0x5c, 0x5c, 0xc5, 0x19, 0xf9, 0x45,
	0x25, 0xf1, 0x60, 0x19, 0x66, 0xb0, 0x0c, 0x27, 0x58, 0xc4, 0x2f, 0x31, 0x37, 0xd5, 0xc9, 0x26,
	0xca, 0x2a, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x00, 0xe1, 0x06,
	0x9f, 0xc4, 0x24, 0x7d, 0x88, 0x13, 0x74, 0x61, 0x4e, 0xd0, 0x4f, 0x2c, 0xc8, 0xd4, 0x07, 0x39,
	0xc3, 0x1a, 0x44, 0xc4, 0x17, 0x24, 0x25, 0xb1, 0x81, 0xdd, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x51, 0x16, 0x5d, 0x46, 0xcf, 0x00, 0x00, 0x00,
}
