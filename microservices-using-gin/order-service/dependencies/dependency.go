package dependencies

import (
	"order-service/controller"
	"order-service/database"
	"order-service/repo"
	"order-service/service"
)

var (
	OrderControllerInstance controller.OrderController
)

func InitializeDependencies() {
	db := database.Connect()
	orderRepo := repo.NewOrderRepo(db)

	orderService := service.NewOrderService(orderRepo)
	OrderControllerInstance = controller.NewOrderController(orderService)
}
