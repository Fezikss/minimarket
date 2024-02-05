package models

import "time"

type Category struct {
	ID        string     `json:"id"`
	ParentID  string     `json:"parent_id"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
}
type CreateCategory struct {
	ParentID  string    `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}
type UpdateCategory struct {
	ID        string    `json:"-"`
	Name      string    `json:"name"`
	ParentID  string    `json:"parent_id"`
	UpdatedAt time.Time `json:"-"`
}
type CategoryResponse struct {
	Categories []Category `json:"categories"`
	Count      int        `json:"count"`
}
