package http

import (
	"ecommerce-api/application"
	"ecommerce-api/domain"
	"ecommerce-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService   *application.OrderService
	ProductService *application.ProductService
	UserService    *application.UserService
}

func NewOrderController(orderService *application.OrderService, productService *application.ProductService, userService *application.UserService) *OrderController {
	return &OrderController{OrderService: orderService, UserService: userService, ProductService: productService}
}

func (ctrl *OrderController) PlaceOrder(c *gin.Context) {
	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "User ID not found", Response: nil})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Invalid user ID type", Response: nil})
		return
	}

	var input struct {
		ProductIDs []uint `json:"product_ids" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: err.Error(), Response: nil})
		return
	}

	if _, err := ctrl.UserService.GetUserByID(userID); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "User not found", Response: nil})
		return
	}

	products, err := ctrl.ProductService.GetProductsById(input.ProductIDs)
	if err != nil || len(products) == 0 {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "One or more products not found", Response: nil})
		return
	}

	order := domain.Order{
		UserID:   userID,
		Products: products,
		Status:   utils.Pending,
	}

	if err := ctrl.OrderService.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Success: false, Message: "Failed to place order", Response: nil})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "Order placed successfully",
		Response: struct {
			OrderID uint `json:"order_id"`
		}{OrderID: order.ID},
	})
}

func (ctrl *OrderController) GetOrdersByUser(c *gin.Context) {
	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "User ID not found", Response: nil})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Invalid user ID type", Response: nil})
		return
	}

	orders, err := ctrl.OrderService.GetOrdersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Success: false, Message: "Failed to fetch orders", Response: nil})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Orders fetched successfully",
		Response: struct {
			Orders interface{} `json:"orders"`
		}{Orders: orders},
	})
}

func (ctrl *OrderController) UpdateOrderStatus(c *gin.Context) {
	var input struct {
		Status utils.OrderStatus `json:"status" binding:"required,oneof=0 1 2 3 4"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: err.Error(), Response: nil})
		return
	}

	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 8)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Invalid order ID", Response: nil})
		return
	}

	order, err := ctrl.OrderService.GetOrderByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.Response{Success: false, Message: "Order not found", Response: nil})
		return
	}


	if order.Status == input.Status {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Order is already in the desired status", Response: nil})
		return
	}

	if err := ctrl.OrderService.UpdateOrderStatus(uint(orderID), input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Success: false, Message: "Failed to update order status", Response: nil})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Order status updated successfully",
		Response: struct {
			OrderID uint64 `json:"order_id"`
		}{OrderID: orderID},
	})
}

func (ctrl *OrderController) CancelOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 8)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Invalid order ID", Response: nil})
		return
	}

	order, err := ctrl.OrderService.GetOrderByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.Response{Success: false, Message: "Order not found", Response: nil})
		return
	}

	if order.Status == utils.Cancelled {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Order already cancelled", Response: nil})
		return
	}

	if order.Status > utils.Pending {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Order cannot be cancelled unless it is Pending", Response: nil})
		return
	}

	if err := ctrl.OrderService.UpdateOrderStatus(uint(orderID), utils.Cancelled); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Success: false, Message: "Failed to cancel order", Response: nil})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Order cancelled successfully",
		Response: struct {
			OrderID uint64 `json:"order_id"`
		}{OrderID: orderID},
	})
}
