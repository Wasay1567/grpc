package main

import (
	"log"
	"net"

	pb "github.com/Wasay1567/GRPC/proto/order"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &OrderServer{})
	log.Println("Order service running on :50051")
	grpcServer.Serve(lis)
}
