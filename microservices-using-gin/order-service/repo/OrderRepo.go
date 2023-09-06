package repo

import (
	"gorm.io/gorm"
	"order-service/model"
	"strconv"
)

type OrderRepo interface {
	SaveOrder(entity model.OrderEntity) (string, error)
}

type orderRepo struct {
	database *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{database: db}
}

func (s orderRepo) SaveOrder(entity model.OrderEntity) (string, error) {
	res := s.database.Create(&entity)
	return strconv.Itoa(entity.Id), res.Error
}
