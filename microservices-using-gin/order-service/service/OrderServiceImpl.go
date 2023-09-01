package service

import "order-service/model"

// declare all dependencies in below struct and link that struct with func
type orderService struct {
	// database
	// kafka
}

// NewOrderService initialize dependencies here
func NewOrderService() OrderService {
	return &orderService{}
}

func (s *orderService) CreateOrder(dto model.OrderDto) (string, error) {

	return "ORDER SAVED SUCCESSFULLY!", nil
}
