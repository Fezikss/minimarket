package models

import "time"

type StorageTransaction struct {
	ID                     string     `json:"id"`
	StaffID                string     `json:"staff_id"`
	ProductID              string     `json:"product_id"`
	StorageTransactionType string     `json:"storagetransaction_type"`
	Price                  int        `json:"price"`
	Quantity               int        `json:"quantity"`
	BranchID               string     `json:"branch_id"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"update_at"`
	DeletedAt              *time.Time `json:"-"`
}

type CreateStorageTransaction struct {
	StaffID                string `json:"staff_id"`
	ProductID              string `json:"product_id"`
	StorageTransactionType string `json:"storagetransaction_type"`
	Price                  int    `json:"price"`
	Quantity               int    `json:"quantity"`
	BranchID               string `json:"branch_id"`
}

type UpdateStorageTransaction struct {
	ID                     string    `json:"-"`
	StaffID                string    `json:"staff_id"`
	ProductID              string    `json:"product_id"`
	StorageTransactionType string    `json:"storagetransaction_type"`
	Price                  int       `json:"price"`
	Quantity               int       `json:"quantity"`
	UpdatedAt              time.Time `json:"update_at"`
}

type StorageTransactionsResponse struct {
	StorageTransactions []StorageTransaction `json:"storagetransaction"`
	Count               int                  `json:"count"`
}
