// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sagas.proto

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

type PutError int32

const (
	PutError_PUT_KEY_EXISTS PutError = 0
)

var PutError_name = map[int32]string{
	0: "PUT_KEY_EXISTS",
}

var PutError_value = map[string]int32{
	"PUT_KEY_EXISTS": 0,
}

func (x PutError) String() string {
	return proto.EnumName(PutError_name, int32(x))
}

func (PutError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_774b0cf37958ccf0, []int{0}
}

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
	return fileDescriptor_774b0cf37958ccf0, []int{1}
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
	return fileDescriptor_774b0cf37958ccf0, []int{0}
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_774b0cf37958ccf0, []int{1}
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
	return fileDescriptor_774b0cf37958ccf0, []int{2}
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
	ErrorType            PutError `protobuf:"varint,2,opt,name=errorType,proto3,enum=PutError" json:"errorType,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,3,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_774b0cf37958ccf0, []int{3}
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

func (m *PutResponse) GetErrorType() PutError {
	if m != nil {
		return m.ErrorType
	}
	return PutError_PUT_KEY_EXISTS
}

func (m *PutResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
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
	return fileDescriptor_774b0cf37958ccf0, []int{4}
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
	ErrorMsg             string      `protobuf:"bytes,3,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_774b0cf37958ccf0, []int{5}
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

