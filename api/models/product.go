package models

import "time"

type Product struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Price      float64    `json:"price"`
	Barcode    int        `json:"barcode"`
	CategoryID string     `json:"category_id"`
	UpdatedAt  time.Time  `json:"updated_at"`
	CreatedAt  time.Time  `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
type CreateProduct struct {
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	Barcode    int       `json:"barcode"`
	CategoryID string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}
type UpdateProduct struct {
	ID         string    `json:"-"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	CategoryID string    `json:"category_id"`
	UpdatedAt  time.Time `json:"-"`
}
type ProductResponse struct {
	Products []Product `json:"products"`
	Count    int       `json:"count"`
}
type GetListRequestProduct struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	Search  string `json:"search"`
	Barcode int    `json:"barcode"`
}
