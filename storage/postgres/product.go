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

type productRepo struct {
	DB *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) storage.IProductStorage {
	return productRepo{DB: db}
}
func (p productRepo) Create(ctx context.Context, product models.CreateProduct) (string, error) {
	id := uuid.New()
	product.CreatedAt = time.Now()
	if _, err := p.DB.Exec(ctx, `insert into product (id, name, price,barcode,category_id, created_at) values($1, $2, $3, $4,$5,$6)`, id, product.Name, product.Price, product.Barcode, product.CategoryID, product.CreatedAt); err != nil {
		fmt.Println("error while inserting data to product")
		return "", err
	}
	return id.String(), nil
}
func (p productRepo) GetById(ctx context.Context, pk models.PrimaryKey) (models.Product, error) {
	product := models.Product{}
	if err := p.DB.QueryRow(ctx, "select id, name, price,barcode, category_id, created_at from product where id = $1", pk.ID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Barcode,
		&product.CategoryID,
		&product.CreatedAt); err != nil {
		fmt.Println("error getting by id product")
		return models.Product{}, err
	}
	return product, nil
}

func (p productRepo) GetList(ctx context.Context, request models.GetListRequestProduct) (models.ProductResponse, error) {
	var (
		products          = []models.Product{}
		count             = 0
		query, countquery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)
	countquery = `select count(1) from product where 1=1`
	if search != "" {
		countquery += fmt.Sprintf(` and name ilike '%%%s%%'`, search)
	}
	if request.Barcode != 0 {
        countquery += fmt.Sprintf(` AND barcode = %d`, request.Barcode)
    }
	if err := p.DB.QueryRow(ctx, countquery).Scan(&count); err != nil {
		fmt.Println("error while counting")
		return models.ProductResponse{}, err
	}
	query = `select id, name , price, barcode, category_id, update_at, created_at from product `

	query += ` LIMIT $1 OFFSET $2`
	rows, err := p.DB.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting product", err.Error())
		return models.ProductResponse{}, err
	}
	for rows.Next() {
		product := models.Product{}
		if err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Barcode, &product.CategoryID, &product.UpdatedAt, &product.CreatedAt); err != nil {
			fmt.Println("error while getting list of products ")
			return models.ProductResponse{}, err
		}
		products = append(products, product)
	}
	return models.ProductResponse{
		Products: products,
		Count:    count,
	}, nil
}

func (p productRepo) Update(ctx context.Context, product models.UpdateProduct) (string, error) {
	product.UpdatedAt = time.Now()

	// Corrected placeholder order: $1, $2, $3, $4, $5, $6
	if _, err := p.DB.Exec(ctx, "UPDATE product SET name=$1, price=$2, category_id=$3, update_at=$4 WHERE id=$5", product.Name, product.Price, product.CategoryID, product.UpdatedAt, product.ID); err != nil {
		return "", err
	}

	return product.ID, nil
}

func (p productRepo) Delete(ctx context.Context, pk models.PrimaryKey) error {
	if _, err := p.DB.Exec(ctx, `delete from storage where product_id=$1`, pk.ID); err != nil {
		return err
	}
	if _, err := p.DB.Exec(ctx, `delete from storagetransaction where product_id=$1`, pk.ID); err != nil {
		return err
	}
	if _, err := p.DB.Exec(ctx, `delete from basket where product_id=$1`, pk.ID); err != nil {
		return err
	}

	if _, err := p.DB.Exec(ctx, `delete from product where id=$1`, pk.ID); err != nil {
		return err
	}
	return nil

}
