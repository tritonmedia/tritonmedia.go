// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1.identify.proto

package api

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

type Identify struct {
	// when this was created
	CreatedAt string `protobuf:"bytes,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// media object
	Media                *Media   `protobuf:"bytes,2,opt,name=media,proto3" json:"media,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identify) Reset()         { *m = Identify{} }
func (m *Identify) String() string { return proto.CompactTextString(m) }
func (*Identify) ProtoMessage()    {}
func (*Identify) Descriptor() ([]byte, []int) {
	return fileDescriptor_d98ff4b80a59c996, []int{0}
}

func (m *Identify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identify.Unmarshal(m, b)
}
func (m *Identify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identify.Marshal(b, m, deterministic)
}
func (m *Identify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identify.Merge(m, src)
}
func (m *Identify) XXX_Size() int {
	return xxx_messageInfo_Identify.Size(m)
}
func (m *Identify) XXX_DiscardUnknown() {
	xxx_messageInfo_Identify.DiscardUnknown(m)
}

var xxx_messageInfo_Identify proto.InternalMessageInfo

func (m *Identify) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Identify) GetMedia() *Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func init() {
	proto.RegisterType((*Identify)(nil), "api.Identify")
}

func init() { proto.RegisterFile("v1.identify.proto", fileDescriptor_d98ff4b80a59c996) }

var fileDescriptor_d98ff4b80a59c996 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x33, 0xd4, 0xcb,
	0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xab, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e,
	0x2c, 0xc8, 0x94, 0xe2, 0x2b, 0x33, 0xd4, 0xcb, 0x4d, 0x4d, 0xc9, 0x4c, 0x84, 0x08, 0x2a, 0x79,
	0x71, 0x71, 0x78, 0x42, 0x95, 0x09, 0xc9, 0x70, 0x71, 0x26, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0xa6,
	0x38, 0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x84, 0x14, 0xb8, 0x58, 0xc1,
	0x1a, 0x25, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xb8, 0xf4, 0x12, 0x0b, 0x32, 0xf5, 0x7c, 0x41,
	0x22, 0x41, 0x10, 0x89, 0x24, 0x36, 0xb0, 0x91, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01,
	0xfd, 0xdf, 0x5d, 0x7c, 0x00, 0x00, 0x00,
}