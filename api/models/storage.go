package models

import "time"

type Storage struct {
	ID        string     `json:"id"`
	ProductID string     `json:"product_id"`
	BranchID  string     `json:"branch_id"`
	Count     int        `json:"count"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
type CreateStorage struct {
	ProductID string    `json:"product_id"`
	BranchID  string    `json:"branch_id"`
	Count     int        `json:"count"`
	CreatedAt time.Time `json:"created_at"`
}
type UpdateStorage struct {
	ID        string    `json:"-"`
	ProductID string    `json:"product_id"`
	BranchID  string    `json:"branch_id"`
	Count     int       `json:"count"`
	UpdatedAt time.Time `json:"-"`
}
type StorageResponse struct {
	Storages []Storage `json:"storages"`
	Count    int       `json:"count"`
}
