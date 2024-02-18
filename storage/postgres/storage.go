package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"main.go/api/models"
	"main.go/storage"
)

type storageRepo struct {
	DB *pgxpool.Pool
}

func NewStorageRepo(db *pgxpool.Pool) storage.IStorageStorage {
	return storageRepo{DB: db}
}

func (s storageRepo) Create(ctx context.Context, storage models.CreateStorage) (string, error) {
	id := uuid.New()
	storage.CreatedAt = time.Now()
	if _, err := s.DB.Exec(ctx, `INSERT INTO storage (id, product_id, branch_id, count,created_at) VALUES ($1, $2, $3, $4,$5)`, id, storage.ProductID, storage.BranchID, storage.Count,storage.CreatedAt); err != nil {
		fmt.Println("error while inserting data to storage")
		return "", err
	}

	return id.String(), nil
}

func (s storageRepo) GetById(ctx context.Context, pk models.PrimaryKey) (models.Storage, error) {
	storage := models.Storage{}
	if err := s.DB.QueryRow(ctx, "select id, product_id, branch_id,count, created_at from storage where id = $1", pk.ID).Scan(
		&storage.ID,
		&storage.ProductID,
		&storage.BranchID,
		&storage.Count,
		&storage.CreatedAt); err != nil {

		fmt.Println("error getting by id storage")
		return models.Storage{}, err
	}
	fmt.Println("storage", storage)
	return storage, nil
}

func (s storageRepo) GetList(ctx context.Context, request models.GetListRequest) (models.StorageResponse, error) {
	var (
		storages          = []models.Storage{}
		count             = 0
		query, countquery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		// search            = request.Search
	)
	countquery = `select count(1) from storage `
	// if search != "" {
	// 	countquery += fmt.Sprintf(`  ='%d'`, search)
	// }
	if err := s.DB.QueryRow(ctx, countquery).Scan(&count); err != nil {
		fmt.Println("error while counting")
		return models.StorageResponse{}, err
	}
	query = `select id, product_id, branch_id, count, update_at, created_at from storage `

	query += ` LIMIT $1 OFFSET $2`
	rows, err := s.DB.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting storages", err.Error())
		return models.StorageResponse{}, err
	}
	for rows.Next() {
		storage := models.Storage{}
		if err = rows.Scan(&storage.ID, &storage.ProductID, &storage.BranchID, &storage.Count, &storage.UpdatedAt, &storage.CreatedAt); err != nil {
			fmt.Println("error while getting list of storage  ")
			return models.StorageResponse{}, err
		}
		storages = append(storages, storage)
	}
	return models.StorageResponse{
		Storages: storages,
		Count:    count,
	}, nil
}
func (s storageRepo) Update(ctx context.Context, storage models.UpdateStorage) (string, error) {

	storage.UpdatedAt = time.Now()
	if _, err := s.DB.Exec(ctx, "update storage set product_id=$1, branch_id=$2,count=$3, update_at=$4 where id=$5", storage.ProductID, storage.BranchID,storage.Count, storage.UpdatedAt, storage.ID); err != nil {
		return "", err
	}
	
	return storage.ID, nil
}

func (s storageRepo) Delete(ctx context.Context, pk models.PrimaryKey) error {

	if _, err := s.DB.Exec(ctx, `delete from storage where id=$1`, pk.ID); err != nil {
		return err
	}
	return nil

}
