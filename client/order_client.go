package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Wasay1576/GRPC/proto/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.PlaceOrder(ctx, &pb.OrderRequest{
		ItemId:   "iphone",
		Quantity: 2,
	})

	if err != nil {
		log.Fatalf("Could not place order: %v", err)
	}

	log.Print(res.Message)
}
