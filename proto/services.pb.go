// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package proto

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

type LoginRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{1}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SignupRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignupRequest) Reset()         { *m = SignupRequest{} }
func (m *SignupRequest) String() string { return proto.CompactTextString(m) }
func (*SignupRequest) ProtoMessage()    {}
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{2}
}

func (m *SignupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupRequest.Unmarshal(m, b)
}
func (m *SignupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupRequest.Marshal(b, m, deterministic)
}
func (m *SignupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupRequest.Merge(m, src)
}
func (m *SignupRequest) XXX_Size() int {
	return xxx_messageInfo_SignupRequest.Size(m)
}
func (m *SignupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignupRequest proto.InternalMessageInfo

func (m *SignupRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignupRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignupRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UsernameUsedRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsernameUsedRequest) Reset()         { *m = UsernameUsedRequest{} }
func (m *UsernameUsedRequest) String() string { return proto.CompactTextString(m) }
func (*UsernameUsedRequest) ProtoMessage()    {}
func (*UsernameUsedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{3}
}

func (m *UsernameUsedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsernameUsedRequest.Unmarshal(m, b)
}
func (m *UsernameUsedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsernameUsedRequest.Marshal(b, m, deterministic)
}
func (m *UsernameUsedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsernameUsedRequest.Merge(m, src)
}
func (m *UsernameUsedRequest) XXX_Size() int {
	return xxx_messageInfo_UsernameUsedRequest.Size(m)
}
func (m *UsernameUsedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsernameUsedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsernameUsedRequest proto.InternalMessageInfo

func (m *UsernameUsedRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type EmailUsedRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailUsedRequest) Reset()         { *m = EmailUsedRequest{} }
func (m *EmailUsedRequest) String() string { return proto.CompactTextString(m) }
func (*EmailUsedRequest) ProtoMessage()    {}
func (*EmailUsedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{4}
}

func (m *EmailUsedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailUsedRequest.Unmarshal(m, b)
}
func (m *EmailUsedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailUsedRequest.Marshal(b, m, deterministic)
}
func (m *EmailUsedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailUsedRequest.Merge(m, src)
}
func (m *EmailUsedRequest) XXX_Size() int {
	return xxx_messageInfo_EmailUsedRequest.Size(m)
}
func (m *EmailUsedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailUsedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailUsedRequest proto.InternalMessageInfo

func (m *EmailUsedRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UsedResponse struct {
	Used                 bool     `protobuf:"varint,1,opt,name=Used,proto3" json:"Used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsedResponse) Reset()         { *m = UsedResponse{} }
func (m *UsedResponse) String() string { return proto.CompactTextString(m) }
func (*UsedResponse) ProtoMessage()    {}
func (*UsedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{5}
}

func (m *UsedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsedResponse.Unmarshal(m, b)
}
func (m *UsedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsedResponse.Marshal(b, m, deterministic)
}
func (m *UsedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsedResponse.Merge(m, src)
}
func (m *UsedResponse) XXX_Size() int {
	return xxx_messageInfo_UsedResponse.Size(m)
}
func (m *UsedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsedResponse proto.InternalMessageInfo

func (m *UsedResponse) GetUsed() bool {
	if m != nil {
		return m.Used
	}
	return false
}

type AuthUserRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUserRequest) Reset()         { *m = AuthUserRequest{} }
func (m *AuthUserRequest) String() string { return proto.CompactTextString(m) }
func (*AuthUserRequest) ProtoMessage()    {}
func (*AuthUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{6}
}

func (m *AuthUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthUserRequest.Unmarshal(m, b)
}
func (m *AuthUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthUserRequest.Marshal(b, m, deterministic)
}
func (m *AuthUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserRequest.Merge(m, src)
}
func (m *AuthUserRequest) XXX_Size() int {
	return xxx_messageInfo_AuthUserRequest.Size(m)
}
func (m *AuthUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserRequest proto.InternalMessageInfo

func (m *AuthUserRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthUserResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUserResponse) Reset()         { *m = AuthUserResponse{} }
func (m *AuthUserResponse) String() string { return proto.CompactTextString(m) }
func (*AuthUserResponse) ProtoMessage()    {}
func (*AuthUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{7}
}

func (m *AuthUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthUserResponse.Unmarshal(m, b)
}
func (m *AuthUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthUserResponse.Marshal(b, m, deterministic)
}
func (m *AuthUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserResponse.Merge(m, src)
}
func (m *AuthUserResponse) XXX_Size() int {
	return xxx_messageInfo_AuthUserResponse.Size(m)
}
func (m *AuthUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserResponse proto.InternalMessageInfo

func (m *AuthUserResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AuthUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthUserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "proto.LoginRequest")
	proto.RegisterType((*AuthResponse)(nil), "proto.AuthResponse")
	proto.RegisterType((*SignupRequest)(nil), "proto.SignupRequest")
	proto.RegisterType((*UsernameUsedRequest)(nil), "proto.UsernameUsedRequest")
	proto.RegisterType((*EmailUsedRequest)(nil), "proto.EmailUsedRequest")
	proto.RegisterType((*UsedResponse)(nil), "proto.UsedResponse")
	proto.RegisterType((*AuthUserRequest)(nil), "proto.AuthUserRequest")
	proto.RegisterType((*AuthUserResponse)(nil), "proto.AuthUserResponse")
}

func init() { proto.RegisterFile("services.proto", fileDescriptor_8e16ccb8c5307b32) }

var fileDescriptor_8e16ccb8c5307b32 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x4f, 0xc2, 0x40,
	0x14, 0x0c, 0x45, 0x48, 0x79, 0x56, 0x24, 0x0b, 0x11, 0xd2, 0x93, 0xd9, 0x98, 0xc8, 0x89, 0x04,
	0x39, 0x19, 0x0f, 0x6a, 0x82, 0x07, 0x12, 0x0f, 0xa6, 0xc0, 0xd1, 0x43, 0x95, 0x17, 0x6c, 0x94,
	0x6e, 0xed, 0xb6, 0xfa, 0xaf, 0xfc, 0x8d, 0x66, 0x3f, 0xba, 0xee, 0x62, 0x4d, 0x3c, 0xc1, 0xbc,
	0x9d, 0x37, 0x6f, 0x3a, 0x03, 0x5d, 0x8e, 0xf9, 0x47, 0xf2, 0x8c, 0x7c, 0x92, 0xe5, 0xac, 0x60,
	0xa4, 0x25, 0x7f, 0xe8, 0x0d, 0x04, 0xf7, 0x6c, 0x9b, 0xa4, 0x11, 0xbe, 0x97, 0xc8, 0x0b, 0x32,
	0x80, 0x96, 0xc4, 0xa3, 0xc6, 0x69, 0x63, 0xdc, 0x89, 0x14, 0x20, 0x21, 0xf8, 0x0f, 0x31, 0xe7,
	0x9f, 0x2c, 0xdf, 0x8c, 0x3c, 0xf9, 0x60, 0x30, 0x3d, 0x83, 0xe0, 0xb6, 0x2c, 0x5e, 0x22, 0xe4,
	0x19, 0x4b, 0x39, 0x0a, 0x85, 0x15, 0x7b, 0x45, 0xa3, 0x20, 0x01, 0x7d, 0x84, 0xa3, 0x65, 0xb2,
	0x4d, 0xcb, 0xac, 0x3a, 0x14, 0x82, 0xbf, 0xe6, 0x98, 0xa7, 0xf1, 0x0e, 0x35, 0xd3, 0x60, 0x21,
	0x71, 0xb7, 0x8b, 0x93, 0x37, 0x7d, 0x4b, 0x01, 0xc7, 0x44, 0x73, 0xcf, 0xc4, 0x14, 0xfa, 0xd5,
	0xf6, 0x9a, 0xe3, 0xe6, 0x1f, 0x47, 0xe8, 0x18, 0x7a, 0x52, 0xd7, 0xe6, 0x9b, 0xc3, 0x0d, 0xeb,
	0x30, 0xa5, 0x10, 0x28, 0x92, 0xfe, 0x42, 0x02, 0x07, 0x02, 0x4b, 0x92, 0x1f, 0xc9, 0xff, 0xf4,
	0x1c, 0x8e, 0x45, 0x0a, 0x42, 0xdd, 0x12, 0xab, 0x09, 0x62, 0x05, 0xbd, 0x1f, 0xa2, 0x16, 0xec,
	0x82, 0xb7, 0x98, 0x6b, 0x9a, 0xb7, 0x98, 0x3b, 0xb6, 0xbd, 0xbf, 0xb2, 0x69, 0x5a, 0x16, 0x2f,
	0xbe, 0x3c, 0x38, 0x14, 0xb2, 0x4b, 0x55, 0x32, 0x99, 0xea, 0x1a, 0x49, 0x5f, 0xd5, 0x3d, 0xb1,
	0x4b, 0x0e, 0xab, 0xa1, 0xd3, 0xdb, 0x0c, 0xda, 0xaa, 0x21, 0x32, 0xd0, 0xcf, 0x4e, 0x61, 0xf5,
	0x4b, 0xd7, 0x32, 0x1a, 0x93, 0x3b, 0x09, 0x35, 0xa9, 0xa6, 0x0c, 0x23, 0xe0, 0x64, 0x79, 0x09,
	0x1d, 0xd3, 0x02, 0x19, 0x6a, 0xc6, 0x7e, 0x2f, 0xf5, 0xab, 0x57, 0xe0, 0x57, 0x49, 0x92, 0x13,
	0xcb, 0x9c, 0xd5, 0x41, 0x38, 0xfc, 0x35, 0x57, 0xcb, 0x4f, 0x6d, 0x39, 0x9f, 0x7d, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0xc4, 0x87, 0xec, 0x17, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	UsernameUsed(ctx context.Context, in *UsernameUsedRequest, opts ...grpc.CallOption) (*UsedResponse, error)
	EmailUsed(ctx context.Context, in *EmailUsedRequest, opts ...grpc.CallOption) (*UsedResponse, error)
	AuthUser(ctx context.Context, in *AuthUserRequest, opts ...grpc.CallOption) (*AuthUserResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UsernameUsed(ctx context.Context, in *UsernameUsedRequest, opts ...grpc.CallOption) (*UsedResponse, error) {
	out := new(UsedResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/UsernameUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) EmailUsed(ctx context.Context, in *EmailUsedRequest, opts ...grpc.CallOption) (*UsedResponse, error) {
	out := new(UsedResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/EmailUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthUser(ctx context.Context, in *AuthUserRequest, opts ...grpc.CallOption) (*AuthUserResponse, error) {
	out := new(AuthUserResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/AuthUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*AuthResponse, error)
	Signup(context.Context, *SignupRequest) (*AuthResponse, error)
	UsernameUsed(context.Context, *UsernameUsedRequest) (*UsedResponse, error)
	EmailUsed(context.Context, *EmailUsedRequest) (*UsedResponse, error)
	AuthUser(context.Context, *AuthUserRequest) (*AuthUserResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) Signup(ctx context.Context, req *SignupRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (*UnimplementedAuthServiceServer) UsernameUsed(ctx context.Context, req *UsernameUsedRequest) (*UsedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameUsed not implemented")
}
func (*UnimplementedAuthServiceServer) EmailUsed(ctx context.Context, req *EmailUsedRequest) (*UsedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailUsed not implemented")
}
func (*UnimplementedAuthServiceServer) AuthUser(ctx context.Context, req *AuthUserRequest) (*AuthUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUser not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UsernameUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsernameUsedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UsernameUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/UsernameUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UsernameUsed(ctx, req.(*UsernameUsedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_EmailUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailUsedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).EmailUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/EmailUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).EmailUsed(ctx, req.(*EmailUsedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/AuthUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthUser(ctx, req.(*AuthUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Signup",
			Handler:    _AuthService_Signup_Handler,
		},
		{
			MethodName: "UsernameUsed",
			Handler:    _AuthService_UsernameUsed_Handler,
		},
		{
			MethodName: "EmailUsed",
			Handler:    _AuthService_EmailUsed_Handler,
		},
		{
			MethodName: "AuthUser",
			Handler:    _AuthService_AuthUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
