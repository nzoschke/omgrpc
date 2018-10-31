// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/users/v1/users.proto

package v1pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               string               `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string               `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_ec943581295b9478, []int{0}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type GetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_ec943581295b9478, []int{1}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateRequest struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_ec943581295b9478, []int{2}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "gomesh.users.v1.User")
	proto.RegisterType((*GetRequest)(nil), "gomesh.users.v1.GetRequest")
	proto.RegisterType((*CreateRequest)(nil), "gomesh.users.v1.CreateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*User, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gomesh.users.v1.Users/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gomesh.users.v1.Users/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Get(context.Context, *GetRequest) (*User, error)
	Create(context.Context, *CreateRequest) (*User, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gomesh.users.v1.Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gomesh.users.v1.Users/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gomesh.users.v1.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Users_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Users_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/users/v1/users.proto",
}

func init() { proto.RegisterFile("proto/users/v1/users.proto", fileDescriptor_users_ec943581295b9478) }

var fileDescriptor_users_ec943581295b9478 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xd9, 0x36, 0x2d, 0x38, 0xf1, 0x0f, 0xac, 0x28, 0x21, 0x82, 0xd6, 0x9e, 0xda, 0xcb,
	0x86, 0xd6, 0x93, 0x78, 0x10, 0xea, 0xa1, 0x37, 0x29, 0x41, 0x2f, 0x5e, 0x4a, 0xd2, 0x8c, 0x31,
	0xd0, 0x74, 0x63, 0x76, 0x13, 0xf0, 0xec, 0x37, 0xf1, 0x93, 0xca, 0xce, 0x36, 0x84, 0xb6, 0x7a,
	0x9b, 0x79, 0xef, 0xed, 0xf0, 0xdb, 0x07, 0x7e, 0x51, 0x4a, 0x2d, 0x83, 0x4a, 0x61, 0xa9, 0x82,
	0x7a, 0x62, 0x07, 0x41, 0x22, 0x3f, 0x4b, 0x65, 0x8e, 0xea, 0x43, 0x58, 0xad, 0x9e, 0xf8, 0x37,
	0xa9, 0x94, 0xe9, 0x1a, 0x03, 0xb2, 0xe3, 0xea, 0x3d, 0xd0, 0x59, 0x8e, 0x4a, 0x47, 0x79, 0x61,
	0x5f, 0x0c, 0x7f, 0x18, 0x38, 0xaf, 0x0a, 0x4b, 0x7e, 0x0a, 0x9d, 0x2c, 0xf1, 0xd8, 0x80, 0x8d,
	0x8e, 0xc2, 0x4e, 0x96, 0xf0, 0x4b, 0xe8, 0x17, 0x51, 0x89, 0x1b, 0xed, 0x75, 0x48, 0xdb, 0x6e,
	0x9c, 0x83, 0xb3, 0x89, 0x72, 0xf4, 0xba, 0xa4, 0xd2, 0xcc, 0x6f, 0xe1, 0x38, 0xc9, 0x54, 0xb1,
	0x8e, 0xbe, 0x96, 0xe4, 0x39, 0xe4, 0xb9, 0x5b, 0xed, 0xd9, 0x44, 0x1e, 0xc0, 0x5d, 0x95, 0x18,
	0x69, 0x5c, 0x1a, 0x02, 0xaf, 0x37, 0x60, 0x23, 0x77, 0xea, 0x0b, 0x8b, 0x27, 0x1a, 0x3c, 0xf1,
	0xd2, 0xe0, 0x85, 0x60, 0xe3, 0x46, 0x18, 0x0e, 0x00, 0xe6, 0xa8, 0x43, 0xfc, 0xac, 0x50, 0xb5,
	0x04, 0xac, 0x25, 0x18, 0x86, 0x70, 0xf2, 0x44, 0xf9, 0x26, 0xd4, 0xe2, 0xb3, 0x1d, 0xfc, 0x31,
	0x38, 0xa6, 0x1c, 0xfa, 0x94, 0x3b, 0xbd, 0x10, 0x7b, 0x85, 0x09, 0xd3, 0x45, 0x48, 0x91, 0xe9,
	0x37, 0x83, 0x9e, 0x59, 0x15, 0xbf, 0x87, 0xee, 0x1c, 0x35, 0xbf, 0x3a, 0x48, 0xb7, 0x54, 0xfe,
	0xdf, 0xa7, 0xf8, 0x23, 0xf4, 0x2d, 0x18, 0xbf, 0x3e, 0x08, 0xec, 0x10, 0xff, 0x73, 0x60, 0x36,
	0x86, 0xf3, 0x95, 0xcc, 0xf7, 0xbd, 0x19, 0x10, 0xd9, 0xc2, 0xf4, 0xb6, 0x60, 0x6f, 0x4e, 0x3d,
	0x29, 0xe2, 0xb8, 0x4f, 0x35, 0xde, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xa6, 0x64, 0xbb,
	0x22, 0x02, 0x00, 0x00,
}
