// Code generated by protoc-gen-go.
// source: github.com/kwokhunglee/micro/server/debug/proto/debug.proto
// DO NOT EDIT!

/*
Package debug is a generated protocol buffer package.

It is generated from these files:
	github.com/kwokhunglee/micro/server/debug/proto/debug.proto

It has these top-level messages:
	HealthRequest
	HealthResponse
	StatsRequest
	StatsResponse
*/
package debug

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

type HealthRequest struct {
}

func (m *HealthRequest) Reset()                    { *m = HealthRequest{} }
func (m *HealthRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()               {}
func (*HealthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HealthResponse struct {
	// default: ok
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *HealthResponse) Reset()                    { *m = HealthResponse{} }
func (m *HealthResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()               {}
func (*HealthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StatsRequest struct {
}

func (m *StatsRequest) Reset()                    { *m = StatsRequest{} }
func (m *StatsRequest) String() string            { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()               {}
func (*StatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StatsResponse struct {
	// unix timestamp
	Started uint64 `protobuf:"varint,1,opt,name=started" json:"started,omitempty"`
	// in seconds
	Uptime uint64 `protobuf:"varint,2,opt,name=uptime" json:"uptime,omitempty"`
	// in bytes
	Memory uint64 `protobuf:"varint,3,opt,name=memory" json:"memory,omitempty"`
	// num threads
	Threads uint64 `protobuf:"varint,4,opt,name=threads" json:"threads,omitempty"`
	// in nanoseconds
	Gc uint64 `protobuf:"varint,5,opt,name=gc" json:"gc,omitempty"`
}

func (m *StatsResponse) Reset()                    { *m = StatsResponse{} }
func (m *StatsResponse) String() string            { return proto.CompactTextString(m) }
func (*StatsResponse) ProtoMessage()               {}
func (*StatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*HealthRequest)(nil), "HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "HealthResponse")
	proto.RegisterType((*StatsRequest)(nil), "StatsRequest")
	proto.RegisterType((*StatsResponse)(nil), "StatsResponse")
}

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x34, 0x8f, 0xbd, 0x6e, 0x83, 0x30,
	0x14, 0x85, 0x05, 0xa5, 0x54, 0xbd, 0x2a, 0x54, 0x62, 0xa8, 0x3c, 0x56, 0x4c, 0x2c, 0xc5, 0x43,
	0x97, 0x3e, 0x42, 0x67, 0xf2, 0x04, 0xfc, 0x5c, 0x19, 0xa4, 0x38, 0x26, 0xbe, 0xd7, 0x91, 0x32,
	0xe7, 0xc5, 0x03, 0xb6, 0xd9, 0xce, 0xf7, 0xd9, 0xe7, 0x48, 0x17, 0xfe, 0xd4, 0xc2, 0xb3, 0x1b,
	0xda, 0xd1, 0x68, 0xa9, 0x97, 0xd1, 0x1a, 0xa9, 0xcc, 0x4f, 0x08, 0x84, 0xf6, 0x86, 0x56, 0x4e,
	0x38, 0x38, 0x25, 0x57, 0x6b, 0xd8, 0x84, 0xdc, 0xfa, 0x5c, 0x7f, 0x42, 0xf1, 0x8f, 0xfd, 0x99,
	0xe7, 0x0e, 0xaf, 0x0e, 0x89, 0xeb, 0x06, 0xca, 0x43, 0xd0, 0x6a, 0x2e, 0x84, 0xd5, 0x17, 0xe4,
	0xc4, 0x3d, 0x3b, 0x12, 0xc9, 0x77, 0xd2, 0xbc, 0x77, 0x91, 0xea, 0x12, 0x3e, 0x4e, 0x5b, 0xa2,
	0xa3, 0xf9, 0x48, 0xa0, 0x88, 0x22, 0x36, 0x05, 0xbc, 0x6d, 0x7f, 0x2d, 0xe3, 0xe4, 0xab, 0x59,
	0x77, 0xe0, 0xbe, 0xe9, 0x56, 0x5e, 0x34, 0x8a, 0xd4, 0x3f, 0x44, 0xda, 0xbd, 0x46, 0x6d, 0xec,
	0x5d, 0xbc, 0x04, 0x1f, 0x68, 0x5f, 0xe2, 0xd9, 0x62, 0x3f, 0x91, 0xc8, 0xc2, 0x52, 0xc4, 0xaa,
	0x84, 0x54, 0x8d, 0xe2, 0xd5, 0xcb, 0x2d, 0x0d, 0xb9, 0xbf, 0xeb, 0xf7, 0x19, 0x00, 0x00, 0xff,
	0xff, 0xc6, 0x75, 0x51, 0x35, 0x13, 0x01, 0x00, 0x00,
}
