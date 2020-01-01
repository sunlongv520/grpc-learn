// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Models.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ProdModel struct {
	ProdId               int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdName             string   `protobuf:"bytes,2,opt,name=prod_name,json=prodName,proto3" json:"prod_name,omitempty"`
	ProdPrice            float32  `protobuf:"fixed32,3,opt,name=prod_price,json=prodPrice,proto3" json:"prod_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdModel) Reset()         { *m = ProdModel{} }
func (m *ProdModel) String() string { return proto.CompactTextString(m) }
func (*ProdModel) ProtoMessage()    {}
func (*ProdModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{0}
}

func (m *ProdModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdModel.Unmarshal(m, b)
}
func (m *ProdModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdModel.Marshal(b, m, deterministic)
}
func (m *ProdModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdModel.Merge(m, src)
}
func (m *ProdModel) XXX_Size() int {
	return xxx_messageInfo_ProdModel.Size(m)
}
func (m *ProdModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdModel.DiscardUnknown(m)
}

var xxx_messageInfo_ProdModel proto.InternalMessageInfo

func (m *ProdModel) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

func (m *ProdModel) GetProdName() string {
	if m != nil {
		return m.ProdName
	}
	return ""
}

func (m *ProdModel) GetProdPrice() float32 {
	if m != nil {
		return m.ProdPrice
	}
	return 0
}

type OrderMain struct {
	OrderId              int32                `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderNo              string               `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	UserId               int32                `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderMoney           float32              `protobuf:"fixed32,4,opt,name=order_money,json=orderMoney,proto3" json:"order_money,omitempty"`
	OrderTime            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=order_time,json=orderTime,proto3" json:"order_time,omitempty"`
	OrderDetails         []*OrderDetail       `protobuf:"bytes,6,rep,name=order_details,json=orderDetails,proto3" json:"order_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderMain) Reset()         { *m = OrderMain{} }
func (m *OrderMain) String() string { return proto.CompactTextString(m) }
func (*OrderMain) ProtoMessage()    {}
func (*OrderMain) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{1}
}

func (m *OrderMain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderMain.Unmarshal(m, b)
}
func (m *OrderMain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderMain.Marshal(b, m, deterministic)
}
func (m *OrderMain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderMain.Merge(m, src)
}
func (m *OrderMain) XXX_Size() int {
	return xxx_messageInfo_OrderMain.Size(m)
}
func (m *OrderMain) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderMain.DiscardUnknown(m)
}

var xxx_messageInfo_OrderMain proto.InternalMessageInfo

func (m *OrderMain) GetOrderId() int32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderMain) GetOrderNo() string {
	if m != nil {
		return m.OrderNo
	}
	return ""
}

func (m *OrderMain) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *OrderMain) GetOrderMoney() float32 {
	if m != nil {
		return m.OrderMoney
	}
	return 0
}

func (m *OrderMain) GetOrderTime() *timestamp.Timestamp {
	if m != nil {
		return m.OrderTime
	}
	return nil
}

func (m *OrderMain) GetOrderDetails() []*OrderDetail {
	if m != nil {
		return m.OrderDetails
	}
	return nil
}

//子订单模型
type OrderDetail struct {
	DetailId             int32    `protobuf:"varint,1,opt,name=detail_id,json=detailId,proto3" json:"detail_id,omitempty"`
	OrderNo              string   `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	ProdId               int32    `protobuf:"varint,3,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdPrice            float32  `protobuf:"fixed32,4,opt,name=prod_price,json=prodPrice,proto3" json:"prod_price,omitempty"`
	ProdNum              int32    `protobuf:"varint,5,opt,name=prod_num,json=prodNum,proto3" json:"prod_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderDetail) Reset()         { *m = OrderDetail{} }
func (m *OrderDetail) String() string { return proto.CompactTextString(m) }
func (*OrderDetail) ProtoMessage()    {}
func (*OrderDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{2}
}

func (m *OrderDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderDetail.Unmarshal(m, b)
}
func (m *OrderDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderDetail.Marshal(b, m, deterministic)
}
func (m *OrderDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderDetail.Merge(m, src)
}
func (m *OrderDetail) XXX_Size() int {
	return xxx_messageInfo_OrderDetail.Size(m)
}
func (m *OrderDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderDetail.DiscardUnknown(m)
}

var xxx_messageInfo_OrderDetail proto.InternalMessageInfo

func (m *OrderDetail) GetDetailId() int32 {
	if m != nil {
		return m.DetailId
	}
	return 0
}

func (m *OrderDetail) GetOrderNo() string {
	if m != nil {
		return m.OrderNo
	}
	return ""
}

func (m *OrderDetail) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

func (m *OrderDetail) GetProdPrice() float32 {
	if m != nil {
		return m.ProdPrice
	}
	return 0
}

func (m *OrderDetail) GetProdNum() int32 {
	if m != nil {
		return m.ProdNum
	}
	return 0
}

