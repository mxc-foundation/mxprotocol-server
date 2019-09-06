// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GatewayMode int32

const (
	GatewayMode_gw_inactive              GatewayMode = 0
	GatewayMode_gw_free_gateways_limited GatewayMode = 1
	GatewayMode_gw_whole_network         GatewayMode = 2
)

var GatewayMode_name = map[int32]string{
	0: "gw_inactive",
	1: "gw_free_gateways_limited",
	2: "gw_whole_network",
}

var GatewayMode_value = map[string]int32{
	"gw_inactive":              0,
	"gw_free_gateways_limited": 1,
	"gw_whole_network":         2,
}

func (x GatewayMode) String() string {
	return proto.EnumName(GatewayMode_name, int32(x))
}

func (GatewayMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

type GetGatewayListRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGatewayListRequest) Reset()         { *m = GetGatewayListRequest{} }
func (m *GetGatewayListRequest) String() string { return proto.CompactTextString(m) }
func (*GetGatewayListRequest) ProtoMessage()    {}
func (*GetGatewayListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

func (m *GetGatewayListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayListRequest.Unmarshal(m, b)
}
func (m *GetGatewayListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayListRequest.Marshal(b, m, deterministic)
}
func (m *GetGatewayListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayListRequest.Merge(m, src)
}
func (m *GetGatewayListRequest) XXX_Size() int {
	return xxx_messageInfo_GetGatewayListRequest.Size(m)
}
func (m *GetGatewayListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayListRequest proto.InternalMessageInfo

func (m *GetGatewayListRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetGatewayListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetGatewayListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetGatewayListResponse struct {
	GwList               string           `protobuf:"bytes,1,opt,name=gw_list,json=gwList,proto3" json:"gw_list,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetGatewayListResponse) Reset()         { *m = GetGatewayListResponse{} }
func (m *GetGatewayListResponse) String() string { return proto.CompactTextString(m) }
func (*GetGatewayListResponse) ProtoMessage()    {}
func (*GetGatewayListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

func (m *GetGatewayListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayListResponse.Unmarshal(m, b)
}
func (m *GetGatewayListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayListResponse.Marshal(b, m, deterministic)
}
func (m *GetGatewayListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayListResponse.Merge(m, src)
}
func (m *GetGatewayListResponse) XXX_Size() int {
	return xxx_messageInfo_GetGatewayListResponse.Size(m)
}
func (m *GetGatewayListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayListResponse proto.InternalMessageInfo

func (m *GetGatewayListResponse) GetGwList() string {
	if m != nil {
		return m.GwList
	}
	return ""
}

func (m *GetGatewayListResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetGatewayProfileRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	GwId                 int64    `protobuf:"varint,2,opt,name=gw_id,json=gwId,proto3" json:"gw_id,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGatewayProfileRequest) Reset()         { *m = GetGatewayProfileRequest{} }
func (m *GetGatewayProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetGatewayProfileRequest) ProtoMessage()    {}
func (*GetGatewayProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}

func (m *GetGatewayProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayProfileRequest.Unmarshal(m, b)
}
func (m *GetGatewayProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetGatewayProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayProfileRequest.Merge(m, src)
}
func (m *GetGatewayProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetGatewayProfileRequest.Size(m)
}
func (m *GetGatewayProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayProfileRequest proto.InternalMessageInfo

func (m *GetGatewayProfileRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetGatewayProfileRequest) GetGwId() int64 {
	if m != nil {
		return m.GwId
	}
	return 0
}

func (m *GetGatewayProfileRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetGatewayProfileRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GatewayProfile struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FkGatewayLa          int64    `protobuf:"varint,2,opt,name=fk_gateway_la,json=fkGatewayLa,proto3" json:"fk_gateway_la,omitempty"`
	Mode                 string   `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayProfile) Reset()         { *m = GatewayProfile{} }
func (m *GatewayProfile) String() string { return proto.CompactTextString(m) }
func (*GatewayProfile) ProtoMessage()    {}
func (*GatewayProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{3}
}

func (m *GatewayProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayProfile.Unmarshal(m, b)
}
func (m *GatewayProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayProfile.Marshal(b, m, deterministic)
}
func (m *GatewayProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayProfile.Merge(m, src)
}
func (m *GatewayProfile) XXX_Size() int {
	return xxx_messageInfo_GatewayProfile.Size(m)
}
func (m *GatewayProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayProfile.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayProfile proto.InternalMessageInfo

func (m *GatewayProfile) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GatewayProfile) GetFkGatewayLa() int64 {
	if m != nil {
		return m.FkGatewayLa
	}
	return 0
}

func (m *GatewayProfile) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

type GetGatewayProfileResponse struct {
	GwProfile            []*GatewayProfile `protobuf:"bytes,1,rep,name=gw_profile,json=gwProfile,proto3" json:"gw_profile,omitempty"`
	UserProfile          *ProfileResponse  `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetGatewayProfileResponse) Reset()         { *m = GetGatewayProfileResponse{} }
func (m *GetGatewayProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetGatewayProfileResponse) ProtoMessage()    {}
func (*GetGatewayProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{4}
}

func (m *GetGatewayProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayProfileResponse.Unmarshal(m, b)
}
func (m *GetGatewayProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetGatewayProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayProfileResponse.Merge(m, src)
}
func (m *GetGatewayProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetGatewayProfileResponse.Size(m)
}
func (m *GetGatewayProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayProfileResponse proto.InternalMessageInfo

func (m *GetGatewayProfileResponse) GetGwProfile() []*GatewayProfile {
	if m != nil {
		return m.GwProfile
	}
	return nil
}

func (m *GetGatewayProfileResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type SetGatewayModeRequest struct {
	OrgId                int64       `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	GwId                 int64       `protobuf:"varint,2,opt,name=gw_id,json=gwId,proto3" json:"gw_id,omitempty"`
	GwMode               GatewayMode `protobuf:"varint,3,opt,name=gw_mode,json=gwMode,proto3,enum=api.GatewayMode" json:"gw_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SetGatewayModeRequest) Reset()         { *m = SetGatewayModeRequest{} }
func (m *SetGatewayModeRequest) String() string { return proto.CompactTextString(m) }
func (*SetGatewayModeRequest) ProtoMessage()    {}
func (*SetGatewayModeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{5}
}

func (m *SetGatewayModeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGatewayModeRequest.Unmarshal(m, b)
}
func (m *SetGatewayModeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGatewayModeRequest.Marshal(b, m, deterministic)
}
func (m *SetGatewayModeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGatewayModeRequest.Merge(m, src)
}
func (m *SetGatewayModeRequest) XXX_Size() int {
	return xxx_messageInfo_SetGatewayModeRequest.Size(m)
}
func (m *SetGatewayModeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGatewayModeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetGatewayModeRequest proto.InternalMessageInfo

func (m *SetGatewayModeRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *SetGatewayModeRequest) GetGwId() int64 {
	if m != nil {
		return m.GwId
	}
	return 0
}

func (m *SetGatewayModeRequest) GetGwMode() GatewayMode {
	if m != nil {
		return m.GwMode
	}
	return GatewayMode_gw_inactive
}

type SetGatewayModeResponse struct {
	Status               bool             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SetGatewayModeResponse) Reset()         { *m = SetGatewayModeResponse{} }
func (m *SetGatewayModeResponse) String() string { return proto.CompactTextString(m) }
func (*SetGatewayModeResponse) ProtoMessage()    {}
func (*SetGatewayModeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{6}
}

func (m *SetGatewayModeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGatewayModeResponse.Unmarshal(m, b)
}
func (m *SetGatewayModeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGatewayModeResponse.Marshal(b, m, deterministic)
}
func (m *SetGatewayModeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGatewayModeResponse.Merge(m, src)
}
func (m *SetGatewayModeResponse) XXX_Size() int {
	return xxx_messageInfo_SetGatewayModeResponse.Size(m)
}
func (m *SetGatewayModeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGatewayModeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetGatewayModeResponse proto.InternalMessageInfo

func (m *SetGatewayModeResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SetGatewayModeResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.GatewayMode", GatewayMode_name, GatewayMode_value)
	proto.RegisterType((*GetGatewayListRequest)(nil), "api.GetGatewayListRequest")
	proto.RegisterType((*GetGatewayListResponse)(nil), "api.GetGatewayListResponse")
	proto.RegisterType((*GetGatewayProfileRequest)(nil), "api.GetGatewayProfileRequest")
	proto.RegisterType((*GatewayProfile)(nil), "api.GatewayProfile")
	proto.RegisterType((*GetGatewayProfileResponse)(nil), "api.GetGatewayProfileResponse")
	proto.RegisterType((*SetGatewayModeRequest)(nil), "api.SetGatewayModeRequest")
	proto.RegisterType((*SetGatewayModeResponse)(nil), "api.SetGatewayModeResponse")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb1, 0xdd, 0x1a, 0x3a, 0x26, 0x21, 0x6c, 0x93, 0xd4, 0x98, 0x16, 0x45, 0x3e, 0x95,
	0x1e, 0x12, 0x29, 0x1c, 0x90, 0x78, 0x81, 0x08, 0x09, 0x24, 0xe4, 0x5e, 0x38, 0x20, 0x59, 0x4b,
	0xbd, 0xde, 0x2e, 0x71, 0xbd, 0xc6, 0xbb, 0xe9, 0x8a, 0x2b, 0x27, 0xee, 0x3c, 0x1a, 0x67, 0x6e,
	0x3c, 0x08, 0xf2, 0x7a, 0x6d, 0xe2, 0xc4, 0x20, 0xd4, 0x93, 0x3d, 0xb3, 0xe3, 0xf9, 0xe6, 0x9f,
	0xd9, 0x31, 0x0c, 0x28, 0x96, 0x44, 0xe1, 0x2f, 0xf3, 0xa2, 0xe4, 0x92, 0x23, 0x07, 0x17, 0x2c,
	0x38, 0xa5, 0x9c, 0xd3, 0x8c, 0x2c, 0x70, 0xc1, 0x16, 0x38, 0xcf, 0xb9, 0xc4, 0x92, 0xf1, 0x5c,
	0xd4, 0x21, 0xc1, 0xa0, 0x28, 0x79, 0xca, 0x32, 0x52, 0x9b, 0xe1, 0x07, 0x98, 0xac, 0x88, 0x5c,
	0xd5, 0x59, 0xde, 0x30, 0x21, 0x23, 0xf2, 0x79, 0x43, 0x84, 0x44, 0x13, 0x70, 0x79, 0x49, 0x63,
	0x96, 0xf8, 0xd6, 0xcc, 0x3a, 0x77, 0xa2, 0x43, 0x5e, 0xd2, 0xd7, 0x09, 0x9a, 0x82, 0xcb, 0xd3,
	0x54, 0x10, 0xe9, 0xdb, 0xda, 0x6d, 0x2c, 0x34, 0x86, 0xc3, 0x8c, 0xdd, 0x30, 0xe9, 0x3b, 0x75,
	0xb4, 0x36, 0xc2, 0x4f, 0x30, 0xdd, 0xcd, 0x2e, 0x0a, 0x9e, 0x0b, 0x82, 0x4e, 0xe0, 0x3e, 0x55,
	0x71, 0xc6, 0x84, 0xd4, 0xf9, 0x8f, 0x22, 0x97, 0xaa, 0x2a, 0x00, 0xbd, 0x84, 0x87, 0x1b, 0x41,
	0xca, 0xd8, 0x94, 0xa9, 0x31, 0xde, 0x72, 0x3c, 0xc7, 0x05, 0x9b, 0xbf, 0xab, 0x7d, 0x4d, 0x92,
	0xc8, 0xab, 0x22, 0x8d, 0x33, 0x94, 0xe0, 0xff, 0x61, 0xb5, 0x91, 0xff, 0x14, 0x73, 0x0c, 0x87,
	0x54, 0x55, 0xde, 0x5a, 0xcb, 0x01, 0x55, 0x1d, 0x85, 0x4e, 0xbf, 0xc2, 0x83, 0x6d, 0x85, 0xef,
	0x61, 0xd8, 0x45, 0xa2, 0x21, 0xd8, 0x2d, 0xc7, 0x66, 0x09, 0x0a, 0x61, 0x90, 0xae, 0x63, 0x33,
	0xa7, 0x38, 0xc3, 0x06, 0xe6, 0xa5, 0xeb, 0xa6, 0x2f, 0x18, 0x21, 0x38, 0xb8, 0xe1, 0x09, 0xd1,
	0xc4, 0xa3, 0x48, 0xbf, 0x87, 0xdf, 0x2c, 0x78, 0xd2, 0x23, 0xc8, 0xf4, 0x6f, 0x09, 0x40, 0x55,
	0xdb, 0x24, 0x6b, 0xe6, 0x9c, 0x7b, 0xcb, 0x63, 0xdd, 0xa4, 0x9d, 0x0f, 0x8e, 0xa8, 0x6a, 0x2a,
	0xbb, 0x73, 0x6b, 0x73, 0x98, 0x5c, 0xb6, 0x95, 0xbc, 0xe5, 0xc9, 0x9d, 0xfa, 0xfa, 0x5c, 0x4f,
	0xbc, 0x95, 0x39, 0x5c, 0x8e, 0xb6, 0xcb, 0xd5, 0x59, 0x5d, 0xaa, 0xaa, 0x67, 0xc8, 0x60, 0xba,
	0xcb, 0x33, 0xb2, 0xa7, 0xe0, 0x0a, 0x89, 0xe5, 0x46, 0x68, 0xe0, 0x83, 0xc8, 0x58, 0x7b, 0xd2,
	0x9c, 0xff, 0x94, 0x76, 0x11, 0x81, 0xb7, 0xc5, 0x41, 0x8f, 0xc0, 0xab, 0x2a, 0xcf, 0xf1, 0x95,
	0x64, 0xb7, 0x64, 0x74, 0x0f, 0x9d, 0x82, 0x4f, 0x55, 0x9c, 0x96, 0x84, 0x34, 0x23, 0x14, 0xb1,
	0x9e, 0x3c, 0x49, 0x46, 0x16, 0x1a, 0xc3, 0x88, 0xaa, 0x58, 0x5d, 0xf3, 0x8c, 0xc4, 0x39, 0x91,
	0x8a, 0x97, 0xeb, 0x91, 0xbd, 0xfc, 0x69, 0xb7, 0x97, 0xe2, 0x92, 0x94, 0xb7, 0xec, 0x8a, 0xa0,
	0x6b, 0x18, 0x76, 0x17, 0x01, 0x05, 0xb5, 0xfa, 0xbe, 0xdd, 0x0b, 0x9e, 0xf6, 0x9e, 0xd5, 0xe5,
	0x87, 0x67, 0x5f, 0x7f, 0xfc, 0xfa, 0x6e, 0x9f, 0xa0, 0x89, 0x5e, 0xf0, 0x84, 0x54, 0xf9, 0x17,
	0x94, 0xc8, 0x55, 0xbd, 0x3f, 0x12, 0x1e, 0xef, 0xdd, 0x1a, 0x74, 0xb6, 0x93, 0xb0, 0xbb, 0x1e,
	0xc1, 0xb3, 0xbf, 0x1d, 0x1b, 0xe4, 0x4c, 0x23, 0x03, 0xe4, 0xef, 0x21, 0x1b, 0xc0, 0x1a, 0x86,
	0xdd, 0x89, 0x19, 0x7d, 0xbd, 0xd7, 0xc6, 0xe8, 0xeb, 0x1f, 0x71, 0x03, 0x0b, 0x3b, 0xfa, 0x44,
	0x05, 0xab, 0xc2, 0x5e, 0x59, 0x17, 0x1f, 0x5d, 0xfd, 0xeb, 0x7a, 0xf1, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xc6, 0x6a, 0x55, 0xf2, 0xfd, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayServiceClient interface {
	GetGatewayList(ctx context.Context, in *GetGatewayListRequest, opts ...grpc.CallOption) (*GetGatewayListResponse, error)
	GetGatewayProfile(ctx context.Context, in *GetGatewayProfileRequest, opts ...grpc.CallOption) (*GetGatewayProfileResponse, error)
	SetGatewayMode(ctx context.Context, in *SetGatewayModeRequest, opts ...grpc.CallOption) (*SetGatewayModeResponse, error)
}

type gatewayServiceClient struct {
	cc *grpc.ClientConn
}

func NewGatewayServiceClient(cc *grpc.ClientConn) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) GetGatewayList(ctx context.Context, in *GetGatewayListRequest, opts ...grpc.CallOption) (*GetGatewayListResponse, error) {
	out := new(GetGatewayListResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayService/GetGatewayList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetGatewayProfile(ctx context.Context, in *GetGatewayProfileRequest, opts ...grpc.CallOption) (*GetGatewayProfileResponse, error) {
	out := new(GetGatewayProfileResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayService/GetGatewayProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) SetGatewayMode(ctx context.Context, in *SetGatewayModeRequest, opts ...grpc.CallOption) (*SetGatewayModeResponse, error) {
	out := new(SetGatewayModeResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayService/SetGatewayMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
type GatewayServiceServer interface {
	GetGatewayList(context.Context, *GetGatewayListRequest) (*GetGatewayListResponse, error)
	GetGatewayProfile(context.Context, *GetGatewayProfileRequest) (*GetGatewayProfileResponse, error)
	SetGatewayMode(context.Context, *SetGatewayModeRequest) (*SetGatewayModeResponse, error)
}

// UnimplementedGatewayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (*UnimplementedGatewayServiceServer) GetGatewayList(ctx context.Context, req *GetGatewayListRequest) (*GetGatewayListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayList not implemented")
}
func (*UnimplementedGatewayServiceServer) GetGatewayProfile(ctx context.Context, req *GetGatewayProfileRequest) (*GetGatewayProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayProfile not implemented")
}
func (*UnimplementedGatewayServiceServer) SetGatewayMode(ctx context.Context, req *SetGatewayModeRequest) (*SetGatewayModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGatewayMode not implemented")
}

func RegisterGatewayServiceServer(s *grpc.Server, srv GatewayServiceServer) {
	s.RegisterService(&_GatewayService_serviceDesc, srv)
}

func _GatewayService_GetGatewayList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetGatewayList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayService/GetGatewayList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetGatewayList(ctx, req.(*GetGatewayListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetGatewayProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetGatewayProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayService/GetGatewayProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetGatewayProfile(ctx, req.(*GetGatewayProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_SetGatewayMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGatewayModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).SetGatewayMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayService/SetGatewayMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).SetGatewayMode(ctx, req.(*SetGatewayModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGatewayList",
			Handler:    _GatewayService_GetGatewayList_Handler,
		},
		{
			MethodName: "GetGatewayProfile",
			Handler:    _GatewayService_GetGatewayProfile_Handler,
		},
		{
			MethodName: "SetGatewayMode",
			Handler:    _GatewayService_SetGatewayMode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}
