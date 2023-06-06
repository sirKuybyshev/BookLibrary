package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "library/proto/pb"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The Server port")
)

type Server struct {
	pb.UnimplementedLibrarianServer
}

func (s *Server) GetBooks(*pb.Author, pb.Librarian_GetBooksServer) error {
	return nil
}

func (s *Server) GetAuthor(context.Context, *pb.Book) (*pb.Author, error) {
	return nil, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLibrarianServer(s, &Server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
