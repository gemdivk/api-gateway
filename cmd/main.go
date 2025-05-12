package main

import (
	"api-gateway/internal/client"
	"api-gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	clients := client.NewServiceClients()

	handler.RegisterInventoryRoutes(r, clients)
	handler.RegisterOrderRoutes(r, clients)

	r.Run(":8080")
}
