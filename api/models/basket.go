package models

import "time"

type Basket struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	SaleID    string    `json:"sale_id"`
	UpdateAt  time.Time `json:"update_at"`
	CreateAt  time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
type CreateBasket struct {
	ID        string    `json:"-"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	SaleID    string    `json:"sale_id"`
	Price     float64   `json:"price"`
	CreateAt  time.Time `json:"-"`
}
type UpdateBasket struct {
	ID        string    `json:"-"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	UpdateAt  time.Time `json:"-"`
}
type BasketResponse struct {
	Basket []Basket `json:"baskets"`
	Count  int      `json:"count"`
}
