package storage

import (
	"context"

	"main.go/api/models"
)

type IStorage interface {
	Close()
	Branch() IBranchStorage
	Sale() ISaleStorage
	Product() IProductStorage
	Category() ICategoryStorage
	Storage() IStorageStorage
}
type IBranchStorage interface {
	Create(context.Context,models.CreateBranch) (string, error)
	GetById(context.Context,models.PrimaryKey) (models.Branch, error)
	GetList(context.Context,models.GetListRequest) (models.BranchResponse, error)
	Update(context.Context,models.UpdateBranch) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}
type ISaleStorage interface {
	Create(context.Context,models.CreateSale) (string, error)
	GetById(context.Context,models.PrimaryKey) (models.Sale, error)
	GetList(context.Context,models.GetListRequest) (models.SaleResponse, error)
	Update(context.Context,models.UpdateSale) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}
type ICategoryStorage interface {
	Create(context.Context,models.CreateCategory) (string, error)
	GetById(context.Context,models.PrimaryKey) (models.Category, error)
	GetList(context.Context,models.GetListRequest) (models.CategoryResponse, error)
	Update(context.Context,models.UpdateCategory) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}
type IProductStorage interface {
	Create(context.Context,models.CreateProduct) (string, error)
	GetById(context.Context,models.PrimaryKey) (models.Product, error)
	GetList(context.Context,models.GetListRequestProduct) (models.ProductResponse, error)
	Update(context.Context,models.UpdateProduct) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}
type IStorageStorage interface {
	Create(context.Context,models.CreateStorage) (string, error)
	GetById(context.Context,models.PrimaryKey) (models.Storage, error)
	GetList(context.Context,models.GetListRequest) (models.StorageResponse, error)
	Update(context.Context,models.UpdateStorage) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}
