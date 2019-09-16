// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inner_device.proto

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

type AppServerDeviceProfile struct {
	DevEui               string   `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastSeenAt           string   `protobuf:"bytes,3,opt,name=last_seen_at,json=lastSeenAt,proto3" json:"last_seen_at,omitempty"`
	ApplicationId        int64    `protobuf:"varint,4,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppServerDeviceProfile) Reset()         { *m = AppServerDeviceProfile{} }
func (m *AppServerDeviceProfile) String() string { return proto.CompactTextString(m) }
func (*AppServerDeviceProfile) ProtoMessage()    {}
func (*AppServerDeviceProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_06b99fe8ee14f928, []int{0}
}

func (m *AppServerDeviceProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppServerDeviceProfile.Unmarshal(m, b)
}
func (m *AppServerDeviceProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppServerDeviceProfile.Marshal(b, m, deterministic)
}
func (m *AppServerDeviceProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppServerDeviceProfile.Merge(m, src)
}
func (m *AppServerDeviceProfile) XXX_Size() int {
	return xxx_messageInfo_AppServerDeviceProfile.Size(m)
}
func (m *AppServerDeviceProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_AppServerDeviceProfile.DiscardUnknown(m)
}

var xxx_messageInfo_AppServerDeviceProfile proto.InternalMessageInfo

func (m *AppServerDeviceProfile) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

func (m *AppServerDeviceProfile) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *AppServerDeviceProfile) GetLastSeenAt() string {
	if m != nil {
		return m.LastSeenAt
	}
	return ""
}

