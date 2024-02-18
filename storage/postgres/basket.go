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

type basketRepo struct {
	DB *pgxpool.Pool
}

func NewBasketRepo(db *pgxpool.Pool) storage.IBasketStorage {
	return basketRepo{DB: db}
}

func (b basketRepo) Create(ctx context.Context, basket models.CreateBasket) (string, error) {
	id := uuid.New()
	basket.CreateAt = time.Now()
	if _, err := b.DB.Exec(ctx, `insert into basket (id, product_id, quantity,price, created_at,sale_id) values($1, $2, $3, $4,$5,$6)`, id, basket.ProductID, basket.Quantity, basket.Price, basket.CreateAt, basket.SaleID); err != nil {
		fmt.Println("error while inserting data to basket")
		return "", err
	}
	return id.String(), nil
}

func (b basketRepo) GetById(ctx context.Context, pk models.PrimaryKey) (models.Basket, error) {
	basket := models.Basket{}
	if err := b.DB.QueryRow(ctx, "select id, product_id, quantity,price, update_at, created_at , sale_id from basket where id = $1", pk.ID).Scan(
		&basket.ID,
		&basket.ProductID,
		&basket.Quantity,
		&basket.Price,
		&basket.UpdateAt,
		&basket.CreateAt,
		&basket.SaleID,
	); err != nil {
		fmt.Println("error getting by id basket")
		return models.Basket{}, err
	}
	return basket, nil
}

func (b basketRepo) GetList(ctx context.Context, request models.GetListRequest) (models.BasketResponse, error) {
	var (
		baskets           = []models.Basket{}
		count             = 0
		query, countquery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)
	countquery = `select count(1) from basket where 1=1`
	
	if search != "" {
		countquery += fmt.Sprintf(` and sale_id = '%s'`, search)
	}
	if err := b.DB.QueryRow(ctx, countquery).Scan(&count); err != nil {
		fmt.Println("error while counting", err.Error())
		return models.BasketResponse{}, err
	}

	query = ` select id, product_id , quantity, price, update_at, created_at , sale_id from basket `

	if search != "" {
		query += fmt.Sprintf(` where sale_id = '%s'`, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := b.DB.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting baskets", err.Error())
		return models.BasketResponse{}, err
	}
	
	for rows.Next() {
		bas := models.Basket{}
		if err = rows.Scan(&bas.ID, &bas.ProductID, &bas.Quantity, &bas.Price, &bas.UpdateAt, &bas.CreateAt, &bas.SaleID); err != nil {
			fmt.Println("error while getting list of baskets ")
			return models.BasketResponse{}, err
		}
		baskets = append(baskets, bas)
	}

	fmt.Println("basket post", baskets)
	return models.BasketResponse{
		Basket: baskets,
		Count:  count,
	}, nil
}

func (b basketRepo) Update(ctx context.Context, basket models.UpdateBasket) (string, error) {
	basket.UpdateAt = time.Now()
	if _, err := b.DB.Exec(ctx, "update basket set product_id=$1, quantity=$2, update_at=$3 ,price=$4 where id=$5", basket.ProductID, basket.Quantity, basket.UpdateAt,basket.Price, basket.ID); err != nil {
		return "", err
	}
	return basket.ID, nil
}

func (b basketRepo) Delete(ctx context.Context, pk models.PrimaryKey) error {
	if _, err := b.DB.Exec(ctx, `delete from basket where id=$1`, pk.ID); err != nil {
		return err
	}
	return nil

}
