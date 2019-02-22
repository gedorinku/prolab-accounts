// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package api_pb // import "github.com/ProgrammingLab/prolab-accounts/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _type "github.com/ProgrammingLab/prolab-accounts/api/type"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/mwitkow/go-proto-validators"
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

type ProfileScope int32

const (
	ProfileScope_MEMBERS_ONLY ProfileScope = 0
	ProfileScope_PUBLIC       ProfileScope = 1
)

var ProfileScope_name = map[int32]string{
	0: "MEMBERS_ONLY",
	1: "PUBLIC",
}
var ProfileScope_value = map[string]int32{
	"MEMBERS_ONLY": 0,
	"PUBLIC":       1,
}

func (x ProfileScope) String() string {
	return proto.EnumName(ProfileScope_name, int32(x))
}
func (ProfileScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_users_b2854ca0ded5a5cb, []int{0}
}

type User struct {
	UserId               uint32            `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string            `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	FullName             string            `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	IconUrl              string            `protobuf:"bytes,5,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	Description          string            `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Grade                int32             `protobuf:"varint,7,opt,name=grade,proto3" json:"grade,omitempty"`
	Left                 bool              `protobuf:"varint,8,opt,name=left,proto3" json:"left,omitempty"`
	Role                 string            `protobuf:"bytes,11,opt,name=role,proto3" json:"role,omitempty"`
	TwitterScreenName    string            `protobuf:"bytes,12,opt,name=twitter_screen_name,json=twitterScreenName,proto3" json:"twitter_screen_name,omitempty"`
	GithubUserName       string            `protobuf:"bytes,13,opt,name=github_user_name,json=githubUserName,proto3" json:"github_user_name,omitempty"`
	Department           *_type.Department `protobuf:"bytes,14,opt,name=department,proto3" json:"department,omitempty"`
	ProfileScope         ProfileScope      `protobuf:"varint,15,opt,name=profile_scope,json=profileScope,proto3,enum=programming_lab.prolab_accounts.ProfileScope" json:"profile_scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_b2854ca0ded5a5cb, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *User) GetIconUrl() string {
	if m != nil {
		return m.IconUrl
	}
	return ""
}

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *User) GetGrade() int32 {
	if m != nil {
		return m.Grade
	}
	return 0
}

func (m *User) GetLeft() bool {
	if m != nil {
		return m.Left
	}
	return false
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetTwitterScreenName() string {
	if m != nil {
		return m.TwitterScreenName
	}
	return ""
}

func (m *User) GetGithubUserName() string {
	if m != nil {
		return m.GithubUserName
	}
	return ""
}

func (m *User) GetDepartment() *_type.Department {
	if m != nil {
		return m.Department
	}
	return nil
}

func (m *User) GetProfileScope() ProfileScope {
	if m != nil {
		return m.ProfileScope
	}
	return ProfileScope_MEMBERS_ONLY
}

type ListUsersRequest struct {
	PageToken            uint32   `protobuf:"varint,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_b2854ca0ded5a5cb, []int{1}
}
func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (dst *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(dst, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetPageToken() uint32 {
	if m != nil {
		return m.PageToken
	}
	return 0
}

func (m *ListUsersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type ListUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	NextPageToken        uint32   `protobuf:"varint,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_b2854ca0ded5a5cb, []int{2}
}
func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (dst *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(dst, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ListUsersResponse) GetNextPageToken() uint32 {
	if m != nil {
		return m.NextPageToken
	}
	return 0
}

type GetUserRequest struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_b2854ca0ded5a5cb, []int{3}
}
func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(dst, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	RegisterationToken   string   `protobuf:"bytes,2,opt,name=registeration_token,json=registerationToken,proto3" json:"registeration_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_b2854ca0ded5a5cb, []int{4}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserRequest) GetRegisterationToken() string {
	if m != nil {
		return m.RegisterationToken
	}
	return ""
}

type GetCurrentUserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentUserRequest) Reset()         { *m = GetCurrentUserRequest{} }
func (m *GetCurrentUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentUserRequest) ProtoMessage()    {}
func (*GetCurrentUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_b2854ca0ded5a5cb, []int{5}
}
func (m *GetCurrentUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentUserRequest.Unmarshal(m, b)
}
func (m *GetCurrentUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetCurrentUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentUserRequest.Merge(dst, src)
}
func (m *GetCurrentUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentUserRequest.Size(m)
}
func (m *GetCurrentUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentUserRequest proto.InternalMessageInfo

type UpdateUserProfileRequest struct {
	FullName             string       `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Grade                int32        `protobuf:"varint,3,opt,name=grade,proto3" json:"grade,omitempty"`
	Left                 bool         `protobuf:"varint,4,opt,name=left,proto3" json:"left,omitempty"`
	RoleId               uint32       `protobuf:"varint,5,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	TwitterScreenName    string       `protobuf:"bytes,6,opt,name=twitter_screen_name,json=twitterScreenName,proto3" json:"twitter_screen_name,omitempty"`
	GithubUserName       string       `protobuf:"bytes,7,opt,name=github_user_name,json=githubUserName,proto3" json:"github_user_name,omitempty"`
	DepartmentId         uint32       `protobuf:"varint,8,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	ProfileScope         ProfileScope `protobuf:"varint,9,opt,name=profile_scope,json=profileScope,proto3,enum=programming_lab.prolab_accounts.ProfileScope" json:"profile_scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateUserProfileRequest) Reset()         { *m = UpdateUserProfileRequest{} }
func (m *UpdateUserProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserProfileRequest) ProtoMessage()    {}
func (*UpdateUserProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_b2854ca0ded5a5cb, []int{6}
}
func (m *UpdateUserProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserProfileRequest.Unmarshal(m, b)
}
func (m *UpdateUserProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserProfileRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserProfileRequest.Merge(dst, src)
}
func (m *UpdateUserProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserProfileRequest.Size(m)
}
func (m *UpdateUserProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserProfileRequest proto.InternalMessageInfo

func (m *UpdateUserProfileRequest) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetGrade() int32 {
	if m != nil {
		return m.Grade
	}
	return 0
}

func (m *UpdateUserProfileRequest) GetLeft() bool {
	if m != nil {
		return m.Left
	}
	return false
}

func (m *UpdateUserProfileRequest) GetRoleId() uint32 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *UpdateUserProfileRequest) GetTwitterScreenName() string {
	if m != nil {
		return m.TwitterScreenName
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetGithubUserName() string {
	if m != nil {
		return m.GithubUserName
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetDepartmentId() uint32 {
	if m != nil {
		return m.DepartmentId
	}
	return 0
}

func (m *UpdateUserProfileRequest) GetProfileScope() ProfileScope {
	if m != nil {
		return m.ProfileScope
	}
	return ProfileScope_MEMBERS_ONLY
}

type UpdateUserIconRequest struct {
	Image                []byte   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserIconRequest) Reset()         { *m = UpdateUserIconRequest{} }
func (m *UpdateUserIconRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserIconRequest) ProtoMessage()    {}
func (*UpdateUserIconRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_b2854ca0ded5a5cb, []int{7}
}
func (m *UpdateUserIconRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserIconRequest.Unmarshal(m, b)
}
func (m *UpdateUserIconRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserIconRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserIconRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserIconRequest.Merge(dst, src)
}
func (m *UpdateUserIconRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserIconRequest.Size(m)
}
func (m *UpdateUserIconRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserIconRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserIconRequest proto.InternalMessageInfo

func (m *UpdateUserIconRequest) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type UpdatePasswordRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	NewPassword          string   `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	NewPasswordRepeat    string   `protobuf:"bytes,3,opt,name=new_password_repeat,json=newPasswordRepeat,proto3" json:"new_password_repeat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePasswordRequest) Reset()         { *m = UpdatePasswordRequest{} }
func (m *UpdatePasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePasswordRequest) ProtoMessage()    {}
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_b2854ca0ded5a5cb, []int{8}
}
func (m *UpdatePasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePasswordRequest.Unmarshal(m, b)
}
func (m *UpdatePasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePasswordRequest.Marshal(b, m, deterministic)
}
func (dst *UpdatePasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePasswordRequest.Merge(dst, src)
}
func (m *UpdatePasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePasswordRequest.Size(m)
}
func (m *UpdatePasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePasswordRequest proto.InternalMessageInfo

func (m *UpdatePasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UpdatePasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *UpdatePasswordRequest) GetNewPasswordRepeat() string {
	if m != nil {
		return m.NewPasswordRepeat
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "programming_lab.prolab_accounts.User")
	proto.RegisterType((*ListUsersRequest)(nil), "programming_lab.prolab_accounts.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "programming_lab.prolab_accounts.ListUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "programming_lab.prolab_accounts.GetUserRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "programming_lab.prolab_accounts.CreateUserRequest")
	proto.RegisterType((*GetCurrentUserRequest)(nil), "programming_lab.prolab_accounts.GetCurrentUserRequest")
	proto.RegisterType((*UpdateUserProfileRequest)(nil), "programming_lab.prolab_accounts.UpdateUserProfileRequest")
	proto.RegisterType((*UpdateUserIconRequest)(nil), "programming_lab.prolab_accounts.UpdateUserIconRequest")
	proto.RegisterType((*UpdatePasswordRequest)(nil), "programming_lab.prolab_accounts.UpdatePasswordRequest")
	proto.RegisterEnum("programming_lab.prolab_accounts.ProfileScope", ProfileScope_name, ProfileScope_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	ListPublicUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUserIcon(ctx context.Context, in *UpdateUserIconRequest, opts ...grpc.CallOption) (*User, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ListPublicUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/ListPublicUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/GetCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserIcon(ctx context.Context, in *UpdateUserIconRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/UpdateUserIcon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	ListPublicUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*User, error)
	UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*User, error)
	UpdateUserIcon(context.Context, *UpdateUserIconRequest) (*User, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*empty.Empty, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_ListPublicUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListPublicUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/ListPublicUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListPublicUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/GetCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, req.(*UpdateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserIcon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserIconRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserIcon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/UpdateUserIcon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserIcon(ctx, req.(*UpdateUserIconRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "programming_lab.prolab_accounts.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPublicUsers",
			Handler:    _UserService_ListPublicUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _UserService_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserService_UpdateUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserIcon",
			Handler:    _UserService_UpdateUserIcon_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserService_UpdatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_users_b2854ca0ded5a5cb) }

var fileDescriptor_users_b2854ca0ded5a5cb = []byte{
	// 1120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x5d, 0x6f, 0x1b, 0x45,
	0x17, 0xee, 0x26, 0xfe, 0x3c, 0xfe, 0x88, 0x33, 0xcd, 0xc7, 0xbe, 0xee, 0x4b, 0x63, 0x96, 0x80,
	0x9c, 0x50, 0x7b, 0x53, 0x03, 0x91, 0xd2, 0x5e, 0xe1, 0x10, 0x95, 0xa0, 0x34, 0x58, 0x1b, 0x72,
	0x41, 0x43, 0xbb, 0x1a, 0x7b, 0x27, 0xcb, 0xaa, 0xeb, 0xdd, 0x65, 0x76, 0x5c, 0xa7, 0x89, 0x02,
	0x88, 0x8b, 0x8a, 0x3b, 0x2e, 0xf8, 0x23, 0xfc, 0x14, 0x24, 0x7e, 0x40, 0x44, 0xc4, 0x0f, 0x41,
	0x33, 0xb3, 0x1b, 0x6f, 0x9c, 0x54, 0x76, 0xb9, 0x9b, 0x39, 0x1f, 0x73, 0x9e, 0x39, 0x67, 0x9e,
	0x47, 0x03, 0x85, 0x41, 0x48, 0x68, 0xd8, 0x0c, 0xa8, 0xcf, 0x7c, 0xb4, 0x12, 0x50, 0xdf, 0xa6,
	0xb8, 0xdf, 0x77, 0x3c, 0xdb, 0x74, 0x71, 0x97, 0x9b, 0x5d, 0xdc, 0x35, 0x71, 0xaf, 0xe7, 0x0f,
	0x3c, 0x16, 0x56, 0x37, 0x6d, 0x87, 0x7d, 0x3f, 0xe8, 0x36, 0x7b, 0x7e, 0x5f, 0xef, 0x0f, 0x1d,
	0xf6, 0xd2, 0x1f, 0xea, 0xb6, 0xdf, 0x10, 0xd9, 0x8d, 0x57, 0xd8, 0x75, 0x2c, 0xcc, 0x7c, 0x1a,
	0xea, 0x57, 0x4b, 0x79, 0x70, 0xf5, 0xff, 0xb6, 0xef, 0xdb, 0x2e, 0xd1, 0x71, 0xe0, 0xe8, 0xd8,
	0xf3, 0x7c, 0x86, 0x99, 0xe3, 0x7b, 0x51, 0xd9, 0xea, 0xbd, 0xc8, 0x2b, 0x76, 0xdd, 0xc1, 0xb1,
	0x4e, 0xfa, 0x01, 0x7b, 0x1d, 0x39, 0x17, 0xd9, 0xeb, 0x80, 0xe8, 0x16, 0x09, 0x30, 0x65, 0x7d,
	0xe2, 0x31, 0x69, 0xd6, 0xfe, 0x9e, 0x85, 0xd4, 0x61, 0x48, 0x28, 0x5a, 0x86, 0x2c, 0xbf, 0x82,
	0xe9, 0x58, 0xaa, 0x52, 0x53, 0xea, 0x25, 0x23, 0xc3, 0xb7, 0xbb, 0x16, 0x42, 0x90, 0xf2, 0x70,
	0x9f, 0xa8, 0x33, 0x35, 0xa5, 0x9e, 0x37, 0xc4, 0x1a, 0x2d, 0x40, 0x9a, 0xf4, 0xb1, 0xe3, 0xaa,
	0xb3, 0xc2, 0x28, 0x37, 0xe8, 0x1e, 0xe4, 0x8f, 0x07, 0xae, 0x6b, 0x8a, 0xf0, 0x94, 0xf0, 0xe4,
	0xb8, 0x61, 0x9f, 0xa7, 0xfc, 0x0f, 0x72, 0x4e, 0xcf, 0xf7, 0xcc, 0x01, 0x75, 0xd5, 0xb4, 0xf0,
	0x65, 0xf9, 0xfe, 0x90, 0xba, 0xa8, 0x06, 0x05, 0x8b, 0x84, 0x3d, 0xea, 0x04, 0xfc, 0x36, 0x6a,
	0x46, 0x78, 0x93, 0x26, 0x5e, 0xcf, 0xa6, 0xd8, 0x22, 0x6a, 0xb6, 0xa6, 0xd4, 0xd3, 0x86, 0xdc,
	0x70, 0x64, 0x2e, 0x39, 0x66, 0x6a, 0xae, 0xa6, 0xd4, 0x73, 0x86, 0x58, 0x73, 0x1b, 0xf5, 0x5d,
	0xa2, 0x16, 0x24, 0x5a, 0xbe, 0x46, 0x4d, 0xb8, 0xcb, 0x86, 0x0e, 0x63, 0x84, 0x9a, 0x61, 0x8f,
	0x12, 0xe2, 0x49, 0x84, 0x45, 0x11, 0x32, 0x1f, 0xb9, 0x0e, 0x84, 0x47, 0x40, 0xad, 0x43, 0x45,
	0xce, 0xc7, 0x14, 0x1d, 0x11, 0xc1, 0x25, 0x11, 0x5c, 0x96, 0x76, 0xde, 0x30, 0x11, 0xd9, 0x01,
	0x18, 0x75, 0x54, 0x2d, 0xd7, 0x94, 0x7a, 0xa1, 0xb5, 0xd1, 0x9c, 0x30, 0xfd, 0x26, 0x9f, 0x44,
	0xf3, 0x8b, 0xab, 0x3c, 0x23, 0x71, 0x06, 0x32, 0xa0, 0x14, 0x50, 0xff, 0xd8, 0x71, 0x89, 0x19,
	0xf6, 0xfc, 0x80, 0xa8, 0x73, 0x35, 0xa5, 0x5e, 0x6e, 0x35, 0x26, 0x1e, 0xda, 0x91, 0x59, 0x07,
	0x3c, 0xc9, 0x28, 0x06, 0x89, 0x9d, 0xb6, 0x0f, 0x95, 0x3d, 0x27, 0x64, 0x1c, 0x75, 0x68, 0x90,
	0x1f, 0x06, 0x24, 0x64, 0xe8, 0x3d, 0x80, 0x00, 0xdb, 0xc4, 0x64, 0xfe, 0x4b, 0xe2, 0x45, 0x13,
	0xcf, 0x73, 0xcb, 0x37, 0xdc, 0xc0, 0x47, 0x29, 0xdc, 0xa1, 0x73, 0x2a, 0x27, 0x9f, 0x36, 0x72,
	0xdc, 0x70, 0xe0, 0x9c, 0x12, 0xed, 0x04, 0xe6, 0x13, 0xe7, 0x85, 0x81, 0xef, 0x85, 0x04, 0x3d,
	0x86, 0xb4, 0xa0, 0x80, 0xaa, 0xd4, 0x66, 0xeb, 0x85, 0xd6, 0x87, 0x13, 0x01, 0xf3, 0x74, 0x43,
	0xe6, 0xa0, 0x8f, 0x60, 0xce, 0x23, 0x27, 0xcc, 0x4c, 0x40, 0x9a, 0x11, 0x90, 0x4a, 0xdc, 0xdc,
	0x89, 0x61, 0x69, 0x6b, 0x50, 0x7e, 0x42, 0x44, 0xe1, 0xf8, 0x1e, 0x6f, 0x7b, 0xb6, 0xda, 0x4f,
	0x30, 0xbf, 0x4d, 0x09, 0x66, 0x24, 0x19, 0xbd, 0x05, 0x29, 0xee, 0x16, 0xa1, 0x53, 0x63, 0x14,
	0x29, 0x48, 0x87, 0xbb, 0x94, 0xd8, 0x4e, 0xc8, 0x08, 0x15, 0xa4, 0x4b, 0xc0, 0xcc, 0x1b, 0xe8,
	0x9a, 0x4b, 0x62, 0x5d, 0x86, 0xc5, 0x27, 0x84, 0x6d, 0x0f, 0x28, 0x25, 0x5e, 0x12, 0xb2, 0xf6,
	0xe7, 0x2c, 0xa8, 0x87, 0x81, 0x15, 0x41, 0x8b, 0xe6, 0x16, 0x23, 0x5c, 0x4d, 0x72, 0x88, 0xc3,
	0xcc, 0xb7, 0xb3, 0x97, 0x17, 0x2b, 0xb3, 0x27, 0x3f, 0x2b, 0x09, 0x32, 0xad, 0x5d, 0x67, 0xcc,
	0x4c, 0x32, 0x2e, 0x77, 0x9d, 0x3a, 0xf7, 0x63, 0xea, 0x70, 0xaa, 0xa6, 0xdb, 0xb9, 0xcb, 0x8b,
	0x95, 0x54, 0xe5, 0x8e, 0x9a, 0x19, 0x27, 0x51, 0x2a, 0x41, 0xa2, 0x65, 0xc8, 0x72, 0xe2, 0xf0,
	0xa6, 0xa6, 0x65, 0x53, 0xf9, 0x76, 0xd7, 0x42, 0x5f, 0xdd, 0xce, 0x24, 0xc1, 0xd8, 0x76, 0xf5,
	0xf2, 0x62, 0x65, 0x09, 0x16, 0x5e, 0x1c, 0x7d, 0xde, 0x78, 0x86, 0x1b, 0xa7, 0x1b, 0x8d, 0x2d,
	0xf3, 0xf9, 0xd9, 0xc6, 0x83, 0x87, 0x9f, 0x9d, 0xaf, 0xde, 0xc6, 0xb2, 0x2f, 0x6f, 0x61, 0x59,
	0x56, 0x1c, 0x74, 0xff, 0xf2, 0x62, 0xa5, 0x0a, 0x4b, 0x2f, 0xea, 0x47, 0xb8, 0x71, 0xfa, 0x9d,
	0xf5, 0xfc, 0xe3, 0xc6, 0xda, 0x7a, 0xbc, 0x5c, 0x3d, 0xa9, 0xdf, 0x60, 0xe1, 0x07, 0x50, 0x1a,
	0x31, 0x88, 0x83, 0xce, 0x09, 0xd0, 0xc5, 0x91, 0x71, 0xd7, 0x42, 0x47, 0xe3, 0xc4, 0xca, 0xff,
	0x07, 0x62, 0xc9, 0x1e, 0xff, 0xaa, 0x28, 0x63, 0x0c, 0x6b, 0xc0, 0xe2, 0x68, 0xa2, 0xbb, 0x3d,
	0xdf, 0x8b, 0xc7, 0xb9, 0x00, 0x69, 0xa7, 0x8f, 0x6d, 0x39, 0xca, 0xa2, 0x21, 0x37, 0xda, 0x1b,
	0x25, 0x8e, 0xef, 0xe0, 0x30, 0x1c, 0xfa, 0xd4, 0x8a, 0xe3, 0xab, 0x90, 0x0b, 0x22, 0x93, 0x9c,
	0xbe, 0x71, 0xb5, 0x47, 0xef, 0x43, 0xd1, 0x23, 0x43, 0xf3, 0xca, 0x2f, 0x9f, 0x5e, 0xc1, 0x23,
	0xc3, 0xf8, 0x14, 0xae, 0x74, 0xc9, 0x10, 0x93, 0x92, 0x80, 0x60, 0x16, 0xa9, 0xf4, 0x7c, 0x22,
	0xd2, 0x10, 0x8e, 0xf5, 0x07, 0x50, 0x4c, 0x5e, 0x0f, 0x55, 0xa0, 0xf8, 0x74, 0xe7, 0x69, 0x7b,
	0xc7, 0x38, 0x30, 0xbf, 0xde, 0xdf, 0xfb, 0xb6, 0x72, 0x07, 0x01, 0x64, 0x3a, 0x87, 0xed, 0xbd,
	0xdd, 0xed, 0x8a, 0xd2, 0xfa, 0x23, 0x0b, 0x05, 0x7e, 0xc1, 0x03, 0x42, 0x5f, 0x39, 0x3d, 0x82,
	0x7e, 0x53, 0x60, 0x8e, 0x0b, 0x41, 0x67, 0xd0, 0x75, 0x9d, 0x9e, 0x90, 0x03, 0xf4, 0x70, 0x62,
	0x3f, 0xc7, 0xa5, 0xa8, 0xda, 0x7a, 0x97, 0x14, 0xa9, 0x36, 0xda, 0xe2, 0x2f, 0x7f, 0xfd, 0xf3,
	0xfb, 0xcc, 0x1c, 0x2a, 0xe9, 0x42, 0x40, 0xf4, 0x40, 0x40, 0x40, 0x67, 0x90, 0x8d, 0xf4, 0x01,
	0xe9, 0x13, 0x4f, 0xbd, 0xae, 0x24, 0xd5, 0xe9, 0xd4, 0x40, 0x53, 0x45, 0x65, 0x84, 0x2a, 0x51,
	0xe5, 0xb3, 0x48, 0x7e, 0xce, 0xd1, 0x39, 0xc0, 0x48, 0x71, 0xd0, 0xe4, 0x5b, 0xdd, 0x90, 0xa7,
	0x69, 0x21, 0x2c, 0x08, 0x08, 0x65, 0x2d, 0x23, 0x21, 0x3c, 0x92, 0x02, 0xf5, 0xa3, 0xd0, 0xc6,
	0x84, 0xde, 0xa0, 0xcd, 0x69, 0x5a, 0x70, 0x53, 0xa0, 0xa6, 0x85, 0x51, 0x12, 0x30, 0xb2, 0x28,
	0x2d, 0x60, 0xf0, 0xd7, 0x30, 0x7f, 0x43, 0xd6, 0xd0, 0xd6, 0xe4, 0xb3, 0xde, 0x22, 0x85, 0xef,
	0x38, 0x90, 0x96, 0x7c, 0x0a, 0x7a, 0xc4, 0xcb, 0x47, 0xca, 0x3a, 0x7a, 0xa3, 0x40, 0xf9, 0x3a,
	0x2d, 0xa7, 0x68, 0xc9, 0xad, 0x3c, 0x9e, 0x16, 0x4b, 0xf4, 0x2c, 0x5b, 0x20, 0xb1, 0xf0, 0x0f,
	0x0e, 0x07, 0x72, 0x12, 0xe3, 0xb8, 0x22, 0xea, 0xb4, 0x38, 0xc6, 0xf4, 0xa1, 0xba, 0xd4, 0x94,
	0x7f, 0xbc, 0x66, 0xfc, 0xc7, 0x6b, 0xee, 0xf0, 0x3f, 0x9e, 0xb6, 0x24, 0x0a, 0x57, 0xaa, 0xe5,
	0xa8, 0x09, 0x51, 0x5a, 0x7b, 0xf3, 0xd9, 0xa7, 0x89, 0xaf, 0x66, 0x67, 0x54, 0x73, 0x0f, 0x77,
	0x75, 0x59, 0xb2, 0x11, 0x97, 0xe4, 0xbf, 0xca, 0xc7, 0x38, 0x70, 0xcc, 0xa0, 0xdb, 0xcd, 0x88,
	0xf3, 0x3f, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x47, 0xa9, 0x3b, 0x9f, 0xd7, 0x0a, 0x00, 0x00,
}
