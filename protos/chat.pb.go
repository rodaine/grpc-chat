// Code generated by protoc-gen-go.
// source: protos/chat.proto
// DO NOT EDIT!

/*
Package chat is a generated protocol buffer package.

It is generated from these files:
	protos/chat.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
	StreamRequest
	StreamResponse
*/
package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutResponse struct {
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type StreamRequest struct {
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StreamRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StreamResponse struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*StreamResponse_ClientLogin
	//	*StreamResponse_ClientLogout
	//	*StreamResponse_ClientMessage
	//	*StreamResponse_ServerShutdown
	Event isStreamResponse_Event `protobuf_oneof:"event"`
}

func (m *StreamResponse) Reset()                    { *m = StreamResponse{} }
func (m *StreamResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()               {}
func (*StreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isStreamResponse_Event interface {
	isStreamResponse_Event()
}

type StreamResponse_ClientLogin struct {
	ClientLogin *StreamResponse_Login `protobuf:"bytes,2,opt,name=client_login,json=clientLogin,oneof"`
}
type StreamResponse_ClientLogout struct {
	ClientLogout *StreamResponse_Logout `protobuf:"bytes,3,opt,name=client_logout,json=clientLogout,oneof"`
}
type StreamResponse_ClientMessage struct {
	ClientMessage *StreamResponse_Message `protobuf:"bytes,4,opt,name=client_message,json=clientMessage,oneof"`
}
type StreamResponse_ServerShutdown struct {
	ServerShutdown *StreamResponse_Shutdown `protobuf:"bytes,5,opt,name=server_shutdown,json=serverShutdown,oneof"`
}

func (*StreamResponse_ClientLogin) isStreamResponse_Event()    {}
func (*StreamResponse_ClientLogout) isStreamResponse_Event()   {}
func (*StreamResponse_ClientMessage) isStreamResponse_Event()  {}
func (*StreamResponse_ServerShutdown) isStreamResponse_Event() {}

func (m *StreamResponse) GetEvent() isStreamResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *StreamResponse) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *StreamResponse) GetClientLogin() *StreamResponse_Login {
	if x, ok := m.GetEvent().(*StreamResponse_ClientLogin); ok {
		return x.ClientLogin
	}
	return nil
}

func (m *StreamResponse) GetClientLogout() *StreamResponse_Logout {
	if x, ok := m.GetEvent().(*StreamResponse_ClientLogout); ok {
		return x.ClientLogout
	}
	return nil
}

func (m *StreamResponse) GetClientMessage() *StreamResponse_Message {
	if x, ok := m.GetEvent().(*StreamResponse_ClientMessage); ok {
		return x.ClientMessage
	}
	return nil
}

func (m *StreamResponse) GetServerShutdown() *StreamResponse_Shutdown {
	if x, ok := m.GetEvent().(*StreamResponse_ServerShutdown); ok {
		return x.ServerShutdown
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamResponse_OneofMarshaler, _StreamResponse_OneofUnmarshaler, _StreamResponse_OneofSizer, []interface{}{
		(*StreamResponse_ClientLogin)(nil),
		(*StreamResponse_ClientLogout)(nil),
		(*StreamResponse_ClientMessage)(nil),
		(*StreamResponse_ServerShutdown)(nil),
	}
}

func _StreamResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamResponse)
	// event
	switch x := m.Event.(type) {
	case *StreamResponse_ClientLogin:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientLogin); err != nil {
			return err
		}
	case *StreamResponse_ClientLogout:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientLogout); err != nil {
			return err
		}
	case *StreamResponse_ClientMessage:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientMessage); err != nil {
			return err
		}
	case *StreamResponse_ServerShutdown:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ServerShutdown); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamResponse.Event has unexpected type %T", x)
	}
	return nil
}

func _StreamResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamResponse)
	switch tag {
	case 2: // event.client_login
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Login)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ClientLogin{msg}
		return true, err
	case 3: // event.client_logout
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Logout)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ClientLogout{msg}
		return true, err
	case 4: // event.client_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Message)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ClientMessage{msg}
		return true, err
	case 5: // event.server_shutdown
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Shutdown)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ServerShutdown{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamResponse)
	// event
	switch x := m.Event.(type) {
	case *StreamResponse_ClientLogin:
		s := proto.Size(x.ClientLogin)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_ClientLogout:
		s := proto.Size(x.ClientLogout)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_ClientMessage:
		s := proto.Size(x.ClientMessage)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_ServerShutdown:
		s := proto.Size(x.ServerShutdown)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamResponse_Login struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StreamResponse_Login) Reset()                    { *m = StreamResponse_Login{} }
