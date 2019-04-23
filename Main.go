package main

import (
	"fmt"
	"github.com/pivotal/test/proto"
	"github.com/pivotal/test/src"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := 10000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := src.Server{}

	grpcServer := grpc.NewServer()

	proto.RegisterValueStoreServer(grpcServer, &s)

	fmt.Println("grpc server running on port: ", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
