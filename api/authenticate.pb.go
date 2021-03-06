// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authenticate.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AuthenticateLoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateLoginRequest) Reset()         { *m = AuthenticateLoginRequest{} }
func (m *AuthenticateLoginRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateLoginRequest) ProtoMessage()    {}
func (*AuthenticateLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_073f4e8cf2d4d6a6, []int{0}
}

func (m *AuthenticateLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateLoginRequest.Unmarshal(m, b)
}
func (m *AuthenticateLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateLoginRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticateLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateLoginRequest.Merge(m, src)
}
func (m *AuthenticateLoginRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateLoginRequest.Size(m)
}
func (m *AuthenticateLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateLoginRequest proto.InternalMessageInfo

func (m *AuthenticateLoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthenticateLoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticateLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticateLoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Duration             int32    `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateLoginResponse) Reset()         { *m = AuthenticateLoginResponse{} }
func (m *AuthenticateLoginResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateLoginResponse) ProtoMessage()    {}
func (*AuthenticateLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_073f4e8cf2d4d6a6, []int{1}
}

func (m *AuthenticateLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateLoginResponse.Unmarshal(m, b)
}
func (m *AuthenticateLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateLoginResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateLoginResponse.Merge(m, src)
}
func (m *AuthenticateLoginResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateLoginResponse.Size(m)
}
func (m *AuthenticateLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateLoginResponse proto.InternalMessageInfo

func (m *AuthenticateLoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthenticateLoginResponse) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type AuthenticateRegisterRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRegisterRequest) Reset()         { *m = AuthenticateRegisterRequest{} }
func (m *AuthenticateRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRegisterRequest) ProtoMessage()    {}
func (*AuthenticateRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_073f4e8cf2d4d6a6, []int{2}
}

func (m *AuthenticateRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRegisterRequest.Unmarshal(m, b)
}
func (m *AuthenticateRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRegisterRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticateRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRegisterRequest.Merge(m, src)
}
func (m *AuthenticateRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRegisterRequest.Size(m)
}
func (m *AuthenticateRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRegisterRequest proto.InternalMessageInfo

func (m *AuthenticateRegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthenticateRegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticateRegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticateCheckRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateCheckRequest) Reset()         { *m = AuthenticateCheckRequest{} }
func (m *AuthenticateCheckRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateCheckRequest) ProtoMessage()    {}
func (*AuthenticateCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_073f4e8cf2d4d6a6, []int{3}
}

func (m *AuthenticateCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateCheckRequest.Unmarshal(m, b)
}
func (m *AuthenticateCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateCheckRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticateCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateCheckRequest.Merge(m, src)
}
func (m *AuthenticateCheckRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateCheckRequest.Size(m)
}
func (m *AuthenticateCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateCheckRequest proto.InternalMessageInfo

func (m *AuthenticateCheckRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthenticateCheckResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateCheckResponse) Reset()         { *m = AuthenticateCheckResponse{} }
func (m *AuthenticateCheckResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateCheckResponse) ProtoMessage()    {}
func (*AuthenticateCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_073f4e8cf2d4d6a6, []int{4}
}

func (m *AuthenticateCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateCheckResponse.Unmarshal(m, b)
}
func (m *AuthenticateCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateCheckResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateCheckResponse.Merge(m, src)
}
func (m *AuthenticateCheckResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateCheckResponse.Size(m)
}
func (m *AuthenticateCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateCheckResponse proto.InternalMessageInfo

func (m *AuthenticateCheckResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AuthenticateCheckResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthenticateCheckResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthenticateLoginRequest)(nil), "api.AuthenticateLoginRequest")
	proto.RegisterType((*AuthenticateLoginResponse)(nil), "api.AuthenticateLoginResponse")
	proto.RegisterType((*AuthenticateRegisterRequest)(nil), "api.AuthenticateRegisterRequest")
	proto.RegisterType((*AuthenticateCheckRequest)(nil), "api.AuthenticateCheckRequest")
	proto.RegisterType((*AuthenticateCheckResponse)(nil), "api.AuthenticateCheckResponse")
}

func init() {
	proto.RegisterFile("authenticate.proto", fileDescriptor_073f4e8cf2d4d6a6)
}

var fileDescriptor_073f4e8cf2d4d6a6 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x4f, 0x4b, 0x85, 0x40,
	0x14, 0xc5, 0x51, 0x31, 0x5e, 0x97, 0x08, 0x1a, 0x5a, 0x98, 0x51, 0x3c, 0x5c, 0xb5, 0x92, 0xa8,
	0x0f, 0x10, 0x11, 0xbc, 0x55, 0x6d, 0x5c, 0xb4, 0x6b, 0x31, 0xe9, 0xe5, 0x79, 0x31, 0x67, 0x6c,
	0x66, 0xa4, 0x8f, 0xdb, 0x57, 0x09, 0xc7, 0x3f, 0x8d, 0x3d, 0x5d, 0xb6, 0x3c, 0xde, 0xeb, 0x39,
	0xbf, 0x7b, 0x14, 0x18, 0x6f, 0x4d, 0x89, 0xc2, 0x50, 0xce, 0x0d, 0xa6, 0x8d, 0x92, 0x46, 0xb2,
	0x80, 0x37, 0x14, 0x43, 0x4d, 0x3a, 0xef, 0x1f, 0x24, 0x25, 0x44, 0x8f, 0xce, 0xda, 0xb3, 0xdc,
	0x93, 0xc8, 0xf0, 0xb3, 0x45, 0x6d, 0xd8, 0x39, 0x84, 0x58, 0x73, 0xfa, 0x88, 0xbc, 0xad, 0x77,
	0x73, 0x9c, 0xf5, 0x82, 0xc5, 0xb0, 0x69, 0x35, 0x2a, 0xc1, 0x6b, 0x8c, 0x7c, 0x3b, 0x98, 0x74,
	0x37, 0x6b, 0xb8, 0xd6, 0x5f, 0x52, 0x15, 0x51, 0xd0, 0xcf, 0x46, 0x9d, 0xbc, 0xc0, 0xc5, 0x42,
	0x92, 0x6e, 0xa4, 0xd0, 0xd8, 0x45, 0x19, 0x59, 0xa1, 0x18, 0xa3, 0xac, 0xe8, 0xec, 0x8a, 0x56,
	0x71, 0x43, 0x52, 0xd8, 0xa8, 0x30, 0x9b, 0x74, 0x52, 0xc1, 0xa5, 0x6b, 0x97, 0xe1, 0x9e, 0xb4,
	0x41, 0xf5, 0x3f, 0xec, 0xb7, 0xf3, 0x96, 0x9e, 0x4a, 0xcc, 0x2b, 0x27, 0xe9, 0x10, 0x3d, 0x79,
	0x9b, 0x5f, 0x3b, 0xbc, 0x31, 0x5c, 0x7b, 0x0a, 0x3e, 0x15, 0xc3, 0xbe, 0x4f, 0xc5, 0x2f, 0xac,
	0xbf, 0x06, 0x1b, 0xcc, 0x61, 0xef, 0xbe, 0x3d, 0x38, 0x71, 0xfd, 0xd9, 0x0e, 0x42, 0xdb, 0x28,
	0xbb, 0x4a, 0x79, 0x43, 0xe9, 0xda, 0x37, 0x8d, 0xaf, 0xd7, 0xc6, 0x03, 0xda, 0x03, 0x6c, 0xc6,
	0x2a, 0xd9, 0xf6, 0x60, 0xf7, 0x4f, 0xcb, 0xf1, 0x99, 0xdd, 0x78, 0x95, 0x54, 0x4c, 0x06, 0x3b,
	0x08, 0xed, 0xb1, 0x0b, 0x20, 0x6e, 0x6d, 0x0b, 0x20, 0xb3, 0x8e, 0xde, 0x8f, 0xec, 0xff, 0x79,
	0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x57, 0x75, 0x2a, 0x3f, 0xc6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticateClient is the client API for Authenticate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticateClient interface {
	Login(ctx context.Context, in *AuthenticateLoginRequest, opts ...grpc.CallOption) (*AuthenticateLoginResponse, error)
	Register(ctx context.Context, in *AuthenticateRegisterRequest, opts ...grpc.CallOption) (*VoidResponse, error)
	Check(ctx context.Context, in *AuthenticateCheckRequest, opts ...grpc.CallOption) (*AuthenticateCheckResponse, error)
}

type authenticateClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticateClient(cc grpc.ClientConnInterface) AuthenticateClient {
	return &authenticateClient{cc}
}

func (c *authenticateClient) Login(ctx context.Context, in *AuthenticateLoginRequest, opts ...grpc.CallOption) (*AuthenticateLoginResponse, error) {
	out := new(AuthenticateLoginResponse)
	err := c.cc.Invoke(ctx, "/api.Authenticate/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateClient) Register(ctx context.Context, in *AuthenticateRegisterRequest, opts ...grpc.CallOption) (*VoidResponse, error) {
	out := new(VoidResponse)
	err := c.cc.Invoke(ctx, "/api.Authenticate/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateClient) Check(ctx context.Context, in *AuthenticateCheckRequest, opts ...grpc.CallOption) (*AuthenticateCheckResponse, error) {
	out := new(AuthenticateCheckResponse)
	err := c.cc.Invoke(ctx, "/api.Authenticate/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticateServer is the server API for Authenticate service.
type AuthenticateServer interface {
	Login(context.Context, *AuthenticateLoginRequest) (*AuthenticateLoginResponse, error)
	Register(context.Context, *AuthenticateRegisterRequest) (*VoidResponse, error)
	Check(context.Context, *AuthenticateCheckRequest) (*AuthenticateCheckResponse, error)
}

// UnimplementedAuthenticateServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticateServer struct {
}

func (*UnimplementedAuthenticateServer) Login(ctx context.Context, req *AuthenticateLoginRequest) (*AuthenticateLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthenticateServer) Register(ctx context.Context, req *AuthenticateRegisterRequest) (*VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedAuthenticateServer) Check(ctx context.Context, req *AuthenticateCheckRequest) (*AuthenticateCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterAuthenticateServer(s *grpc.Server, srv AuthenticateServer) {
	s.RegisterService(&_Authenticate_serviceDesc, srv)
}

func _Authenticate_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Authenticate/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).Login(ctx, req.(*AuthenticateLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticate_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Authenticate/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).Register(ctx, req.(*AuthenticateRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticate_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Authenticate/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).Check(ctx, req.(*AuthenticateCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authenticate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Authenticate",
	HandlerType: (*AuthenticateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authenticate_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Authenticate_Register_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Authenticate_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authenticate.proto",
}
