// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1.convert.proto

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

type Convert struct {
	// when this job was created
	CreatedAt string `protobuf:"bytes,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// media is the media we need to convert
	Media                *Media   `protobuf:"bytes,2,opt,name=media,proto3" json:"media,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Convert) Reset()         { *m = Convert{} }
func (m *Convert) String() string { return proto.CompactTextString(m) }
func (*Convert) ProtoMessage()    {}
func (*Convert) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcc17f438f84db89, []int{0}
}

func (m *Convert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Convert.Unmarshal(m, b)
}
func (m *Convert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Convert.Marshal(b, m, deterministic)
}
func (m *Convert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Convert.Merge(m, src)
}
func (m *Convert) XXX_Size() int {
	return xxx_messageInfo_Convert.Size(m)
}
func (m *Convert) XXX_DiscardUnknown() {
	xxx_messageInfo_Convert.DiscardUnknown(m)
}

var xxx_messageInfo_Convert proto.InternalMessageInfo

func (m *Convert) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Convert) GetMedia() *Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func init() {
	proto.RegisterType((*Convert)(nil), "api.Convert")
}

func init() { proto.RegisterFile("v1.convert.proto", fileDescriptor_bcc17f438f84db89) }

var fileDescriptor_bcc17f438f84db89 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x33, 0xd4, 0x4b,
	0xce, 0xcf, 0x2b, 0x4b, 0x2d, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c,
	0xc8, 0x94, 0xe2, 0x2b, 0x33, 0xd4, 0xcb, 0x4d, 0x4d, 0xc9, 0x4c, 0x84, 0x08, 0x2a, 0x79, 0x72,
	0xb1, 0x3b, 0x43, 0x54, 0x09, 0xc9, 0x70, 0x71, 0x26, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0xa6, 0x38,
	0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x84, 0x14, 0xb8, 0x58, 0xc1, 0xfa,
	0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xb8, 0xf4, 0x12, 0x0b, 0x32, 0xf5, 0x7c, 0x41, 0x22,
	0x41, 0x10, 0x89, 0x24, 0x36, 0xb0, 0x89, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xf9,
	0x35, 0xe0, 0x7a, 0x00, 0x00, 0x00,
}
