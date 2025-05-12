package handler

import (
	"net/http"

	"github.com/gemdivk/api-gateway/internal/client"

	"github.com/gin-gonic/gin"
)

func RegisterInventoryRoutes(r *gin.Engine, clients *client.ServiceClients) {
	inventory := clients.Inventory

	r.GET("/products", func(c *gin.Context) {
		resp, err := inventory.ListProducts(c, &pb.ListProductsRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Products)
	})

	r.POST("/products", func(c *gin.Context) {
		var req pb.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		resp, err := inventory.CreateProduct(c, &pb.CreateProductRequest{Product: &req})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp.Product)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := inventory.GetProductByID(c, &pb.GetProductRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Product)
	})

	r.PATCH("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var req pb.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		resp, err := inventory.UpdateProduct(c, &pb.UpdateProductRequest{
			Id:      id,
			Product: &req,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Product)
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := inventory.DeleteProduct(c, &pb.DeleteProductRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
	})
}
