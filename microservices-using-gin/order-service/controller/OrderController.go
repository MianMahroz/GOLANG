package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/model"
	"order-service/service"
)

// OrderController  Method Declaration
type OrderController interface {
	SaveOrder(context *gin.Context)
}

// creating struct to link with implementation
type orderController struct {
	orderService service.OrderService
}

// NewOrderController creating a function to initialize linked struct
func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{orderService: orderService}
}

func (c *orderController) SaveOrder(context *gin.Context) {

	var input model.OrderDto

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	res, err := c.orderService.CreateOrder(input)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	context.JSON(http.StatusOK, gin.H{"response": res})
}
