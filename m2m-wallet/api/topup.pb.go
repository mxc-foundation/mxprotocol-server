// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topup.proto

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

type GetTopUpHistoryRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopUpHistoryRequest) Reset()         { *m = GetTopUpHistoryRequest{} }
func (m *GetTopUpHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopUpHistoryRequest) ProtoMessage()    {}
func (*GetTopUpHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{0}
}

func (m *GetTopUpHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUpHistoryRequest.Unmarshal(m, b)
}
func (m *GetTopUpHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUpHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetTopUpHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUpHistoryRequest.Merge(m, src)
}
func (m *GetTopUpHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopUpHistoryRequest.Size(m)
}
func (m *GetTopUpHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUpHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUpHistoryRequest proto.InternalMessageInfo

func (m *GetTopUpHistoryRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetTopUpHistoryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetTopUpHistoryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type TopUpHistory struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	MoneyType            string   `protobuf:"bytes,5,opt,name=money_type,json=moneyType,proto3" json:"money_type,omitempty"`
	TxHash               string   `protobuf:"bytes,6,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopUpHistory) Reset()         { *m = TopUpHistory{} }
func (m *TopUpHistory) String() string { return proto.CompactTextString(m) }
func (*TopUpHistory) ProtoMessage()    {}
func (*TopUpHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{1}
}

func (m *TopUpHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopUpHistory.Unmarshal(m, b)
}
func (m *TopUpHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopUpHistory.Marshal(b, m, deterministic)
}
func (m *TopUpHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopUpHistory.Merge(m, src)
}
func (m *TopUpHistory) XXX_Size() int {
	return xxx_messageInfo_TopUpHistory.Size(m)
}
func (m *TopUpHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_TopUpHistory.DiscardUnknown(m)
}

var xxx_messageInfo_TopUpHistory proto.InternalMessageInfo

func (m *TopUpHistory) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TopUpHistory) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TopUpHistory) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TopUpHistory) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *TopUpHistory) GetMoneyType() string {
	if m != nil {
		return m.MoneyType
	}
	return ""
}

func (m *TopUpHistory) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

type GetTopUpHistoryResponse struct {
	Count                int64            `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	TopupHistory         []*TopUpHistory  `protobuf:"bytes,2,rep,name=topup_history,json=topupHistory,proto3" json:"topup_history,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetTopUpHistoryResponse) Reset()         { *m = GetTopUpHistoryResponse{} }
func (m *GetTopUpHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetTopUpHistoryResponse) ProtoMessage()    {}
func (*GetTopUpHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{2}
}

func (m *GetTopUpHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUpHistoryResponse.Unmarshal(m, b)
}
func (m *GetTopUpHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUpHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetTopUpHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUpHistoryResponse.Merge(m, src)
}
func (m *GetTopUpHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetTopUpHistoryResponse.Size(m)
}
func (m *GetTopUpHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUpHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUpHistoryResponse proto.InternalMessageInfo

func (m *GetTopUpHistoryResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetTopUpHistoryResponse) GetTopupHistory() []*TopUpHistory {
	if m != nil {
		return m.TopupHistory
	}
	return nil
}

func (m *GetTopUpHistoryResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTopUpHistoryRequest)(nil), "api.GetTopUpHistoryRequest")
	proto.RegisterType((*TopUpHistory)(nil), "api.TopUpHistory")
	proto.RegisterType((*GetTopUpHistoryResponse)(nil), "api.GetTopUpHistoryResponse")
}

func init() { proto.RegisterFile("topup.proto", fileDescriptor_8eec749941d0cb6c) }

var fileDescriptor_8eec749941d0cb6c = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0xce, 0xd3, 0x30,
	0x14, 0x54, 0x92, 0xaf, 0x41, 0x75, 0xfa, 0x81, 0x30, 0x5f, 0xdb, 0xa8, 0x2d, 0x52, 0x95, 0x55,
	0x57, 0xad, 0x54, 0x24, 0x58, 0xb3, 0xa2, 0xec, 0x50, 0x28, 0x4b, 0x14, 0x99, 0xd6, 0x49, 0x2c,
	0x35, 0x79, 0xc6, 0x7e, 0x41, 0xcd, 0x96, 0x2b, 0x70, 0x01, 0x38, 0x13, 0x57, 0xe0, 0x20, 0xa8,
	0xcf, 0xae, 0xc4, 0x4f, 0x77, 0x99, 0x79, 0x93, 0x79, 0x9e, 0xb1, 0x59, 0x82, 0xa0, 0x3b, 0xbd,
	0xd6, 0x06, 0x10, 0x78, 0x24, 0xb4, 0x9a, 0x2d, 0x2a, 0x80, 0xea, 0x24, 0x37, 0x42, 0xab, 0x8d,
	0x68, 0x5b, 0x40, 0x81, 0x0a, 0x5a, 0xeb, 0x24, 0xb3, 0x7b, 0x6d, 0xa0, 0x54, 0x27, 0xe9, 0x60,
	0xf6, 0x91, 0x4d, 0xde, 0x48, 0xdc, 0x83, 0xfe, 0xa0, 0x77, 0xca, 0x22, 0x98, 0x3e, 0x97, 0x9f,
	0x3b, 0x69, 0x91, 0x8f, 0x59, 0x0c, 0xa6, 0x2a, 0xd4, 0x31, 0x0d, 0x96, 0xc1, 0x2a, 0xca, 0x07,
	0x60, 0xaa, 0xb7, 0x47, 0x3e, 0x61, 0x31, 0x94, 0xa5, 0x95, 0x98, 0x86, 0x44, 0x7b, 0xc4, 0x1f,
	0xd8, 0xe0, 0xa4, 0x1a, 0x85, 0x69, 0xe4, 0xd4, 0x04, 0xb2, 0x1f, 0x01, 0x1b, 0xfd, 0x69, 0xce,
	0x39, 0xbb, 0x2b, 0x0d, 0x34, 0xe4, 0x39, 0xcc, 0xe9, 0x9b, 0x3f, 0x66, 0x21, 0x02, 0xd9, 0x0d,
	0xf3, 0x10, 0xe1, 0xb2, 0x42, 0x34, 0xd0, 0xb5, 0xce, 0x2b, 0xc8, 0x3d, 0xe2, 0xcf, 0x19, 0x3b,
	0x18, 0x29, 0x50, 0x1e, 0x0b, 0x81, 0xe9, 0x1d, 0xe9, 0x87, 0x9e, 0x79, 0x4d, 0xe3, 0x06, 0x5a,
	0xd9, 0x17, 0xd8, 0x6b, 0x99, 0x0e, 0xdc, 0x98, 0x98, 0x7d, 0xaf, 0x25, 0x9f, 0xb2, 0x47, 0x78,
	0x2e, 0x6a, 0x61, 0xeb, 0x34, 0xa6, 0x59, 0x8c, 0xe7, 0x9d, 0xb0, 0x75, 0xf6, 0x3d, 0x60, 0xd3,
	0xff, 0x3a, 0xb0, 0x1a, 0x5a, 0x2b, 0x2f, 0xa9, 0x0e, 0x74, 0x12, 0xdf, 0x01, 0x01, 0xfe, 0x92,
	0xdd, 0x53, 0xeb, 0x45, 0xed, 0xe4, 0x69, 0xb8, 0x8c, 0x56, 0xc9, 0xf6, 0xe9, 0x5a, 0x68, 0xb5,
	0xfe, 0xcb, 0x67, 0x44, 0xba, 0x6b, 0xf8, 0x57, 0x6c, 0xd4, 0x59, 0x69, 0x0a, 0x7f, 0x05, 0x14,
	0x2f, 0xd9, 0x3e, 0xd0, 0x6f, 0xef, 0x1c, 0x77, 0xdd, 0x9c, 0x27, 0x17, 0xa5, 0x27, 0xb7, 0xbd,
	0x6f, 0xf1, 0xbd, 0x34, 0x5f, 0xd4, 0x41, 0x72, 0xc5, 0x9e, 0xfc, 0x73, 0x62, 0x3e, 0x27, 0x97,
	0xdb, 0x77, 0x39, 0x5b, 0xdc, 0x1e, 0xba, 0x55, 0xd9, 0xfc, 0xeb, 0xcf, 0x5f, 0xdf, 0xc2, 0x31,
	0x7f, 0x46, 0x4f, 0x06, 0x41, 0x17, 0x9d, 0xde, 0xf8, 0x68, 0x9f, 0x62, 0x7a, 0x27, 0x2f, 0x7e,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x50, 0x92, 0x55, 0xf0, 0x68, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopUpServiceClient is the client API for TopUpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopUpServiceClient interface {
	GetTopUpHistory(ctx context.Context, in *GetTopUpHistoryRequest, opts ...grpc.CallOption) (*GetTopUpHistoryResponse, error)
}

type topUpServiceClient struct {
	cc *grpc.ClientConn
}

func NewTopUpServiceClient(cc *grpc.ClientConn) TopUpServiceClient {
	return &topUpServiceClient{cc}
}

func (c *topUpServiceClient) GetTopUpHistory(ctx context.Context, in *GetTopUpHistoryRequest, opts ...grpc.CallOption) (*GetTopUpHistoryResponse, error) {
	out := new(GetTopUpHistoryResponse)
	err := c.cc.Invoke(ctx, "/api.TopUpService/GetTopUpHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopUpServiceServer is the server API for TopUpService service.
type TopUpServiceServer interface {
	GetTopUpHistory(context.Context, *GetTopUpHistoryRequest) (*GetTopUpHistoryResponse, error)
}

func RegisterTopUpServiceServer(s *grpc.Server, srv TopUpServiceServer) {
	s.RegisterService(&_TopUpService_serviceDesc, srv)
}

func _TopUpService_GetTopUpHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopUpHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopUpServiceServer).GetTopUpHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TopUpService/GetTopUpHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopUpServiceServer).GetTopUpHistory(ctx, req.(*GetTopUpHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopUpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TopUpService",
	HandlerType: (*TopUpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopUpHistory",
			Handler:    _TopUpService_GetTopUpHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topup.proto",
}
