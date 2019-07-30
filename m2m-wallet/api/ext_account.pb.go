// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ext_account.proto

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

type Money int32

const (
	Money_Ether Money = 0
)

var Money_name = map[int32]string{
	0: "Ether",
}

var Money_value = map[string]int32{
	"Ether": 0,
}

func (x Money) String() string {
	return proto.EnumName(Money_name, int32(x))
}

func (Money) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{0}
}

type ModifyMoneyAccountRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	MoneyAbbr            Money    `protobuf:"varint,2,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	CurrentAccount       string   `protobuf:"bytes,3,opt,name=current_account,json=currentAccount,proto3" json:"current_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyMoneyAccountRequest) Reset()         { *m = ModifyMoneyAccountRequest{} }
func (m *ModifyMoneyAccountRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyMoneyAccountRequest) ProtoMessage()    {}
func (*ModifyMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{0}
}

func (m *ModifyMoneyAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyMoneyAccountRequest.Unmarshal(m, b)
}
func (m *ModifyMoneyAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyMoneyAccountRequest.Marshal(b, m, deterministic)
}
func (m *ModifyMoneyAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyMoneyAccountRequest.Merge(m, src)
}
func (m *ModifyMoneyAccountRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyMoneyAccountRequest.Size(m)
}
func (m *ModifyMoneyAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyMoneyAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyMoneyAccountRequest proto.InternalMessageInfo

func (m *ModifyMoneyAccountRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *ModifyMoneyAccountRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

func (m *ModifyMoneyAccountRequest) GetCurrentAccount() string {
	if m != nil {
		return m.CurrentAccount
	}
	return ""
}

type ModifyMoneyAccountResponse struct {
	Status               bool             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ModifyMoneyAccountResponse) Reset()         { *m = ModifyMoneyAccountResponse{} }
func (m *ModifyMoneyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyMoneyAccountResponse) ProtoMessage()    {}
func (*ModifyMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{1}
}

func (m *ModifyMoneyAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyMoneyAccountResponse.Unmarshal(m, b)
}
func (m *ModifyMoneyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyMoneyAccountResponse.Marshal(b, m, deterministic)
}
func (m *ModifyMoneyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyMoneyAccountResponse.Merge(m, src)
}
func (m *ModifyMoneyAccountResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyMoneyAccountResponse.Size(m)
}
func (m *ModifyMoneyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyMoneyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyMoneyAccountResponse proto.InternalMessageInfo

func (m *ModifyMoneyAccountResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ModifyMoneyAccountResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetMoneyAccountChangeHistoryRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	MoneyAbbr            Money    `protobuf:"varint,4,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMoneyAccountChangeHistoryRequest) Reset()         { *m = GetMoneyAccountChangeHistoryRequest{} }
func (m *GetMoneyAccountChangeHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetMoneyAccountChangeHistoryRequest) ProtoMessage()    {}
func (*GetMoneyAccountChangeHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{2}
}

func (m *GetMoneyAccountChangeHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryRequest.Unmarshal(m, b)
}
func (m *GetMoneyAccountChangeHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetMoneyAccountChangeHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMoneyAccountChangeHistoryRequest.Merge(m, src)
}
func (m *GetMoneyAccountChangeHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryRequest.Size(m)
}
func (m *GetMoneyAccountChangeHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMoneyAccountChangeHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMoneyAccountChangeHistoryRequest proto.InternalMessageInfo

func (m *GetMoneyAccountChangeHistoryRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetMoneyAccountChangeHistoryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetMoneyAccountChangeHistoryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetMoneyAccountChangeHistoryRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

type MoneyAccountChangeHistory struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoneyAccountChangeHistory) Reset()         { *m = MoneyAccountChangeHistory{} }
func (m *MoneyAccountChangeHistory) String() string { return proto.CompactTextString(m) }
func (*MoneyAccountChangeHistory) ProtoMessage()    {}
func (*MoneyAccountChangeHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{3}
}

func (m *MoneyAccountChangeHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoneyAccountChangeHistory.Unmarshal(m, b)
}
func (m *MoneyAccountChangeHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoneyAccountChangeHistory.Marshal(b, m, deterministic)
}
func (m *MoneyAccountChangeHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoneyAccountChangeHistory.Merge(m, src)
}
func (m *MoneyAccountChangeHistory) XXX_Size() int {
	return xxx_messageInfo_MoneyAccountChangeHistory.Size(m)
}
func (m *MoneyAccountChangeHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_MoneyAccountChangeHistory.DiscardUnknown(m)
}

var xxx_messageInfo_MoneyAccountChangeHistory proto.InternalMessageInfo

func (m *MoneyAccountChangeHistory) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *MoneyAccountChangeHistory) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *MoneyAccountChangeHistory) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type GetMoneyAccountChangeHistoryResponse struct {
	Count                int64                        `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	ChangeHistory        []*MoneyAccountChangeHistory `protobuf:"bytes,2,rep,name=change_history,json=changeHistory,proto3" json:"change_history,omitempty"`
	UserProfile          *ProfileResponse             `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetMoneyAccountChangeHistoryResponse) Reset()         { *m = GetMoneyAccountChangeHistoryResponse{} }
func (m *GetMoneyAccountChangeHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetMoneyAccountChangeHistoryResponse) ProtoMessage()    {}
func (*GetMoneyAccountChangeHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{4}
}

func (m *GetMoneyAccountChangeHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryResponse.Unmarshal(m, b)
}
func (m *GetMoneyAccountChangeHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetMoneyAccountChangeHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMoneyAccountChangeHistoryResponse.Merge(m, src)
}
func (m *GetMoneyAccountChangeHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryResponse.Size(m)
}
func (m *GetMoneyAccountChangeHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMoneyAccountChangeHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMoneyAccountChangeHistoryResponse proto.InternalMessageInfo

func (m *GetMoneyAccountChangeHistoryResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetMoneyAccountChangeHistoryResponse) GetChangeHistory() []*MoneyAccountChangeHistory {
	if m != nil {
		return m.ChangeHistory
	}
	return nil
}

func (m *GetMoneyAccountChangeHistoryResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetActiveMoneyAccountRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	MoneyAbbr            Money    `protobuf:"varint,2,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetActiveMoneyAccountRequest) Reset()         { *m = GetActiveMoneyAccountRequest{} }
func (m *GetActiveMoneyAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetActiveMoneyAccountRequest) ProtoMessage()    {}
func (*GetActiveMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{5}
}

func (m *GetActiveMoneyAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetActiveMoneyAccountRequest.Unmarshal(m, b)
}
func (m *GetActiveMoneyAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetActiveMoneyAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetActiveMoneyAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetActiveMoneyAccountRequest.Merge(m, src)
}
func (m *GetActiveMoneyAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetActiveMoneyAccountRequest.Size(m)
}
func (m *GetActiveMoneyAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetActiveMoneyAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetActiveMoneyAccountRequest proto.InternalMessageInfo

func (m *GetActiveMoneyAccountRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetActiveMoneyAccountRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

type GetActiveMoneyAccountResponse struct {
	ActiveAccount        string           `protobuf:"bytes,1,opt,name=active_account,json=activeAccount,proto3" json:"active_account,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetActiveMoneyAccountResponse) Reset()         { *m = GetActiveMoneyAccountResponse{} }
func (m *GetActiveMoneyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*GetActiveMoneyAccountResponse) ProtoMessage()    {}
func (*GetActiveMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{6}
}

func (m *GetActiveMoneyAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetActiveMoneyAccountResponse.Unmarshal(m, b)
}
func (m *GetActiveMoneyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetActiveMoneyAccountResponse.Marshal(b, m, deterministic)
}
func (m *GetActiveMoneyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetActiveMoneyAccountResponse.Merge(m, src)
}
func (m *GetActiveMoneyAccountResponse) XXX_Size() int {
	return xxx_messageInfo_GetActiveMoneyAccountResponse.Size(m)
}
func (m *GetActiveMoneyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetActiveMoneyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetActiveMoneyAccountResponse proto.InternalMessageInfo

func (m *GetActiveMoneyAccountResponse) GetActiveAccount() string {
	if m != nil {
		return m.ActiveAccount
	}
	return ""
}

func (m *GetActiveMoneyAccountResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.Money", Money_name, Money_value)
	proto.RegisterType((*ModifyMoneyAccountRequest)(nil), "api.ModifyMoneyAccountRequest")
	proto.RegisterType((*ModifyMoneyAccountResponse)(nil), "api.ModifyMoneyAccountResponse")
	proto.RegisterType((*GetMoneyAccountChangeHistoryRequest)(nil), "api.GetMoneyAccountChangeHistoryRequest")
	proto.RegisterType((*MoneyAccountChangeHistory)(nil), "api.MoneyAccountChangeHistory")
	proto.RegisterType((*GetMoneyAccountChangeHistoryResponse)(nil), "api.GetMoneyAccountChangeHistoryResponse")
	proto.RegisterType((*GetActiveMoneyAccountRequest)(nil), "api.GetActiveMoneyAccountRequest")
	proto.RegisterType((*GetActiveMoneyAccountResponse)(nil), "api.GetActiveMoneyAccountResponse")
}

func init() { proto.RegisterFile("ext_account.proto", fileDescriptor_d6ddbf9312a3f483) }

var fileDescriptor_d6ddbf9312a3f483 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x6e, 0x12, 0x4d,
	0x14, 0xfe, 0x87, 0x2d, 0xe4, 0xe7, 0x50, 0x50, 0x27, 0xb4, 0xc1, 0x4d, 0xab, 0xb8, 0x6a, 0xa4,
	0x8d, 0x81, 0x88, 0x26, 0x4d, 0xbc, 0x23, 0xa6, 0xa9, 0x5e, 0x34, 0x31, 0xeb, 0x03, 0xac, 0xc3,
	0xee, 0x00, 0x93, 0xc0, 0xce, 0x3a, 0x3b, 0x34, 0x12, 0x63, 0x4c, 0xbc, 0xf0, 0xd2, 0x1b, 0xf5,
	0x21, 0x7c, 0x07, 0xdf, 0xc2, 0x37, 0x30, 0x3e, 0x88, 0xe1, 0xcc, 0x2c, 0x85, 0x16, 0x28, 0x31,
	0xf1, 0x6e, 0xcf, 0x99, 0x6f, 0xcf, 0xf9, 0xce, 0xf7, 0x9d, 0x19, 0xb8, 0xc1, 0xdf, 0xea, 0x80,
	0x85, 0xa1, 0x1c, 0xc7, 0xba, 0x99, 0x28, 0xa9, 0x25, 0x75, 0x58, 0x22, 0xdc, 0xbd, 0xbe, 0x94,
	0xfd, 0x21, 0x6f, 0xb1, 0x44, 0xb4, 0x58, 0x1c, 0x4b, 0xcd, 0xb4, 0x90, 0x71, 0x6a, 0x20, 0x6e,
	0x39, 0x51, 0xb2, 0x27, 0x86, 0xdc, 0x84, 0xde, 0x27, 0x02, 0x37, 0x4f, 0x65, 0x24, 0x7a, 0x93,
	0x53, 0x19, 0xf3, 0x49, 0xc7, 0x94, 0xf3, 0xf9, 0x9b, 0x31, 0x4f, 0x35, 0xdd, 0x81, 0x82, 0x54,
	0xfd, 0x40, 0x44, 0x35, 0x52, 0x27, 0x0d, 0xc7, 0xcf, 0x4b, 0xd5, 0x7f, 0x11, 0xd1, 0x03, 0x80,
	0xd1, 0x14, 0x1d, 0xb0, 0x6e, 0x57, 0xd5, 0x72, 0x75, 0xd2, 0xa8, 0xb4, 0xa1, 0xc9, 0x12, 0xd1,
	0xc4, 0x22, 0x7e, 0x11, 0x4f, 0x3b, 0xdd, 0xae, 0xa2, 0x0f, 0xe0, 0x5a, 0x38, 0x56, 0x8a, 0xc7,
	0x33, 0xaa, 0x35, 0xa7, 0x4e, 0x1a, 0x45, 0xbf, 0x62, 0xd3, 0xb6, 0xa3, 0x37, 0x02, 0x77, 0x19,
	0x8f, 0x34, 0x91, 0x71, 0xca, 0xe9, 0x2e, 0x14, 0x52, 0xcd, 0xf4, 0x38, 0x45, 0x22, 0xff, 0xfb,
	0x36, 0xa2, 0x47, 0xb0, 0x3d, 0x4e, 0xb9, 0x0a, 0xec, 0x50, 0xc8, 0xa5, 0xd4, 0xae, 0x22, 0x97,
	0x97, 0x26, 0x97, 0xd5, 0xf0, 0x4b, 0x53, 0xa4, 0x4d, 0x7a, 0xdf, 0x08, 0xdc, 0x3d, 0xe1, 0x7a,
	0xbe, 0xd9, 0xb3, 0x01, 0x8b, 0xfb, 0xfc, 0xb9, 0x48, 0xb5, 0x54, 0x93, 0x2b, 0x14, 0xd8, 0x85,
	0x82, 0xec, 0xf5, 0x52, 0xae, 0xb1, 0xa3, 0xe3, 0xdb, 0x88, 0x56, 0x21, 0x3f, 0x14, 0x23, 0x61,
	0x86, 0x74, 0x7c, 0x13, 0x5c, 0xd0, 0x6b, 0x6b, 0x8d, 0x5e, 0x5e, 0x6f, 0x6a, 0xc7, 0x0a, 0x4e,
	0x94, 0xc2, 0x16, 0x8b, 0x22, 0x85, 0x54, 0x8a, 0x3e, 0x7e, 0xcf, 0x29, 0x93, 0xc3, 0x6c, 0xa6,
	0xcc, 0x3e, 0x40, 0xa8, 0x38, 0xd3, 0x3c, 0x0a, 0x58, 0xa6, 0x79, 0xd1, 0x66, 0x3a, 0xda, 0xfb,
	0x41, 0xe0, 0xde, 0xfa, 0xf9, 0xad, 0xf2, 0x55, 0xc8, 0x1b, 0xdb, 0xec, 0xfc, 0x18, 0xd0, 0x63,
	0xa8, 0x84, 0x08, 0x0f, 0x06, 0x06, 0x5f, 0xcb, 0xd5, 0x9d, 0x46, 0xa9, 0x7d, 0xeb, 0x7c, 0xaa,
	0xa5, 0x55, 0xcb, 0xe1, 0xc2, 0x40, 0x17, 0xed, 0x73, 0x36, 0xb5, 0xef, 0x35, 0xec, 0x9d, 0x70,
	0xdd, 0x09, 0xb5, 0x38, 0xe3, 0xff, 0x64, 0x71, 0xbd, 0x0f, 0xb0, 0xbf, 0xa2, 0x83, 0x15, 0xe6,
	0x3e, 0x54, 0x18, 0x9e, 0xce, 0x16, 0xdb, 0xd8, 0x52, 0x36, 0x59, 0x0b, 0xff, 0xeb, 0x0d, 0x3d,
	0xa4, 0x90, 0xc7, 0xbe, 0xb4, 0x08, 0xf9, 0x63, 0x3d, 0xe0, 0xea, 0xfa, 0x7f, 0xed, 0x5f, 0x0e,
	0x6c, 0x63, 0xf2, 0x15, 0x57, 0x67, 0x22, 0xe4, 0xf4, 0x33, 0x01, 0x7a, 0xf9, 0xda, 0xd0, 0xcc,
	0x86, 0x15, 0xf7, 0xda, 0xbd, 0xbd, 0xf2, 0xdc, 0x30, 0xf1, 0x8e, 0x3e, 0xfe, 0xfc, 0xfd, 0x25,
	0xf7, 0xc8, 0x7d, 0x88, 0xaf, 0xc8, 0xdc, 0x43, 0xd3, 0x1a, 0xe1, 0x4f, 0xb3, 0xf0, 0xdd, 0xb9,
	0x9e, 0xef, 0x9f, 0x92, 0x43, 0xfa, 0x9d, 0xa0, 0x33, 0xc6, 0xf5, 0xf9, 0xd2, 0x99, 0xe5, 0x0d,
	0x6c, 0xbd, 0xc1, 0xd5, 0x73, 0x0f, 0x36, 0x40, 0x5a, 0xba, 0x4f, 0x90, 0x6e, 0x93, 0x5e, 0xa6,
	0xbb, 0xb8, 0xa5, 0x0b, 0x74, 0xe9, 0x57, 0x02, 0x3b, 0x4b, 0x3d, 0xa6, 0x77, 0xb2, 0xd6, 0x2b,
	0x37, 0xcc, 0xf5, 0xd6, 0x41, 0xae, 0xa4, 0xb5, 0xb8, 0x39, 0x0b, 0xb4, 0xba, 0x05, 0x7c, 0x99,
	0x1f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x6b, 0x5e, 0x8d, 0xe0, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MoneyServiceClient is the client API for MoneyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoneyServiceClient interface {
	ModifyMoneyAccount(ctx context.Context, in *ModifyMoneyAccountRequest, opts ...grpc.CallOption) (*ModifyMoneyAccountResponse, error)
	GetChangeMoneyAccountHistory(ctx context.Context, in *GetMoneyAccountChangeHistoryRequest, opts ...grpc.CallOption) (*GetMoneyAccountChangeHistoryResponse, error)
	GetActiveMoneyAccount(ctx context.Context, in *GetActiveMoneyAccountRequest, opts ...grpc.CallOption) (*GetActiveMoneyAccountResponse, error)
}

type moneyServiceClient struct {
	cc *grpc.ClientConn
}

func NewMoneyServiceClient(cc *grpc.ClientConn) MoneyServiceClient {
	return &moneyServiceClient{cc}
}

func (c *moneyServiceClient) ModifyMoneyAccount(ctx context.Context, in *ModifyMoneyAccountRequest, opts ...grpc.CallOption) (*ModifyMoneyAccountResponse, error) {
	out := new(ModifyMoneyAccountResponse)
	err := c.cc.Invoke(ctx, "/api.MoneyService/ModifyMoneyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moneyServiceClient) GetChangeMoneyAccountHistory(ctx context.Context, in *GetMoneyAccountChangeHistoryRequest, opts ...grpc.CallOption) (*GetMoneyAccountChangeHistoryResponse, error) {
	out := new(GetMoneyAccountChangeHistoryResponse)
	err := c.cc.Invoke(ctx, "/api.MoneyService/GetChangeMoneyAccountHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moneyServiceClient) GetActiveMoneyAccount(ctx context.Context, in *GetActiveMoneyAccountRequest, opts ...grpc.CallOption) (*GetActiveMoneyAccountResponse, error) {
	out := new(GetActiveMoneyAccountResponse)
	err := c.cc.Invoke(ctx, "/api.MoneyService/GetActiveMoneyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoneyServiceServer is the server API for MoneyService service.
type MoneyServiceServer interface {
	ModifyMoneyAccount(context.Context, *ModifyMoneyAccountRequest) (*ModifyMoneyAccountResponse, error)
	GetChangeMoneyAccountHistory(context.Context, *GetMoneyAccountChangeHistoryRequest) (*GetMoneyAccountChangeHistoryResponse, error)
	GetActiveMoneyAccount(context.Context, *GetActiveMoneyAccountRequest) (*GetActiveMoneyAccountResponse, error)
}

func RegisterMoneyServiceServer(s *grpc.Server, srv MoneyServiceServer) {
	s.RegisterService(&_MoneyService_serviceDesc, srv)
}

func _MoneyService_ModifyMoneyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyMoneyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyServiceServer).ModifyMoneyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MoneyService/ModifyMoneyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyServiceServer).ModifyMoneyAccount(ctx, req.(*ModifyMoneyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoneyService_GetChangeMoneyAccountHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoneyAccountChangeHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyServiceServer).GetChangeMoneyAccountHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MoneyService/GetChangeMoneyAccountHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyServiceServer).GetChangeMoneyAccountHistory(ctx, req.(*GetMoneyAccountChangeHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoneyService_GetActiveMoneyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveMoneyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyServiceServer).GetActiveMoneyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MoneyService/GetActiveMoneyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyServiceServer).GetActiveMoneyAccount(ctx, req.(*GetActiveMoneyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MoneyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MoneyService",
	HandlerType: (*MoneyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModifyMoneyAccount",
			Handler:    _MoneyService_ModifyMoneyAccount_Handler,
		},
		{
			MethodName: "GetChangeMoneyAccountHistory",
			Handler:    _MoneyService_GetChangeMoneyAccountHistory_Handler,
		},
		{
			MethodName: "GetActiveMoneyAccount",
			Handler:    _MoneyService_GetActiveMoneyAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ext_account.proto",
}
