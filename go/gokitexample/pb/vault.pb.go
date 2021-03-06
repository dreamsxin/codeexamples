// Code generated by protoc-gen-go.
// source: gokitexample.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	gokitexample.proto

It has these top-level messages:
	HashRequest
	HashResponse
	ValidateRequest
	ValidateResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HashRequest struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
}

func (m *HashRequest) Reset()                    { *m = HashRequest{} }
func (m *HashRequest) String() string            { return proto.CompactTextString(m) }
func (*HashRequest) ProtoMessage()               {}
func (*HashRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HashResponse struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *HashResponse) Reset()                    { *m = HashResponse{} }
func (m *HashResponse) String() string            { return proto.CompactTextString(m) }
func (*HashResponse) ProtoMessage()               {}
func (*HashResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ValidateRequest struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Hash     string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
}

func (m *ValidateRequest) Reset()                    { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string            { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()               {}
func (*ValidateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ValidateResponse struct {
	Valid bool `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
}

func (m *ValidateResponse) Reset()                    { *m = ValidateResponse{} }
func (m *ValidateResponse) String() string            { return proto.CompactTextString(m) }
func (*ValidateResponse) ProtoMessage()               {}
func (*ValidateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*HashRequest)(nil), "pb.HashRequest")
	proto.RegisterType((*HashResponse)(nil), "pb.HashResponse")
	proto.RegisterType((*ValidateRequest)(nil), "pb.ValidateRequest")
	proto.RegisterType((*ValidateResponse)(nil), "pb.ValidateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Vault service

type VaultClient interface {
	Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type vaultClient struct {
	cc *grpc.ClientConn
}

func NewVaultClient(cc *grpc.ClientConn) VaultClient {
	return &vaultClient{cc}
}

func (c *vaultClient) Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := grpc.Invoke(ctx, "/pb.Vault/Hash", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vaultClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := grpc.Invoke(ctx, "/pb.Vault/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Vault service

type VaultServer interface {
	Hash(context.Context, *HashRequest) (*HashResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
}

func RegisterVaultServer(s *grpc.Server, srv VaultServer) {
	s.RegisterService(&_Vault_serviceDesc, srv)
}

func _Vault_Hash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).Hash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Vault/Hash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).Hash(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vault_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Vault/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Vault_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Vault",
	HandlerType: (*VaultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hash",
			Handler:    _Vault_Hash_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Vault_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("gokitexample.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4b, 0x2c, 0xcd,
	0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe4, 0xe2, 0xf6,
	0x48, 0x2c, 0xce, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x28, 0x48,
	0x2c, 0x2e, 0x2e, 0xcf, 0x2f, 0x4a, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95,
	0x4c, 0xb8, 0x78, 0x20, 0x4a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x32,
	0x12, 0x8b, 0x33, 0xa0, 0xea, 0xc0, 0x6c, 0x21, 0x01, 0x2e, 0xe6, 0xd4, 0xa2, 0x22, 0x09, 0x26,
	0xb0, 0x10, 0x88, 0xa9, 0xe4, 0xc8, 0xc5, 0x1f, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58, 0x92, 0x4a,
	0x84, 0x25, 0x70, 0x43, 0x99, 0x10, 0x86, 0x2a, 0x69, 0x70, 0x09, 0x20, 0x8c, 0x80, 0x5a, 0x2e,
	0xc2, 0xc5, 0x5a, 0x06, 0x12, 0x03, 0x1b, 0xc0, 0x11, 0x04, 0xe1, 0x18, 0xe5, 0x72, 0xb1, 0x86,
	0x81, 0x3c, 0x28, 0xa4, 0xcd, 0xc5, 0x02, 0x72, 0xab, 0x10, 0xbf, 0x5e, 0x41, 0x92, 0x1e, 0x92,
	0x07, 0xa5, 0x04, 0x10, 0x02, 0x10, 0x93, 0x94, 0x18, 0x84, 0xcc, 0xb9, 0x38, 0x60, 0xe6, 0x0b,
	0x09, 0x83, 0xe4, 0xd1, 0x1c, 0x2c, 0x25, 0x82, 0x2a, 0x08, 0xd3, 0x98, 0xc4, 0x06, 0x0e, 0x47,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x61, 0x09, 0xbd, 0x56, 0x01, 0x00, 0x00,
}
