package dependencies

import (
	"order-service/controller"
	"order-service/service"
)

var (
	OrderControllerInstance controller.OrderController
)

func InitializeDependencies() {
	orderService := service.NewOrderService()
	OrderControllerInstance = controller.NewOrderController(orderService)
}
