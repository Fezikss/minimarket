package models

import "time"

type Sale struct {
	ID              string     `json:"id"`
	BranchID        string     `json:"branch_id"`
	ShopAssistantID string     `json:"shop_assistant_id"`
	Cashier         string     `json:"cashier"`
	PaymentType     string     `json:"payment_type"`
	Price           float64    `json:"price"`
	Status          string     `json:"status"`
	ClientName      string     `json:"client_name"`
	UpdatedAt       time.Time  `json:"updated_at"`
	CreatedAt       time.Time  `json:"created_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}
type CreateSale struct {
	BranchID        string  `json:"branch_id"`
	ShopAssistantID string  `json:"shop_assistant_id"`
	Cashier         string  `json:"cashier"`
	PaymentType     string  `json:"payment_type"`
	Price           float64 `json:"price"`
	Status          string  `json:"status"`
	ClientName      string  `json:"client_name"`
	CreatedAt       time.Time  `json:"created_at"`
}
type UpdateSale struct {
	ID              string  `json:"-"`
	BranchID        string  `json:"branch_id"`
	ShopAssistantID string  `json:"shop_assistant_id"`
	Cashier         string  `json:"cashier"`
	PaymentType     string  `json:"payment_type"`
	Price           float64 `json:"price"`
}
type SaleResponse struct {
	Sales []Sale `json:"sales"`
	Count int    `json:"count"`
}
