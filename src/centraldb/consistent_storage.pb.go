// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consistent_storage.proto

package main

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

type RemoveError int32

const (
	RemoveError_REMOVE_KEY_ERROR RemoveError = 0
	RemoveError_REMOVE_NOT_OWNER RemoveError = 1
)

var RemoveError_name = map[int32]string{
	0: "REMOVE_KEY_ERROR",
	1: "REMOVE_NOT_OWNER",
}

var RemoveError_value = map[string]int32{
	"REMOVE_KEY_ERROR": 0,
	"REMOVE_NOT_OWNER": 1,
}

func (x RemoveError) String() string {
	return proto.EnumName(RemoveError_name, int32(x))
}

func (RemoveError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{0}
}

type TransferError int32

const (
	TransferError_TRANSFER_KEY_ERROR TransferError = 0
	TransferError_TRANSFER_NOT_OWNER TransferError = 1
)

var TransferError_name = map[int32]string{
	0: "TRANSFER_KEY_ERROR",
	1: "TRANSFER_NOT_OWNER",
}

var TransferError_value = map[string]int32{
	"TRANSFER_KEY_ERROR": 0,
	"TRANSFER_NOT_OWNER": 1,
}

func (x TransferError) String() string {
	return proto.EnumName(TransferError_name, int32(x))
}

func (TransferError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{1}
}

type WrappedRequest struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*WrappedRequest_Get
	//	*WrappedRequest_Put
	//	*WrappedRequest_Remove
	//	*WrappedRequest_Transfer
	Request              isWrappedRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WrappedRequest) Reset()         { *m = WrappedRequest{} }
func (m *WrappedRequest) String() string { return proto.CompactTextString(m) }
func (*WrappedRequest) ProtoMessage()    {}
func (*WrappedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{0}
}

func (m *WrappedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WrappedRequest.Unmarshal(m, b)
}
func (m *WrappedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WrappedRequest.Marshal(b, m, deterministic)
}
func (m *WrappedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedRequest.Merge(m, src)
}
func (m *WrappedRequest) XXX_Size() int {
	return xxx_messageInfo_WrappedRequest.Size(m)
}
func (m *WrappedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedRequest proto.InternalMessageInfo

func (m *WrappedRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WrappedRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type isWrappedRequest_Request interface {
	isWrappedRequest_Request()
}

type WrappedRequest_Get struct {
	Get *GetRequest `protobuf:"bytes,10,opt,name=get,proto3,oneof"`
}

type WrappedRequest_Put struct {
	Put *PutRequest `protobuf:"bytes,11,opt,name=put,proto3,oneof"`
}

type WrappedRequest_Remove struct {
	Remove *RemoveRequest `protobuf:"bytes,12,opt,name=remove,proto3,oneof"`
}

type WrappedRequest_Transfer struct {
	Transfer *TransferRequest `protobuf:"bytes,13,opt,name=transfer,proto3,oneof"`
}

func (*WrappedRequest_Get) isWrappedRequest_Request() {}

func (*WrappedRequest_Put) isWrappedRequest_Request() {}

func (*WrappedRequest_Remove) isWrappedRequest_Request() {}

func (*WrappedRequest_Transfer) isWrappedRequest_Request() {}

func (m *WrappedRequest) GetRequest() isWrappedRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *WrappedRequest) GetGet() *GetRequest {
	if x, ok := m.GetRequest().(*WrappedRequest_Get); ok {
		return x.Get
	}
	return nil
}

func (m *WrappedRequest) GetPut() *PutRequest {
	if x, ok := m.GetRequest().(*WrappedRequest_Put); ok {
		return x.Put
	}
	return nil
}

func (m *WrappedRequest) GetRemove() *RemoveRequest {
	if x, ok := m.GetRequest().(*WrappedRequest_Remove); ok {
		return x.Remove
	}
	return nil
}

func (m *WrappedRequest) GetTransfer() *TransferRequest {
	if x, ok := m.GetRequest().(*WrappedRequest_Transfer); ok {
		return x.Transfer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WrappedRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WrappedRequest_Get)(nil),
		(*WrappedRequest_Put)(nil),
		(*WrappedRequest_Remove)(nil),
		(*WrappedRequest_Transfer)(nil),
	}
}

type WrappedResponse struct {
	// Types that are valid to be assigned to Response:
	//	*WrappedResponse_Get
	//	*WrappedResponse_Put
	//	*WrappedResponse_Remove
	//	*WrappedResponse_Transfer
	Response             isWrappedResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *WrappedResponse) Reset()         { *m = WrappedResponse{} }
