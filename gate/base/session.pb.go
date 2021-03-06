// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session/session.proto

/*
Package gate is a generated protocol buffer package.

It is generated from these files:
	session/session.proto

It has these top-level messages:
	Session
*/
package basegate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import (
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type session struct {
	IP        string            `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	Network   string            `protobuf:"bytes,2,opt,name=Network" json:"Network,omitempty"`
	Userid    string            `protobuf:"bytes,3,opt,name=Userid" json:"Userid,omitempty"`
	Sessionid string            `protobuf:"bytes,4,opt,name=Sessionid" json:"Sessionid,omitempty"`
	Serverid  string            `protobuf:"bytes,5,opt,name=Serverid" json:"Serverid,omitempty"`
	Settings  map[string]string `protobuf:"bytes,6,rep,name=Settings" json:"Settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Carrier   map[string]string `protobuf:"bytes,7,rep,name=Carrier" json:"Carrier,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *session) Reset()                    { *m = session{} }
func (m *session) String() string            { return proto.CompactTextString(m) }
func (*session) ProtoMessage()               {}
func (*session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *session) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *session) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *session) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *session) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

func (m *session) GetServerid() string {
	if m != nil {
		return m.Serverid
	}
	return ""
}

func (m *session) GetSettings() map[string]string {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *session) GetCarrier() map[string]string {
	if m != nil {
		return m.Carrier
	}
	return nil
}

func init() {
	proto.RegisterType((*session)(nil), "gate.Session")
}

func init() { proto.RegisterFile("session/session.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0xe9, 0x89,
	0x25, 0xa9, 0x4a, 0x2f, 0x99, 0xb8, 0xd8, 0x83, 0x21, 0xe2, 0x42, 0x7c, 0x5c, 0x4c, 0x9e, 0x01,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x9e, 0x01, 0x42, 0x12, 0x5c, 0xec, 0x7e, 0xa9,
	0x25, 0xe5, 0xf9, 0x45, 0xd9, 0x12, 0x4c, 0x60, 0x41, 0x18, 0x57, 0x48, 0x8c, 0x8b, 0x2d, 0xb4,
	0x38, 0xb5, 0x28, 0x33, 0x45, 0x82, 0x19, 0x2c, 0x01, 0xe5, 0x09, 0xc9, 0x70, 0x71, 0x42, 0x0d,
	0xcb, 0x4c, 0x91, 0x60, 0x01, 0x4b, 0x21, 0x04, 0x84, 0xa4, 0xb8, 0x38, 0x82, 0x53, 0x8b, 0xca,
	0xc0, 0xfa, 0x58, 0xc1, 0x92, 0x70, 0xbe, 0x90, 0x39, 0x48, 0xae, 0xa4, 0x24, 0x33, 0x2f, 0xbd,
	0x58, 0x82, 0x4d, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5a, 0x0f, 0xe4, 0x40, 0x3d, 0xa8, 0x76, 0x3d,
	0x98, 0xac, 0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0x5c, 0xb1, 0x90, 0x21, 0x17, 0x5b, 0x48, 0x51,
	0x62, 0x72, 0x6a, 0x91, 0x04, 0x3b, 0x58, 0x9b, 0x24, 0xaa, 0x36, 0x88, 0x1c, 0x44, 0x13, 0x54,
	0xa1, 0x94, 0x35, 0x17, 0x2f, 0x8a, 0x69, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x50, 0x9f,
	0x83, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x50, 0x8f, 0x43, 0x38, 0x56,
	0x4c, 0x16, 0x8c, 0x52, 0x96, 0x5c, 0xdc, 0x48, 0x66, 0x92, 0xa2, 0x35, 0x89, 0x0d, 0x1c, 0xf0,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0xd1, 0x71, 0xc8, 0x91, 0x01, 0x00, 0x00,
}
