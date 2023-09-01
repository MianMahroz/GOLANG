package service

import (
	"github.com/dranikpg/dto-mapper"
	"order-service/model"
	"order-service/repo"
)

// declare all dependencies in below struct and link that struct with func
type orderService struct {
	orderRepo repo.OrderRepo
}

// NewOrderService initialize dependencies here
func NewOrderService(orderRepo repo.OrderRepo) OrderService {
	return &orderService{orderRepo: orderRepo}
}

func (s *orderService) CreateOrder(orderDto model.OrderDto) (string, error) {
	var entity = model.OrderEntity{}
	err := dto.Map(&entity, orderDto)
	if err != nil {
		return "", err
	}
	return s.orderRepo.SaveOrder(entity)
}
