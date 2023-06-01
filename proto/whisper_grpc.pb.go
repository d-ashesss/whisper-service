// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: whisper.proto

package whisperpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	WhisperService_Transcribe_FullMethodName = "/proto.WhisperService/Transcribe"
)

// WhisperServiceClient is the client API for WhisperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WhisperServiceClient interface {
	Transcribe(ctx context.Context, opts ...grpc.CallOption) (WhisperService_TranscribeClient, error)
}

type whisperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWhisperServiceClient(cc grpc.ClientConnInterface) WhisperServiceClient {
	return &whisperServiceClient{cc}
}

func (c *whisperServiceClient) Transcribe(ctx context.Context, opts ...grpc.CallOption) (WhisperService_TranscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &WhisperService_ServiceDesc.Streams[0], WhisperService_Transcribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &whisperServiceTranscribeClient{stream}
	return x, nil
}

type WhisperService_TranscribeClient interface {
	Send(*TranscribeRequest) error
	CloseAndRecv() (*TranscribeResponse, error)
	grpc.ClientStream
}

type whisperServiceTranscribeClient struct {
	grpc.ClientStream
}

func (x *whisperServiceTranscribeClient) Send(m *TranscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *whisperServiceTranscribeClient) CloseAndRecv() (*TranscribeResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TranscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WhisperServiceServer is the server API for WhisperService service.
// All implementations must embed UnimplementedWhisperServiceServer
// for forward compatibility
type WhisperServiceServer interface {
	Transcribe(WhisperService_TranscribeServer) error
	mustEmbedUnimplementedWhisperServiceServer()
}

// UnimplementedWhisperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWhisperServiceServer struct {
}

func (UnimplementedWhisperServiceServer) Transcribe(WhisperService_TranscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Transcribe not implemented")
}
func (UnimplementedWhisperServiceServer) mustEmbedUnimplementedWhisperServiceServer() {}

// UnsafeWhisperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WhisperServiceServer will
// result in compilation errors.
type UnsafeWhisperServiceServer interface {
	mustEmbedUnimplementedWhisperServiceServer()
}

func RegisterWhisperServiceServer(s grpc.ServiceRegistrar, srv WhisperServiceServer) {
	s.RegisterService(&WhisperService_ServiceDesc, srv)
}

func _WhisperService_Transcribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WhisperServiceServer).Transcribe(&whisperServiceTranscribeServer{stream})
}

type WhisperService_TranscribeServer interface {
	SendAndClose(*TranscribeResponse) error
	Recv() (*TranscribeRequest, error)
	grpc.ServerStream
}

type whisperServiceTranscribeServer struct {
	grpc.ServerStream
}

func (x *whisperServiceTranscribeServer) SendAndClose(m *TranscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *whisperServiceTranscribeServer) Recv() (*TranscribeRequest, error) {
	m := new(TranscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WhisperService_ServiceDesc is the grpc.ServiceDesc for WhisperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WhisperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WhisperService",
	HandlerType: (*WhisperServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transcribe",
			Handler:       _WhisperService_Transcribe_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "whisper.proto",
}
