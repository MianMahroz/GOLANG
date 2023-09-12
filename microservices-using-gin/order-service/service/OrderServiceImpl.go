package service

import (
	"encoding/hex"
	"github.com/dranikpg/dto-mapper"
	"io"
	"log"
	"net/http"
	"order-service/model"
	"order-service/repo"
	"strconv"
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

	// Inter-service communication
	res, err := s.FetchUserDetails(orderDto.UserId)

	log.Println(res)

	return s.orderRepo.SaveOrder(entity)
}

func (s *orderService) FetchUserDetails(userId int) (string, error) {
	res, err := http.Get("http://user-service:8002/user?id=" + strconv.Itoa(userId))

	if err != nil {
		return "SOMETHING WENT WRONG", err
	}

	data, _ := io.ReadAll(res.Body)

	response := hex.EncodeToString(data)

	return response, nil
}
