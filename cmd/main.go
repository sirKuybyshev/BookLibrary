package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	ls "library/pkg"
	"library/proto/pb"
)

var (
	c = flag.String("config", "cmd/config/serverConfig.yaml",
		"Path to serverConfig config, example is shown in cmd/config/serverConfig.yaml")
)

func main() {
	flag.Parse()
	server := &ls.Server{}
	err := server.Construct(*c)
	if err != nil {
		log.Fatalf("failed to establish db connection, %e", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server.GetPort()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer server.StopServer()
	grpcServer := grpc.NewServer()
	pb.RegisterLibrarianServer(grpcServer, server)
	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
