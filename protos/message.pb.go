// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/message.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message struct {
	Id         int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FromUserId int64 `protobuf:"varint,2,opt,name=fromUserId" json:"fromUserId,omitempty"`
	ToUserId   int64 `protobuf:"varint,3,opt,name=toUserId" json:"toUserId,omitempty"`
	// 时间戳ms
	CreatedAt int64  `protobuf:"varint,4,opt,name=createdAt" json:"createdAt,omitempty"`
	Title     string `protobuf:"bytes,5,opt,name=title" json:"title,omitempty"`
	Message   string `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	Status    string `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	IsDelete  bool   `protobuf:"varint,8,opt,name=isDelete" json:"isDelete,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Message) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message) GetFromUserId() int64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func (m *Message) GetToUserId() int64 {
	if m != nil {
		return m.ToUserId
	}
	return 0
}

func (m *Message) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Message) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Message) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Message) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

type MessageResponse struct {
	// 创建成功id
	Mid int64 `protobuf:"varint,1,opt,name=mid" json:"mid,omitempty"`
}

func (m *MessageResponse) Reset()                    { *m = MessageResponse{} }
func (m *MessageResponse) String() string            { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()               {}
func (*MessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *MessageResponse) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

type MessageReadRequest struct {
	// 分页
	PageNumber int32 `protobuf:"varint,1,opt,name=pageNumber" json:"pageNumber,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=pageSize" json:"pageSize,omitempty"`
	// msgType: UNSEEN/SEEN/DETAIL/ALL
	MsgType    string `protobuf:"bytes,3,opt,name=msgType" json:"msgType,omitempty"`
	Id         int64  `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	FromUserId int64  `protobuf:"varint,5,opt,name=FromUserId" json:"FromUserId,omitempty"`
	ToUserId   int64  `protobuf:"varint,6,opt,name=toUserId" json:"toUserId,omitempty"`
	Title      string `protobuf:"bytes,7,opt,name=Title" json:"Title,omitempty"`
	Message    string `protobuf:"bytes,8,opt,name=Message" json:"Message,omitempty"`
	Status     string `protobuf:"bytes,9,opt,name=status" json:"status,omitempty"`
	IsDelete   string `protobuf:"bytes,10,opt,name=isDelete" json:"isDelete,omitempty"`
}

func (m *MessageReadRequest) Reset()                    { *m = MessageReadRequest{} }
func (m *MessageReadRequest) String() string            { return proto.CompactTextString(m) }
func (*MessageReadRequest) ProtoMessage()               {}
func (*MessageReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *MessageReadRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *MessageReadRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *MessageReadRequest) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MessageReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MessageReadRequest) GetFromUserId() int64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func (m *MessageReadRequest) GetToUserId() int64 {
	if m != nil {
		return m.ToUserId
	}
	return 0
}

func (m *MessageReadRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MessageReadRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MessageReadRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *MessageReadRequest) GetIsDelete() string {
	if m != nil {
		return m.IsDelete
	}
	return ""
}

type MessageListResponse struct {
	// 数量总数
	PageNumber int32      `protobuf:"varint,1,opt,name=pageNumber" json:"pageNumber,omitempty"`
	PageSize   int32      `protobuf:"varint,2,opt,name=pageSize" json:"pageSize,omitempty"`
	Nums       int64      `protobuf:"varint,3,opt,name=nums" json:"nums,omitempty"`
	MsgList    []*Message `protobuf:"bytes,4,rep,name=msgList" json:"msgList,omitempty"`
}

func (m *MessageListResponse) Reset()                    { *m = MessageListResponse{} }
func (m *MessageListResponse) String() string            { return proto.CompactTextString(m) }
func (*MessageListResponse) ProtoMessage()               {}
func (*MessageListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *MessageListResponse) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *MessageListResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *MessageListResponse) GetNums() int64 {
	if m != nil {
		return m.Nums
	}
	return 0
}

func (m *MessageListResponse) GetMsgList() []*Message {
	if m != nil {
		return m.MsgList
	}
	return nil
}

