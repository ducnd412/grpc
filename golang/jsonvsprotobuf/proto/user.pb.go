// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.ProtoPackageIsVersion1

type User struct {
	FirstName  string          `protobuf:"bytes,1,opt,name=firstName" json:"firstName,omitempty"`
	LastName   string          `protobuf:"bytes,2,opt,name=lastName" json:"lastName,omitempty"`
	BirthDay   int64           `protobuf:"varint,3,opt,name=birthDay" json:"birthDay,omitempty"`
	Phone      string          `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	Gender     int32           `protobuf:"varint,5,opt,name=gender" json:"gender,omitempty"`
	IsEmployed bool            `protobuf:"varint,6,opt,name=isEmployed" json:"isEmployed,omitempty"`
	Salary     float64         `protobuf:"fixed64,7,opt,name=salary" json:"salary,omitempty"`
	Bio        string          `protobuf:"bytes,8,opt,name=bio" json:"bio,omitempty"`
	WebSites   []*User_WebSite `protobuf:"bytes,9,rep,name=webSites" json:"webSites,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetWebSites() []*User_WebSite {
	if m != nil {
		return m.WebSites
	}
	return nil
}

type User_WebSite struct {
	Url      string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Snippets []string `protobuf:"bytes,3,rep,name=snippets" json:"snippets,omitempty"`
}

func (m *User_WebSite) Reset()                    { *m = User_WebSite{} }
func (m *User_WebSite) String() string            { return proto1.CompactTextString(m) }
func (*User_WebSite) ProtoMessage()               {}
func (*User_WebSite) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func init() {
	proto1.RegisterType((*User)(nil), "proto.User")
	proto1.RegisterType((*User_WebSite)(nil), "proto.User.WebSite")
}

var fileDescriptor0 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xc9, 0x66, 0xdb, 0x6d, 0xc7, 0x8b, 0x44, 0x91, 0xb0, 0x88, 0x04, 0x4f, 0x39, 0x55,
	0xd0, 0x57, 0xd0, 0xa3, 0x1e, 0x22, 0xe2, 0xb9, 0x65, 0x47, 0x37, 0x90, 0x6d, 0x42, 0x92, 0x45,
	0xfa, 0x9e, 0x3e, 0x90, 0x24, 0xcd, 0xd6, 0x3d, 0x65, 0xbe, 0xff, 0xcf, 0xcc, 0xf0, 0x0f, 0xc0,
	0x31, 0xa0, 0xef, 0x9c, 0xb7, 0xd1, 0xb2, 0x2a, 0x3f, 0xf7, 0xbf, 0x2b, 0x58, 0x7f, 0x04, 0xf4,
	0xec, 0x16, 0xda, 0x2f, 0xed, 0x43, 0x7c, 0xeb, 0x0f, 0xc8, 0x89, 0x20, 0xb2, 0x55, 0xff, 0x02,
	0xdb, 0x42, 0x63, 0xfa, 0x62, 0xae, 0xb2, 0xb9, 0x70, 0xf2, 0x06, 0xed, 0xe3, 0xfe, 0xb9, 0x9f,
	0x38, 0x15, 0x44, 0x52, 0xb5, 0x30, 0xbb, 0x86, 0xca, 0xed, 0xed, 0x88, 0x7c, 0x9d, 0x9b, 0x66,
	0x60, 0x37, 0x50, 0x7f, 0xe3, 0xb8, 0x43, 0xcf, 0x2b, 0x41, 0x64, 0xa5, 0x0a, 0xb1, 0x3b, 0x00,
	0x1d, 0x5e, 0x0e, 0xce, 0xd8, 0x09, 0x77, 0xbc, 0x16, 0x44, 0x36, 0xea, 0x4c, 0x49, 0x7d, 0xa1,
	0x37, 0xbd, 0x9f, 0xf8, 0x46, 0x10, 0x49, 0x54, 0x21, 0x76, 0x09, 0x74, 0xd0, 0x96, 0x37, 0x79,
	0x47, 0x2a, 0xd9, 0x03, 0x34, 0x3f, 0x38, 0xbc, 0xeb, 0x88, 0x81, 0xb7, 0x82, 0xca, 0x8b, 0xc7,
	0xab, 0x39, 0x77, 0x97, 0xc2, 0x76, 0x9f, 0xb3, 0xa7, 0x96, 0x4f, 0xdb, 0x57, 0xd8, 0x14, 0x31,
	0x4d, 0x3b, 0x7a, 0x53, 0x6e, 0x90, 0xca, 0x94, 0x22, 0xea, 0x68, 0x4e, 0xd1, 0x67, 0x48, 0xb9,
	0xc3, 0xa8, 0x9d, 0xc3, 0x18, 0x38, 0x15, 0x34, 0xdd, 0xe4, 0xc4, 0x43, 0x9d, 0x97, 0x3d, 0xfd,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x93, 0xe9, 0x43, 0x72, 0x01, 0x00, 0x00,
}