func (m *StreamResponse_Login) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse_Login) ProtoMessage()               {}
func (*StreamResponse_Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *StreamResponse_Login) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamResponse_Logout struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StreamResponse_Logout) Reset()                    { *m = StreamResponse_Logout{} }
func (m *StreamResponse_Logout) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse_Logout) ProtoMessage()               {}
func (*StreamResponse_Logout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 1} }

func (m *StreamResponse_Logout) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamResponse_Message struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *StreamResponse_Message) Reset()                    { *m = StreamResponse_Message{} }
func (m *StreamResponse_Message) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse_Message) ProtoMessage()               {}
func (*StreamResponse_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 2} }

func (m *StreamResponse_Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamResponse_Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StreamResponse_Shutdown struct {
}

func (m *StreamResponse_Shutdown) Reset()                    { *m = StreamResponse_Shutdown{} }
func (m *StreamResponse_Shutdown) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse_Shutdown) ProtoMessage()               {}
func (*StreamResponse_Shutdown) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 3} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "chat.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "chat.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "chat.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "chat.LogoutResponse")
	proto.RegisterType((*StreamRequest)(nil), "chat.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "chat.StreamResponse")
	proto.RegisterType((*StreamResponse_Login)(nil), "chat.StreamResponse.Login")
	proto.RegisterType((*StreamResponse_Logout)(nil), "chat.StreamResponse.Logout")
	proto.RegisterType((*StreamResponse_Message)(nil), "chat.StreamResponse.Message")
	proto.RegisterType((*StreamResponse_Shutdown)(nil), "chat.StreamResponse.Shutdown")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chat service

type ChatClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Chat_StreamClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/chat.Chat/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := grpc.Invoke(ctx, "/chat.Chat/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Chat_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Chat_serviceDesc.Streams[0], c.cc, "/chat.Chat/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatStreamClient{stream}
	return x, nil
}

type Chat_StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type chatStreamClient struct {
	grpc.ClientStream
}

func (x *chatStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Chat service

type ChatServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Stream(Chat_StreamServer) error
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Stream(&chatStreamServer{stream})
}

type Chat_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type chatStreamServer struct {
	grpc.ServerStream
}

func (x *chatStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Chat_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Chat_Logout_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Chat_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/chat.proto",
}

func init() { proto.RegisterFile("protos/chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xb5, 0x8b, 0xcd, 0xc7, 0x00, 0x6e, 0xbb, 0xe5, 0x60, 0x2d, 0x54, 0xad, 0x2c, 0x55, 0xa2,
	0x17, 0x53, 0x51, 0x55, 0x6d, 0x2f, 0x89, 0x44, 0x14, 0xc9, 0x87, 0xe4, 0x62, 0x72, 0x47, 0x06,
	0x36, 0x06, 0x05, 0x7b, 0x1d, 0xef, 0x1a, 0xfe, 0x54, 0xfe, 0x58, 0xfe, 0x45, 0xc4, 0x7e, 0xd8,
	0xb1, 0x64, 0x6e, 0x3b, 0xe3, 0xf7, 0xde, 0xcc, 0x3c, 0x3f, 0xf8, 0x9c, 0xe5, 0x94, 0x53, 0x36,
	0xdb, 0xec, 0x22, 0xee, 0x8b, 0x37, 0xb2, 0xce, 0x6f, 0xfc, 0x2d, 0xa6, 0x34, 0x3e, 0x90, 0x99,
	0xe8, 0xad, 0x8b, 0xc7, 0x19, 0xdf, 0x27, 0x84, 0xf1, 0x28, 0xc9, 0x24, 0xcc, 0xbb, 0x82, 0xc1,
	0x1d, 0x8d, 0xf7, 0x69, 0x48, 0x9e, 0x0b, 0xc2, 0x38, 0xc2, 0xd0, 0xcd, 0x22, 0xc6, 0x4e, 0x34,
	0xdf, 0xba, 0xe6, 0x77, 0x73, 0xda, 0x0b, 0xcb, 0x1a, 0x21, 0xb0, 0xd2, 0x28, 0x21, 0xee, 0x07,
	0xd1, 0x17, 0x6f, 0xef, 0x07, 0x0c, 0x15, 0x9f, 0x65, 0x34, 0x65, 0x04, 0x8d, 0xc0, 0xe6, 0xf4,
	0x89, 0xa4, 0x8a, 0x2d, 0x0b, 0x05, 0xa3, 0x05, 0xd7, 0x73, 0x9a, 0x61, 0x9f, 0xc0, 0xd1, 0x30,
	0x29, 0xe7, 0xfd, 0x84, 0xe1, 0x92, 0xe7, 0x24, 0x4a, 0x34, 0xd1, 0x85, 0x4e, 0x42, 0x18, 0x8b,
	0x62, 0xbd, 0x87, 0x2e, 0xbd, 0xd7, 0x16, 0x38, 0x1a, 0xab, 0x96, 0xf9, 0x07, 0xbd, 0xf2, 0x60,
	0x31, 0xa9, 0x3f, 0xc7, 0xbe, 0xb4, 0xc4, 0xd7, 0x96, 0xf8, 0x0f, 0x1a, 0x11, 0x56, 0x60, 0x74,
	0x0d, 0x83, 0xcd, 0x61, 0x4f, 0x52, 0xbe, 0x3a, 0x9c, 0xcf, 0x13, 0xb3, 0xce, 0x64, 0xe1, 0x70,
	0x7d, 0x8a, 0x2f, 0x0c, 0x08, 0x8c, 0xb0, 0x2f, 0x19, 0xa2, 0x44, 0x0b, 0x18, 0x56, 0x02, 0xb4,
	0xe0, 0x6e, 0x4b, 0x28, 0x8c, 0x2f, 0x29, 0xd0, 0x82, 0x07, 0x46, 0x38, 0x28, 0x25, 0x68, 0xc1,
	0xd1, 0x2d, 0x38, 0x4a, 0x43, 0x9f, 0x6c, 0x09, 0x91, 0x49, 0xa3, 0xc8, 0xbd, 0xc4, 0x04, 0x46,
	0xa8, 0x26, 0xab, 0x06, 0x0a, 0xe0, 0x23, 0x23, 0xf9, 0x91, 0xe4, 0x2b, 0xb6, 0x2b, 0xf8, 0x96,
	0x9e, 0x52, 0xd7, 0x16, 0x3a, 0x5f, 0x1b, 0x75, 0x96, 0x0a, 0x14, 0x18, 0xa1, 0x23, 0x79, 0xba,
	0x83, 0xc7, 0x60, 0xcb, 0xeb, 0x74, 0x14, 0xcc, 0x2a, 0x0a, 0x78, 0x02, 0x6d, 0xb5, 0x77, 0xd3,
	0xd7, 0xbf, 0xd0, 0xd1, 0xfb, 0x34, 0x7c, 0xbe, 0xfc, 0x5b, 0x31, 0x40, 0x57, 0xcf, 0x5f, 0x74,
	0xc0, 0x26, 0x47, 0x92, 0xf2, 0xf9, 0x8b, 0x09, 0xd6, 0xcd, 0x2e, 0xe2, 0x68, 0x5e, 0x6e, 0x24,
	0x6f, 0x79, 0x1f, 0x66, 0xfc, 0xa5, 0xd6, 0x53, 0x89, 0x32, 0xd0, 0x9f, 0x72, 0xd1, 0x0a, 0x50,
	0x45, 0x13, 0x8f, 0xea, 0xcd, 0x92, 0xf6, 0x1f, 0xda, 0xd2, 0x29, 0x4d, 0xab, 0x05, 0x53, 0xd3,
	0xea, 0x66, 0x7a, 0xc6, 0xd4, 0xfc, 0x65, 0xae, 0xdb, 0x22, 0x6c, 0xbf, 0xdf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x46, 0x94, 0x28, 0x61, 0xa8, 0x03, 0x00, 0x00,
}