func (m *WrappedResponse) String() string { return proto.CompactTextString(m) }
func (*WrappedResponse) ProtoMessage()    {}
func (*WrappedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{1}
}

func (m *WrappedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WrappedResponse.Unmarshal(m, b)
}
func (m *WrappedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WrappedResponse.Marshal(b, m, deterministic)
}
func (m *WrappedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedResponse.Merge(m, src)
}
func (m *WrappedResponse) XXX_Size() int {
	return xxx_messageInfo_WrappedResponse.Size(m)
}
func (m *WrappedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedResponse proto.InternalMessageInfo

type isWrappedResponse_Response interface {
	isWrappedResponse_Response()
}

type WrappedResponse_Get struct {
	Get *GetResponse `protobuf:"bytes,10,opt,name=get,proto3,oneof"`
}

type WrappedResponse_Put struct {
	Put *PutResponse `protobuf:"bytes,11,opt,name=put,proto3,oneof"`
}

type WrappedResponse_Remove struct {
	Remove *RemoveResponse `protobuf:"bytes,12,opt,name=remove,proto3,oneof"`
}

type WrappedResponse_Transfer struct {
	Transfer *TransferResponse `protobuf:"bytes,13,opt,name=transfer,proto3,oneof"`
}

func (*WrappedResponse_Get) isWrappedResponse_Response() {}

func (*WrappedResponse_Put) isWrappedResponse_Response() {}

func (*WrappedResponse_Remove) isWrappedResponse_Response() {}

func (*WrappedResponse_Transfer) isWrappedResponse_Response() {}

func (m *WrappedResponse) GetResponse() isWrappedResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *WrappedResponse) GetGet() *GetResponse {
	if x, ok := m.GetResponse().(*WrappedResponse_Get); ok {
		return x.Get
	}
	return nil
}

func (m *WrappedResponse) GetPut() *PutResponse {
	if x, ok := m.GetResponse().(*WrappedResponse_Put); ok {
		return x.Put
	}
	return nil
}

func (m *WrappedResponse) GetRemove() *RemoveResponse {
	if x, ok := m.GetResponse().(*WrappedResponse_Remove); ok {
		return x.Remove
	}
	return nil
}

func (m *WrappedResponse) GetTransfer() *TransferResponse {
	if x, ok := m.GetResponse().(*WrappedResponse_Transfer); ok {
		return x.Transfer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WrappedResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WrappedResponse_Get)(nil),
		(*WrappedResponse_Put)(nil),
		(*WrappedResponse_Remove)(nil),
		(*WrappedResponse_Transfer)(nil),
	}
}

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{2}
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

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResponse struct {
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	IsOwner              bool     `protobuf:"varint,3,opt,name=isOwner,proto3" json:"isOwner,omitempty"`
	Owner                string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{3}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *GetResponse) GetIsOwner() bool {
	if m != nil {
		return m.IsOwner
	}
	return false
}

func (m *GetResponse) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type PutRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{4}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PutResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{5}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *PutResponse) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type RemoveRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{6}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type RemoveResponse struct {
	Removed              bool        `protobuf:"varint,1,opt,name=removed,proto3" json:"removed,omitempty"`
	ErrorType            RemoveError `protobuf:"varint,2,opt,name=errorType,proto3,enum=RemoveError" json:"errorType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{7}
}

func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResponse.Unmarshal(m, b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
}
func (m *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(m, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveResponse.Size(m)
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

func (m *RemoveResponse) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}

func (m *RemoveResponse) GetErrorType() RemoveError {
	if m != nil {
		return m.ErrorType
	}
	return RemoveError_REMOVE_KEY_ERROR
}

type TransferRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	NewOwner             string   `protobuf:"bytes,2,opt,name=newOwner,proto3" json:"newOwner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{8}
}