func (m *RemoveResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func init() {
	proto.RegisterEnum("PutError", PutError_name, PutError_value)
	proto.RegisterEnum("RemoveError", RemoveError_name, RemoveError_value)
	proto.RegisterType((*GetRequest)(nil), "GetRequest")
	proto.RegisterType((*GetResponse)(nil), "GetResponse")
	proto.RegisterType((*PutRequest)(nil), "PutRequest")
	proto.RegisterType((*PutResponse)(nil), "PutResponse")
	proto.RegisterType((*RemoveRequest)(nil), "RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "RemoveResponse")
}

func init() {
	proto.RegisterFile("sagas.proto", fileDescriptor_774b0cf37958ccf0)
}

var fileDescriptor_774b0cf37958ccf0 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x6a, 0xf2, 0x40,
	0x14, 0x4d, 0xf4, 0xfb, 0x52, 0xbd, 0x49, 0xd3, 0x30, 0x88, 0x84, 0x2c, 0xc4, 0xce, 0xa6, 0x62,
	0x61, 0x16, 0xb6, 0x9b, 0xd2, 0x5d, 0x4b, 0x90, 0x52, 0x34, 0x61, 0x92, 0xfe, 0x6e, 0x24, 0xd2,
	0x41, 0xc4, 0x9a, 0xb1, 0x99, 0x89, 0xd4, 0x87, 0xe8, 0x3b, 0x97, 0xfc, 0x58, 0x23, 0xc5, 0x45,
	0x77, 0x73, 0xee, 0x9c, 0x33, 0xf7, 0xcc, 0x3d, 0x17, 0x74, 0x11, 0xcd, 0x22, 0x41, 0x56, 0x09,
	0x97, 0x1c, 0x77, 0x00, 0x86, 0x4c, 0x52, 0xf6, 0x91, 0x32, 0x21, 0x91, 0x05, 0xf5, 0x05, 0xdb,
	0xd8, 0x6a, 0x57, 0xed, 0x35, 0x69, 0x76, 0xc4, 0xd7, 0xa0, 0xe7, 0xf7, 0x62, 0xc5, 0x63, 0xc1,
	0x50, 0x1b, 0x34, 0xf6, 0x39, 0x17, 0x52, 0xe4, 0x9c, 0x06, 0x2d, 0x11, 0x6a, 0xc1, 0xff, 0x75,
	0xf4, 0x9e, 0x32, 0xbb, 0xd6, 0x55, 0x7b, 0x06, 0x2d, 0x00, 0xbe, 0x04, 0xf0, 0xd3, 0xc3, 0x8f,
	0x1f, 0x50, 0x4d, 0x41, 0xcf, 0x55, 0x65, 0x4b, 0x13, 0x6a, 0x7c, 0x51, 0xb6, 0xab, 0xf1, 0x05,
	0x3a, 0x83, 0x26, 0x4b, 0x12, 0x9e, 0x84, 0x9b, 0x55, 0x21, 0x34, 0x07, 0x4d, 0xe2, 0xa7, 0xd2,
	0xcd, 0x8a, 0x74, 0x77, 0x87, 0x1c, 0x68, 0xe4, 0x60, 0x24, 0x66, 0x76, 0x3d, 0x6f, 0xfa, 0x83,
	0xf1, 0x29, 0x1c, 0x53, 0xb6, 0xe4, 0x6b, 0x76, 0xf8, 0xe7, 0x09, 0x98, 0x5b, 0x4a, 0xe9, 0xc4,
	0x86, 0xa3, 0x24, 0xaf, 0xbc, 0x95, 0x76, 0xb6, 0x10, 0xf5, 0x7f, 0x7b, 0x32, 0x48, 0xa1, 0xfe,
	0x8b, 0xad, 0x7e, 0x07, 0x1a, 0xdb, 0x9f, 0x20, 0x04, 0xa6, 0xff, 0x10, 0x4e, 0xee, 0xdd, 0x97,
	0x89, 0xfb, 0x7c, 0x17, 0x84, 0x81, 0xa5, 0xf4, 0xaf, 0x40, 0xaf, 0xbc, 0x8a, 0x5a, 0x60, 0x51,
	0x77, 0xe4, 0x3d, 0xba, 0x05, 0x8b, 0x52, 0x8f, 0x5a, 0x4a, 0xa5, 0x3a, 0xf6, 0xc2, 0x89, 0xf7,
	0x34, 0x76, 0xa9, 0xa5, 0x0e, 0xbe, 0x54, 0x68, 0x07, 0x59, 0xf0, 0xb7, 0x3c, 0x16, 0x73, 0x21,
	0x59, 0x2c, 0x03, 0xc9, 0x93, 0x68, 0xc6, 0x10, 0x86, 0xfa, 0x90, 0x49, 0xa4, 0x93, 0xdd, 0x26,
	0x38, 0x06, 0xa9, 0xc4, 0x8e, 0x95, 0x8c, 0xe3, 0xa7, 0x19, 0x67, 0x17, 0xa8, 0x63, 0x90, 0x4a,
	0x4e, 0x58, 0x41, 0xe7, 0xa0, 0x15, 0xee, 0x90, 0x49, 0xf6, 0xa6, 0xeb, 0x9c, 0x90, 0xfd, 0x51,
	0x62, 0xe5, 0x46, 0x7b, 0xfd, 0xb7, 0x8c, 0xe6, 0xf1, 0x54, 0xcb, 0xf7, 0xf0, 0xe2, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0xb5, 0x9f, 0x9e, 0x43, 0x96, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SagasConsistentStorageClient is the client API for SagasConsistentStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SagasConsistentStorageClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
}

type sagasConsistentStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewSagasConsistentStorageClient(cc grpc.ClientConnInterface) SagasConsistentStorageClient {
	return &sagasConsistentStorageClient{cc}
}

func (c *sagasConsistentStorageClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/SagasConsistentStorage/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sagasConsistentStorageClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/SagasConsistentStorage/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sagasConsistentStorageClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/SagasConsistentStorage/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SagasConsistentStorageServer is the server API for SagasConsistentStorage service.
type SagasConsistentStorageServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
}

// UnimplementedSagasConsistentStorageServer can be embedded to have forward compatible implementations.
type UnimplementedSagasConsistentStorageServer struct {
}

func (*UnimplementedSagasConsistentStorageServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSagasConsistentStorageServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedSagasConsistentStorageServer) Remove(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func RegisterSagasConsistentStorageServer(s *grpc.Server, srv SagasConsistentStorageServer) {
	s.RegisterService(&_SagasConsistentStorage_serviceDesc, srv)
}

func _SagasConsistentStorage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SagasConsistentStorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SagasConsistentStorage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SagasConsistentStorageServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SagasConsistentStorage_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SagasConsistentStorageServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SagasConsistentStorage/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SagasConsistentStorageServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SagasConsistentStorage_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SagasConsistentStorageServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SagasConsistentStorage/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SagasConsistentStorageServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SagasConsistentStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SagasConsistentStorage",
	HandlerType: (*SagasConsistentStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SagasConsistentStorage_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _SagasConsistentStorage_Put_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _SagasConsistentStorage_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sagas.proto",
}
