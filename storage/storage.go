package storage

import "main.go/api/models"

type IStorage interface {
	Close()
	Branch() IBranchStorage
	Sale() ISaleStorage
}
type IBranchStorage interface {
	Create(models.CreateBranch) (string, error)
	GetById(models.PrimaryKey) (models.Branch, error)
	GetList(models.GetListRequest) (models.BranchResponse, error)
	Update(models.UpdateBranch) (string, error)
	Delete(models.PrimaryKey) error
}
type ISaleStorage interface {
	Create(models.CreateSale) (string, error)
	GetById(models.PrimaryKey) (models.Sale, error)
	GetList(models.GetListRequest) (models.SaleResponse, error)
	Update(models.UpdateSale) (string, error)
	Delete(models.PrimaryKey) error
}
