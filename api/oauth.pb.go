// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oauth.proto

package api_pb // import "github.com/ProgrammingLab/prolab-accounts/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _type "github.com/ProgrammingLab/prolab-accounts/api/type"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type StartOAuthLoginRequest struct {
	LoginChallenge       string   `protobuf:"bytes,1,opt,name=login_challenge,json=loginChallenge,proto3" json:"login_challenge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartOAuthLoginRequest) Reset()         { *m = StartOAuthLoginRequest{} }
func (m *StartOAuthLoginRequest) String() string { return proto.CompactTextString(m) }
func (*StartOAuthLoginRequest) ProtoMessage()    {}
func (*StartOAuthLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_ad7e20b6ef80c2ba, []int{0}
}
func (m *StartOAuthLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartOAuthLoginRequest.Unmarshal(m, b)
}
func (m *StartOAuthLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartOAuthLoginRequest.Marshal(b, m, deterministic)
}
func (dst *StartOAuthLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartOAuthLoginRequest.Merge(dst, src)
}
func (m *StartOAuthLoginRequest) XXX_Size() int {
	return xxx_messageInfo_StartOAuthLoginRequest.Size(m)
}
func (m *StartOAuthLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartOAuthLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartOAuthLoginRequest proto.InternalMessageInfo

func (m *StartOAuthLoginRequest) GetLoginChallenge() string {
	if m != nil {
		return m.LoginChallenge
	}
	return ""
}

type StartOAuthLoginResponse struct {
	Skip                 bool     `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	RedirectUrl          string   `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartOAuthLoginResponse) Reset()         { *m = StartOAuthLoginResponse{} }
func (m *StartOAuthLoginResponse) String() string { return proto.CompactTextString(m) }
func (*StartOAuthLoginResponse) ProtoMessage()    {}
func (*StartOAuthLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_ad7e20b6ef80c2ba, []int{1}
}
func (m *StartOAuthLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartOAuthLoginResponse.Unmarshal(m, b)
}
func (m *StartOAuthLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartOAuthLoginResponse.Marshal(b, m, deterministic)
}
func (dst *StartOAuthLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartOAuthLoginResponse.Merge(dst, src)
}
func (m *StartOAuthLoginResponse) XXX_Size() int {
	return xxx_messageInfo_StartOAuthLoginResponse.Size(m)
}
func (m *StartOAuthLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartOAuthLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartOAuthLoginResponse proto.InternalMessageInfo

func (m *StartOAuthLoginResponse) GetSkip() bool {
	if m != nil {
		return m.Skip
	}
	return false
}

func (m *StartOAuthLoginResponse) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

type OAuthLoginRequest struct {
	LoginChallenge       string   `protobuf:"bytes,1,opt,name=login_challenge,json=loginChallenge,proto3" json:"login_challenge,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Remember             bool     `protobuf:"varint,4,opt,name=remember,proto3" json:"remember,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthLoginRequest) Reset()         { *m = OAuthLoginRequest{} }
