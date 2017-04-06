// Code generated by protoc-gen-gogo.
// source: client/version/versionpb/version.proto
// DO NOT EDIT!

/*
Package versionpb is a generated protocol buffer package.

It is generated from these files:
	client/version/versionpb/version.proto

It has these top-level messages:
	Version
*/
package versionpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Version struct {
	Major      uint32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor      uint32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Micro      uint32 `protobuf:"varint,3,opt,name=micro,proto3" json:"micro,omitempty"`
	Additional string `protobuf:"bytes,4,opt,name=additional,proto3" json:"additional,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptorVersion, []int{0} }

func (m *Version) GetMajor() uint32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *Version) GetMinor() uint32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *Version) GetMicro() uint32 {
	if m != nil {
		return m.Micro
	}
	return 0
}

func (m *Version) GetAdditional() string {
	if m != nil {
		return m.Additional
	}
	return ""
}

func init() {
	proto.RegisterType((*Version)(nil), "versionpb.Version")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	GetVersion(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Version, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetVersion(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := grpc.Invoke(ctx, "/versionpb.API/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	GetVersion(context.Context, *google_protobuf1.Empty) (*Version, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versionpb.API/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetVersion(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "versionpb.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _API_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/version/versionpb/version.proto",
}

func init() { proto.RegisterFile("client/version/versionpb/version.proto", fileDescriptorVersion) }

var fileDescriptorVersion = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x83, 0xd1, 0x05, 0x49, 0x30, 0x96,
	0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x10, 0x27, 0x5c, 0x42, 0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d,
	0x27, 0x55, 0x3f, 0xb1, 0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f,
	0xaf, 0x18, 0xa2, 0x50, 0x4a, 0x1a, 0x2a, 0x0b, 0xe6, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16,
	0x94, 0x54, 0x42, 0x24, 0x95, 0xb2, 0xb9, 0xd8, 0xc3, 0x20, 0xe6, 0x08, 0x89, 0x70, 0xb1, 0xe6,
	0x26, 0x66, 0xe5, 0x17, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0x41, 0x38, 0x60, 0xd1, 0xcc,
	0xbc, 0xfc, 0x22, 0x09, 0x26, 0xa8, 0x28, 0x88, 0x03, 0x11, 0x4d, 0x2e, 0xca, 0x97, 0x60, 0x86,
	0x89, 0x26, 0x17, 0xe5, 0x0b, 0xc9, 0x71, 0x71, 0x25, 0xa6, 0xa4, 0x64, 0x82, 0x2c, 0x4f, 0xcc,
	0x91, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x12, 0x31, 0x0a, 0xe4, 0x62, 0x76, 0x0c, 0xf0,
	0x14, 0xf2, 0xe2, 0xe2, 0x72, 0x4f, 0x2d, 0x81, 0x59, 0x2b, 0xa6, 0x07, 0x71, 0x9f, 0x1e, 0xcc,
	0x7d, 0x7a, 0xae, 0x20, 0xf7, 0x49, 0x09, 0xe9, 0xc1, 0x3d, 0xa8, 0x07, 0x55, 0xab, 0x24, 0xd0,
	0x74, 0xf9, 0xc9, 0x64, 0x26, 0x2e, 0x21, 0x0e, 0x58, 0x58, 0x24, 0xb1, 0x81, 0x75, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x18, 0x96, 0x30, 0x36, 0x01, 0x00, 0x00,
}
