// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/librarian.proto

package pb

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

// LibrarianClient is the client API for Librarian service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibrarianClient interface {
	GetBooks(ctx context.Context, in *Author, opts ...grpc.CallOption) (Librarian_GetBooksClient, error)
	GetAuthor(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Author, error)
}

type librarianClient struct {
	cc grpc.ClientConnInterface
}

func NewLibrarianClient(cc grpc.ClientConnInterface) LibrarianClient {
	return &librarianClient{cc}
}

func (c *librarianClient) GetBooks(ctx context.Context, in *Author, opts ...grpc.CallOption) (Librarian_GetBooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &Librarian_ServiceDesc.Streams[0], "/librarian.Librarian/GetBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &librarianGetBooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Librarian_GetBooksClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type librarianGetBooksClient struct {
	grpc.ClientStream
}

func (x *librarianGetBooksClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *librarianClient) GetAuthor(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/librarian.Librarian/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibrarianServer is the server API for Librarian service.
// All implementations must embed UnimplementedLibrarianServer
// for forward compatibility
type LibrarianServer interface {
	GetBooks(*Author, Librarian_GetBooksServer) error
	GetAuthor(context.Context, *Book) (*Author, error)
	mustEmbedUnimplementedLibrarianServer()
}

// UnimplementedLibrarianServer must be embedded to have forward compatible implementations.
type UnimplementedLibrarianServer struct {
}

func (UnimplementedLibrarianServer) GetBooks(*Author, Librarian_GetBooksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedLibrarianServer) GetAuthor(context.Context, *Book) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}
func (UnimplementedLibrarianServer) mustEmbedUnimplementedLibrarianServer() {}

// UnsafeLibrarianServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibrarianServer will
// result in compilation errors.
type UnsafeLibrarianServer interface {
	mustEmbedUnimplementedLibrarianServer()
}

func RegisterLibrarianServer(s grpc.ServiceRegistrar, srv LibrarianServer) {
	s.RegisterService(&Librarian_ServiceDesc, srv)
}

func _Librarian_GetBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Author)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LibrarianServer).GetBooks(m, &librarianGetBooksServer{stream})
}

type Librarian_GetBooksServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type librarianGetBooksServer struct {
	grpc.ServerStream
}

func (x *librarianGetBooksServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

func _Librarian_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibrarianServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/librarian.Librarian/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibrarianServer).GetAuthor(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

// Librarian_ServiceDesc is the grpc.ServiceDesc for Librarian service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Librarian_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "librarian.Librarian",
	HandlerType: (*LibrarianServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthor",
			Handler:    _Librarian_GetAuthor_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBooks",
			Handler:       _Librarian_GetBooks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/librarian.proto",
}
