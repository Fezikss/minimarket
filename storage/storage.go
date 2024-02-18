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
	Basket() IBasketStorage
	StorageTransaction() IStorageTransaction
	Staff() IStaffStorage
	StaffTariff() IStaffTariffStorage
	Transaction() ITransactionStorage
}
type IBranchStorage interface {
	Create(context.Context, models.CreateBranch) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.Branch, error)
	GetList(context.Context, models.GetListRequest) (models.BranchResponse, error)
	Update(context.Context, models.UpdateBranch) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}
type ISaleStorage interface {
	Create(context.Context, models.CreateSale) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.Sale, error)
	GetList(context.Context, models.GetListRequest) (models.SaleResponse, error)
	Update(context.Context, models.UpdateSale) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	UpdatePriceSale(context.Context, models.EndSaleUpdate) (string, error)
}
type ICategoryStorage interface {
	Create(context.Context, models.CreateCategory) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.Category, error)
	GetList(context.Context, models.GetListRequest) (models.CategoryResponse, error)
	Update(context.Context, models.UpdateCategory) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}
type IProductStorage interface {
	Create(context.Context, models.CreateProduct) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.Product, error)
	GetList(context.Context, models.GetListRequestProduct) (models.ProductResponse, error)
	Update(context.Context, models.UpdateProduct) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}
type IStorageStorage interface {
	Create(context.Context, models.CreateStorage) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.Storage, error)
	GetList(context.Context, models.GetListRequest) (models.StorageResponse, error)
	Update(context.Context, models.UpdateStorage) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}
type IBasketStorage interface {
	Create(context.Context, models.CreateBasket) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.Basket, error)
	GetList(context.Context, models.GetListRequest) (models.BasketResponse, error)
	Update(context.Context, models.UpdateBasket) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}
type IStorageTransaction interface {
	Create(context.Context, models.CreateStorageTransaction) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.StorageTransaction, error)
	GetList(context.Context, models.GetListRequest) (models.StorageTransactionsResponse, error)
	Update(context.Context, models.UpdateStorageTransaction) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}
type IStaffStorage interface {
	Create(context.Context, models.CreateStaff) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.Staff, error)
	GetList(context.Context, models.GetListRequest) (models.StaffsResponse, error)
	Update(context.Context, models.UpdateStaff) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	GetPassword(context.Context, string) (string, error)
	UpdatePassword(context.Context, models.UpdateStaffPassword) error
}
type IStaffTariffStorage interface {
	Create(context.Context, models.CreateStaffTariff) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.StaffTariff, error)
	GetList(context.Context, models.GetListRequest) (models.StaffTariffResponse, error)
	Update(context.Context, models.UpdateStaffTariff) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}
type ITransactionStorage interface {
	Create(context.Context, models.CreateTransaction) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.Transaction, error)
	GetList(context.Context, models.TransactionGetListRequest) (models.TransactionResponse, error)
	Update(context.Context, models.UpdateTransaction) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}
