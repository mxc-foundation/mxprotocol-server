// Code generated by protoc-gen-go. DO NOT EDIT.
// source: staking.proto

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

type StakeRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakeRequest) Reset()         { *m = StakeRequest{} }
func (m *StakeRequest) String() string { return proto.CompactTextString(m) }
func (*StakeRequest) ProtoMessage()    {}
func (*StakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{0}
}

func (m *StakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakeRequest.Unmarshal(m, b)
}
func (m *StakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakeRequest.Marshal(b, m, deterministic)
}
func (m *StakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeRequest.Merge(m, src)
}
func (m *StakeRequest) XXX_Size() int {
	return xxx_messageInfo_StakeRequest.Size(m)
}
func (m *StakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StakeRequest proto.InternalMessageInfo

func (m *StakeRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *StakeRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type StakeResponse struct {
	Status               string           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StakeResponse) Reset()         { *m = StakeResponse{} }
func (m *StakeResponse) String() string { return proto.CompactTextString(m) }
func (*StakeResponse) ProtoMessage()    {}
func (*StakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{1}
}

func (m *StakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakeResponse.Unmarshal(m, b)
}
func (m *StakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakeResponse.Marshal(b, m, deterministic)
}
func (m *StakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeResponse.Merge(m, src)
}
func (m *StakeResponse) XXX_Size() int {
	return xxx_messageInfo_StakeResponse.Size(m)
}
func (m *StakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StakeResponse proto.InternalMessageInfo

func (m *StakeResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *StakeResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type UnstakeRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnstakeRequest) Reset()         { *m = UnstakeRequest{} }
func (m *UnstakeRequest) String() string { return proto.CompactTextString(m) }
func (*UnstakeRequest) ProtoMessage()    {}
func (*UnstakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{2}
}

func (m *UnstakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnstakeRequest.Unmarshal(m, b)
}
func (m *UnstakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnstakeRequest.Marshal(b, m, deterministic)
}
func (m *UnstakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnstakeRequest.Merge(m, src)
}
func (m *UnstakeRequest) XXX_Size() int {
	return xxx_messageInfo_UnstakeRequest.Size(m)
}
func (m *UnstakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnstakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnstakeRequest proto.InternalMessageInfo

func (m *UnstakeRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

type UnstakeResponse struct {
	Status               string           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UnstakeResponse) Reset()         { *m = UnstakeResponse{} }
func (m *UnstakeResponse) String() string { return proto.CompactTextString(m) }
func (*UnstakeResponse) ProtoMessage()    {}
func (*UnstakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{3}
}

func (m *UnstakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnstakeResponse.Unmarshal(m, b)
}
func (m *UnstakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnstakeResponse.Marshal(b, m, deterministic)
}
func (m *UnstakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnstakeResponse.Merge(m, src)
}
func (m *UnstakeResponse) XXX_Size() int {
	return xxx_messageInfo_UnstakeResponse.Size(m)
}
func (m *UnstakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnstakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnstakeResponse proto.InternalMessageInfo

func (m *UnstakeResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UnstakeResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetActiveStakesRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetActiveStakesRequest) Reset()         { *m = GetActiveStakesRequest{} }
func (m *GetActiveStakesRequest) String() string { return proto.CompactTextString(m) }
func (*GetActiveStakesRequest) ProtoMessage()    {}
func (*GetActiveStakesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{4}
}

func (m *GetActiveStakesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetActiveStakesRequest.Unmarshal(m, b)
}
func (m *GetActiveStakesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetActiveStakesRequest.Marshal(b, m, deterministic)
}
func (m *GetActiveStakesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetActiveStakesRequest.Merge(m, src)
}
func (m *GetActiveStakesRequest) XXX_Size() int {
	return xxx_messageInfo_GetActiveStakesRequest.Size(m)
}
func (m *GetActiveStakesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetActiveStakesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetActiveStakesRequest proto.InternalMessageInfo

func (m *GetActiveStakesRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

type ActiveStake struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FkWallet             int64    `protobuf:"varint,2,opt,name=FkWallet,proto3" json:"FkWallet,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	StakeStatus          string   `protobuf:"bytes,4,opt,name=StakeStatus,proto3" json:"StakeStatus,omitempty"`
	StartStakeTime       string   `protobuf:"bytes,5,opt,name=StartStakeTime,proto3" json:"StartStakeTime,omitempty"`
	UnstakeTime          string   `protobuf:"bytes,6,opt,name=UnstakeTime,proto3" json:"UnstakeTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActiveStake) Reset()         { *m = ActiveStake{} }
func (m *ActiveStake) String() string { return proto.CompactTextString(m) }
func (*ActiveStake) ProtoMessage()    {}
func (*ActiveStake) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{5}
}

func (m *ActiveStake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActiveStake.Unmarshal(m, b)
}
func (m *ActiveStake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActiveStake.Marshal(b, m, deterministic)
}
func (m *ActiveStake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActiveStake.Merge(m, src)
}
func (m *ActiveStake) XXX_Size() int {
	return xxx_messageInfo_ActiveStake.Size(m)
}
func (m *ActiveStake) XXX_DiscardUnknown() {
	xxx_messageInfo_ActiveStake.DiscardUnknown(m)
}

var xxx_messageInfo_ActiveStake proto.InternalMessageInfo

func (m *ActiveStake) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActiveStake) GetFkWallet() int64 {
	if m != nil {
		return m.FkWallet
	}
	return 0
}

func (m *ActiveStake) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ActiveStake) GetStakeStatus() string {
	if m != nil {
		return m.StakeStatus
	}
	return ""
}

func (m *ActiveStake) GetStartStakeTime() string {
	if m != nil {
		return m.StartStakeTime
	}
	return ""
}

func (m *ActiveStake) GetUnstakeTime() string {
	if m != nil {
		return m.UnstakeTime
	}
	return ""
}

type GetActiveStakesResponse struct {
	ActStake             *ActiveStake     `protobuf:"bytes,1,opt,name=act_stake,json=actStake,proto3" json:"act_stake,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetActiveStakesResponse) Reset()         { *m = GetActiveStakesResponse{} }
func (m *GetActiveStakesResponse) String() string { return proto.CompactTextString(m) }
func (*GetActiveStakesResponse) ProtoMessage()    {}
func (*GetActiveStakesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{6}
}

func (m *GetActiveStakesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetActiveStakesResponse.Unmarshal(m, b)
}
func (m *GetActiveStakesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetActiveStakesResponse.Marshal(b, m, deterministic)
}
func (m *GetActiveStakesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetActiveStakesResponse.Merge(m, src)
}
func (m *GetActiveStakesResponse) XXX_Size() int {
	return xxx_messageInfo_GetActiveStakesResponse.Size(m)
}
func (m *GetActiveStakesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetActiveStakesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetActiveStakesResponse proto.InternalMessageInfo

func (m *GetActiveStakesResponse) GetActStake() *ActiveStake {
	if m != nil {
		return m.ActStake
	}
	return nil
}

func (m *GetActiveStakesResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type StakingHistoryRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakingHistoryRequest) Reset()         { *m = StakingHistoryRequest{} }
func (m *StakingHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*StakingHistoryRequest) ProtoMessage()    {}
func (*StakingHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{7}
}

func (m *StakingHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakingHistoryRequest.Unmarshal(m, b)
}
func (m *StakingHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakingHistoryRequest.Marshal(b, m, deterministic)
}
func (m *StakingHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingHistoryRequest.Merge(m, src)
}
func (m *StakingHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_StakingHistoryRequest.Size(m)
}
func (m *StakingHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StakingHistoryRequest proto.InternalMessageInfo

func (m *StakingHistoryRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *StakingHistoryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *StakingHistoryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetStakingHistory struct {
	StakeAmount          float64  `protobuf:"fixed64,1,opt,name=stake_amount,json=stakeAmount,proto3" json:"stake_amount,omitempty"`
	Start                string   `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  string   `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	RevMonth             string   `protobuf:"bytes,4,opt,name=rev_month,json=revMonth,proto3" json:"rev_month,omitempty"`
	NetworkIncome        float64  `protobuf:"fixed64,5,opt,name=network_income,json=networkIncome,proto3" json:"network_income,omitempty"`
	MonthlyRate          float64  `protobuf:"fixed64,6,opt,name=monthly_rate,json=monthlyRate,proto3" json:"monthly_rate,omitempty"`
	Revenue              float64  `protobuf:"fixed64,7,opt,name=revenue,proto3" json:"revenue,omitempty"`
	Balance              float64  `protobuf:"fixed64,8,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStakingHistory) Reset()         { *m = GetStakingHistory{} }
func (m *GetStakingHistory) String() string { return proto.CompactTextString(m) }
func (*GetStakingHistory) ProtoMessage()    {}
func (*GetStakingHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{8}
}

func (m *GetStakingHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStakingHistory.Unmarshal(m, b)
}
func (m *GetStakingHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStakingHistory.Marshal(b, m, deterministic)
}
func (m *GetStakingHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStakingHistory.Merge(m, src)
}
func (m *GetStakingHistory) XXX_Size() int {
	return xxx_messageInfo_GetStakingHistory.Size(m)
}
func (m *GetStakingHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStakingHistory.DiscardUnknown(m)
}

var xxx_messageInfo_GetStakingHistory proto.InternalMessageInfo

func (m *GetStakingHistory) GetStakeAmount() float64 {
	if m != nil {
		return m.StakeAmount
	}
	return 0
}

func (m *GetStakingHistory) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *GetStakingHistory) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *GetStakingHistory) GetRevMonth() string {
	if m != nil {
		return m.RevMonth
	}
	return ""
}

func (m *GetStakingHistory) GetNetworkIncome() float64 {
	if m != nil {
		return m.NetworkIncome
	}
	return 0
}

func (m *GetStakingHistory) GetMonthlyRate() float64 {
	if m != nil {
		return m.MonthlyRate
	}
	return 0
}

func (m *GetStakingHistory) GetRevenue() float64 {
	if m != nil {
		return m.Revenue
	}
	return 0
}

func (m *GetStakingHistory) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type StakingHistoryResponse struct {
	UserProfile          *ProfileResponse     `protobuf:"bytes,1,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	StakingHist          []*GetStakingHistory `protobuf:"bytes,2,rep,name=staking_hist,json=stakingHist,proto3" json:"staking_hist,omitempty"`
	Count                int64                `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StakingHistoryResponse) Reset()         { *m = StakingHistoryResponse{} }
func (m *StakingHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*StakingHistoryResponse) ProtoMessage()    {}
func (*StakingHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{9}
}

func (m *StakingHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakingHistoryResponse.Unmarshal(m, b)
}
func (m *StakingHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakingHistoryResponse.Marshal(b, m, deterministic)
}
func (m *StakingHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingHistoryResponse.Merge(m, src)
}
func (m *StakingHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_StakingHistoryResponse.Size(m)
}
func (m *StakingHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StakingHistoryResponse proto.InternalMessageInfo

func (m *StakingHistoryResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func (m *StakingHistoryResponse) GetStakingHist() []*GetStakingHistory {
	if m != nil {
		return m.StakingHist
	}
	return nil
}

func (m *StakingHistoryResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*StakeRequest)(nil), "api.StakeRequest")
	proto.RegisterType((*StakeResponse)(nil), "api.StakeResponse")
	proto.RegisterType((*UnstakeRequest)(nil), "api.UnstakeRequest")
	proto.RegisterType((*UnstakeResponse)(nil), "api.UnstakeResponse")
	proto.RegisterType((*GetActiveStakesRequest)(nil), "api.GetActiveStakesRequest")
	proto.RegisterType((*ActiveStake)(nil), "api.ActiveStake")
	proto.RegisterType((*GetActiveStakesResponse)(nil), "api.GetActiveStakesResponse")
	proto.RegisterType((*StakingHistoryRequest)(nil), "api.StakingHistoryRequest")
	proto.RegisterType((*GetStakingHistory)(nil), "api.GetStakingHistory")
	proto.RegisterType((*StakingHistoryResponse)(nil), "api.StakingHistoryResponse")
}

func init() { proto.RegisterFile("staking.proto", fileDescriptor_289e7c8aea278311) }

var fileDescriptor_289e7c8aea278311 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0x13, 0x3b,
	0x14, 0xd6, 0x64, 0x6e, 0xd2, 0xe4, 0x4c, 0x93, 0xb6, 0xbe, 0x69, 0x6e, 0x34, 0x69, 0x75, 0xcb,
	0x40, 0xa1, 0x42, 0xa2, 0x91, 0xc2, 0x02, 0xb1, 0x60, 0xd1, 0x0d, 0x25, 0x0b, 0x24, 0xe4, 0x80,
	0x58, 0x80, 0x34, 0xb8, 0x13, 0x37, 0xb5, 0x32, 0xb1, 0xc3, 0xd8, 0x09, 0x54, 0x88, 0x05, 0xbc,
	0x02, 0x6f, 0xc0, 0x1b, 0xf0, 0x00, 0x3c, 0x05, 0xaf, 0xc0, 0x4b, 0xb0, 0x43, 0x73, 0xec, 0x44,
	0x69, 0x1a, 0xa9, 0x15, 0x12, 0xbb, 0x7c, 0x9f, 0x3f, 0x9f, 0x9f, 0xef, 0x9c, 0x71, 0xa0, 0xaa,
	0x0d, 0x1b, 0x0a, 0x39, 0x38, 0x1c, 0x67, 0xca, 0x28, 0xe2, 0xb3, 0xb1, 0x08, 0x77, 0x06, 0x4a,
	0x0d, 0x52, 0xde, 0x66, 0x63, 0xd1, 0x66, 0x52, 0x2a, 0xc3, 0x8c, 0x50, 0x52, 0x5b, 0x49, 0x58,
	0x1d, 0x67, 0xea, 0x54, 0xa4, 0xdc, 0xc2, 0xe8, 0x11, 0xac, 0xf7, 0x0c, 0x1b, 0x72, 0xca, 0xdf,
	0x4e, 0xb8, 0x36, 0x64, 0x1b, 0x4a, 0x2a, 0x1b, 0xc4, 0xa2, 0xdf, 0xf4, 0xf6, 0xbc, 0x03, 0x9f,
	0x16, 0x55, 0x36, 0xe8, 0xf6, 0x49, 0x03, 0x4a, 0x6c, 0xa4, 0x26, 0xd2, 0x34, 0x0b, 0x7b, 0xde,
	0x81, 0x47, 0x1d, 0x8a, 0xde, 0x40, 0xd5, 0x5d, 0xd7, 0x63, 0x25, 0x35, 0xcf, 0x85, 0xda, 0x30,
	0x33, 0xd1, 0x78, 0xbf, 0x42, 0x1d, 0x22, 0x0f, 0x60, 0x7d, 0xa2, 0x79, 0x16, 0xbb, 0xec, 0x18,
	0x26, 0xe8, 0xd4, 0x0f, 0xd9, 0x58, 0x1c, 0x3e, 0xb3, 0xdc, 0x2c, 0x06, 0x0d, 0x72, 0xa5, 0x23,
	0xa3, 0x3b, 0x50, 0x7b, 0x21, 0xf5, 0xd5, 0x25, 0x46, 0x27, 0xb0, 0x31, 0x17, 0xfe, 0xad, 0x62,
	0xda, 0xd0, 0x38, 0xe6, 0xe6, 0x28, 0x31, 0x62, 0xca, 0xb1, 0x6f, 0x7d, 0x45, 0x51, 0xdf, 0x3d,
	0x08, 0x16, 0xe4, 0xa4, 0x06, 0x85, 0xee, 0x4c, 0x52, 0xe8, 0xf6, 0x49, 0x08, 0xe5, 0xc7, 0xc3,
	0x97, 0x2c, 0x4d, 0xb9, 0x75, 0xd6, 0xa7, 0x73, 0x9c, 0x57, 0x7f, 0x64, 0x3d, 0xf7, 0xad, 0xe7,
	0x16, 0x91, 0x3d, 0x08, 0x30, 0x58, 0xcf, 0xb6, 0xf6, 0x0f, 0xb6, 0xb6, 0x48, 0x91, 0xdb, 0x50,
	0xeb, 0x19, 0x96, 0x19, 0xe4, 0x9e, 0x8b, 0x11, 0x6f, 0x16, 0x51, 0xb4, 0xc4, 0xe6, 0x91, 0x9c,
	0x65, 0x28, 0x2a, 0xd9, 0x48, 0x0b, 0x54, 0xf4, 0xc9, 0x83, 0xff, 0x2e, 0x75, 0xec, 0xdc, 0xbd,
	0x07, 0x15, 0x96, 0x98, 0x18, 0xc5, 0xd8, 0x52, 0xd0, 0xd9, 0x44, 0x0b, 0x17, 0xd4, 0xb4, 0xcc,
	0x12, 0x9b, 0xf0, 0xcf, 0x4d, 0x7f, 0x0d, 0xdb, 0x3d, 0xbb, 0xe5, 0x4f, 0x84, 0x36, 0x2a, 0x3b,
	0xbf, 0x7a, 0x57, 0xd5, 0xe9, 0xa9, 0x9e, 0x3b, 0xea, 0x10, 0xa9, 0x43, 0x31, 0x15, 0x23, 0x61,
	0xed, 0xf4, 0xa9, 0x05, 0xd1, 0x2f, 0x0f, 0xb6, 0x8e, 0xb9, 0xb9, 0x98, 0x81, 0xdc, 0x80, 0x75,
	0xec, 0x2b, 0x76, 0x5b, 0xef, 0xe1, 0x04, 0x02, 0xe4, 0xdc, 0x18, 0xea, 0x50, 0xd4, 0xb9, 0x9d,
	0x98, 0xa5, 0x42, 0x2d, 0x20, 0x9b, 0xe0, 0x73, 0xd9, 0xc7, 0x14, 0x15, 0x9a, 0xff, 0x24, 0x2d,
	0xa8, 0x64, 0x7c, 0x1a, 0x8f, 0x94, 0x34, 0x67, 0x6e, 0x58, 0xe5, 0x8c, 0x4f, 0x9f, 0xe6, 0x98,
	0xec, 0x43, 0x4d, 0x72, 0xf3, 0x4e, 0x65, 0xc3, 0x58, 0xc8, 0x44, 0xb9, 0x49, 0x79, 0xb4, 0xea,
	0xd8, 0x2e, 0x92, 0x79, 0x39, 0x78, 0x3f, 0x3d, 0x8f, 0x33, 0x66, 0xec, 0xa4, 0x3c, 0x1a, 0x38,
	0x8e, 0x32, 0xc3, 0x49, 0x13, 0xd6, 0x32, 0x3e, 0xe5, 0x72, 0xc2, 0x9b, 0x6b, 0x78, 0x3a, 0x83,
	0xf9, 0xc9, 0x09, 0x4b, 0x99, 0x4c, 0x78, 0xb3, 0x6c, 0x4f, 0x1c, 0x8c, 0xbe, 0x7a, 0xd0, 0x58,
	0xb6, 0xd6, 0x0d, 0x77, 0x79, 0x5a, 0xde, 0x35, 0xa7, 0x45, 0x1e, 0x5a, 0xe7, 0x84, 0x1c, 0xc4,
	0x67, 0x42, 0xe7, 0xee, 0xf8, 0x07, 0x41, 0xa7, 0x81, 0x17, 0x2f, 0xf9, 0x6c, 0x1d, 0x75, 0x38,
	0x77, 0x34, 0x99, 0xef, 0xbb, 0x4f, 0x2d, 0xe8, 0x7c, 0xf3, 0x71, 0x9b, 0x73, 0x55, 0x8f, 0x67,
	0x53, 0x91, 0x70, 0xd2, 0x83, 0xa2, 0xdd, 0xa9, 0x2d, 0x0c, 0xbb, 0xf8, 0x80, 0x85, 0x64, 0x91,
	0xb2, 0x05, 0x46, 0x37, 0x3f, 0xff, 0xf8, 0xf9, 0xa5, 0xb0, 0x4b, 0x5a, 0xf8, 0x26, 0xba, 0x94,
	0xed, 0x0f, 0x76, 0x79, 0x3e, 0x22, 0xc1, 0xc9, 0x2b, 0x58, 0x73, 0x9b, 0x4f, 0xfe, 0xc5, 0x18,
	0x17, 0x9f, 0x9d, 0xb0, 0x7e, 0x91, 0x74, 0xa1, 0xf7, 0x31, 0xf4, 0xff, 0x64, 0x77, 0x75, 0xe8,
	0x89, 0x8b, 0xf8, 0x1e, 0x36, 0x96, 0x3e, 0x23, 0xd2, 0x9a, 0x59, 0xb2, 0xe2, 0x39, 0x09, 0x77,
	0x56, 0x1f, 0xba, 0xa4, 0x77, 0x31, 0xe9, 0x2d, 0x12, 0xad, 0x4e, 0xca, 0xf0, 0x8e, 0xb6, 0x69,
	0xf4, 0xaa, 0xf5, 0x0e, 0xe7, 0x26, 0x5d, 0xfa, 0xaa, 0xc2, 0xd6, 0xca, 0xb3, 0xeb, 0xb5, 0x7b,
	0x66, 0xe5, 0x27, 0x25, 0xfc, 0x73, 0xb9, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x61, 0x43, 0x00,
	0x64, 0x9f, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StakingServiceClient is the client API for StakingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StakingServiceClient interface {
	Stake(ctx context.Context, in *StakeRequest, opts ...grpc.CallOption) (*StakeResponse, error)
	Unstake(ctx context.Context, in *UnstakeRequest, opts ...grpc.CallOption) (*UnstakeResponse, error)
	GetActiveStakes(ctx context.Context, in *GetActiveStakesRequest, opts ...grpc.CallOption) (*GetActiveStakesResponse, error)
	GetStakingHistory(ctx context.Context, in *StakingHistoryRequest, opts ...grpc.CallOption) (*StakingHistoryResponse, error)
}

type stakingServiceClient struct {
	cc *grpc.ClientConn
}

func NewStakingServiceClient(cc *grpc.ClientConn) StakingServiceClient {
	return &stakingServiceClient{cc}
}

func (c *stakingServiceClient) Stake(ctx context.Context, in *StakeRequest, opts ...grpc.CallOption) (*StakeResponse, error) {
	out := new(StakeResponse)
	err := c.cc.Invoke(ctx, "/api.StakingService/Stake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakingServiceClient) Unstake(ctx context.Context, in *UnstakeRequest, opts ...grpc.CallOption) (*UnstakeResponse, error) {
	out := new(UnstakeResponse)
	err := c.cc.Invoke(ctx, "/api.StakingService/Unstake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakingServiceClient) GetActiveStakes(ctx context.Context, in *GetActiveStakesRequest, opts ...grpc.CallOption) (*GetActiveStakesResponse, error) {
	out := new(GetActiveStakesResponse)
	err := c.cc.Invoke(ctx, "/api.StakingService/GetActiveStakes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakingServiceClient) GetStakingHistory(ctx context.Context, in *StakingHistoryRequest, opts ...grpc.CallOption) (*StakingHistoryResponse, error) {
	out := new(StakingHistoryResponse)
	err := c.cc.Invoke(ctx, "/api.StakingService/GetStakingHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StakingServiceServer is the server API for StakingService service.
type StakingServiceServer interface {
	Stake(context.Context, *StakeRequest) (*StakeResponse, error)
	Unstake(context.Context, *UnstakeRequest) (*UnstakeResponse, error)
	GetActiveStakes(context.Context, *GetActiveStakesRequest) (*GetActiveStakesResponse, error)
	GetStakingHistory(context.Context, *StakingHistoryRequest) (*StakingHistoryResponse, error)
}

// UnimplementedStakingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStakingServiceServer struct {
}

func (*UnimplementedStakingServiceServer) Stake(ctx context.Context, req *StakeRequest) (*StakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stake not implemented")
}
func (*UnimplementedStakingServiceServer) Unstake(ctx context.Context, req *UnstakeRequest) (*UnstakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unstake not implemented")
}
func (*UnimplementedStakingServiceServer) GetActiveStakes(ctx context.Context, req *GetActiveStakesRequest) (*GetActiveStakesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveStakes not implemented")
}
func (*UnimplementedStakingServiceServer) GetStakingHistory(ctx context.Context, req *StakingHistoryRequest) (*StakingHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStakingHistory not implemented")
}

func RegisterStakingServiceServer(s *grpc.Server, srv StakingServiceServer) {
	s.RegisterService(&_StakingService_serviceDesc, srv)
}

func _StakingService_Stake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakingServiceServer).Stake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StakingService/Stake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakingServiceServer).Stake(ctx, req.(*StakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakingService_Unstake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnstakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakingServiceServer).Unstake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StakingService/Unstake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakingServiceServer).Unstake(ctx, req.(*UnstakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakingService_GetActiveStakes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveStakesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakingServiceServer).GetActiveStakes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StakingService/GetActiveStakes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakingServiceServer).GetActiveStakes(ctx, req.(*GetActiveStakesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakingService_GetStakingHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StakingHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakingServiceServer).GetStakingHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StakingService/GetStakingHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakingServiceServer).GetStakingHistory(ctx, req.(*StakingHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StakingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.StakingService",
	HandlerType: (*StakingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stake",
			Handler:    _StakingService_Stake_Handler,
		},
		{
			MethodName: "Unstake",
			Handler:    _StakingService_Unstake_Handler,
		},
		{
			MethodName: "GetActiveStakes",
			Handler:    _StakingService_GetActiveStakes_Handler,
		},
		{
			MethodName: "GetStakingHistory",
			Handler:    _StakingService_GetStakingHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staking.proto",
}
