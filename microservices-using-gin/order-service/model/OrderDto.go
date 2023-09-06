package model

type OrderDto struct {
	OrderNumber string `json:"orderNumber"`
	ProductId   int    `json:"productId"`
	Qty         int    `json:"qty"`
	UserId      int    `json:"userId"`
}
