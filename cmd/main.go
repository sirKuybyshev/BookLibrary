package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	yaml "gopkg.in/yaml.v2"
	ls "library/pkg"
	"library/proto/pb"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	sc = flag.String("serverConfig", "cmd/config/serverConfig.yaml",
		"Path to serverConfig config, example is shown in cmd/config/serverConfig.yaml")
)

type serverConfig struct {
	Port int `yaml:"port"`
}

func main() {
	flag.Parse()
	data, err := os.ReadFile(*sc)
	config := &struct {
		Server serverConfig `yaml:"server"`
		DB     ls.Database  `yaml:"db"`
	}{}
	if err = yaml.Unmarshal(data, config); err != nil {
		log.Fatalf("%e", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := &ls.Server{}
	err = server.ConnectToDB(config.DB)
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
