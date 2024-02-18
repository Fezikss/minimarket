package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"main.go/api/models"
	"main.go/storage"
)

type storageTransactionRepo struct {
	DB *pgxpool.Pool
}

func NewStorageTransactionRepo(DB *pgxpool.Pool) storage.IStorageTransaction {
	return &storageTransactionRepo{
		DB: DB,
	}
}

func (s storageTransactionRepo) Create(ctx context.Context, stransaction models.CreateStorageTransaction) (string, error) {
	id := uuid.New().String()

	if _, err := s.DB.Exec(ctx, `INSERT INTO storagetransaction
		(id, staff_id, product_id, storagetransactiontype, price, quantity, branch_id)
			VALUES($1, $2, $3, $4, $5, $6,$7)`,
		id,
		stransaction.StaffID,
		stransaction.ProductID,
		stransaction.StorageTransactionType,
		stransaction.Price,
		stransaction.Quantity,
		stransaction.BranchID,
	); err != nil {
		log.Println("Error while inserting data:", err)
		return "", err
	}

	return id, nil
}

func (s storageTransactionRepo) GetById(ctx context.Context, id models.PrimaryKey) (models.StorageTransaction, error) {
	stransaction := models.StorageTransaction{}
	query := `SELECT id, staff_id, product_id, storagetransactiontype, price, quantity, branch_id, created_at, update_at 
							FROM storagetransaction WHERE id = $1 and deleted_at is null
`

	err := s.DB.QueryRow(ctx, query, id.ID).Scan(
		&stransaction.ID,
		&stransaction.StaffID,
		&stransaction.ProductID,
		&stransaction.StorageTransactionType,
		&stransaction.Price,
		&stransaction.Quantity,
		&stransaction.CreatedAt,
		&stransaction.UpdatedAt,
	)
	if err != nil {
		log.Println("Error while selecting repository by ID:", err)
		return models.StorageTransaction{}, err
	}

	return stransaction, nil
}

func (s storageTransactionRepo) GetList(ctx context.Context, req models.GetListRequest) (models.StorageTransactionsResponse, error) {
	var (
		rtransactions []models.StorageTransaction
		count         int
	)

	countQuery := `SELECT COUNT(*) FROM storagetransaction where deleted_at is null `
	if req.Search != "" {
		countQuery += fmt.Sprintf(` and quantity = %s`, req.Search)
	}

	err := s.DB.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		log.Println("Error while scanning count of repository_transactions:", err)
		return models.StorageTransactionsResponse{}, err
	}

	query := `SELECT id, staff_id, product_id, storagetransactiontype, price, quantity, created_at, update_at 
							FROM storagetransaction where deleted_at is null
`
	if req.Search != "" {
		query += fmt.Sprintf(` and quantity = %s`, req.Search)
	}
	query += ` order by created_at desc LIMIT $1 OFFSET $2 `

	rows, err := s.DB.Query(ctx, query, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		log.Println("Error while querying storagetransaction:", err)
		return models.StorageTransactionsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		stransaction := models.StorageTransaction{}
		err := rows.Scan(
			&stransaction.ID,
			&stransaction.StaffID,
			&stransaction.ProductID,
			&stransaction.StorageTransactionType,
			&stransaction.Price,
			&stransaction.Quantity,
			&stransaction.CreatedAt,
			&stransaction.UpdatedAt,
		)
		if err != nil {
			log.Println("Error while scanning row of repository_transactions:", err)
			return models.StorageTransactionsResponse{}, err
		}
		rtransactions = append(rtransactions, stransaction)
	}

	return models.StorageTransactionsResponse{
		StorageTransactions: rtransactions,
		Count:               count,
	}, nil
}

func (s storageTransactionRepo) Update(ctx context.Context, transaction models.UpdateStorageTransaction) (string, error) {
	query := `UPDATE storagetransaction SET staff_id = $1, product_id = $2, storagetransactiontype = $3, 
                                   price = $4, quantity = $5, update_at = NOW() WHERE id = $6
`

	_, err := s.DB.Exec(ctx, query,
		&transaction.StaffID,
		&transaction.ProductID,
		&transaction.StorageTransactionType,
		&transaction.Price,
		&transaction.Quantity,
		&transaction.ID,
	)
	if err != nil {
		log.Println("Error while storagetransactions Repository :", err)
		return "", err
	}

	return transaction.ID, nil
}

func (s storageTransactionRepo) Delete(ctx context.Context, id models.PrimaryKey) error {
	query := `UPDATE storagetransaction SET deleted_at = NOW() WHERE id = $1`

	_, err := s.DB.Exec(ctx, query, id.ID)
	if err != nil {
		log.Println("Error while deleting storagetransactions ", err)
		return err
	}

	return nil
}
