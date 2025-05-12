package client

import (
	"context"
	"log"
	"time"

	inventorypb "github.com/gemdivk/api-gateway/proto/inventory"
	orderpb "github.com/gemdivk/api-gateway/proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClients struct {
	Inventory inventorypb.InventoryServiceClient
	Order     orderpb.OrderServiceClient
}

func InitServiceClients() *ServiceClients {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	invConn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to inventory-service: %v", err)
	}
	ordConn, err := grpc.DialContext(ctx, "localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to order-service: %v", err)
	}

	return &ServiceClients{
		Inventory: inventorypb.NewInventoryServiceClient(invConn),
		Order:     orderpb.NewOrderServiceClient(ordConn),
	}
}
