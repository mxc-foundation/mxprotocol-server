// Code generated by protoc-gen-go. DO NOT EDIT.
// source: super_node.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetSuperNodeActiveMoneyAccountRequest struct {
	MoneyAbbr            Money    `protobuf:"varint,1,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSuperNodeActiveMoneyAccountRequest) Reset()         { *m = GetSuperNodeActiveMoneyAccountRequest{} }
func (m *GetSuperNodeActiveMoneyAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetSuperNodeActiveMoneyAccountRequest) ProtoMessage()    {}
func (*GetSuperNodeActiveMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e142dc5bc4ebd3, []int{0}
}

func (m *GetSuperNodeActiveMoneyAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest.Unmarshal(m, b)
}
func (m *GetSuperNodeActiveMoneyAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetSuperNodeActiveMoneyAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest.Merge(m, src)
}
func (m *GetSuperNodeActiveMoneyAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest.Size(m)
}
func (m *GetSuperNodeActiveMoneyAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest proto.InternalMessageInfo

func (m *GetSuperNodeActiveMoneyAccountRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

type GetSuperNodeActiveMoneyAccountResponse struct {
	SupernodeActiveAccount string           `protobuf:"bytes,1,opt,name=supernode_active_account,json=supernodeActiveAccount,proto3" json:"supernode_active_account,omitempty"`
	UserProfile            *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}         `json:"-"`
	XXX_unrecognized       []byte           `json:"-"`
	XXX_sizecache          int32            `json:"-"`
}

func (m *GetSuperNodeActiveMoneyAccountResponse) Reset() {
	*m = GetSuperNodeActiveMoneyAccountResponse{}
}
func (m *GetSuperNodeActiveMoneyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*GetSuperNodeActiveMoneyAccountResponse) ProtoMessage()    {}
func (*GetSuperNodeActiveMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e142dc5bc4ebd3, []int{1}
}

func (m *GetSuperNodeActiveMoneyAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse.Unmarshal(m, b)
}
func (m *GetSuperNodeActiveMoneyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse.Marshal(b, m, deterministic)
}
func (m *GetSuperNodeActiveMoneyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse.Merge(m, src)
}
func (m *GetSuperNodeActiveMoneyAccountResponse) XXX_Size() int {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse.Size(m)
}
func (m *GetSuperNodeActiveMoneyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse proto.InternalMessageInfo

func (m *GetSuperNodeActiveMoneyAccountResponse) GetSupernodeActiveAccount() string {
	if m != nil {
		return m.SupernodeActiveAccount
	}
	return ""
}

func (m *GetSuperNodeActiveMoneyAccountResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type AddSuperNodeMoneyAccountRequest struct {
	MoneyAbbr            Money    `protobuf:"varint,1,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	AccountAddr          string   `protobuf:"bytes,2,opt,name=account_addr,json=accountAddr,proto3" json:"account_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSuperNodeMoneyAccountRequest) Reset()         { *m = AddSuperNodeMoneyAccountRequest{} }
func (m *AddSuperNodeMoneyAccountRequest) String() string { return proto.CompactTextString(m) }
func (*AddSuperNodeMoneyAccountRequest) ProtoMessage()    {}
func (*AddSuperNodeMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e142dc5bc4ebd3, []int{2}
}

func (m *AddSuperNodeMoneyAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSuperNodeMoneyAccountRequest.Unmarshal(m, b)
}
func (m *AddSuperNodeMoneyAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSuperNodeMoneyAccountRequest.Marshal(b, m, deterministic)
}
func (m *AddSuperNodeMoneyAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSuperNodeMoneyAccountRequest.Merge(m, src)
}
func (m *AddSuperNodeMoneyAccountRequest) XXX_Size() int {
	return xxx_messageInfo_AddSuperNodeMoneyAccountRequest.Size(m)
}
func (m *AddSuperNodeMoneyAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSuperNodeMoneyAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddSuperNodeMoneyAccountRequest proto.InternalMessageInfo

func (m *AddSuperNodeMoneyAccountRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

func (m *AddSuperNodeMoneyAccountRequest) GetAccountAddr() string {
	if m != nil {
		return m.AccountAddr
	}
	return ""
}

type AddSuperNodeMoneyAccountResponse struct {
	Status               bool             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AddSuperNodeMoneyAccountResponse) Reset()         { *m = AddSuperNodeMoneyAccountResponse{} }
func (m *AddSuperNodeMoneyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*AddSuperNodeMoneyAccountResponse) ProtoMessage()    {}
func (*AddSuperNodeMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e142dc5bc4ebd3, []int{3}
}

func (m *AddSuperNodeMoneyAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSuperNodeMoneyAccountResponse.Unmarshal(m, b)
}
func (m *AddSuperNodeMoneyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSuperNodeMoneyAccountResponse.Marshal(b, m, deterministic)
}
func (m *AddSuperNodeMoneyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSuperNodeMoneyAccountResponse.Merge(m, src)
}
func (m *AddSuperNodeMoneyAccountResponse) XXX_Size() int {
	return xxx_messageInfo_AddSuperNodeMoneyAccountResponse.Size(m)
}
func (m *AddSuperNodeMoneyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSuperNodeMoneyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddSuperNodeMoneyAccountResponse proto.InternalMessageInfo

func (m *AddSuperNodeMoneyAccountResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *AddSuperNodeMoneyAccountResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterType((*GetSuperNodeActiveMoneyAccountRequest)(nil), "api.GetSuperNodeActiveMoneyAccountRequest")
	proto.RegisterType((*GetSuperNodeActiveMoneyAccountResponse)(nil), "api.GetSuperNodeActiveMoneyAccountResponse")
	proto.RegisterType((*AddSuperNodeMoneyAccountRequest)(nil), "api.AddSuperNodeMoneyAccountRequest")
	proto.RegisterType((*AddSuperNodeMoneyAccountResponse)(nil), "api.AddSuperNodeMoneyAccountResponse")
}

func init() { proto.RegisterFile("super_node.proto", fileDescriptor_02e142dc5bc4ebd3) }

var fileDescriptor_02e142dc5bc4ebd3 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x4a, 0xfb, 0x40,
	0x14, 0xc6, 0x49, 0xff, 0x50, 0xfe, 0x7d, 0xad, 0x52, 0x07, 0x29, 0x21, 0x88, 0xd6, 0x60, 0xb5,
	0x56, 0x6d, 0x20, 0x2e, 0x14, 0x77, 0x59, 0xb9, 0x52, 0x24, 0x3d, 0x40, 0x98, 0x64, 0xc6, 0x12,
	0xa8, 0x99, 0x38, 0x33, 0x29, 0x8a, 0xb8, 0xf1, 0x0a, 0x2e, 0x7b, 0x06, 0x4f, 0xe3, 0x05, 0x5c,
	0x78, 0x10, 0xc9, 0x64, 0x1a, 0xad, 0x50, 0x5b, 0x74, 0x99, 0x37, 0x5f, 0xbe, 0xef, 0xf7, 0xde,
	0x9b, 0x81, 0xa6, 0xc8, 0x52, 0xca, 0x83, 0x84, 0x11, 0xda, 0x4f, 0x39, 0x93, 0x0c, 0xfd, 0xc3,
	0x69, 0x6c, 0x6d, 0x0c, 0x19, 0x1b, 0x8e, 0xa8, 0x83, 0xd3, 0xd8, 0xc1, 0x49, 0xc2, 0x24, 0x96,
	0x31, 0x4b, 0x44, 0x21, 0xb1, 0xd6, 0xe8, 0x9d, 0x0c, 0x70, 0x14, 0xb1, 0x2c, 0x91, 0xba, 0xb4,
	0x92, 0x72, 0x76, 0x1d, 0x8f, 0xb4, 0x89, 0xed, 0x43, 0xe7, 0x9c, 0xca, 0x41, 0xee, 0x7d, 0xc9,
	0x08, 0xf5, 0x22, 0x19, 0x8f, 0xe9, 0x05, 0x4b, 0xe8, 0xbd, 0x57, 0xfc, 0xe6, 0xd3, 0xdb, 0x8c,
	0x0a, 0x89, 0xf6, 0x01, 0x6e, 0xf2, 0x72, 0x80, 0xc3, 0x90, 0x9b, 0x46, 0xdb, 0xe8, 0xae, 0xba,
	0xd0, 0xc7, 0x69, 0xdc, 0x57, 0x6a, 0xbf, 0xa6, 0x4e, 0xbd, 0x30, 0xe4, 0xf6, 0xc4, 0x80, 0xdd,
	0x45, 0xa6, 0x22, 0x65, 0x89, 0xa0, 0xe8, 0x14, 0x4c, 0xd5, 0x57, 0xde, 0x56, 0x80, 0x95, 0x6e,
	0xca, 0xab, 0x32, 0x6a, 0x7e, 0xab, 0x3c, 0x2f, 0x6c, 0xb4, 0x03, 0x3a, 0x81, 0x46, 0x26, 0x28,
	0x0f, 0x74, 0x3b, 0x66, 0xa5, 0x6d, 0x74, 0xeb, 0xee, 0xba, 0x22, 0xba, 0x2a, 0x6a, 0xd3, 0x14,
	0xbf, 0x9e, 0x2b, 0x75, 0xd1, 0x66, 0xb0, 0xe5, 0x11, 0x52, 0xc2, 0xfd, 0xad, 0x57, 0xb4, 0x0d,
	0x0d, 0xcd, 0x1b, 0x60, 0x42, 0xb8, 0xc2, 0xa8, 0xf9, 0x75, 0x5d, 0xf3, 0x08, 0xe1, 0xb6, 0x80,
	0xf6, 0xfc, 0x40, 0x3d, 0x87, 0x16, 0x54, 0x85, 0xc4, 0x32, 0x13, 0x2a, 0xed, 0xbf, 0xaf, 0xbf,
	0x7e, 0xdd, 0xa5, 0xfb, 0x56, 0x81, 0x66, 0x19, 0x39, 0xa0, 0x7c, 0x1c, 0x47, 0x14, 0xbd, 0x18,
	0xb0, 0xf9, 0xf3, 0x62, 0x50, 0x4f, 0x59, 0x2f, 0x75, 0x25, 0xac, 0x83, 0xa5, 0xb4, 0x05, 0x9d,
	0xed, 0x3e, 0xbd, 0xbe, 0x3f, 0x57, 0x0e, 0x51, 0x4f, 0x5d, 0xd5, 0x72, 0xa9, 0xce, 0xec, 0xd2,
	0x9d, 0x87, 0xcf, 0xc1, 0x3f, 0xa2, 0x89, 0x01, 0xe6, 0xbc, 0xd1, 0xa1, 0x1d, 0x95, 0xbe, 0x60,
	0x95, 0x56, 0x67, 0x81, 0x6a, 0x96, 0xce, 0xde, 0xfb, 0x46, 0xf7, 0x15, 0xc7, 0xc1, 0x84, 0x1c,
	0x69, 0xce, 0x33, 0xa3, 0x17, 0x56, 0xd5, 0x0b, 0x3a, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6a,
	0xc9, 0x5f, 0xb6, 0x9a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SuperNodeServiceClient is the client API for SuperNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SuperNodeServiceClient interface {
	GetSuperNodeActiveMoneyAccount(ctx context.Context, in *GetSuperNodeActiveMoneyAccountRequest, opts ...grpc.CallOption) (*GetSuperNodeActiveMoneyAccountResponse, error)
	AddSuperNodeMoneyAccount(ctx context.Context, in *AddSuperNodeMoneyAccountRequest, opts ...grpc.CallOption) (*AddSuperNodeMoneyAccountResponse, error)
}

type superNodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewSuperNodeServiceClient(cc *grpc.ClientConn) SuperNodeServiceClient {
	return &superNodeServiceClient{cc}
}

func (c *superNodeServiceClient) GetSuperNodeActiveMoneyAccount(ctx context.Context, in *GetSuperNodeActiveMoneyAccountRequest, opts ...grpc.CallOption) (*GetSuperNodeActiveMoneyAccountResponse, error) {
	out := new(GetSuperNodeActiveMoneyAccountResponse)
	err := c.cc.Invoke(ctx, "/api.SuperNodeService/GetSuperNodeActiveMoneyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superNodeServiceClient) AddSuperNodeMoneyAccount(ctx context.Context, in *AddSuperNodeMoneyAccountRequest, opts ...grpc.CallOption) (*AddSuperNodeMoneyAccountResponse, error) {
	out := new(AddSuperNodeMoneyAccountResponse)
	err := c.cc.Invoke(ctx, "/api.SuperNodeService/AddSuperNodeMoneyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuperNodeServiceServer is the server API for SuperNodeService service.
type SuperNodeServiceServer interface {
	GetSuperNodeActiveMoneyAccount(context.Context, *GetSuperNodeActiveMoneyAccountRequest) (*GetSuperNodeActiveMoneyAccountResponse, error)
	AddSuperNodeMoneyAccount(context.Context, *AddSuperNodeMoneyAccountRequest) (*AddSuperNodeMoneyAccountResponse, error)
}

func RegisterSuperNodeServiceServer(s *grpc.Server, srv SuperNodeServiceServer) {
	s.RegisterService(&_SuperNodeService_serviceDesc, srv)
}

func _SuperNodeService_GetSuperNodeActiveMoneyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSuperNodeActiveMoneyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperNodeServiceServer).GetSuperNodeActiveMoneyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SuperNodeService/GetSuperNodeActiveMoneyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperNodeServiceServer).GetSuperNodeActiveMoneyAccount(ctx, req.(*GetSuperNodeActiveMoneyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperNodeService_AddSuperNodeMoneyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSuperNodeMoneyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperNodeServiceServer).AddSuperNodeMoneyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SuperNodeService/AddSuperNodeMoneyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperNodeServiceServer).AddSuperNodeMoneyAccount(ctx, req.(*AddSuperNodeMoneyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SuperNodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SuperNodeService",
	HandlerType: (*SuperNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSuperNodeActiveMoneyAccount",
			Handler:    _SuperNodeService_GetSuperNodeActiveMoneyAccount_Handler,
		},
		{
			MethodName: "AddSuperNodeMoneyAccount",
			Handler:    _SuperNodeService_AddSuperNodeMoneyAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "super_node.proto",
}
