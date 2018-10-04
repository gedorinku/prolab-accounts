// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sessions.proto

package api_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
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

type Session struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_cad5307e8204e36b, []int{0}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (dst *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(dst, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type ListSessionsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSessionsRequest) Reset()         { *m = ListSessionsRequest{} }
func (m *ListSessionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListSessionsRequest) ProtoMessage()    {}
func (*ListSessionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_cad5307e8204e36b, []int{1}
}
func (m *ListSessionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSessionsRequest.Unmarshal(m, b)
}
func (m *ListSessionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSessionsRequest.Marshal(b, m, deterministic)
}
func (dst *ListSessionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSessionsRequest.Merge(dst, src)
}
func (m *ListSessionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListSessionsRequest.Size(m)
}
func (m *ListSessionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSessionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSessionsRequest proto.InternalMessageInfo

type ListSessionsResponse struct {
	Sessions             []*Session `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListSessionsResponse) Reset()         { *m = ListSessionsResponse{} }
func (m *ListSessionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListSessionsResponse) ProtoMessage()    {}
func (*ListSessionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_cad5307e8204e36b, []int{2}
}
func (m *ListSessionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSessionsResponse.Unmarshal(m, b)
}
func (m *ListSessionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSessionsResponse.Marshal(b, m, deterministic)
}
func (dst *ListSessionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSessionsResponse.Merge(dst, src)
}
func (m *ListSessionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListSessionsResponse.Size(m)
}
func (m *ListSessionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSessionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSessionsResponse proto.InternalMessageInfo

func (m *ListSessionsResponse) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

type GetSessionRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionRequest) Reset()         { *m = GetSessionRequest{} }
func (m *GetSessionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionRequest) ProtoMessage()    {}
func (*GetSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_cad5307e8204e36b, []int{3}
}
func (m *GetSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionRequest.Unmarshal(m, b)
}
func (m *GetSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionRequest.Marshal(b, m, deterministic)
}
func (dst *GetSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionRequest.Merge(dst, src)
}
func (m *GetSessionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionRequest.Size(m)
}
func (m *GetSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionRequest proto.InternalMessageInfo

func (m *GetSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type CreateSessionRequest struct {
	Session              *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSessionRequest) Reset()         { *m = CreateSessionRequest{} }
func (m *CreateSessionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSessionRequest) ProtoMessage()    {}
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_cad5307e8204e36b, []int{4}
}
func (m *CreateSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionRequest.Unmarshal(m, b)
}
func (m *CreateSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionRequest.Marshal(b, m, deterministic)
}
func (dst *CreateSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionRequest.Merge(dst, src)
}
func (m *CreateSessionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSessionRequest.Size(m)
}
func (m *CreateSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionRequest proto.InternalMessageInfo

func (m *CreateSessionRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type UpdateSessionRequest struct {
	Session              *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSessionRequest) Reset()         { *m = UpdateSessionRequest{} }
func (m *UpdateSessionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSessionRequest) ProtoMessage()    {}
func (*UpdateSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_cad5307e8204e36b, []int{5}
}
func (m *UpdateSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSessionRequest.Unmarshal(m, b)
}
func (m *UpdateSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSessionRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSessionRequest.Merge(dst, src)
}
func (m *UpdateSessionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSessionRequest.Size(m)
}
func (m *UpdateSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSessionRequest proto.InternalMessageInfo

func (m *UpdateSessionRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type DeleteSessionRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSessionRequest) Reset()         { *m = DeleteSessionRequest{} }
func (m *DeleteSessionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSessionRequest) ProtoMessage()    {}
func (*DeleteSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_cad5307e8204e36b, []int{6}
}
func (m *DeleteSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSessionRequest.Unmarshal(m, b)
}
func (m *DeleteSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSessionRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSessionRequest.Merge(dst, src)
}
func (m *DeleteSessionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSessionRequest.Size(m)
}
func (m *DeleteSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSessionRequest proto.InternalMessageInfo

func (m *DeleteSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func init() {
	proto.RegisterType((*Session)(nil), "com.github.ProgrammingLab.prolab_accounts.api.Session")
	proto.RegisterType((*ListSessionsRequest)(nil), "com.github.ProgrammingLab.prolab_accounts.api.ListSessionsRequest")
	proto.RegisterType((*ListSessionsResponse)(nil), "com.github.ProgrammingLab.prolab_accounts.api.ListSessionsResponse")
	proto.RegisterType((*GetSessionRequest)(nil), "com.github.ProgrammingLab.prolab_accounts.api.GetSessionRequest")
	proto.RegisterType((*CreateSessionRequest)(nil), "com.github.ProgrammingLab.prolab_accounts.api.CreateSessionRequest")
	proto.RegisterType((*UpdateSessionRequest)(nil), "com.github.ProgrammingLab.prolab_accounts.api.UpdateSessionRequest")
	proto.RegisterType((*DeleteSessionRequest)(nil), "com.github.ProgrammingLab.prolab_accounts.api.DeleteSessionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionServiceClient interface {
	ListSessions(ctx context.Context, in *ListSessionsRequest, opts ...grpc.CallOption) (*ListSessionsResponse, error)
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*Session, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error)
	UpdateSession(ctx context.Context, in *UpdateSessionRequest, opts ...grpc.CallOption) (*Session, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type sessionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSessionServiceClient(cc *grpc.ClientConn) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) ListSessions(ctx context.Context, in *ListSessionsRequest, opts ...grpc.CallOption) (*ListSessionsResponse, error) {
	out := new(ListSessionsResponse)
	err := c.cc.Invoke(ctx, "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/ListSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) UpdateSession(ctx context.Context, in *UpdateSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/UpdateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
type SessionServiceServer interface {
	ListSessions(context.Context, *ListSessionsRequest) (*ListSessionsResponse, error)
	GetSession(context.Context, *GetSessionRequest) (*Session, error)
	CreateSession(context.Context, *CreateSessionRequest) (*Session, error)
	UpdateSession(context.Context, *UpdateSessionRequest) (*Session, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*empty.Empty, error)
}

func RegisterSessionServiceServer(s *grpc.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_ListSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).ListSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/ListSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).ListSessions(ctx, req.(*ListSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_UpdateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).UpdateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/UpdateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).UpdateSession(ctx, req.(*UpdateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.ProgrammingLab.prolab_accounts.api.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSessions",
			Handler:    _SessionService_ListSessions_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _SessionService_GetSession_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _SessionService_CreateSession_Handler,
		},
		{
			MethodName: "UpdateSession",
			Handler:    _SessionService_UpdateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _SessionService_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sessions.proto",
}

func init() { proto.RegisterFile("sessions.proto", fileDescriptor_sessions_cad5307e8204e36b) }

var fileDescriptor_sessions_cad5307e8204e36b = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x19, 0x85, 0xde, 0xdb, 0x73, 0xed, 0x85, 0x3b, 0xc6, 0x4b, 0x89, 0x7a, 0x29, 0x59,
	0x55, 0xc1, 0x09, 0x44, 0x74, 0xe1, 0x4a, 0x5a, 0x45, 0x84, 0x2e, 0x4a, 0x8a, 0x1b, 0x37, 0x65,
	0x92, 0x8e, 0xe9, 0x48, 0x93, 0x19, 0x33, 0x13, 0x41, 0xc4, 0x8d, 0x6b, 0x77, 0xbe, 0x83, 0x88,
	0xef, 0xe0, 0x53, 0xf8, 0x02, 0x2e, 0x7c, 0x10, 0x69, 0x32, 0x69, 0x1a, 0x1b, 0x95, 0x44, 0x70,
	0x99, 0xcc, 0x9c, 0x73, 0xbe, 0xf9, 0xff, 0x7f, 0x06, 0x4e, 0x15, 0x53, 0x8a, 0x8b, 0x44, 0x11,
	0x99, 0x0a, 0x2d, 0xf0, 0x9d, 0x50, 0xc4, 0x24, 0xe2, 0x7a, 0x9d, 0x05, 0x64, 0x9e, 0x8a, 0x28,
	0xa5, 0x71, 0xcc, 0x93, 0x68, 0x46, 0x83, 0xed, 0x86, 0x0d, 0x0d, 0x96, 0x34, 0x0c, 0x45, 0x96,
	0x68, 0x45, 0xa8, 0xe4, 0xf6, 0x8d, 0x48, 0x88, 0x68, 0xc3, 0x5c, 0x2a, 0xb9, 0x4b, 0x93, 0x44,
	0x68, 0xaa, 0xab, 0x66, 0xf6, 0x75, 0xb3, 0x9a, 0x7f, 0x05, 0xd9, 0x0b, 0x97, 0xc5, 0x52, 0xbf,
	0x29, 0x16, 0x9d, 0x31, 0x1c, 0x2d, 0x8a, 0xd9, 0xf8, 0x26, 0x80, 0xc1, 0x58, 0xf2, 0xd5, 0x10,
	0x8d, 0xd0, 0xb8, 0xef, 0xf7, 0xcd, 0x9f, 0xa7, 0x2b, 0xe7, 0x1a, 0x5c, 0x9d, 0x71, 0xa5, 0xcd,
	0x6e, 0xe5, 0xb3, 0x57, 0x19, 0x53, 0xda, 0x79, 0x09, 0x56, 0xfd, 0xb7, 0x92, 0x22, 0x51, 0x0c,
	0xfb, 0x70, 0x5c, 0x1e, 0x6a, 0x88, 0x46, 0x97, 0xc7, 0x27, 0xde, 0x7d, 0xd2, 0xea, 0x54, 0xc4,
	0xb4, 0xf4, 0x77, 0x7d, 0x1c, 0x0f, 0xce, 0x9e, 0xb0, 0x72, 0x94, 0x01, 0xf8, 0x1b, 0xf6, 0x1a,
	0xac, 0x69, 0xca, 0xa8, 0x66, 0xbf, 0x94, 0xcd, 0xe1, 0xc8, 0x6c, 0xca, 0x6b, 0xba, 0xe3, 0x95,
	0x6d, 0xb6, 0x93, 0x9e, 0xc9, 0xd5, 0xff, 0x98, 0x74, 0x0f, 0xac, 0x47, 0x6c, 0xc3, 0x0e, 0x26,
	0xfd, 0x59, 0x0a, 0xef, 0x7b, 0x0f, 0x4e, 0x4d, 0xc5, 0x82, 0xa5, 0xaf, 0x79, 0xc8, 0xf0, 0x17,
	0x04, 0x57, 0xf6, 0xed, 0xc3, 0x93, 0x96, 0x6c, 0x0d, 0x91, 0xb0, 0xa7, 0xff, 0xd4, 0xa3, 0xc8,
	0x8f, 0x73, 0xf6, 0xfe, 0xdb, 0x8f, 0x8f, 0x97, 0x4e, 0x70, 0xdf, 0x2d, 0xed, 0xc7, 0x9f, 0x10,
	0x40, 0xe5, 0x3f, 0x7e, 0xd8, 0x72, 0xcc, 0x41, 0x74, 0xec, 0x8e, 0x46, 0x38, 0x17, 0x39, 0xdb,
	0x10, 0x9f, 0xef, 0xd8, 0xdc, 0xb7, 0x95, 0xf0, 0xef, 0xf0, 0x67, 0x04, 0x83, 0x5a, 0xe8, 0x70,
	0x5b, 0x49, 0x9a, 0x22, 0xdb, 0x19, 0xd7, 0xce, 0x71, 0x2d, 0xa7, 0x92, 0xf2, 0x41, 0x19, 0x25,
	0xfc, 0x15, 0xc1, 0xa0, 0x96, 0xda, 0xd6, 0xa8, 0x4d, 0x99, 0xef, 0x8c, 0xea, 0xe6, 0xa8, 0xb7,
	0xbc, 0x8b, 0x43, 0x65, 0xc9, 0x9e, 0xc2, 0x15, 0xff, 0x07, 0x04, 0x83, 0xda, 0x5d, 0x68, 0xcd,
	0xdf, 0x74, 0x93, 0xec, 0x73, 0x52, 0x3c, 0x9a, 0xa4, 0x7c, 0x34, 0xc9, 0xe3, 0xed, 0xa3, 0x59,
	0x3a, 0x7f, 0xfb, 0x37, 0xce, 0x4f, 0x8e, 0x9f, 0xf7, 0xa8, 0xe4, 0x4b, 0x19, 0x04, 0xbd, 0xbc,
	0xf2, 0xee, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x46, 0x8c, 0xd2, 0xdb, 0x05, 0x00, 0x00,
}