type MessageDeleteRequest struct {
	// 修改id对应的msg.is_delete true
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *MessageDeleteRequest) Reset()                    { *m = MessageDeleteRequest{} }
func (m *MessageDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*MessageDeleteRequest) ProtoMessage()               {}
func (*MessageDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *MessageDeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "protos.Message")
	proto.RegisterType((*MessageResponse)(nil), "protos.MessageResponse")
	proto.RegisterType((*MessageReadRequest)(nil), "protos.MessageReadRequest")
	proto.RegisterType((*MessageListResponse)(nil), "protos.MessageListResponse")
	proto.RegisterType((*MessageDeleteRequest)(nil), "protos.MessageDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MessageService service

type MessageServiceClient interface {
	MessageCreate(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponse, error)
	MessageUpdate(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	MessageRead(ctx context.Context, in *MessageReadRequest, opts ...grpc.CallOption) (*MessageListResponse, error)
	MessageDelete(ctx context.Context, in *MessageDeleteRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) MessageCreate(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := grpc.Invoke(ctx, "/protos.MessageService/MessageCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) MessageUpdate(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/protos.MessageService/MessageUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) MessageRead(ctx context.Context, in *MessageReadRequest, opts ...grpc.CallOption) (*MessageListResponse, error) {
	out := new(MessageListResponse)
	err := grpc.Invoke(ctx, "/protos.MessageService/MessageRead", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) MessageDelete(ctx context.Context, in *MessageDeleteRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := grpc.Invoke(ctx, "/protos.MessageService/MessageDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessageService service

type MessageServiceServer interface {
	MessageCreate(context.Context, *Message) (*MessageResponse, error)
	MessageUpdate(context.Context, *Message) (*Message, error)
	MessageRead(context.Context, *MessageReadRequest) (*MessageListResponse, error)
	MessageDelete(context.Context, *MessageDeleteRequest) (*MessageResponse, error)
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_MessageCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MessageCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MessageService/MessageCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MessageCreate(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_MessageUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MessageUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MessageService/MessageUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MessageUpdate(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_MessageRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MessageRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MessageService/MessageRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MessageRead(ctx, req.(*MessageReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_MessageDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MessageDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MessageService/MessageDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MessageDelete(ctx, req.(*MessageDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageCreate",
			Handler:    _MessageService_MessageCreate_Handler,
		},
		{
			MethodName: "MessageUpdate",
			Handler:    _MessageService_MessageUpdate_Handler,
		},
		{
			MethodName: "MessageRead",
			Handler:    _MessageService_MessageRead_Handler,
		},
		{
			MethodName: "MessageDelete",
			Handler:    _MessageService_MessageDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/message.proto",
}

func init() { proto.RegisterFile("protos/message.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x9b, 0xec, 0x66, 0xff, 0x4c, 0xd5, 0x16, 0x86, 0x15, 0x58, 0x4b, 0x85, 0x56, 0xa9,
	0x84, 0x96, 0xcb, 0x22, 0xb5, 0x47, 0x4e, 0x14, 0x54, 0x81, 0x04, 0x1c, 0xdc, 0xf6, 0x01, 0xd2,
	0xcd, 0xb0, 0x8a, 0x68, 0x9a, 0x10, 0x3b, 0x48, 0xf0, 0x14, 0x9c, 0x38, 0xf0, 0x62, 0xbc, 0x0e,
	0xf2, 0xd8, 0xeb, 0xa4, 0xa6, 0x70, 0xe1, 0x96, 0xef, 0x1b, 0x3b, 0xe3, 0xef, 0xe7, 0x91, 0x61,
	0x56, 0x37, 0x95, 0xae, 0xd4, 0xf3, 0x92, 0x94, 0xca, 0x36, 0xb4, 0x62, 0x89, 0x23, 0xeb, 0xa6,
	0xbf, 0x22, 0x18, 0xbf, 0xb7, 0x15, 0xdc, 0x87, 0xb8, 0xc8, 0x45, 0xb4, 0x88, 0x96, 0x03, 0x19,
	0x17, 0x39, 0x3e, 0x01, 0xf8, 0xd8, 0x54, 0xe5, 0xa5, 0xa2, 0xe6, 0x6d, 0x2e, 0x62, 0xf6, 0x7b,
	0x0e, 0xce, 0x61, 0xa2, 0x2b, 0x57, 0x1d, 0x70, 0xd5, 0x6b, 0x3c, 0x84, 0xe9, 0xba, 0xa1, 0x4c,
	0x53, 0xfe, 0x52, 0x8b, 0x21, 0x17, 0x3b, 0x03, 0x67, 0x90, 0xe8, 0x42, 0x5f, 0x93, 0x48, 0x16,
	0xd1, 0x72, 0x2a, 0xad, 0x40, 0x01, 0x63, 0x77, 0x48, 0x31, 0x62, 0x7f, 0x2b, 0xf1, 0x21, 0x8c,
	0x94, 0xce, 0x74, 0xab, 0xc4, 0x98, 0x0b, 0x4e, 0x99, 0x13, 0x14, 0xea, 0x35, 0x5d, 0x93, 0x26,
	0x31, 0x59, 0x44, 0xcb, 0x89, 0xf4, 0x3a, 0x3d, 0x82, 0x03, 0x17, 0x4c, 0x92, 0xaa, 0xab, 0x1b,
	0x45, 0x78, 0x0f, 0x06, 0xa5, 0x4f, 0x68, 0x3e, 0xd3, 0x9f, 0x31, 0xa0, 0x5f, 0x95, 0xe5, 0x92,
	0x3e, 0xb7, 0xa4, 0xb4, 0x49, 0x5e, 0x67, 0x1b, 0xfa, 0xd0, 0x96, 0x57, 0xd4, 0xf0, 0xfa, 0x44,
	0xf6, 0x1c, 0xd3, 0xd7, 0xa8, 0xf3, 0xe2, 0x1b, 0x31, 0x97, 0x44, 0x7a, 0xcd, 0x29, 0xd4, 0xe6,
	0xe2, 0x6b, 0x4d, 0x0c, 0xc5, 0xa4, 0xb0, 0xd2, 0xf1, 0x1d, 0xf6, 0xf9, 0x9e, 0x75, 0x7c, 0x13,
	0xcb, 0xf7, 0xec, 0x6e, 0xbe, 0xa3, 0x80, 0xef, 0x0c, 0x92, 0x0b, 0x26, 0x68, 0x81, 0x58, 0x61,
	0x7a, 0xbb, 0x34, 0x8c, 0x63, 0x2a, 0xfd, 0xdd, 0x76, 0x04, 0xa7, 0x7f, 0x25, 0x08, 0x5c, 0xe9,
	0x08, 0x7e, 0x8f, 0xe0, 0x81, 0xdb, 0xff, 0xae, 0x50, 0xda, 0x63, 0xfc, 0x1f, 0x3a, 0x08, 0xc3,
	0x9b, 0xb6, 0x54, 0x6e, 0x5e, 0xf8, 0x1b, 0x9f, 0x31, 0x31, 0xd3, 0x42, 0x0c, 0x17, 0x83, 0xe5,
	0xee, 0xf1, 0x81, 0x1d, 0x52, 0xb5, 0xda, 0x5e, 0xcd, 0xb6, 0x9e, 0x3e, 0x85, 0x99, 0xf3, 0xec,
	0x19, 0xb7, 0x17, 0x16, 0x8c, 0xee, 0xf1, 0x8f, 0x18, 0xf6, 0xdd, 0xc2, 0x73, 0x6a, 0xbe, 0x14,
	0x6b, 0xc2, 0x17, 0xb0, 0xe7, 0x9c, 0x57, 0x3c, 0x87, 0x18, 0x76, 0x99, 0x3f, 0x0a, 0xdb, 0xba,
	0xc0, 0xe9, 0x0e, 0x9e, 0xf8, 0xcd, 0x97, 0x75, 0x7e, 0xe7, 0xe6, 0xd0, 0x48, 0x77, 0xf0, 0x0d,
	0xec, 0xf6, 0x66, 0x0b, 0xe7, 0x7f, 0xfc, 0xde, 0x0f, 0xdc, 0xfc, 0x71, 0x50, 0xeb, 0xf3, 0xe6,
	0x3f, 0xed, 0xdd, 0x8a, 0x8d, 0x87, 0xc1, 0xfa, 0x5b, 0x34, 0xfe, 0x11, 0xe4, 0xf4, 0x08, 0xee,
	0xaf, 0xab, 0x72, 0x95, 0xd5, 0x9f, 0xea, 0xb6, 0x71, 0x8f, 0xc1, 0x69, 0x80, 0xea, 0xca, 0x3e,
	0x0e, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xc0, 0x55, 0x54, 0x3b, 0x04, 0x00, 0x00,
}
