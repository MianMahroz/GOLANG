package service

import "order-service/model"

type OrderService interface {
	CreateOrder(dto model.OrderDto) (string, error)
}
