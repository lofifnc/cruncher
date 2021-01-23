// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// DocumentSummarizerClient is the client API for DocumentSummarizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentSummarizerClient interface {
	SummarizeDocument(ctx context.Context, in *SummarizeDocumentRequest, opts ...grpc.CallOption) (*SummarizeDocumentReply, error)
}

type documentSummarizerClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentSummarizerClient(cc grpc.ClientConnInterface) DocumentSummarizerClient {
	return &documentSummarizerClient{cc}
}

func (c *documentSummarizerClient) SummarizeDocument(ctx context.Context, in *SummarizeDocumentRequest, opts ...grpc.CallOption) (*SummarizeDocumentReply, error) {
	out := new(SummarizeDocumentReply)
	err := c.cc.Invoke(ctx, "/statistics.DocumentSummarizer/SummarizeDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentSummarizerServer is the server API for DocumentSummarizer service.
// All implementations must embed UnimplementedDocumentSummarizerServer
// for forward compatibility
type DocumentSummarizerServer interface {
	SummarizeDocument(context.Context, *SummarizeDocumentRequest) (*SummarizeDocumentReply, error)
	mustEmbedUnimplementedDocumentSummarizerServer()
}

// UnimplementedDocumentSummarizerServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentSummarizerServer struct {
}

func (UnimplementedDocumentSummarizerServer) SummarizeDocument(context.Context, *SummarizeDocumentRequest) (*SummarizeDocumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SummarizeDocument not implemented")
}
func (UnimplementedDocumentSummarizerServer) mustEmbedUnimplementedDocumentSummarizerServer() {}

// UnsafeDocumentSummarizerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentSummarizerServer will
// result in compilation errors.
type UnsafeDocumentSummarizerServer interface {
	mustEmbedUnimplementedDocumentSummarizerServer()
}

func RegisterDocumentSummarizerServer(s grpc.ServiceRegistrar, srv DocumentSummarizerServer) {
	s.RegisterService(&DocumentSummarizer_ServiceDesc, srv)
}

func _DocumentSummarizer_SummarizeDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummarizeDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentSummarizerServer).SummarizeDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statistics.DocumentSummarizer/SummarizeDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentSummarizerServer).SummarizeDocument(ctx, req.(*SummarizeDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentSummarizer_ServiceDesc is the grpc.ServiceDesc for DocumentSummarizer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentSummarizer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "statistics.DocumentSummarizer",
	HandlerType: (*DocumentSummarizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SummarizeDocument",
			Handler:    _DocumentSummarizer_SummarizeDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// DocumentSummarizerBackendClient is the client API for DocumentSummarizerBackend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentSummarizerBackendClient interface {
	SummarizeDocument(ctx context.Context, in *SummarizeDocumentWorkload, opts ...grpc.CallOption) (*SummarizeDocumentReply, error)
}

type documentSummarizerBackendClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentSummarizerBackendClient(cc grpc.ClientConnInterface) DocumentSummarizerBackendClient {
	return &documentSummarizerBackendClient{cc}
}

func (c *documentSummarizerBackendClient) SummarizeDocument(ctx context.Context, in *SummarizeDocumentWorkload, opts ...grpc.CallOption) (*SummarizeDocumentReply, error) {
	out := new(SummarizeDocumentReply)
	err := c.cc.Invoke(ctx, "/statistics.DocumentSummarizerBackend/SummarizeDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentSummarizerBackendServer is the server API for DocumentSummarizerBackend service.
// All implementations must embed UnimplementedDocumentSummarizerBackendServer
// for forward compatibility
type DocumentSummarizerBackendServer interface {
	SummarizeDocument(context.Context, *SummarizeDocumentWorkload) (*SummarizeDocumentReply, error)
	mustEmbedUnimplementedDocumentSummarizerBackendServer()
}

// UnimplementedDocumentSummarizerBackendServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentSummarizerBackendServer struct {
}

func (UnimplementedDocumentSummarizerBackendServer) SummarizeDocument(context.Context, *SummarizeDocumentWorkload) (*SummarizeDocumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SummarizeDocument not implemented")
}
func (UnimplementedDocumentSummarizerBackendServer) mustEmbedUnimplementedDocumentSummarizerBackendServer() {
}

// UnsafeDocumentSummarizerBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentSummarizerBackendServer will
// result in compilation errors.
type UnsafeDocumentSummarizerBackendServer interface {
	mustEmbedUnimplementedDocumentSummarizerBackendServer()
}

func RegisterDocumentSummarizerBackendServer(s grpc.ServiceRegistrar, srv DocumentSummarizerBackendServer) {
	s.RegisterService(&DocumentSummarizerBackend_ServiceDesc, srv)
}

func _DocumentSummarizerBackend_SummarizeDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummarizeDocumentWorkload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentSummarizerBackendServer).SummarizeDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statistics.DocumentSummarizerBackend/SummarizeDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentSummarizerBackendServer).SummarizeDocument(ctx, req.(*SummarizeDocumentWorkload))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentSummarizerBackend_ServiceDesc is the grpc.ServiceDesc for DocumentSummarizerBackend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentSummarizerBackend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "statistics.DocumentSummarizerBackend",
	HandlerType: (*DocumentSummarizerBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SummarizeDocument",
			Handler:    _DocumentSummarizerBackend_SummarizeDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