func (m *AppServerDeviceProfile) GetApplicationId() int64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *AppServerDeviceProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddDeviceInM2MServerRequest struct {
	OrgId                int64                     `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	DevProfile           []*AppServerDeviceProfile `protobuf:"bytes,2,rep,name=dev_profile,json=devProfile,proto3" json:"dev_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AddDeviceInM2MServerRequest) Reset()         { *m = AddDeviceInM2MServerRequest{} }
func (m *AddDeviceInM2MServerRequest) String() string { return proto.CompactTextString(m) }
func (*AddDeviceInM2MServerRequest) ProtoMessage()    {}
func (*AddDeviceInM2MServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06b99fe8ee14f928, []int{1}
}

func (m *AddDeviceInM2MServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDeviceInM2MServerRequest.Unmarshal(m, b)
}
func (m *AddDeviceInM2MServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDeviceInM2MServerRequest.Marshal(b, m, deterministic)
}
func (m *AddDeviceInM2MServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDeviceInM2MServerRequest.Merge(m, src)
}
func (m *AddDeviceInM2MServerRequest) XXX_Size() int {
	return xxx_messageInfo_AddDeviceInM2MServerRequest.Size(m)
}
func (m *AddDeviceInM2MServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDeviceInM2MServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddDeviceInM2MServerRequest proto.InternalMessageInfo

func (m *AddDeviceInM2MServerRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *AddDeviceInM2MServerRequest) GetDevProfile() []*AppServerDeviceProfile {
	if m != nil {
		return m.DevProfile
	}
	return nil
}

type AddDeviceInM2MServerResponse struct {
	DevId                int64            `protobuf:"varint,1,opt,name=dev_id,json=devId,proto3" json:"dev_id,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AddDeviceInM2MServerResponse) Reset()         { *m = AddDeviceInM2MServerResponse{} }
func (m *AddDeviceInM2MServerResponse) String() string { return proto.CompactTextString(m) }
func (*AddDeviceInM2MServerResponse) ProtoMessage()    {}
func (*AddDeviceInM2MServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06b99fe8ee14f928, []int{2}
}

func (m *AddDeviceInM2MServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDeviceInM2MServerResponse.Unmarshal(m, b)
}
func (m *AddDeviceInM2MServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDeviceInM2MServerResponse.Marshal(b, m, deterministic)
}
func (m *AddDeviceInM2MServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDeviceInM2MServerResponse.Merge(m, src)
}
func (m *AddDeviceInM2MServerResponse) XXX_Size() int {
	return xxx_messageInfo_AddDeviceInM2MServerResponse.Size(m)
}
func (m *AddDeviceInM2MServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDeviceInM2MServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddDeviceInM2MServerResponse proto.InternalMessageInfo

func (m *AddDeviceInM2MServerResponse) GetDevId() int64 {
	if m != nil {
		return m.DevId
	}
	return 0
}

func (m *AddDeviceInM2MServerResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type DeleteDeviceInM2MServerRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	DevEui               string   `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDeviceInM2MServerRequest) Reset()         { *m = DeleteDeviceInM2MServerRequest{} }
func (m *DeleteDeviceInM2MServerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDeviceInM2MServerRequest) ProtoMessage()    {}
func (*DeleteDeviceInM2MServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06b99fe8ee14f928, []int{3}
}

func (m *DeleteDeviceInM2MServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeviceInM2MServerRequest.Unmarshal(m, b)
}
func (m *DeleteDeviceInM2MServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeviceInM2MServerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDeviceInM2MServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeviceInM2MServerRequest.Merge(m, src)
}
func (m *DeleteDeviceInM2MServerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDeviceInM2MServerRequest.Size(m)
}
func (m *DeleteDeviceInM2MServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeviceInM2MServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeviceInM2MServerRequest proto.InternalMessageInfo

func (m *DeleteDeviceInM2MServerRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *DeleteDeviceInM2MServerRequest) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

type DeleteDeviceInM2MServerResponse struct {
	Status               bool             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DeleteDeviceInM2MServerResponse) Reset()         { *m = DeleteDeviceInM2MServerResponse{} }
func (m *DeleteDeviceInM2MServerResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteDeviceInM2MServerResponse) ProtoMessage()    {}
func (*DeleteDeviceInM2MServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06b99fe8ee14f928, []int{4}
}

func (m *DeleteDeviceInM2MServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeviceInM2MServerResponse.Unmarshal(m, b)
}
func (m *DeleteDeviceInM2MServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeviceInM2MServerResponse.Marshal(b, m, deterministic)
}
func (m *DeleteDeviceInM2MServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeviceInM2MServerResponse.Merge(m, src)
}
func (m *DeleteDeviceInM2MServerResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteDeviceInM2MServerResponse.Size(m)
}
func (m *DeleteDeviceInM2MServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeviceInM2MServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeviceInM2MServerResponse proto.InternalMessageInfo

func (m *DeleteDeviceInM2MServerResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *DeleteDeviceInM2MServerResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterType((*AppServerDeviceProfile)(nil), "api.AppServerDeviceProfile")
	proto.RegisterType((*AddDeviceInM2MServerRequest)(nil), "api.AddDeviceInM2MServerRequest")
	proto.RegisterType((*AddDeviceInM2MServerResponse)(nil), "api.AddDeviceInM2MServerResponse")
	proto.RegisterType((*DeleteDeviceInM2MServerRequest)(nil), "api.DeleteDeviceInM2MServerRequest")
	proto.RegisterType((*DeleteDeviceInM2MServerResponse)(nil), "api.DeleteDeviceInM2MServerResponse")
}

func init() { proto.RegisterFile("inner_device.proto", fileDescriptor_06b99fe8ee14f928) }

var fileDescriptor_06b99fe8ee14f928 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x8f, 0x94, 0x30,
	0x18, 0xc6, 0xc3, 0xb0, 0x8b, 0xee, 0xcb, 0xae, 0x31, 0xcd, 0xba, 0x4b, 0x66, 0xfd, 0x83, 0xa8,
	0x09, 0xa7, 0x39, 0xe0, 0xc1, 0x8b, 0x17, 0x92, 0xf5, 0xc0, 0x81, 0x64, 0xc2, 0x1c, 0x3d, 0x90,
	0x4a, 0x5f, 0x27, 0x4d, 0xb0, 0xd4, 0xb6, 0xf0, 0xad, 0xfc, 0x38, 0x7e, 0x1f, 0xd3, 0xd2, 0x89,
	0x1c, 0x66, 0x26, 0xd1, 0x5b, 0x79, 0x79, 0x78, 0x7e, 0x7d, 0xde, 0x27, 0x00, 0xe1, 0x42, 0xa0,
	0x6a, 0x19, 0x4e, 0xbc, 0xc3, 0x8d, 0x54, 0x83, 0x19, 0x48, 0x48, 0x25, 0x5f, 0xdf, 0x48, 0x35,
	0x7c, 0xe7, 0xbd, 0x9f, 0x65, 0xbf, 0x02, 0xb8, 0x2b, 0xa5, 0xdc, 0xa1, 0x9a, 0x50, 0x3d, 0x3a,
	0xf5, 0x76, 0x16, 0x90, 0x7b, 0x78, 0xc2, 0x70, 0x6a, 0x71, 0xe4, 0x49, 0x90, 0x06, 0xf9, 0x55,
	0x13, 0x31, 0x9c, 0xbe, 0x8c, 0x9c, 0xbc, 0x02, 0xe8, 0x14, 0x52, 0x83, 0xac, 0xa5, 0x26, 0x59,
	0xb9, 0x77, 0x57, 0x7e, 0x52, 0x1a, 0x92, 0xc2, 0x75, 0x4f, 0xb5, 0x69, 0x35, 0xa2, 0xb0, 0x82,
	0xd0, 0x09, 0xc0, 0xce, 0x76, 0x88, 0xa2, 0x34, 0xe4, 0x03, 0x3c, 0xa3, 0x52, 0xf6, 0xbc, 0xa3,
	0x86, 0x0f, 0xa2, 0xe5, 0x2c, 0xb9, 0x48, 0x83, 0x3c, 0x6c, 0x6e, 0x16, 0xd3, 0x8a, 0x11, 0x02,
	0x17, 0x82, 0xfe, 0xc0, 0xe4, 0xd2, 0x19, 0xb8, 0x73, 0xa6, 0xe0, 0xa1, 0x64, 0x6c, 0xbe, 0x68,
	0x25, 0xea, 0xa2, 0x9e, 0x6f, 0xde, 0xe0, 0xcf, 0x11, 0xb5, 0x21, 0x2f, 0x20, 0x1a, 0xd4, 0xde,
	0x3a, 0x06, 0xce, 0xf1, 0x72, 0x50, 0xfb, 0x8a, 0x91, 0xcf, 0x10, 0xdb, 0x28, 0x3e, 0x7a, 0xb2,
	0x4a, 0xc3, 0x3c, 0x2e, 0x1e, 0x36, 0x54, 0xf2, 0xcd, 0xf1, 0xf0, 0x0d, 0x30, 0x9c, 0xfc, 0x39,
	0x13, 0xf0, 0xf2, 0x38, 0x53, 0xcb, 0x41, 0x68, 0xb4, 0x50, 0xeb, 0xfe, 0x17, 0xca, 0x70, 0xaa,
	0x18, 0xf9, 0x04, 0xd7, 0xa3, 0x46, 0xb5, 0xa0, 0x06, 0x79, 0x5c, 0xdc, 0x3a, 0xea, 0x01, 0xe3,
	0x2d, 0x9a, 0xd8, 0x2a, 0x0f, 0xbc, 0x2d, 0xbc, 0x7e, 0xc4, 0x1e, 0x0d, 0xfe, 0x6b, 0xcc, 0x45,
	0x63, 0xab, 0x65, 0x63, 0x99, 0x82, 0x37, 0x27, 0x1d, 0x7d, 0x88, 0x3b, 0x88, 0xb4, 0xa1, 0x66,
	0xd4, 0xce, 0xf2, 0x69, 0xe3, 0x9f, 0xfe, 0x3b, 0x45, 0xf1, 0x3b, 0x80, 0xe7, 0x75, 0x51, 0xcf,
	0x44, 0x0b, 0xe3, 0x1d, 0x92, 0xaf, 0x70, 0x7b, 0x6c, 0x95, 0x24, 0x9d, 0xbb, 0x38, 0xdd, 0xec,
	0xfa, 0xed, 0x19, 0x85, 0x8f, 0xc0, 0xe0, 0xfe, 0x44, 0x4a, 0xf2, 0xce, 0x7d, 0x7d, 0x7e, 0xab,
	0xeb, 0xf7, 0xe7, 0x45, 0x33, 0xe5, 0x5b, 0xe4, 0x7e, 0x9c, 0x8f, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0x17, 0xd1, 0xf8, 0x62, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// M2MDeviceServiceClient is the client API for M2MDeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type M2MDeviceServiceClient interface {
	AddDeviceInM2MServer(ctx context.Context, in *AddDeviceInM2MServerRequest, opts ...grpc.CallOption) (*AddDeviceInM2MServerResponse, error)
	DeleteDeviceInM2MServer(ctx context.Context, in *DeleteDeviceInM2MServerRequest, opts ...grpc.CallOption) (*DeleteDeviceInM2MServerResponse, error)
}

type m2MDeviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewM2MDeviceServiceClient(cc *grpc.ClientConn) M2MDeviceServiceClient {
	return &m2MDeviceServiceClient{cc}
}

func (c *m2MDeviceServiceClient) AddDeviceInM2MServer(ctx context.Context, in *AddDeviceInM2MServerRequest, opts ...grpc.CallOption) (*AddDeviceInM2MServerResponse, error) {
	out := new(AddDeviceInM2MServerResponse)
	err := c.cc.Invoke(ctx, "/api.M2MDeviceService/AddDeviceInM2MServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *m2MDeviceServiceClient) DeleteDeviceInM2MServer(ctx context.Context, in *DeleteDeviceInM2MServerRequest, opts ...grpc.CallOption) (*DeleteDeviceInM2MServerResponse, error) {
	out := new(DeleteDeviceInM2MServerResponse)
	err := c.cc.Invoke(ctx, "/api.M2MDeviceService/DeleteDeviceInM2MServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// M2MDeviceServiceServer is the server API for M2MDeviceService service.
type M2MDeviceServiceServer interface {
	AddDeviceInM2MServer(context.Context, *AddDeviceInM2MServerRequest) (*AddDeviceInM2MServerResponse, error)
	DeleteDeviceInM2MServer(context.Context, *DeleteDeviceInM2MServerRequest) (*DeleteDeviceInM2MServerResponse, error)
}

// UnimplementedM2MDeviceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedM2MDeviceServiceServer struct {
}

func (*UnimplementedM2MDeviceServiceServer) AddDeviceInM2MServer(ctx context.Context, req *AddDeviceInM2MServerRequest) (*AddDeviceInM2MServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDeviceInM2MServer not implemented")
}
func (*UnimplementedM2MDeviceServiceServer) DeleteDeviceInM2MServer(ctx context.Context, req *DeleteDeviceInM2MServerRequest) (*DeleteDeviceInM2MServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeviceInM2MServer not implemented")
}

func RegisterM2MDeviceServiceServer(s *grpc.Server, srv M2MDeviceServiceServer) {
	s.RegisterService(&_M2MDeviceService_serviceDesc, srv)
}

func _M2MDeviceService_AddDeviceInM2MServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceInM2MServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(M2MDeviceServiceServer).AddDeviceInM2MServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.M2MDeviceService/AddDeviceInM2MServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(M2MDeviceServiceServer).AddDeviceInM2MServer(ctx, req.(*AddDeviceInM2MServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _M2MDeviceService_DeleteDeviceInM2MServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceInM2MServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(M2MDeviceServiceServer).DeleteDeviceInM2MServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.M2MDeviceService/DeleteDeviceInM2MServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(M2MDeviceServiceServer).DeleteDeviceInM2MServer(ctx, req.(*DeleteDeviceInM2MServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _M2MDeviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.M2MDeviceService",
	HandlerType: (*M2MDeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDeviceInM2MServer",
			Handler:    _M2MDeviceService_AddDeviceInM2MServer_Handler,
		},
		{
			MethodName: "DeleteDeviceInM2MServer",
			Handler:    _M2MDeviceService_DeleteDeviceInM2MServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inner_device.proto",
}
