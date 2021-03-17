// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/rpc/proto/auth/auth.proto

package auth

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

type TokenRequest struct {
	// token
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenRequest) Reset()         { *m = TokenRequest{} }
func (m *TokenRequest) String() string { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()    {}
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a551c8d06eb20ad, []int{0}
}

func (m *TokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenRequest.Unmarshal(m, b)
}
func (m *TokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenRequest.Marshal(b, m, deterministic)
}
func (m *TokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenRequest.Merge(m, src)
}
func (m *TokenRequest) XXX_Size() int {
	return xxx_messageInfo_TokenRequest.Size(m)
}
func (m *TokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TokenRequest proto.InternalMessageInfo

func (m *TokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type TokenResponse struct {
	// 状态码 200-正常 2-未创建文化云 0-错误 401-未授权 405-无权限
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 用户类型 文化云-culture 用户-user 运营后台-admin
	UserType string `protobuf:"bytes,2,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
	// 用户类型对应的用户ID
	UserId               uint64   `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenResponse) Reset()         { *m = TokenResponse{} }
func (m *TokenResponse) String() string { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()    {}
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a551c8d06eb20ad, []int{1}
}

func (m *TokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenResponse.Unmarshal(m, b)
}
func (m *TokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenResponse.Marshal(b, m, deterministic)
}
func (m *TokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenResponse.Merge(m, src)
}
func (m *TokenResponse) XXX_Size() int {
	return xxx_messageInfo_TokenResponse.Size(m)
}
func (m *TokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenResponse proto.InternalMessageInfo

func (m *TokenResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TokenResponse) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

func (m *TokenResponse) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*TokenRequest)(nil), "auth.TokenRequest")
	proto.RegisterType((*TokenResponse)(nil), "auth.TokenResponse")
}

func init() { proto.RegisterFile("server/rpc/proto/auth/auth.proto", fileDescriptor_6a551c8d06eb20ad) }

var fileDescriptor_6a551c8d06eb20ad = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2c, 0x2d,
	0xc9, 0x00, 0x13, 0x7a, 0x60, 0xbe, 0x10, 0x0b, 0x88, 0xad, 0xa4, 0xc2, 0xc5, 0x13, 0x92, 0x9f,
	0x9d, 0x9a, 0x17, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc2, 0xc5, 0x5a, 0x02, 0xe2,
	0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x4a, 0x91, 0x5c, 0xbc, 0x50, 0x55, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0xc9, 0xf9, 0x29, 0xa9, 0x60, 0x55, 0xcc,
	0x41, 0x60, 0xb6, 0x90, 0x34, 0x17, 0x67, 0x69, 0x71, 0x6a, 0x51, 0x7c, 0x49, 0x65, 0x41, 0xaa,
	0x04, 0x13, 0x58, 0x3b, 0x07, 0x48, 0x20, 0xa4, 0xb2, 0x20, 0x55, 0x48, 0x9c, 0x8b, 0x1d, 0x2c,
	0x99, 0x99, 0x22, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x06, 0xe2, 0x7a, 0xa6, 0x18, 0x95,
	0x73, 0xb1, 0x38, 0x96, 0x96, 0x64, 0x08, 0x99, 0x72, 0x71, 0x85, 0x25, 0xe6, 0x64, 0xa6, 0x80,
	0xed, 0x11, 0x12, 0xd2, 0x03, 0xbb, 0x14, 0xd9, 0x69, 0x52, 0xc2, 0x28, 0x62, 0x50, 0x87, 0x58,
	0x70, 0xf1, 0x82, 0xb5, 0xf9, 0xe5, 0x97, 0x90, 0xa6, 0x33, 0x89, 0x0d, 0x1c, 0x0c, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xd9, 0x11, 0x03, 0x2a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	// ValidToken 验证token
	ValidToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	// ValidNotToken 验证通用token
	ValidNotToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) ValidToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/ValidToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidNotToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/ValidNotToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	// ValidToken 验证token
	ValidToken(context.Context, *TokenRequest) (*TokenResponse, error)
	// ValidNotToken 验证通用token
	ValidNotToken(context.Context, *TokenRequest) (*TokenResponse, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) ValidToken(ctx context.Context, req *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidToken not implemented")
}
func (*UnimplementedAuthServer) ValidNotToken(ctx context.Context, req *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidNotToken not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_ValidToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ValidToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ValidToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ValidToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ValidNotToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ValidNotToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ValidNotToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ValidNotToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidToken",
			Handler:    _Auth_ValidToken_Handler,
		},
		{
			MethodName: "ValidNotToken",
			Handler:    _Auth_ValidNotToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/rpc/proto/auth/auth.proto",
}