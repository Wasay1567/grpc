package main

import (
	"context"

	pb "github.com/Wasay1576/GRPC/proto/inventory"
)

var inventory = map[string]int32{
	"iphone": 20,
	"Camera": 10,
}

type InventoryServer struct {
	pb.UnimplementedInventoryServiceServer
}

func (s *InventoryServer) CheckStore(ctx context.Context, req *pb.StockRequest) (*pb.StockResponse, error) {
	qty, ok := inventory[req.ItemId]
	if !ok || qty < req.Quantity {
		return &pb.StockResponse{Available: false, Message: "Out of stock"}, nil
	}
	return &pb.StockResponse{Available: true, Message: "Stock available"}, nil

}

func (s *InventoryServer) DecreaseStock(ctx context.Context, req *pb.StockRequest) (*pb.StockResponse, error) {
	if qty, ok := inventory[req.ItemId]; ok && qty >= req.Quantity {
		inventory[req.ItemId] -= req.Quantity
		return &pb.StockResponse{Available: true, Message: "Stock updated"}, nil
	}
	return &pb.StockResponse{Available: false, Message: "Insufficient Stock"}, nil
}
