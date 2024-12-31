package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCserver struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCserver {
	return &gRPCserver{
		addr: addr,
	}
}

func (s *gRPCserver) Run() error {
	lis, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatal("Fatal to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}