// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gescache.proto

package geecachepb

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

type Request struct {
	Group                string   `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c8c347a8e93050, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Request) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Response struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c8c347a8e93050, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "geecachepb.Request")
	proto.RegisterType((*Response)(nil), "geecachepb.Response")
}

func init() { proto.RegisterFile("gescache.proto", fileDescriptor_e3c8c347a8e93050) }

var fileDescriptor_e3c8c347a8e93050 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4f, 0x2d, 0x4e,
	0x4e, 0x4c, 0xce, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x4f, 0x4d, 0x05,
	0xf3, 0x0b, 0x92, 0x94, 0x0c, 0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44,
	0xb8, 0x58, 0xd3, 0x8b, 0xf2, 0x4b, 0x0b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c,
	0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x26, 0xb0, 0x18, 0x88, 0xa9, 0xa4, 0xc0, 0xc5,
	0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a, 0xd2, 0x53, 0x96, 0x98, 0x53, 0x9a, 0x0a,
	0xd6, 0xc3, 0x13, 0x04, 0xe1, 0x18, 0xd9, 0x71, 0x71, 0xb9, 0x83, 0x34, 0x3b, 0x83, 0x2c, 0x11,
	0x32, 0xe0, 0x62, 0x76, 0x4f, 0x2d, 0x11, 0x12, 0xd6, 0x43, 0x58, 0xab, 0x07, 0xb5, 0x53, 0x4a,
	0x04, 0x55, 0x10, 0x62, 0x6a, 0x12, 0x1b, 0xd8, 0x9d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x9d, 0xf6, 0xc8, 0x30, 0xb9, 0x00, 0x00, 0x00,
}
