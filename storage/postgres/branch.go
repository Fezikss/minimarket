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

type branchRepo struct {
	DB *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.IBranchStorage {
	return branchRepo{DB: db}
}

func (b branchRepo) Create(ctx context.Context, branch models.CreateBranch) (string, error) {
	id := uuid.New()
	branch.CreatedAt = time.Now()
	if _, err := b.DB.Exec(ctx, `insert into branch (id, name, address, created_at) values($1, $2, $3, $4)`, id, branch.Name, branch.Address, branch.CreatedAt); err != nil {
		fmt.Println("error while inserting data to branch")
		return "", err
	}
	return id.String(), nil
}

func (b branchRepo) GetById(ctx context.Context, pk models.PrimaryKey) (models.Branch, error) {
	branch := models.Branch{}
	if err := b.DB.QueryRow(ctx, "select id, name, address, update_at, created_at from branch where id = $1", pk.ID).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.UpdatedAt,
		&branch.CreatedAt); err != nil {
		fmt.Println("error getting by id branch")
		return models.Branch{}, err
	}
	return branch, nil
}

func (b branchRepo) GetList(ctx context.Context, request models.GetListRequest) (models.BranchResponse, error) {
	var (
		branchs           = []models.Branch{}
		count             = 0
		query, countquery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)
	countquery = `select count(1) from branch `
	if search != "" {
		countquery += fmt.Sprintf(` and name ilike '%%%s'`, search)
	}
	if err := b.DB.QueryRow(ctx, countquery).Scan(&count); err != nil {
		fmt.Println("error while counting")
		return models.BranchResponse{}, err
	}
	query = `select id, name , address, update_at, created_at from branch `

	query += ` LIMIT $1 OFFSET $2`
	rows, err := b.DB.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting branchs", err.Error())
		return models.BranchResponse{}, err
	}
	for rows.Next() {
		b := models.Branch{}
		if err = rows.Scan(&b.ID, &b.Name, &b.Address, &b.UpdatedAt, &b.CreatedAt); err != nil {
			fmt.Println("error while getting list of branchs ")
			return models.BranchResponse{}, err
		}
		branchs = append(branchs, b)
	}
	return models.BranchResponse{
		Branchs: branchs,
		Count:   count,
	}, nil
}
func (b branchRepo) Update(ctx context.Context, branch models.UpdateBranch) (string, error) {
	branchs := models.Branch{}
	branch.UpdatedAt = time.Now()
	if _, err := b.DB.Exec(ctx, "update branch set name=$1, address=$2, update_at=$3 where id=$4", branch.Name, branch.Address, branch.UpdatedAt, branch.ID); err != nil {
		return "", err
	}
	if err := b.DB.QueryRow(ctx, "select id, name, address, update_at, created_at from branch where id=$1", branch.ID).Scan(&branchs.ID, &branchs.Name, &branchs.Address, &branchs.UpdatedAt, &branchs.CreatedAt); err != nil {
		fmt.Println("error while updating branch")
		return "", err
	}
	return branchs.ID, nil
}

func (b branchRepo) Delete(ctx context.Context, pk models.PrimaryKey) error {
	if _, err := b.DB.Exec(ctx, `delete from storage where branch_id=$1`, pk.ID); err != nil {
		return err
	}
	if _, err := b.DB.Exec(ctx, `delete from storagetransaction where branch_id=$1`, pk.ID); err != nil {
		return err
	}

	if _, err := b.DB.Exec(ctx, `delete from branch where id=$1`, pk.ID); err != nil {
		return err
	}
	return nil

}
