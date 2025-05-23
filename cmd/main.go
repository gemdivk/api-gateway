package main

import (
	"github.com/gemdivk/api-gateway/internal/client"
	"github.com/gemdivk/api-gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	clients := client.InitServiceClients()

	handler.RegisterInventoryRoutes(r, clients)
	handler.RegisterOrderRoutes(r, clients)

	r.Run(":8080")
}
