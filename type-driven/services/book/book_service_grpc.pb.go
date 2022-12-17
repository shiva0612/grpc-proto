// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: book_service.proto

package book

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	book "shiva/type-driven/models/book"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookApiClient is the client API for BookApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookApiClient interface {
	AddBook(ctx context.Context, in *book.Book, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetBook(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*book.Book, error)
}

type bookApiClient struct {
	cc grpc.ClientConnInterface
}

func NewBookApiClient(cc grpc.ClientConnInterface) BookApiClient {
	return &bookApiClient{cc}
}

func (c *bookApiClient) AddBook(ctx context.Context, in *book.Book, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/book.BookApi/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookApiClient) GetBook(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*book.Book, error) {
	out := new(book.Book)
	err := c.cc.Invoke(ctx, "/book.BookApi/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookApiServer is the server API for BookApi service.
// All implementations must embed UnimplementedBookApiServer
// for forward compatibility
type BookApiServer interface {
	AddBook(context.Context, *book.Book) (*emptypb.Empty, error)
	GetBook(context.Context, *emptypb.Empty) (*book.Book, error)
	mustEmbedUnimplementedBookApiServer()
}

// UnimplementedBookApiServer must be embedded to have forward compatible implementations.
type UnimplementedBookApiServer struct {
}

func (UnimplementedBookApiServer) AddBook(context.Context, *book.Book) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedBookApiServer) GetBook(context.Context, *emptypb.Empty) (*book.Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookApiServer) mustEmbedUnimplementedBookApiServer() {}

// UnsafeBookApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookApiServer will
// result in compilation errors.
type UnsafeBookApiServer interface {
	mustEmbedUnimplementedBookApiServer()
}

func RegisterBookApiServer(s grpc.ServiceRegistrar, srv BookApiServer) {
	s.RegisterService(&BookApi_ServiceDesc, srv)
}

func _BookApi_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(book.Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookApiServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookApi/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookApiServer).AddBook(ctx, req.(*book.Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookApi_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookApiServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookApi/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookApiServer).GetBook(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BookApi_ServiceDesc is the grpc.ServiceDesc for BookApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookApi",
	HandlerType: (*BookApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBook",
			Handler:    _BookApi_AddBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookApi_GetBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book_service.proto",
}
