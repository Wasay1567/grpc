package main

import (
	"log"
	"net"

	pb "github.com/Wasay1567/GRPC/proto/inventory"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, &InventoryServer{})
	log.Println("Inventory Server running on :50052")
	grpcServer.Serve(lis)

}
