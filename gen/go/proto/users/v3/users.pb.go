// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/users/v3/users.proto

package v3pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/lyft/protoc-gen-validate/validate"
	v2 "github.com/nzoschke/gomesh/gen/go/proto/widgets/v2"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string               `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Widgets              []*v2.Widget         `protobuf:"bytes,6,rep,name=widgets,proto3" json:"widgets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bc6d07c35d93e81, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

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

func (m *User) GetWidgets() []*v2.Widget {
	if m != nil {
		return m.Widgets
	}
	return nil
}

type GetRequest struct {
	// Name
	//
	// Format "users/{user}"
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bc6d07c35d93e81, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
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

func init() {
	proto.RegisterType((*User)(nil), "gomesh.users.v3.User")
	proto.RegisterType((*GetRequest)(nil), "gomesh.users.v3.GetRequest")
}

func init() { proto.RegisterFile("proto/users/v3/users.proto", fileDescriptor_5bc6d07c35d93e81) }

var fileDescriptor_5bc6d07c35d93e81 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xc1, 0x4a, 0xeb, 0x40,
	0x00, 0x24, 0x4d, 0x5e, 0x1f, 0x6f, 0xf3, 0x40, 0x58, 0x2b, 0xb6, 0xa9, 0x68, 0x2c, 0x1e, 0xa2,
	0xd2, 0x8d, 0x24, 0x5e, 0x44, 0xbc, 0xe4, 0xd2, 0x9b, 0x94, 0x60, 0x11, 0x14, 0x5b, 0xb6, 0xed,
	0x1a, 0x03, 0x4d, 0x36, 0x66, 0xb7, 0x11, 0x15, 0x41, 0xfc, 0x05, 0x3f, 0xc4, 0xbb, 0x27, 0xff,
	0xc1, 0x5f, 0xf0, 0xe2, 0x5f, 0xc8, 0xee, 0x26, 0x28, 0xf5, 0x36, 0x99, 0x99, 0xcd, 0xcc, 0xce,
	0x02, 0x2b, 0xcb, 0x29, 0xa7, 0xee, 0x9c, 0x91, 0x9c, 0xb9, 0x85, 0xaf, 0x00, 0x92, 0x24, 0x5c,
	0x8a, 0x68, 0x42, 0xd8, 0x15, 0x52, 0x5c, 0xe1, 0x5b, 0x6b, 0x11, 0xa5, 0xd1, 0x8c, 0xb8, 0x38,
	0x8b, 0x5d, 0x9c, 0xa6, 0x94, 0x63, 0x1e, 0xd3, 0xb4, 0xb4, 0x5b, 0x1b, 0xa5, 0x2a, 0xbf, 0xc6,
	0xf3, 0x4b, 0x97, 0xc7, 0x09, 0x61, 0x1c, 0x27, 0x59, 0x69, 0x58, 0x57, 0x59, 0x37, 0xf1, 0x34,
	0x22, 0x9c, 0xb9, 0x85, 0x57, 0xc1, 0x52, 0x5f, 0x2d, 0xf0, 0x2c, 0x9e, 0x62, 0x4e, 0xdc, 0x0a,
	0x28, 0xa1, 0xf3, 0xa2, 0x01, 0x63, 0xc0, 0x48, 0x0e, 0x21, 0x30, 0x52, 0x9c, 0x90, 0xa6, 0x66,
	0x6b, 0xce, 0xbf, 0x50, 0x62, 0xb8, 0x09, 0xfe, 0x4f, 0x63, 0x96, 0xcd, 0xf0, 0xed, 0x48, 0x6a,
	0x35, 0xa9, 0x99, 0x25, 0x77, 0x2c, 0x2c, 0x87, 0xc0, 0x9c, 0xe4, 0x04, 0x73, 0x32, 0x12, 0x95,
	0x9a, 0xba, 0xad, 0x39, 0xa6, 0x67, 0x21, 0xd5, 0x17, 0x55, 0x7d, 0xd1, 0x49, 0xd5, 0x37, 0x04,
	0xca, 0x2e, 0x08, 0xe8, 0x83, 0xbf, 0x65, 0xcd, 0x66, 0xdd, 0xd6, 0x1d, 0xd3, 0x6b, 0xa1, 0x72,
	0x97, 0xaa, 0x7d, 0xe1, 0xa1, 0x53, 0x09, 0xc3, 0xca, 0xd9, 0x09, 0x00, 0xe8, 0x11, 0x1e, 0x92,
	0xeb, 0x39, 0x61, 0x1c, 0xee, 0xff, 0xac, 0x1d, 0xd8, 0xaf, 0x9f, 0x6f, 0x7a, 0x3b, 0x6f, 0x39,
	0x8f, 0x86, 0xd7, 0x18, 0xaa, 0xe5, 0xcf, 0x71, 0xf7, 0x6e, 0xaf, 0x7b, 0x80, 0x46, 0xdd, 0x8b,
	0xdd, 0x2d, 0x75, 0x31, 0x6f, 0x08, 0xfe, 0x88, 0x4b, 0x33, 0x38, 0x00, 0x7a, 0x8f, 0x70, 0xd8,
	0x46, 0x0b, 0xef, 0x81, 0xbe, 0x23, 0xac, 0x95, 0x5f, 0xa2, 0x38, 0xdb, 0xb1, 0x9e, 0xde, 0x3f,
	0x9e, 0x6b, 0x0d, 0x08, 0xc5, 0xdb, 0xde, 0x8b, 0xdf, 0x1e, 0xa9, 0xc0, 0x9d, 0x87, 0x60, 0x1b,
	0x2c, 0x4f, 0x68, 0xb2, 0x78, 0x2e, 0x00, 0x32, 0xb4, 0x2f, 0x46, 0xe9, 0x6b, 0x67, 0x46, 0xe1,
	0x67, 0xe3, 0x71, 0x5d, 0x6e, 0xe4, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xf3, 0x5e, 0x5e,
	0x2e, 0x02, 0x00, 0x00,
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
	// Get User
	//
	// Takes User name in path
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gomesh.users.v3.Users/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	// Get User
	//
	// Takes User name in path
	Get(context.Context, *GetRequest) (*User, error)
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
		FullMethod: "/gomesh.users.v3.Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gomesh.users.v3.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Users_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/users/v3/users.proto",
}
