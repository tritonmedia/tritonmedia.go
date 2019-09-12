// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v2.convert.proto

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

type ConvertV2Media struct {
	// key is the object to download out of S3 for this operation
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConvertV2Media) Reset()         { *m = ConvertV2Media{} }
func (m *ConvertV2Media) String() string { return proto.CompactTextString(m) }
func (*ConvertV2Media) ProtoMessage()    {}
func (*ConvertV2Media) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9415610f56b8116, []int{0}
}

func (m *ConvertV2Media) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertV2Media.Unmarshal(m, b)
}
func (m *ConvertV2Media) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertV2Media.Marshal(b, m, deterministic)
}
func (m *ConvertV2Media) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertV2Media.Merge(m, src)
}
func (m *ConvertV2Media) XXX_Size() int {
	return xxx_messageInfo_ConvertV2Media.Size(m)
}
func (m *ConvertV2Media) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertV2Media.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertV2Media proto.InternalMessageInfo

func (m *ConvertV2Media) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ConvertV2 struct {
	// when this job was created
	CreatedAt string `protobuf:"bytes,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// id of this convert operation
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// parent is the media object that this convert operation is related too
	Parent *Media `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	// media is the convertmedia operation details
	Media                *ConvertV2Media `protobuf:"bytes,4,opt,name=media,proto3" json:"media,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ConvertV2) Reset()         { *m = ConvertV2{} }
func (m *ConvertV2) String() string { return proto.CompactTextString(m) }
func (*ConvertV2) ProtoMessage()    {}
func (*ConvertV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9415610f56b8116, []int{1}
}

func (m *ConvertV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertV2.Unmarshal(m, b)
}
func (m *ConvertV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertV2.Marshal(b, m, deterministic)
}
func (m *ConvertV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertV2.Merge(m, src)
}
func (m *ConvertV2) XXX_Size() int {
	return xxx_messageInfo_ConvertV2.Size(m)
}
func (m *ConvertV2) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertV2.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertV2 proto.InternalMessageInfo

func (m *ConvertV2) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ConvertV2) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConvertV2) GetParent() *Media {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *ConvertV2) GetMedia() *ConvertV2Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func init() {
	proto.RegisterType((*ConvertV2Media)(nil), "api.ConvertV2Media")
	proto.RegisterType((*ConvertV2)(nil), "api.ConvertV2")
}

func init() { proto.RegisterFile("v2.convert.proto", fileDescriptor_f9415610f56b8116) }

var fileDescriptor_f9415610f56b8116 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x33, 0xd2, 0x4b,
	0xce, 0xcf, 0x2b, 0x4b, 0x2d, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c,
	0xc8, 0x94, 0xe2, 0x2b, 0x33, 0xd4, 0xcb, 0x4d, 0x4d, 0xc9, 0x4c, 0x84, 0x08, 0x2a, 0x29, 0x71,
	0xf1, 0x39, 0x43, 0x54, 0x85, 0x19, 0xf9, 0x82, 0xc4, 0x85, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0xa5, 0x0e, 0x46, 0x2e, 0x4e, 0xb8, 0x22,
	0x21, 0x19, 0x2e, 0xce, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xd4, 0x14, 0xc7, 0x12, 0xa8, 0x2a, 0x84,
	0x80, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0x13, 0x58, 0x98, 0x29, 0x33, 0x45, 0x48, 0x89,
	0x8b, 0xad, 0x20, 0xb1, 0x28, 0x35, 0xaf, 0x44, 0x82, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x4b,
	0x2f, 0xb1, 0x20, 0x53, 0x0f, 0x6c, 0x53, 0x10, 0x54, 0x46, 0x48, 0x93, 0x8b, 0x15, 0xec, 0x24,
	0x09, 0x16, 0xb0, 0x12, 0x61, 0xb0, 0x12, 0x54, 0x57, 0x05, 0x41, 0x54, 0x24, 0xb1, 0x81, 0x5d,
	0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x78, 0x86, 0x03, 0xe3, 0xde, 0x00, 0x00, 0x00,
}