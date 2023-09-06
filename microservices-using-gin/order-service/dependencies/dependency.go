package dependencies

import (
	"log"
	"order-service/controller"
	"order-service/database"
	"order-service/model"
	"order-service/repo"
	"order-service/service"
)

var (
	OrderControllerInstance controller.OrderController
)

func InitializeDependencies() {
	db := database.Connect()
	orderRepo := repo.NewOrderRepo(db)

	err := db.AutoMigrate(model.OrderEntity{})
	if err != nil {
		log.Fatal("AUTO MIGRATE FAILED")
	}
	orderService := service.NewOrderService(orderRepo)
	OrderControllerInstance = controller.NewOrderController(orderService)
}
