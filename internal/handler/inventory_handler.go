package handler

import (
	"context"
	"net/http"

	"github.com/gemdivk/api-gateway/internal/client"

	inventorypb "github.com/gemdivk/api-gateway/proto/inventory"

	"github.com/gin-gonic/gin"
)

func RegisterInventoryRoutes(r *gin.Engine, sc *client.ServiceClients) {
	r.GET("/products", func(c *gin.Context) {
		res, err := sc.Inventory.ListProducts(context.Background(), &inventorypb.ListProductsRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res.Products)
	})

	r.POST("/products", func(c *gin.Context) {
		var p inventorypb.Product
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product"})
			return
		}
		res, err := sc.Inventory.CreateProduct(context.Background(), &inventorypb.CreateProductRequest{Product: &p})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, res.Product)
	})

	r.PATCH("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var p inventorypb.Product
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product"})
			return
		}
		res, err := sc.Inventory.UpdateProduct(context.Background(), &inventorypb.UpdateProductRequest{
			Id:      id,
			Product: &p,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res.Product)
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := sc.Inventory.DeleteProduct(context.Background(), &inventorypb.DeleteProductRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	})
}
