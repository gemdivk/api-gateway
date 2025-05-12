package handler

import (
	"context"
	"net/http"

	"github.com/gemdivk/api-gateway/internal/client"
	orderpb "github.com/gemdivk/api-gateway/proto/order"
	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine, sc *client.ServiceClients) {
	r.GET("/orders", func(c *gin.Context) {
		res, err := sc.Order.ListOrders(context.Background(), &orderpb.ListOrdersRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res.Orders)
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		res, err := sc.Order.GetOrderByID(context.Background(), &orderpb.GetOrderByIDRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusOK, res.Order)
	})

	r.POST("/orders", func(c *gin.Context) {
		var req orderpb.CreateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order"})
			return
		}
		res, err := sc.Order.CreateOrder(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, res.Order)
	})

	r.PATCH("/orders/:id/status", func(c *gin.Context) {
		id := c.Param("id")
		var body struct {
			Status string `json:"status"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
			return
		}
		req := &orderpb.UpdateOrderStatusRequest{
			Id:     id,
			Status: body.Status,
		}
		res, err := sc.Order.UpdateOrderStatus(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.DELETE("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := sc.Order.DeleteOrder(context.Background(), &orderpb.DeleteOrderRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	})
}