func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (m *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(m, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TransferRequest) GetNewOwner() string {
	if m != nil {
		return m.NewOwner
	}
	return ""
}

type TransferResponse struct {
	Transferred          bool          `protobuf:"varint,1,opt,name=transferred,proto3" json:"transferred,omitempty"`
	ErrorType            TransferError `protobuf:"varint,2,opt,name=errorType,proto3,enum=TransferError" json:"errorType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TransferResponse) Reset()         { *m = TransferResponse{} }
func (m *TransferResponse) String() string { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()    {}
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74470502155bf339, []int{9}
}

func (m *TransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferResponse.Unmarshal(m, b)
}
func (m *TransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferResponse.Marshal(b, m, deterministic)
}
func (m *TransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferResponse.Merge(m, src)
}
func (m *TransferResponse) XXX_Size() int {
	return xxx_messageInfo_TransferResponse.Size(m)
}
func (m *TransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferResponse proto.InternalMessageInfo

func (m *TransferResponse) GetTransferred() bool {
	if m != nil {
		return m.Transferred
	}
	return false
}

func (m *TransferResponse) GetErrorType() TransferError {
	if m != nil {
		return m.ErrorType
	}
	return TransferError_TRANSFER_KEY_ERROR
}

func init() {
	proto.RegisterEnum("RemoveError", RemoveError_name, RemoveError_value)
	proto.RegisterEnum("TransferError", TransferError_name, TransferError_value)
	proto.RegisterType((*WrappedRequest)(nil), "WrappedRequest")
	proto.RegisterType((*WrappedResponse)(nil), "WrappedResponse")
	proto.RegisterType((*GetRequest)(nil), "GetRequest")
	proto.RegisterType((*GetResponse)(nil), "GetResponse")
	proto.RegisterType((*PutRequest)(nil), "PutRequest")
	proto.RegisterType((*PutResponse)(nil), "PutResponse")
	proto.RegisterType((*RemoveRequest)(nil), "RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "RemoveResponse")
	proto.RegisterType((*TransferRequest)(nil), "TransferRequest")
	proto.RegisterType((*TransferResponse)(nil), "TransferResponse")
}

func init() {
	proto.RegisterFile("consistent_storage.proto", fileDescriptor_74470502155bf339)
}

var fileDescriptor_74470502155bf339 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x93, 0x76, 0x74, 0xed, 0x49, 0x97, 0x66, 0xd6, 0x34, 0x45, 0x11, 0x82, 0xe0, 0xab,
	0x32, 0x50, 0x10, 0x1d, 0x37, 0x5c, 0x4d, 0x6c, 0x0a, 0x43, 0x20, 0x9a, 0xc9, 0xab, 0x36, 0xc1,
	0x4d, 0x95, 0x51, 0x53, 0x45, 0xdd, 0x92, 0x60, 0x3b, 0x1b, 0x7b, 0x36, 0xc4, 0xe3, 0xf0, 0x1e,
	0x28, 0xdf, 0x4e, 0xd2, 0xbb, 0x9c, 0xe3, 0x7f, 0x8e, 0x7f, 0x3e, 0x5f, 0x60, 0xfe, 0x88, 0x42,
	0x1e, 0x70, 0x41, 0x43, 0xb1, 0xe4, 0x22, 0x62, 0xfe, 0x9a, 0x3a, 0x31, 0x8b, 0x44, 0x84, 0xff,
	0xa9, 0xa0, 0x5f, 0x33, 0x3f, 0x8e, 0xe9, 0x8a, 0xd0, 0x5f, 0x09, 0xe5, 0x02, 0xe9, 0xd0, 0x0b,
	0x56, 0xa6, 0x6a, 0xab, 0xd3, 0x11, 0xe9, 0x05, 0x2b, 0xf4, 0x14, 0x46, 0x3c, 0x58, 0x87, 0xbe,
	0x48, 0x18, 0x35, 0x7b, 0xb6, 0x3a, 0x1d, 0x93, 0xda, 0x81, 0x9e, 0x43, 0x7f, 0x4d, 0x85, 0x09,
	0xb6, 0x3a, 0xd5, 0x66, 0x9a, 0x73, 0x4e, 0x45, 0x11, 0xe7, 0x93, 0x42, 0xd2, 0x93, 0x54, 0x10,
	0x27, 0xc2, 0xd4, 0x0a, 0xc1, 0x45, 0x22, 0x0b, 0xe2, 0x44, 0xa0, 0x29, 0x0c, 0x18, 0xbd, 0x8b,
	0xee, 0xa9, 0x39, 0xce, 0x34, 0xba, 0x43, 0x32, 0xb3, 0x96, 0x15, 0xe7, 0xc8, 0x81, 0xa1, 0x60,
	0x7e, 0xc8, 0x7f, 0x52, 0x66, 0xee, 0x65, 0x5a, 0xc3, 0x59, 0x14, 0x8e, 0x5a, 0x5d, 0x69, 0x4e,
	0x47, 0xb0, 0xcb, 0x72, 0x37, 0xfe, 0xab, 0xc2, 0xa4, 0x7a, 0x27, 0x8f, 0xa3, 0x90, 0x53, 0x64,
	0xcb, 0xe8, 0xe3, 0x1c, 0x3d, 0x3f, 0x2a, 0xd9, 0x6d, 0x99, 0x7d, 0x9c, 0xb3, 0xd7, 0x8a, 0x14,
	0xfe, 0x65, 0x0b, 0x7e, 0x52, 0xc1, 0x57, 0xba, 0x92, 0xfe, 0x4d, 0x87, 0x7e, 0x5f, 0xa2, 0xaf,
	0xe4, 0x35, 0x3e, 0xc0, 0x90, 0x15, 0x7e, 0xfc, 0x0c, 0xa0, 0x4e, 0x2d, 0x32, 0xa0, 0xbf, 0xa1,
	0x8f, 0x45, 0x8d, 0xd2, 0x4f, 0xbc, 0x01, 0x4d, 0xe2, 0x47, 0x87, 0x30, 0xa0, 0xbf, 0x03, 0x2e,
	0x78, 0xa6, 0x19, 0x92, 0xc2, 0x42, 0x07, 0xf0, 0xe4, 0xde, 0xbf, 0x4d, 0xca, 0x3a, 0xe6, 0x06,
	0x32, 0x61, 0x37, 0xe0, 0xde, 0x43, 0x48, 0x99, 0xd9, 0xcf, 0xe4, 0xa5, 0x99, 0xea, 0xa3, 0xcc,
	0xbf, 0x93, 0x5d, 0x95, 0x1b, 0xf8, 0x1d, 0x40, 0x5d, 0xc6, 0x2e, 0xcc, 0xf6, 0x5b, 0xf0, 0x31,
	0x68, 0x52, 0x02, 0xd3, 0x36, 0x8b, 0x36, 0x05, 0x5e, 0x2f, 0xda, 0xd4, 0x57, 0xf5, 0xe4, 0xab,
	0x5e, 0xc0, 0x5e, 0xa3, 0x1b, 0xb6, 0x3c, 0xfd, 0x0a, 0xf4, 0x66, 0xce, 0xd3, 0xf7, 0xe4, 0x39,
	0x5f, 0x15, 0xf1, 0x4b, 0x13, 0x1d, 0xc1, 0x88, 0x32, 0x16, 0xb1, 0xc5, 0x63, 0x9c, 0xd3, 0xe9,
	0xb3, 0x71, 0x51, 0x31, 0x37, 0xf5, 0x93, 0xfa, 0x18, 0x9f, 0xc0, 0xa4, 0xd5, 0x5c, 0x5b, 0x9e,
	0x6a, 0xc1, 0x30, 0xa4, 0x0f, 0x9e, 0x04, 0x5e, 0xd9, 0xf8, 0x06, 0x8c, 0x76, 0x7d, 0x91, 0x0d,
	0x5a, 0x59, 0x5f, 0x56, 0xe1, 0xc9, 0x2e, 0xf4, 0xba, 0x8b, 0xa8, 0x57, 0x7d, 0xd2, 0x86, 0x3c,
	0x7a, 0x0f, 0x9a, 0x84, 0x8f, 0x0e, 0xc0, 0x20, 0xee, 0x57, 0xef, 0xca, 0x5d, 0x7e, 0x71, 0xbf,
	0x2d, 0x5d, 0x42, 0x3c, 0x62, 0x28, 0x92, 0x77, 0xee, 0x2d, 0x96, 0xde, 0xf5, 0xdc, 0x25, 0x86,
	0x7a, 0x74, 0x02, 0x7b, 0x8d, 0xb0, 0xe8, 0x10, 0xd0, 0x82, 0x7c, 0x98, 0x5f, 0x7e, 0x74, 0x49,
	0xe3, 0x77, 0xd9, 0x2f, 0x05, 0x98, 0x7d, 0x06, 0xf3, 0x8c, 0x86, 0x82, 0xf9, 0xb7, 0x67, 0xd5,
	0x7a, 0xb9, 0xcc, 0xb7, 0x0b, 0x72, 0x60, 0xb7, 0x4c, 0xda, 0xc4, 0x69, 0x2e, 0x18, 0xcb, 0x70,
	0x5a, 0x93, 0x88, 0x95, 0xd9, 0x1f, 0x15, 0xf6, 0xbb, 0x51, 0x30, 0xf4, 0xcf, 0xa9, 0x40, 0xf2,
	0x5a, 0xb1, 0x1a, 0x83, 0x8a, 0x95, 0x54, 0x73, 0x91, 0xa4, 0x9a, 0xba, 0x25, 0xad, 0xc6, 0xa8,
	0x62, 0x05, 0xbd, 0x82, 0x41, 0x9e, 0x25, 0xd4, 0x5a, 0x2e, 0x56, 0x7b, 0x5e, 0xb1, 0x82, 0xde,
	0xc2, 0xb0, 0xcc, 0x0b, 0xea, 0xec, 0x17, 0xab, 0x3b, 0xb3, 0x58, 0x39, 0x1d, 0x7c, 0xdf, 0xb9,
	0xf3, 0x83, 0xf0, 0x66, 0x90, 0x2d, 0xd5, 0xe3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11, 0x6f,
	0xae, 0xf9, 0x70, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CentralConsistentStorageClient is the client API for CentralConsistentStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CentralConsistentStorageClient interface {
	Request(ctx context.Context, in *WrappedRequest, opts ...grpc.CallOption) (*WrappedResponse, error)
}

type centralConsistentStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewCentralConsistentStorageClient(cc grpc.ClientConnInterface) CentralConsistentStorageClient {
	return &centralConsistentStorageClient{cc}
}

func (c *centralConsistentStorageClient) Request(ctx context.Context, in *WrappedRequest, opts ...grpc.CallOption) (*WrappedResponse, error) {
	out := new(WrappedResponse)
	err := c.cc.Invoke(ctx, "/CentralConsistentStorage/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CentralConsistentStorageServer is the server API for CentralConsistentStorage service.
type CentralConsistentStorageServer interface {
	Request(context.Context, *WrappedRequest) (*WrappedResponse, error)
}

// UnimplementedCentralConsistentStorageServer can be embedded to have forward compatible implementations.
type UnimplementedCentralConsistentStorageServer struct {
}

func (*UnimplementedCentralConsistentStorageServer) Request(ctx context.Context, req *WrappedRequest) (*WrappedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}

func RegisterCentralConsistentStorageServer(s *grpc.Server, srv CentralConsistentStorageServer) {
	s.RegisterService(&_CentralConsistentStorage_serviceDesc, srv)
}

func _CentralConsistentStorage_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WrappedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralConsistentStorageServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CentralConsistentStorage/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralConsistentStorageServer).Request(ctx, req.(*WrappedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CentralConsistentStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CentralConsistentStorage",
	HandlerType: (*CentralConsistentStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _CentralConsistentStorage_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consistent_storage.proto",
}

// ConsistentStorageClient is the client API for ConsistentStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsistentStorageClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
}

type consistentStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewConsistentStorageClient(cc grpc.ClientConnInterface) ConsistentStorageClient {
	return &consistentStorageClient{cc}
}

func (c *consistentStorageClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/ConsistentStorage/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consistentStorageClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/ConsistentStorage/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consistentStorageClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/ConsistentStorage/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consistentStorageClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, "/ConsistentStorage/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsistentStorageServer is the server API for ConsistentStorage service.
type ConsistentStorageServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	Transfer(context.Context, *TransferRequest) (*TransferResponse, error)
}

// UnimplementedConsistentStorageServer can be embedded to have forward compatible implementations.
type UnimplementedConsistentStorageServer struct {
}

func (*UnimplementedConsistentStorageServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedConsistentStorageServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedConsistentStorageServer) Remove(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedConsistentStorageServer) Transfer(ctx context.Context, req *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}

func RegisterConsistentStorageServer(s *grpc.Server, srv ConsistentStorageServer) {
	s.RegisterService(&_ConsistentStorage_serviceDesc, srv)
}

func _ConsistentStorage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentStorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConsistentStorage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentStorageServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsistentStorage_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentStorageServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConsistentStorage/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentStorageServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsistentStorage_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentStorageServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConsistentStorage/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentStorageServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsistentStorage_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentStorageServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConsistentStorage/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentStorageServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsistentStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ConsistentStorage",
	HandlerType: (*ConsistentStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ConsistentStorage_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _ConsistentStorage_Put_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _ConsistentStorage_Remove_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _ConsistentStorage_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consistent_storage.proto",
}
