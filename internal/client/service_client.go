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

func NewServiceClients() *ServiceClients {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	invConn, err := grpc.DialContext(ctx, "localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Inventory Service: %v", err)
	}

	ordConn, err := grpc.DialContext(ctx, "localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Order Service: %v", err)
	}

	return &ServiceClients{
		Inventory: inventorypb.NewInventoryServiceClient(invConn),
		Order:     orderpb.NewOrderServiceClient(ordConn),
	}
}
