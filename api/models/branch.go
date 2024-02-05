package models

import "time"

type Branch struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Address   string     `json:"address"`
	UpdatedAt time.Time  `json:"update_at"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
type CreateBranch struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}
type UpdateBranch struct {
	ID        string    `json:"-"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	UpdatedAt time.Time `json:"update_at"`
}
type BranchResponse struct {
	Branchs []Branch `json:"branchs"`
	Count   int      `json:"count"`
}
