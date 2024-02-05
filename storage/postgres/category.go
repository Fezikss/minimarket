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

type categoryRepo struct {
	DB *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) storage.ICategoryStorage {
	return categoryRepo{DB: db}
}
func (c categoryRepo) Create(ctx context.Context, category models.CreateCategory) (string, error) {
	id := uuid.New()
	category.CreatedAt = time.Now()
	if _, err := c.DB.Exec(ctx, `INSERT INTO category (id, parent_id, created_at, name) VALUES ($1, $2, $3, $4)`, id, category.ParentID, category.CreatedAt, category.Name); err != nil {
		fmt.Println("error while inserting data to category")
		return "", err
	}
	return id.String(), nil
}

func (c categoryRepo) GetById(ctx context.Context, pk models.PrimaryKey) (models.Category, error) {
	category := models.Category{}
	if err := c.DB.QueryRow(ctx, "select id, parent_id, update_at,created_at, name  from category where id = $1", pk.ID).Scan(
		&category.ID,
		&category.ParentID,
		&category.UpdatedAt,
		&category.CreatedAt,
		&category.Name); err != nil {
		fmt.Println("error getting by id category")
		return models.Category{}, err
	}
	return category, nil
}

func (c categoryRepo) GetList(ctx context.Context, request models.GetListRequest) (models.CategoryResponse, error) {
	var (
		categories        = []models.Category{}
		count             = 0
		query, countquery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)
	countquery = `select count(1) from category `
	if search != "" {
		countquery += fmt.Sprintf(` and name ilike '%%%s'`, search)
	}
	if err := c.DB.QueryRow(ctx, countquery).Scan(&count); err != nil {
		fmt.Println("error while counting")
		return models.CategoryResponse{}, err
	}
	query = `select id, parent_id, update_at, created_at, name from category `

	query += ` LIMIT $1 OFFSET $2`
	rows, err := c.DB.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting categories", err.Error())
		return models.CategoryResponse{}, err
	}
	for rows.Next() {
		category := models.Category{}
		if err = rows.Scan(&category.ID, &category.ParentID, &category.UpdatedAt, &category.CreatedAt, &category.Name); err != nil {
			fmt.Println("error while getting list of categories ")
			return models.CategoryResponse{}, err
		}
		categories = append(categories, category)
	}
	return models.CategoryResponse{
		Categories: categories,
		Count:      count,
	}, nil
}
func (c categoryRepo) Update(ctx context.Context, category models.UpdateCategory) (string, error) {
	category.UpdatedAt = time.Now()

	_, err := c.DB.Exec(
		ctx,
		"UPDATE category SET name=$1, parent_id=$2, update_at=$3 WHERE id=$4",
		category.Name, category.ParentID, category.UpdatedAt, category.ID,
	)
	
	if err != nil {
		return "", err
	}

	return category.ID, nil
}

func (c categoryRepo) Delete(ctx context.Context, pk models.PrimaryKey) error {

	if _, err := c.DB.Exec(ctx, `delete from product where category_id=$1`, pk.ID); err != nil {
		return err
	}

	if _, err := c.DB.Exec(ctx, `delete from category where id=$1`, pk.ID); err != nil {
		return err
	}
	return nil

}
