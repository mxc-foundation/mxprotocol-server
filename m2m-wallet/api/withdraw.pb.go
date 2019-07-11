// Code generated by protoc-gen-go. DO NOT EDIT.
// source: withdraw.proto

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

type GetWithdrawFeeRequest struct {
	// type of crypto currency
	MoneyAbbr            Money    `protobuf:"varint,1,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWithdrawFeeRequest) Reset()         { *m = GetWithdrawFeeRequest{} }
func (m *GetWithdrawFeeRequest) String() string { return proto.CompactTextString(m) }
func (*GetWithdrawFeeRequest) ProtoMessage()    {}
func (*GetWithdrawFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{0}
}

func (m *GetWithdrawFeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWithdrawFeeRequest.Unmarshal(m, b)
}
func (m *GetWithdrawFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWithdrawFeeRequest.Marshal(b, m, deterministic)
}
func (m *GetWithdrawFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWithdrawFeeRequest.Merge(m, src)
}
func (m *GetWithdrawFeeRequest) XXX_Size() int {
	return xxx_messageInfo_GetWithdrawFeeRequest.Size(m)
}
func (m *GetWithdrawFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWithdrawFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWithdrawFeeRequest proto.InternalMessageInfo

func (m *GetWithdrawFeeRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

type GetWithdrawFeeResponse struct {
	// Withdraw object.
	WithdrawFee          float64          `protobuf:"fixed64,1,opt,name=withdrawFee,proto3" json:"withdrawFee,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetWithdrawFeeResponse) Reset()         { *m = GetWithdrawFeeResponse{} }
func (m *GetWithdrawFeeResponse) String() string { return proto.CompactTextString(m) }
func (*GetWithdrawFeeResponse) ProtoMessage()    {}
func (*GetWithdrawFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{1}
}

func (m *GetWithdrawFeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWithdrawFeeResponse.Unmarshal(m, b)
}
func (m *GetWithdrawFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWithdrawFeeResponse.Marshal(b, m, deterministic)
}
func (m *GetWithdrawFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWithdrawFeeResponse.Merge(m, src)
}
func (m *GetWithdrawFeeResponse) XXX_Size() int {
	return xxx_messageInfo_GetWithdrawFeeResponse.Size(m)
}
func (m *GetWithdrawFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWithdrawFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWithdrawFeeResponse proto.InternalMessageInfo

func (m *GetWithdrawFeeResponse) GetWithdrawFee() float64 {
	if m != nil {
		return m.WithdrawFee
	}
	return 0
}

func (m *GetWithdrawFeeResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type WithdrawReqRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	MoneyAbbr            Money    `protobuf:"varint,2,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawReqRequest) Reset()         { *m = WithdrawReqRequest{} }
func (m *WithdrawReqRequest) String() string { return proto.CompactTextString(m) }
func (*WithdrawReqRequest) ProtoMessage()    {}
func (*WithdrawReqRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{2}
}

func (m *WithdrawReqRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawReqRequest.Unmarshal(m, b)
}
func (m *WithdrawReqRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawReqRequest.Marshal(b, m, deterministic)
}
func (m *WithdrawReqRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawReqRequest.Merge(m, src)
}
func (m *WithdrawReqRequest) XXX_Size() int {
	return xxx_messageInfo_WithdrawReqRequest.Size(m)
}
func (m *WithdrawReqRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawReqRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawReqRequest proto.InternalMessageInfo

func (m *WithdrawReqRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *WithdrawReqRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

func (m *WithdrawReqRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type WithdrawReqResponse struct {
	Status               bool             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *WithdrawReqResponse) Reset()         { *m = WithdrawReqResponse{} }
func (m *WithdrawReqResponse) String() string { return proto.CompactTextString(m) }
func (*WithdrawReqResponse) ProtoMessage()    {}
func (*WithdrawReqResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{3}
}

func (m *WithdrawReqResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawReqResponse.Unmarshal(m, b)
}
func (m *WithdrawReqResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawReqResponse.Marshal(b, m, deterministic)
}
func (m *WithdrawReqResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawReqResponse.Merge(m, src)
}
func (m *WithdrawReqResponse) XXX_Size() int {
	return xxx_messageInfo_WithdrawReqResponse.Size(m)
}
func (m *WithdrawReqResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawReqResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawReqResponse proto.InternalMessageInfo

func (m *WithdrawReqResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *WithdrawReqResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetWithdrawHistoryRequest struct {
	OrgId                string   `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Offset               string   `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	MoneyAbbr            Money    `protobuf:"varint,4,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWithdrawHistoryRequest) Reset()         { *m = GetWithdrawHistoryRequest{} }
func (m *GetWithdrawHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetWithdrawHistoryRequest) ProtoMessage()    {}
func (*GetWithdrawHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{4}
}

func (m *GetWithdrawHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWithdrawHistoryRequest.Unmarshal(m, b)
}
func (m *GetWithdrawHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWithdrawHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetWithdrawHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWithdrawHistoryRequest.Merge(m, src)
}
func (m *GetWithdrawHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetWithdrawHistoryRequest.Size(m)
}
func (m *GetWithdrawHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWithdrawHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWithdrawHistoryRequest proto.InternalMessageInfo

func (m *GetWithdrawHistoryRequest) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *GetWithdrawHistoryRequest) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *GetWithdrawHistoryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetWithdrawHistoryRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

type WithdrawHistory struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	MoneyType            string   `protobuf:"bytes,3,opt,name=money_type,json=moneyType,proto3" json:"money_type,omitempty"`
	Amount               float64  `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawHistory) Reset()         { *m = WithdrawHistory{} }
func (m *WithdrawHistory) String() string { return proto.CompactTextString(m) }
func (*WithdrawHistory) ProtoMessage()    {}
func (*WithdrawHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{5}
}

func (m *WithdrawHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawHistory.Unmarshal(m, b)
}
func (m *WithdrawHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawHistory.Marshal(b, m, deterministic)
}
func (m *WithdrawHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawHistory.Merge(m, src)
}
func (m *WithdrawHistory) XXX_Size() int {
	return xxx_messageInfo_WithdrawHistory.Size(m)
}
func (m *WithdrawHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawHistory.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawHistory proto.InternalMessageInfo

func (m *WithdrawHistory) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *WithdrawHistory) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *WithdrawHistory) GetMoneyType() string {
	if m != nil {
		return m.MoneyType
	}
	return ""
}

func (m *WithdrawHistory) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *WithdrawHistory) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type GetWithdrawHistoryResponse struct {
	Count                int64              `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	WithdrawHistory      []*WithdrawHistory `protobuf:"bytes,2,rep,name=withdraw_history,json=withdrawHistory,proto3" json:"withdraw_history,omitempty"`
	UserProfile          *ProfileResponse   `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetWithdrawHistoryResponse) Reset()         { *m = GetWithdrawHistoryResponse{} }
func (m *GetWithdrawHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetWithdrawHistoryResponse) ProtoMessage()    {}
func (*GetWithdrawHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{6}
}

func (m *GetWithdrawHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWithdrawHistoryResponse.Unmarshal(m, b)
}
func (m *GetWithdrawHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWithdrawHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetWithdrawHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWithdrawHistoryResponse.Merge(m, src)
}
func (m *GetWithdrawHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetWithdrawHistoryResponse.Size(m)
}
func (m *GetWithdrawHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWithdrawHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWithdrawHistoryResponse proto.InternalMessageInfo

func (m *GetWithdrawHistoryResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetWithdrawHistoryResponse) GetWithdrawHistory() []*WithdrawHistory {
	if m != nil {
		return m.WithdrawHistory
	}
	return nil
}

func (m *GetWithdrawHistoryResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type ModifyWithdrawFeeRequest struct {
	MoneyAbbr            Money    `protobuf:"varint,1,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	WithdrawFee          float64  `protobuf:"fixed64,2,opt,name=withdraw_fee,json=withdrawFee,proto3" json:"withdraw_fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyWithdrawFeeRequest) Reset()         { *m = ModifyWithdrawFeeRequest{} }
func (m *ModifyWithdrawFeeRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyWithdrawFeeRequest) ProtoMessage()    {}
func (*ModifyWithdrawFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{7}
}

func (m *ModifyWithdrawFeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyWithdrawFeeRequest.Unmarshal(m, b)
}
func (m *ModifyWithdrawFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyWithdrawFeeRequest.Marshal(b, m, deterministic)
}
func (m *ModifyWithdrawFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyWithdrawFeeRequest.Merge(m, src)
}
func (m *ModifyWithdrawFeeRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyWithdrawFeeRequest.Size(m)
}
func (m *ModifyWithdrawFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyWithdrawFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyWithdrawFeeRequest proto.InternalMessageInfo

func (m *ModifyWithdrawFeeRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

func (m *ModifyWithdrawFeeRequest) GetWithdrawFee() float64 {
	if m != nil {
		return m.WithdrawFee
	}
	return 0
}

type ModifyWithdrawFeeResponse struct {
	Status               bool             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ModifyWithdrawFeeResponse) Reset()         { *m = ModifyWithdrawFeeResponse{} }
func (m *ModifyWithdrawFeeResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyWithdrawFeeResponse) ProtoMessage()    {}
func (*ModifyWithdrawFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{8}
}

func (m *ModifyWithdrawFeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyWithdrawFeeResponse.Unmarshal(m, b)
}
func (m *ModifyWithdrawFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyWithdrawFeeResponse.Marshal(b, m, deterministic)
}
func (m *ModifyWithdrawFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyWithdrawFeeResponse.Merge(m, src)
}
func (m *ModifyWithdrawFeeResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyWithdrawFeeResponse.Size(m)
}
func (m *ModifyWithdrawFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyWithdrawFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyWithdrawFeeResponse proto.InternalMessageInfo

func (m *ModifyWithdrawFeeResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ModifyWithdrawFeeResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterType((*GetWithdrawFeeRequest)(nil), "api.GetWithdrawFeeRequest")
	proto.RegisterType((*GetWithdrawFeeResponse)(nil), "api.GetWithdrawFeeResponse")
	proto.RegisterType((*WithdrawReqRequest)(nil), "api.WithdrawReqRequest")
	proto.RegisterType((*WithdrawReqResponse)(nil), "api.WithdrawReqResponse")
	proto.RegisterType((*GetWithdrawHistoryRequest)(nil), "api.GetWithdrawHistoryRequest")
	proto.RegisterType((*WithdrawHistory)(nil), "api.WithdrawHistory")
	proto.RegisterType((*GetWithdrawHistoryResponse)(nil), "api.GetWithdrawHistoryResponse")
	proto.RegisterType((*ModifyWithdrawFeeRequest)(nil), "api.ModifyWithdrawFeeRequest")
	proto.RegisterType((*ModifyWithdrawFeeResponse)(nil), "api.ModifyWithdrawFeeResponse")
}

func init() { proto.RegisterFile("withdraw.proto", fileDescriptor_b0dd7acb611886fa) }

var fileDescriptor_b0dd7acb611886fa = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0xed, 0x24, 0xfa, 0x65, 0x52, 0xa5, 0xbf, 0x0e, 0x6d, 0x70, 0x0d, 0x6d, 0x83, 0x05,
	0xa8, 0x14, 0x29, 0x15, 0xe1, 0x80, 0xd4, 0x0b, 0x2a, 0x07, 0xfe, 0x1c, 0x2a, 0x21, 0x83, 0xc4,
	0xd1, 0xda, 0x24, 0xeb, 0x64, 0xa5, 0xc4, 0xeb, 0xae, 0x37, 0x04, 0x0b, 0x71, 0x28, 0x12, 0xa7,
	0x1e, 0x79, 0x0d, 0xde, 0x86, 0x57, 0xe0, 0x41, 0x50, 0xd6, 0xeb, 0xfc, 0x71, 0x9c, 0xa8, 0x02,
	0x71, 0xcb, 0xcc, 0x78, 0xe7, 0x9b, 0x6f, 0xbe, 0x6f, 0x37, 0x50, 0x9f, 0x30, 0x39, 0xe8, 0x09,
	0x32, 0x69, 0x45, 0x82, 0x4b, 0x8e, 0x16, 0x89, 0x98, 0x73, 0xb7, 0xcf, 0x79, 0x7f, 0x48, 0x4f,
	0x49, 0xc4, 0x4e, 0x49, 0x18, 0x72, 0x49, 0x24, 0xe3, 0x61, 0x9c, 0x7e, 0xe2, 0xec, 0xd0, 0x4f,
	0xd2, 0x27, 0xdd, 0x2e, 0x1f, 0x87, 0x52, 0xa7, 0xea, 0x2c, 0x94, 0x54, 0x84, 0x64, 0x98, 0xc6,
	0xee, 0x0b, 0xd8, 0x7b, 0x45, 0xe5, 0x07, 0xdd, 0xfa, 0x25, 0xa5, 0x1e, 0xbd, 0x1c, 0xd3, 0x58,
	0xe2, 0x23, 0x80, 0x11, 0x0f, 0x69, 0xe2, 0x93, 0x4e, 0x47, 0xd8, 0x46, 0xd3, 0x38, 0xae, 0xb7,
	0xa1, 0x45, 0x22, 0xd6, 0xba, 0x98, 0xa6, 0xbd, 0xaa, 0xaa, 0x9e, 0x77, 0x3a, 0xc2, 0x8d, 0xa1,
	0x91, 0xef, 0x11, 0x47, 0x3c, 0x8c, 0x29, 0x36, 0xa1, 0x36, 0x99, 0xa7, 0x55, 0x17, 0xc3, 0x5b,
	0x4c, 0xe1, 0x33, 0xd8, 0x1a, 0xc7, 0x54, 0xf8, 0x91, 0xe0, 0x01, 0x1b, 0x52, 0xdb, 0x6c, 0x1a,
	0xc7, 0xb5, 0xf6, 0xae, 0x02, 0x7a, 0x9b, 0xe6, 0xb2, 0x6e, 0x5e, 0x6d, 0xfa, 0xa5, 0x4e, 0xba,
	0x21, 0x60, 0x86, 0xe8, 0xd1, 0xcb, 0x6c, 0xea, 0x3d, 0xa8, 0x70, 0xd1, 0xf7, 0x59, 0x4f, 0x61,
	0x59, 0x5e, 0x99, 0x8b, 0xfe, 0x9b, 0x5e, 0x8e, 0x8c, 0xb9, 0x81, 0x0c, 0x36, 0xa0, 0x42, 0x46,
	0xd3, 0x85, 0xd9, 0x96, 0x9a, 0x56, 0x47, 0x6e, 0x00, 0xb7, 0x96, 0xf0, 0x34, 0xc3, 0x06, 0x54,
	0x62, 0x49, 0xe4, 0x38, 0x56, 0x80, 0xff, 0x79, 0x3a, 0xfa, 0x73, 0x5e, 0xd7, 0x06, 0xec, 0x2f,
	0x6c, 0xf3, 0x35, 0x8b, 0x25, 0x17, 0x49, 0x31, 0xbf, 0x6a, 0xc6, 0xaf, 0x01, 0x15, 0x1e, 0x04,
	0x31, 0x95, 0x0a, 0xa7, 0xea, 0xe9, 0x08, 0x77, 0xa1, 0x3c, 0x64, 0x23, 0x96, 0x72, 0xb1, 0xbc,
	0x34, 0xc8, 0x6d, 0xa3, 0xb4, 0x49, 0xda, 0x6b, 0x03, 0xb6, 0x73, 0xa3, 0x20, 0x42, 0x29, 0x10,
	0x7c, 0xa4, 0x27, 0x50, 0xbf, 0xb1, 0x0e, 0xa6, 0xe4, 0x1a, 0xdc, 0x94, 0x1c, 0x0f, 0x32, 0x08,
	0x99, 0x44, 0x54, 0xa1, 0x57, 0x75, 0xdb, 0xf7, 0x49, 0x44, 0x17, 0x96, 0x5c, 0x5a, 0x5c, 0xf2,
	0xf4, 0x58, 0x57, 0x50, 0x22, 0x69, 0xcf, 0x27, 0xd2, 0x2e, 0xa7, 0xc7, 0x74, 0xe6, 0x5c, 0xba,
	0x3f, 0x0c, 0x70, 0x8a, 0x76, 0xa3, 0xb5, 0xd8, 0x85, 0xb2, 0xb2, 0x7a, 0xa6, 0xbd, 0x0a, 0xf0,
	0x39, 0xfc, 0x9f, 0x19, 0xce, 0x1f, 0xa4, 0x27, 0x6c, 0xb3, 0x69, 0xcd, 0xd4, 0xc8, 0x77, 0xdb,
	0x9e, 0xe4, 0xf8, 0xe6, 0xa5, 0xb4, 0x6e, 0x2a, 0xe5, 0x00, 0xec, 0x0b, 0xde, 0x63, 0x41, 0xf2,
	0x57, 0xd7, 0x0b, 0xef, 0xc1, 0xd6, 0x8c, 0x40, 0x40, 0x53, 0x2b, 0x2d, 0xdf, 0x22, 0x77, 0x08,
	0xfb, 0x05, 0x48, 0xff, 0xc8, 0xa2, 0xed, 0xab, 0xd2, 0xdc, 0x14, 0xef, 0xa8, 0xf8, 0xc8, 0xba,
	0x14, 0x13, 0xa8, 0x2f, 0xbf, 0x01, 0xe8, 0xa8, 0x46, 0x85, 0x8f, 0x8b, 0x73, 0xa7, 0xb0, 0x96,
	0x62, 0xb9, 0xad, 0xaf, 0x3f, 0x7f, 0x7d, 0x37, 0x8f, 0xf1, 0xa1, 0x7a, 0xd5, 0x32, 0x9a, 0xa7,
	0x9f, 0xe7, 0xfb, 0xfa, 0x32, 0xcb, 0x06, 0x94, 0x62, 0x1f, 0x6a, 0x0b, 0x37, 0x13, 0x6f, 0x2f,
	0xa9, 0x3a, 0x7f, 0x1b, 0x1c, 0x7b, 0xb5, 0xa0, 0x11, 0x1f, 0x28, 0xc4, 0x23, 0xd7, 0x59, 0x8f,
	0x78, 0x66, 0x9c, 0xe0, 0x95, 0x01, 0xb8, 0x6a, 0x3f, 0x3c, 0xcc, 0x93, 0x59, 0xbe, 0xb3, 0xce,
	0xd1, 0xda, 0xba, 0x86, 0x3f, 0x51, 0xf0, 0xf7, 0xd1, 0xdd, 0x40, 0x58, 0x3b, 0x17, 0xbf, 0x19,
	0xb0, 0xb3, 0x22, 0x35, 0x1e, 0x68, 0xe7, 0x14, 0x9b, 0xcd, 0x39, 0x5c, 0x57, 0xd6, 0x03, 0x3c,
	0x51, 0x03, 0x3c, 0x76, 0x6e, 0xb8, 0xf1, 0x33, 0xe3, 0xa4, 0x53, 0x51, 0x7f, 0x1f, 0x4f, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x0f, 0xb7, 0x9f, 0x96, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WithdrawServiceClient is the client API for WithdrawService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WithdrawServiceClient interface {
	// Get data for current withdraw fee
	GetWithdrawFee(ctx context.Context, in *GetWithdrawFeeRequest, opts ...grpc.CallOption) (*GetWithdrawFeeResponse, error)
	WithdrawReq(ctx context.Context, in *WithdrawReqRequest, opts ...grpc.CallOption) (*WithdrawReqResponse, error)
	GetWithdrawHistory(ctx context.Context, in *GetWithdrawHistoryRequest, opts ...grpc.CallOption) (*GetWithdrawHistoryResponse, error)
	ModifyWithdrawFee(ctx context.Context, in *ModifyWithdrawFeeRequest, opts ...grpc.CallOption) (*ModifyWithdrawFeeResponse, error)
}

type withdrawServiceClient struct {
	cc *grpc.ClientConn
}

func NewWithdrawServiceClient(cc *grpc.ClientConn) WithdrawServiceClient {
	return &withdrawServiceClient{cc}
}

func (c *withdrawServiceClient) GetWithdrawFee(ctx context.Context, in *GetWithdrawFeeRequest, opts ...grpc.CallOption) (*GetWithdrawFeeResponse, error) {
	out := new(GetWithdrawFeeResponse)
	err := c.cc.Invoke(ctx, "/api.WithdrawService/GetWithdrawFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawServiceClient) WithdrawReq(ctx context.Context, in *WithdrawReqRequest, opts ...grpc.CallOption) (*WithdrawReqResponse, error) {
	out := new(WithdrawReqResponse)
	err := c.cc.Invoke(ctx, "/api.WithdrawService/WithdrawReq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawServiceClient) GetWithdrawHistory(ctx context.Context, in *GetWithdrawHistoryRequest, opts ...grpc.CallOption) (*GetWithdrawHistoryResponse, error) {
	out := new(GetWithdrawHistoryResponse)
	err := c.cc.Invoke(ctx, "/api.WithdrawService/GetWithdrawHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawServiceClient) ModifyWithdrawFee(ctx context.Context, in *ModifyWithdrawFeeRequest, opts ...grpc.CallOption) (*ModifyWithdrawFeeResponse, error) {
	out := new(ModifyWithdrawFeeResponse)
	err := c.cc.Invoke(ctx, "/api.WithdrawService/ModifyWithdrawFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WithdrawServiceServer is the server API for WithdrawService service.
type WithdrawServiceServer interface {
	// Get data for current withdraw fee
	GetWithdrawFee(context.Context, *GetWithdrawFeeRequest) (*GetWithdrawFeeResponse, error)
	WithdrawReq(context.Context, *WithdrawReqRequest) (*WithdrawReqResponse, error)
	GetWithdrawHistory(context.Context, *GetWithdrawHistoryRequest) (*GetWithdrawHistoryResponse, error)
	ModifyWithdrawFee(context.Context, *ModifyWithdrawFeeRequest) (*ModifyWithdrawFeeResponse, error)
}

func RegisterWithdrawServiceServer(s *grpc.Server, srv WithdrawServiceServer) {
	s.RegisterService(&_WithdrawService_serviceDesc, srv)
}

func _WithdrawService_GetWithdrawFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServiceServer).GetWithdrawFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WithdrawService/GetWithdrawFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServiceServer).GetWithdrawFee(ctx, req.(*GetWithdrawFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WithdrawService_WithdrawReq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawReqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServiceServer).WithdrawReq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WithdrawService/WithdrawReq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServiceServer).WithdrawReq(ctx, req.(*WithdrawReqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WithdrawService_GetWithdrawHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServiceServer).GetWithdrawHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WithdrawService/GetWithdrawHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServiceServer).GetWithdrawHistory(ctx, req.(*GetWithdrawHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WithdrawService_ModifyWithdrawFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyWithdrawFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServiceServer).ModifyWithdrawFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WithdrawService/ModifyWithdrawFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServiceServer).ModifyWithdrawFee(ctx, req.(*ModifyWithdrawFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WithdrawService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WithdrawService",
	HandlerType: (*WithdrawServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWithdrawFee",
			Handler:    _WithdrawService_GetWithdrawFee_Handler,
		},
		{
			MethodName: "WithdrawReq",
			Handler:    _WithdrawService_WithdrawReq_Handler,
		},
		{
			MethodName: "GetWithdrawHistory",
			Handler:    _WithdrawService_GetWithdrawHistory_Handler,
		},
		{
			MethodName: "ModifyWithdrawFee",
			Handler:    _WithdrawService_ModifyWithdrawFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "withdraw.proto",
}
