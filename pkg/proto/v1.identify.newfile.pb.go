// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1.identify.newfile.proto

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

type IdentifyNewFile struct {
	// when this was created
	CreatedAt string `protobuf:"bytes,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// media object
	Media *Media `protobuf:"bytes,2,opt,name=media,proto3" json:"media,omitempty"`
	// key of the file
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// quality of the file
	Quality string `protobuf:"bytes,4,opt,name=quality,proto3" json:"quality,omitempty"`
	// episode of this media,
	Episode              int64    `protobuf:"varint,5,opt,name=episode,proto3" json:"episode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentifyNewFile) Reset()         { *m = IdentifyNewFile{} }
func (m *IdentifyNewFile) String() string { return proto.CompactTextString(m) }
func (*IdentifyNewFile) ProtoMessage()    {}
func (*IdentifyNewFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_f727e4443d5a0d91, []int{0}
}

func (m *IdentifyNewFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentifyNewFile.Unmarshal(m, b)
}
func (m *IdentifyNewFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentifyNewFile.Marshal(b, m, deterministic)
}
func (m *IdentifyNewFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifyNewFile.Merge(m, src)
}
func (m *IdentifyNewFile) XXX_Size() int {
	return xxx_messageInfo_IdentifyNewFile.Size(m)
}
func (m *IdentifyNewFile) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifyNewFile.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifyNewFile proto.InternalMessageInfo

func (m *IdentifyNewFile) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *IdentifyNewFile) GetMedia() *Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func (m *IdentifyNewFile) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *IdentifyNewFile) GetQuality() string {
	if m != nil {
		return m.Quality
	}
	return ""
}

func (m *IdentifyNewFile) GetEpisode() int64 {
	if m != nil {
		return m.Episode
	}
	return 0
}

func init() {
	proto.RegisterType((*IdentifyNewFile)(nil), "api.IdentifyNewFile")
}

func init() { proto.RegisterFile("v1.identify.newfile.proto", fileDescriptor_f727e4443d5a0d91) }

var fileDescriptor_f727e4443d5a0d91 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x33, 0xd4, 0xcb,
	0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xab, 0xd4, 0xcb, 0x4b, 0x2d, 0x4f, 0xcb, 0xcc, 0x49, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0xe2, 0x2b, 0x33, 0xd4, 0xcb,
	0x4d, 0x4d, 0xc9, 0x4c, 0x84, 0x08, 0x2a, 0x4d, 0x67, 0xe4, 0xe2, 0xf7, 0x84, 0xaa, 0xf7, 0x4b,
	0x2d, 0x77, 0xcb, 0xcc, 0x49, 0x15, 0x92, 0xe1, 0xe2, 0x4c, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x4d,
	0x71, 0x2c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0x29, 0x70, 0xb1, 0x82,
	0x0d, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd2, 0x4b, 0x2c, 0xc8, 0xd4, 0xf3, 0x05,
	0x89, 0x04, 0x41, 0x24, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x98, 0xc1, 0x3a, 0x41,
	0x4c, 0x21, 0x09, 0x2e, 0xf6, 0xc2, 0xd2, 0xc4, 0x9c, 0xcc, 0x92, 0x4a, 0x09, 0x16, 0xb0, 0x28,
	0x8c, 0x0b, 0x92, 0x49, 0x2d, 0xc8, 0x2c, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x55, 0x60, 0xd4, 0x60,
	0x0e, 0x82, 0x71, 0x93, 0xd8, 0xc0, 0x0e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xf7,
	0xa4, 0x9c, 0xd2, 0x00, 0x00, 0x00,
}
