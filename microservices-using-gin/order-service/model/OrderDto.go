package model

type OrderDto struct {
	OrderId   int `json:"orderId"`
	ProductId int `json:"productId"`
	Qty       int `json:"qty"`
}
