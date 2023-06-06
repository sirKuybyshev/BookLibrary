package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	librarianServer "library/pkg"
	"library/proto/pb"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
)

var (
	port = flag.Int("port", 50051, "The Server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := &librarianServer.Server{}
	err = server.ConnectToDB("user:password@tcp(localhost:3306)/Library")
	if err != nil {
		log.Fatalf("failed to establish db connection, %e", err)
	}
	defer server.DisconnectDB()
	grpcServer := grpc.NewServer()
	pb.RegisterLibrarianServer(grpcServer, server)
	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