//用户模型
type UserInfo struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserScore            int32    `protobuf:"varint,2,opt,name=user_score,json=userScore,proto3" json:"user_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{3}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfo) GetUserScore() int32 {
	if m != nil {
		return m.UserScore
	}
	return 0
}

func init() {
	proto.RegisterType((*ProdModel)(nil), "services.ProdModel")
	proto.RegisterType((*OrderMain)(nil), "services.OrderMain")
	proto.RegisterType((*OrderDetail)(nil), "services.OrderDetail")
	proto.RegisterType((*UserInfo)(nil), "services.UserInfo")
}

func init() { proto.RegisterFile("Models.proto", fileDescriptor_96b05f67b8e9f86a) }

var fileDescriptor_96b05f67b8e9f86a = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0xa6, 0x69, 0x93, 0x4b, 0x59, 0x2c, 0x21, 0x42, 0x51, 0xd5, 0x28, 0x53, 0xa6,
	0x54, 0x2a, 0x13, 0x8c, 0x88, 0xa5, 0x43, 0x4b, 0x15, 0x60, 0x2e, 0x69, 0xed, 0x56, 0x91, 0xea,
	0x38, 0xb2, 0x13, 0x24, 0xde, 0x84, 0xe7, 0xe4, 0x09, 0x90, 0xcf, 0x44, 0x21, 0x2c, 0x8c, 0xff,
	0xff, 0xfb, 0xee, 0x7c, 0xdf, 0xc1, 0x64, 0x2d, 0x19, 0x3f, 0xeb, 0xb4, 0x52, 0xb2, 0x96, 0xd4,
	0xd3, 0x5c, 0xbd, 0x17, 0x07, 0xae, 0xa7, 0xf3, 0x93, 0x94, 0xa7, 0x33, 0x5f, 0xa0, 0xbf, 0x6f,
	0x8e, 0x8b, 0xba, 0x10, 0x5c, 0xd7, 0xb9, 0xa8, 0xec, 0xd3, 0xf8, 0x0d, 0xfc, 0xad, 0x92, 0x0c,
	0xcb, 0xe9, 0x15, 0x8c, 0x2b, 0x25, 0xd9, 0xae, 0x60, 0x21, 0x89, 0x48, 0xe2, 0x66, 0x23, 0x23,
	0x57, 0x8c, 0xde, 0x80, 0x8f, 0x41, 0x99, 0x0b, 0x1e, 0x0e, 0x22, 0x92, 0xf8, 0x99, 0x67, 0x8c,
	0x4d, 0x2e, 0x38, 0x9d, 0x01, 0x60, 0x58, 0xa9, 0xe2, 0xc0, 0x43, 0x27, 0x22, 0xc9, 0x20, 0xc3,
	0xe7, 0x5b, 0x63, 0xc4, 0x5f, 0x04, 0xfc, 0x27, 0xc5, 0xb8, 0x5a, 0xe7, 0x45, 0x49, 0xaf, 0xc1,
	0x93, 0x46, 0x74, 0x33, 0xc6, 0xa8, 0x57, 0xac, 0x8b, 0x4a, 0xf9, 0x33, 0xc3, 0x46, 0x1b, 0x69,
	0x3e, 0xd6, 0x68, 0x5b, 0xe4, 0xd8, 0x8f, 0x19, 0xb9, 0x62, 0x74, 0x0e, 0x81, 0xad, 0x11, 0xb2,
	0xe4, 0x1f, 0xe1, 0x10, 0x87, 0x03, 0x5a, 0x6b, 0xe3, 0xd0, 0x3b, 0xb0, 0x6a, 0x67, 0x16, 0x0f,
	0xdd, 0x88, 0x24, 0xc1, 0x72, 0x9a, 0x5a, 0x2a, 0x69, 0x4b, 0x25, 0x7d, 0x69, 0xa9, 0x64, 0x3e,
	0xbe, 0x36, 0x9a, 0xde, 0xc3, 0x85, 0x2d, 0x65, 0xbc, 0xce, 0x8b, 0xb3, 0x0e, 0x47, 0x91, 0x93,
	0x04, 0xcb, 0xcb, 0xb4, 0xa5, 0x9b, 0xe2, 0x5a, 0x8f, 0x98, 0x66, 0x13, 0xd9, 0x09, 0x1d, 0x7f,
	0x12, 0x08, 0x7e, 0xa5, 0x06, 0xa0, 0xed, 0xd2, 0xed, 0xed, 0x59, 0xe3, 0xdf, 0xc5, 0xdb, 0x8b,
	0x38, 0xbd, 0x8b, 0xf4, 0xa1, 0x0f, 0xff, 0x40, 0x37, 0x2d, 0xed, 0xc1, 0x1a, 0x81, 0x4b, 0xbb,
	0x19, 0xf6, 0xd9, 0x34, 0x22, 0x7e, 0x00, 0xef, 0xd5, 0xc0, 0x2b, 0x8f, 0x3d, 0xae, 0xa4, 0xc7,
	0x75, 0x06, 0x80, 0x81, 0x3e, 0x48, 0x65, 0x2f, 0xee, 0x66, 0xbe, 0x71, 0x9e, 0x8d, 0xb1, 0x1f,
	0x21, 0xb9, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x45, 0xfa, 0x0c, 0x77, 0x02, 0x00,
	0x00,
}