func (m *OAuthLoginRequest) String() string { return proto.CompactTextString(m) }
func (*OAuthLoginRequest) ProtoMessage()    {}
func (*OAuthLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_ad7e20b6ef80c2ba, []int{2}
}
func (m *OAuthLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthLoginRequest.Unmarshal(m, b)
}
func (m *OAuthLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthLoginRequest.Marshal(b, m, deterministic)
}
func (dst *OAuthLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthLoginRequest.Merge(dst, src)
}
func (m *OAuthLoginRequest) XXX_Size() int {
	return xxx_messageInfo_OAuthLoginRequest.Size(m)
}
func (m *OAuthLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthLoginRequest proto.InternalMessageInfo

func (m *OAuthLoginRequest) GetLoginChallenge() string {
	if m != nil {
		return m.LoginChallenge
	}
	return ""
}

func (m *OAuthLoginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OAuthLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *OAuthLoginRequest) GetRemember() bool {
	if m != nil {
		return m.Remember
	}
	return false
}

type OAuthLoginResponse struct {
	RedirectUrl          string   `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthLoginResponse) Reset()         { *m = OAuthLoginResponse{} }
func (m *OAuthLoginResponse) String() string { return proto.CompactTextString(m) }
func (*OAuthLoginResponse) ProtoMessage()    {}
func (*OAuthLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_ad7e20b6ef80c2ba, []int{3}
}
func (m *OAuthLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthLoginResponse.Unmarshal(m, b)
}
func (m *OAuthLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthLoginResponse.Marshal(b, m, deterministic)
}
func (dst *OAuthLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthLoginResponse.Merge(dst, src)
}
func (m *OAuthLoginResponse) XXX_Size() int {
	return xxx_messageInfo_OAuthLoginResponse.Size(m)
}
func (m *OAuthLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthLoginResponse proto.InternalMessageInfo

func (m *OAuthLoginResponse) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

type StartOAuthConsentRequest struct {
	ConsentChallenge     string   `protobuf:"bytes,1,opt,name=consent_challenge,json=consentChallenge,proto3" json:"consent_challenge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartOAuthConsentRequest) Reset()         { *m = StartOAuthConsentRequest{} }
func (m *StartOAuthConsentRequest) String() string { return proto.CompactTextString(m) }
func (*StartOAuthConsentRequest) ProtoMessage()    {}
func (*StartOAuthConsentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_ad7e20b6ef80c2ba, []int{4}
}
func (m *StartOAuthConsentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartOAuthConsentRequest.Unmarshal(m, b)
}
func (m *StartOAuthConsentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartOAuthConsentRequest.Marshal(b, m, deterministic)
}
func (dst *StartOAuthConsentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartOAuthConsentRequest.Merge(dst, src)
}
func (m *StartOAuthConsentRequest) XXX_Size() int {
	return xxx_messageInfo_StartOAuthConsentRequest.Size(m)
}
func (m *StartOAuthConsentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartOAuthConsentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartOAuthConsentRequest proto.InternalMessageInfo

func (m *StartOAuthConsentRequest) GetConsentChallenge() string {
	if m != nil {
		return m.ConsentChallenge
	}
	return ""
}

type StartOAuthConsentResponse struct {
	Skip                 bool          `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	RedirectUrl          string        `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	RequestedScopes      []string      `protobuf:"bytes,3,rep,name=requested_scopes,json=requestedScopes,proto3" json:"requested_scopes,omitempty"`
	Client               *_type.Client `protobuf:"bytes,4,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StartOAuthConsentResponse) Reset()         { *m = StartOAuthConsentResponse{} }
func (m *StartOAuthConsentResponse) String() string { return proto.CompactTextString(m) }
func (*StartOAuthConsentResponse) ProtoMessage()    {}
func (*StartOAuthConsentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_ad7e20b6ef80c2ba, []int{5}
}
func (m *StartOAuthConsentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartOAuthConsentResponse.Unmarshal(m, b)
}
func (m *StartOAuthConsentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartOAuthConsentResponse.Marshal(b, m, deterministic)
}
func (dst *StartOAuthConsentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartOAuthConsentResponse.Merge(dst, src)
}
func (m *StartOAuthConsentResponse) XXX_Size() int {
	return xxx_messageInfo_StartOAuthConsentResponse.Size(m)
}
func (m *StartOAuthConsentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartOAuthConsentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartOAuthConsentResponse proto.InternalMessageInfo

func (m *StartOAuthConsentResponse) GetSkip() bool {
	if m != nil {
		return m.Skip
	}
	return false
}

func (m *StartOAuthConsentResponse) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

func (m *StartOAuthConsentResponse) GetRequestedScopes() []string {
	if m != nil {
		return m.RequestedScopes
	}
	return nil
}

func (m *StartOAuthConsentResponse) GetClient() *_type.Client {
	if m != nil {
		return m.Client
	}
	return nil
}

type OAuthConsentRequest struct {
	ConsentChallenge     string   `protobuf:"bytes,1,opt,name=consent_challenge,json=consentChallenge,proto3" json:"consent_challenge,omitempty"`
	Accept               bool     `protobuf:"varint,2,opt,name=accept,proto3" json:"accept,omitempty"`
	GrantScopes          []string `protobuf:"bytes,3,rep,name=grant_scopes,json=grantScopes,proto3" json:"grant_scopes,omitempty"`
	Remember             bool     `protobuf:"varint,4,opt,name=remember,proto3" json:"remember,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthConsentRequest) Reset()         { *m = OAuthConsentRequest{} }
func (m *OAuthConsentRequest) String() string { return proto.CompactTextString(m) }
func (*OAuthConsentRequest) ProtoMessage()    {}
func (*OAuthConsentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_ad7e20b6ef80c2ba, []int{6}
}
func (m *OAuthConsentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthConsentRequest.Unmarshal(m, b)
}
func (m *OAuthConsentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthConsentRequest.Marshal(b, m, deterministic)
}
func (dst *OAuthConsentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthConsentRequest.Merge(dst, src)
}
func (m *OAuthConsentRequest) XXX_Size() int {
	return xxx_messageInfo_OAuthConsentRequest.Size(m)
}
func (m *OAuthConsentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthConsentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthConsentRequest proto.InternalMessageInfo

func (m *OAuthConsentRequest) GetConsentChallenge() string {
	if m != nil {
		return m.ConsentChallenge
	}
	return ""
}

func (m *OAuthConsentRequest) GetAccept() bool {
	if m != nil {
		return m.Accept
	}
	return false
}

func (m *OAuthConsentRequest) GetGrantScopes() []string {
	if m != nil {
		return m.GrantScopes
	}
	return nil
}

func (m *OAuthConsentRequest) GetRemember() bool {
	if m != nil {
		return m.Remember
	}
	return false
}

type OAuthConsentResponse struct {
	RedirectUrl          string   `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthConsentResponse) Reset()         { *m = OAuthConsentResponse{} }
func (m *OAuthConsentResponse) String() string { return proto.CompactTextString(m) }
func (*OAuthConsentResponse) ProtoMessage()    {}
func (*OAuthConsentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_ad7e20b6ef80c2ba, []int{7}
}
func (m *OAuthConsentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthConsentResponse.Unmarshal(m, b)
}
func (m *OAuthConsentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthConsentResponse.Marshal(b, m, deterministic)
}
func (dst *OAuthConsentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthConsentResponse.Merge(dst, src)
}
func (m *OAuthConsentResponse) XXX_Size() int {
	return xxx_messageInfo_OAuthConsentResponse.Size(m)
}
func (m *OAuthConsentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthConsentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthConsentResponse proto.InternalMessageInfo

func (m *OAuthConsentResponse) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*StartOAuthLoginRequest)(nil), "programming_lab.prolab_accounts.StartOAuthLoginRequest")
	proto.RegisterType((*StartOAuthLoginResponse)(nil), "programming_lab.prolab_accounts.StartOAuthLoginResponse")
	proto.RegisterType((*OAuthLoginRequest)(nil), "programming_lab.prolab_accounts.OAuthLoginRequest")
	proto.RegisterType((*OAuthLoginResponse)(nil), "programming_lab.prolab_accounts.OAuthLoginResponse")
	proto.RegisterType((*StartOAuthConsentRequest)(nil), "programming_lab.prolab_accounts.StartOAuthConsentRequest")
	proto.RegisterType((*StartOAuthConsentResponse)(nil), "programming_lab.prolab_accounts.StartOAuthConsentResponse")
	proto.RegisterType((*OAuthConsentRequest)(nil), "programming_lab.prolab_accounts.OAuthConsentRequest")
	proto.RegisterType((*OAuthConsentResponse)(nil), "programming_lab.prolab_accounts.OAuthConsentResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OAuthServiceClient is the client API for OAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OAuthServiceClient interface {
	StartOAuthLogin(ctx context.Context, in *StartOAuthLoginRequest, opts ...grpc.CallOption) (*StartOAuthLoginResponse, error)
	OAuthLogin(ctx context.Context, in *OAuthLoginRequest, opts ...grpc.CallOption) (*OAuthLoginResponse, error)
	StartOAuthConsent(ctx context.Context, in *StartOAuthConsentRequest, opts ...grpc.CallOption) (*StartOAuthConsentResponse, error)
	OAuthConsent(ctx context.Context, in *OAuthConsentRequest, opts ...grpc.CallOption) (*OAuthConsentResponse, error)
}

type oAuthServiceClient struct {
	cc *grpc.ClientConn
}

func NewOAuthServiceClient(cc *grpc.ClientConn) OAuthServiceClient {
	return &oAuthServiceClient{cc}
}

func (c *oAuthServiceClient) StartOAuthLogin(ctx context.Context, in *StartOAuthLoginRequest, opts ...grpc.CallOption) (*StartOAuthLoginResponse, error) {
	out := new(StartOAuthLoginResponse)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.OAuthService/StartOAuthLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) OAuthLogin(ctx context.Context, in *OAuthLoginRequest, opts ...grpc.CallOption) (*OAuthLoginResponse, error) {
	out := new(OAuthLoginResponse)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.OAuthService/OAuthLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) StartOAuthConsent(ctx context.Context, in *StartOAuthConsentRequest, opts ...grpc.CallOption) (*StartOAuthConsentResponse, error) {
	out := new(StartOAuthConsentResponse)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.OAuthService/StartOAuthConsent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) OAuthConsent(ctx context.Context, in *OAuthConsentRequest, opts ...grpc.CallOption) (*OAuthConsentResponse, error) {
	out := new(OAuthConsentResponse)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.OAuthService/OAuthConsent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthServiceServer is the server API for OAuthService service.
type OAuthServiceServer interface {
	StartOAuthLogin(context.Context, *StartOAuthLoginRequest) (*StartOAuthLoginResponse, error)
	OAuthLogin(context.Context, *OAuthLoginRequest) (*OAuthLoginResponse, error)
	StartOAuthConsent(context.Context, *StartOAuthConsentRequest) (*StartOAuthConsentResponse, error)
	OAuthConsent(context.Context, *OAuthConsentRequest) (*OAuthConsentResponse, error)
}

func RegisterOAuthServiceServer(s *grpc.Server, srv OAuthServiceServer) {
	s.RegisterService(&_OAuthService_serviceDesc, srv)
}

func _OAuthService_StartOAuthLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartOAuthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).StartOAuthLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.OAuthService/StartOAuthLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).StartOAuthLogin(ctx, req.(*StartOAuthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_OAuthLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).OAuthLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.OAuthService/OAuthLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).OAuthLogin(ctx, req.(*OAuthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_StartOAuthConsent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartOAuthConsentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).StartOAuthConsent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.OAuthService/StartOAuthConsent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).StartOAuthConsent(ctx, req.(*StartOAuthConsentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_OAuthConsent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthConsentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).OAuthConsent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.OAuthService/OAuthConsent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).OAuthConsent(ctx, req.(*OAuthConsentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OAuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "programming_lab.prolab_accounts.OAuthService",
	HandlerType: (*OAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartOAuthLogin",
			Handler:    _OAuthService_StartOAuthLogin_Handler,
		},
		{
			MethodName: "OAuthLogin",
			Handler:    _OAuthService_OAuthLogin_Handler,
		},
		{
			MethodName: "StartOAuthConsent",
			Handler:    _OAuthService_StartOAuthConsent_Handler,
		},
		{
			MethodName: "OAuthConsent",
			Handler:    _OAuthService_OAuthConsent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth.proto",
}

func init() { proto.RegisterFile("oauth.proto", fileDescriptor_oauth_ad7e20b6ef80c2ba) }

var fileDescriptor_oauth_ad7e20b6ef80c2ba = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x35, 0x6d, 0x55, 0xa5, 0x37, 0x51, 0x7e, 0xa6, 0x55, 0xea, 0x06, 0x24, 0x8a, 0x37,
	0x94, 0x02, 0xb1, 0x94, 0x16, 0x4a, 0xc3, 0xaa, 0x04, 0x89, 0x4d, 0x25, 0x2a, 0x47, 0x6c, 0xd8,
	0x58, 0x63, 0x67, 0x70, 0x2c, 0xec, 0x99, 0x61, 0x3c, 0x06, 0x75, 0xcb, 0x8a, 0x15, 0x1b, 0x24,
	0x24, 0x58, 0xf2, 0x2a, 0x3c, 0x02, 0xaf, 0xc0, 0x53, 0xb0, 0x42, 0x1e, 0x4f, 0x12, 0xf2, 0x23,
	0x42, 0xda, 0x9d, 0xef, 0x19, 0x9d, 0x9b, 0xef, 0x5c, 0xdf, 0x71, 0xa0, 0xcc, 0x49, 0xa6, 0x86,
	0x6d, 0x21, 0xb9, 0xe2, 0xf8, 0x96, 0x90, 0x3c, 0x94, 0x24, 0x49, 0x22, 0x16, 0x7a, 0x31, 0xf1,
	0x73, 0x39, 0x26, 0xbe, 0x47, 0x82, 0x80, 0x67, 0x4c, 0xa5, 0xad, 0x9b, 0x21, 0xe7, 0x61, 0x4c,
	0x1d, 0x22, 0x22, 0x87, 0x30, 0xc6, 0x15, 0x51, 0x11, 0x67, 0x69, 0x61, 0x6f, 0xdd, 0x30, 0xa7,
	0xba, 0xf2, 0xb3, 0xd7, 0x0e, 0x4d, 0x84, 0xba, 0x34, 0x87, 0x0d, 0x75, 0x29, 0xa8, 0x13, 0xc4,
	0x11, 0x65, 0xaa, 0x90, 0xec, 0x33, 0x68, 0xf6, 0x15, 0x91, 0xea, 0xc5, 0x59, 0xa6, 0x86, 0xe7,
	0x3c, 0x8c, 0x98, 0x4b, 0xdf, 0x66, 0x34, 0x55, 0xf8, 0x0e, 0xd4, 0xe2, 0xbc, 0xf6, 0x82, 0x21,
	0x89, 0x63, 0xca, 0x42, 0x6a, 0xa1, 0x7d, 0x74, 0xb0, 0xe5, 0x56, 0xb5, 0xdc, 0x1b, 0xa9, 0xf6,
	0x05, 0xec, 0xce, 0xb5, 0x48, 0x05, 0x67, 0x29, 0xc5, 0x18, 0x36, 0xd2, 0x37, 0x91, 0xd0, 0xc6,
	0x92, 0xab, 0x9f, 0xf1, 0x6d, 0xa8, 0x48, 0x3a, 0x88, 0x24, 0x0d, 0x94, 0x97, 0xc9, 0xd8, 0x5a,
	0xd3, 0x4d, 0xcb, 0x23, 0xed, 0xa5, 0x8c, 0xed, 0x8f, 0x08, 0x1a, 0x57, 0x07, 0xca, 0x7f, 0x95,
	0x91, 0x84, 0x9a, 0xce, 0xfa, 0x19, 0xb7, 0xa0, 0x24, 0x48, 0x9a, 0xbe, 0xe7, 0x72, 0x60, 0xad,
	0x6b, 0x7d, 0x5c, 0xe7, 0x67, 0x92, 0x26, 0x34, 0xf1, 0xa9, 0xb4, 0x36, 0x34, 0xe9, 0xb8, 0xb6,
	0x4f, 0x00, 0x2f, 0xc8, 0x35, 0x9b, 0x01, 0xcd, 0x67, 0x78, 0x0e, 0xd6, 0x64, 0x2a, 0xbd, 0xdc,
	0xc5, 0xd4, 0x28, 0xc9, 0x3d, 0x68, 0x04, 0x85, 0x32, 0x97, 0xa5, 0x6e, 0x0e, 0x26, 0xe3, 0xfd,
	0x81, 0x60, 0x6f, 0x41, 0xa7, 0x6b, 0x4d, 0x18, 0xdf, 0x85, 0xba, 0x2c, 0x60, 0xe8, 0xc0, 0x4b,
	0x03, 0x2e, 0x68, 0x6a, 0xad, 0xef, 0xaf, 0x1f, 0x6c, 0xb9, 0xb5, 0xb1, 0xde, 0xd7, 0x32, 0x7e,
	0x06, 0x9b, 0xc5, 0xc6, 0xe8, 0xd9, 0x94, 0x3b, 0xf7, 0xdb, 0x4b, 0x36, 0xb4, 0x9d, 0x6f, 0x59,
	0xbb, 0xa7, 0x3d, 0xae, 0xf1, 0xda, 0x5f, 0x11, 0x6c, 0x5f, 0x77, 0x14, 0xb8, 0x09, 0x9b, 0x24,
	0x08, 0xa8, 0x50, 0x3a, 0x52, 0xc9, 0x35, 0x55, 0x1e, 0x38, 0x94, 0x84, 0xa9, 0xe9, 0x24, 0x65,
	0xad, 0x99, 0x14, 0xff, 0x7a, 0xc7, 0xa7, 0xb0, 0xb3, 0x70, 0xb6, 0xcb, 0xdf, 0x72, 0xe7, 0xf7,
	0x06, 0x54, 0xb4, 0xb7, 0x4f, 0xe5, 0xbb, 0x28, 0xa0, 0xf8, 0x1b, 0x82, 0xda, 0xcc, 0x6d, 0xc0,
	0x27, 0x4b, 0x27, 0xb6, 0xf8, 0x0a, 0xb6, 0x1e, 0xaf, 0x6e, 0x2c, 0xd0, 0xed, 0x9d, 0x0f, 0x3f,
	0x7f, 0x7d, 0x5e, 0xab, 0xe2, 0x8a, 0xa3, 0xbf, 0x2d, 0x8e, 0xbe, 0x21, 0xf8, 0x13, 0x02, 0xf8,
	0x8b, 0xab, 0xb3, 0xb4, 0xfd, 0x3c, 0xd2, 0xd1, 0x4a, 0x1e, 0x43, 0xb3, 0xab, 0x69, 0x1a, 0xf6,
	0x14, 0x4d, 0x17, 0x1d, 0xe2, 0xef, 0x08, 0x1a, 0x73, 0xbb, 0x8d, 0x4f, 0x57, 0x88, 0x3d, 0xbd,
	0x4e, 0xad, 0xee, 0x55, 0xac, 0x86, 0xb2, 0xa9, 0x29, 0xeb, 0xb8, 0x6a, 0x28, 0xcd, 0xfa, 0xe1,
	0x2f, 0xc8, 0xbc, 0xe3, 0x11, 0xdf, 0xf1, 0xff, 0xcd, 0x60, 0x06, 0xed, 0xe1, 0x8a, 0x2e, 0x43,
	0xb5, 0xa7, 0xa9, 0xb6, 0xed, 0x19, 0xaa, 0x2e, 0x3a, 0x7c, 0xfa, 0xe8, 0xd5, 0x71, 0x18, 0xa9,
	0x61, 0xe6, 0xb7, 0x03, 0x9e, 0x38, 0x17, 0x93, 0xee, 0xe7, 0xc4, 0x77, 0x8a, 0xe6, 0x0f, 0x46,
	0xcd, 0xf3, 0xbf, 0x8b, 0x27, 0x44, 0x44, 0x9e, 0xf0, 0xfd, 0x4d, 0xfd, 0xe9, 0x3f, 0xfa, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x7e, 0x7e, 0xb1, 0x04, 0x78, 0x06, 0x00, 0x00,
}