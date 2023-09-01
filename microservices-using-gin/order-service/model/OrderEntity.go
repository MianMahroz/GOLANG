package model

import "gorm.io/gorm"

type OrderEntity struct {
	gorm.Model
	Id          int
	UserId      int
	ProductId   int
	Qty         int
	OrderNumber string
}
