package handler

import (
	"api-gateway/internal/client"
	pb "api-gateway/proto/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine, clients *client.ServiceClients) {
	order := clients.Order

	r.POST("/orders", func(c *gin.Context) {
		var req pb.CreateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		resp, err := order.CreateOrder(c, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp.Order)
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := order.GetOrderByID(c, &pb.GetOrderByIDRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Order)
	})

	r.GET("/orders", func(c *gin.Context) {
		resp, err := order.ListOrders(c, &pb.ListOrdersRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Orders)
	})

	r.PATCH("/orders/:id/status", func(c *gin.Context) {
		id := c.Param("id")
		var body struct {
			Status string `json:"status"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		resp, err := order.UpdateOrderStatus(c, &pb.UpdateOrderStatusRequest{
			Id:     id,
			Status: body.Status,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": resp.Status})
	})

	r.DELETE("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := order.DeleteOrder(c, &pb.DeleteOrderRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
	})
}
