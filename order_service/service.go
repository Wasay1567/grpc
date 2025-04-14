package main

import (
	"context"
	"log"

	inventorypb "github.com/Wasay1567/GRPC/proto/inventory"
	orderpb "github.com/Wasay1567/GRPC/proto/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderServer struct {
	orderpb.UnimplementedOrderServiceServer
}

func (s *OrderServer) PlaceOrder(ctx context.Context, req *orderpb.OrderRequest) (*orderpb.OrderResponse, error) {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect %v", err)
	}
	defer conn.Close()

	inventoryClient := inventorypb.NewInventoryServiceClient(conn)
	checkResp, _ := inventoryClient.CheckStore(ctx, &inventorypb.StockRequest{
		ItemId:   req.ItemId,
		Quantity: req.Quantity,
	})
	if !checkResp.Available {
		return &orderpb.OrderResponse{Success: false, Message: "Order failed : " + checkResp.Message}, nil
	}
	inventoryClient.DecreaseStock(ctx, &inventorypb.StockRequest{
		ItemId:   req.ItemId,
		Quantity: req.Quantity,
	})
	return &orderpb.OrderResponse{Success: true, Message: "Order placed successfully"}, nil
}
