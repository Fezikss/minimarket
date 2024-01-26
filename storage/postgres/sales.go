package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"main.go/api/models"
	"main.go/storage"
)

type saleRepo struct {
	DB *sql.DB
}

func NewSaleRepo(db *sql.DB) storage.ISaleStorage {
	return saleRepo{DB: db}

}

func (s saleRepo) Create(sale models.CreateSale) (string, error) {
	id := uuid.New()
	sale.CreatedAt = time.Now()
	if _, err := s.DB.Exec(`insert into sale values($1,$2,&3,&4,&5,&6,&7,&8,&9)`, id, sale.BranchID, sale.ShopAssistantID, sale.Cashier, sale.PaymentType, sale.Price, sale.Status,sale.ClientName, sale.CreatedAt); err != nil {
		fmt.Println("error while inserting data to sale")
		return "", err
	}
	return id.String(), nil
}
func (s saleRepo) GetById(pk models.PrimaryKey) (models.Sale, error) {
	sale:=models.Sale{}
	if err := s.DB.QueryRow(`select id, branch_id , shop_assistant_id, cashier,payment_type,price,status,client_name,created_at from sale where id=$1 `, pk.ID).Scan(
		&sale.ID,
		&sale.BranchID,
		&sale.ShopAssistantID,
		&sale.Cashier,
		&sale.PaymentType,
		&sale.Price,
		&sale.Status,
		&sale.ClientName,
		&sale.CreatedAt); err != nil {
		fmt.Println("error getting by id sale")
		return models.Sale{}, err
	}
	return sale, nil
}
func (s saleRepo) GetList(request models.GetListRequest) (models.SaleResponse, error) {
	var (
		sales           = []models.Sale{}
		count             = 0
		query, countquery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)
	countquery = `select count(1) from sale `
	if search != "" {
		countquery += fmt.Sprintf(` and client_name ilike '%%%s'`, search)
	}
	if err := s.DB.QueryRow(countquery).Scan(&count); err != nil {
		fmt.Println("error while counting")
		return models.SaleResponse{}, err
	}
	query = `select id,branch_id,shop_assistant_id,cashier,payment_type,price,status,client_name,created_at from sale `
	if search != "" {
		query += ` LIMIT $1 OFFSET $2`

	}
	rows, err := s.DB.Query(query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting sales", err.Error())
		return models.SaleResponse{}, err
	}
	for rows.Next() {
		sale:=models.Sale{}
		if err = rows.Scan(&sale.ID,&sale.BranchID,&sale.ShopAssistantID,&sale.Cashier,&sale.PaymentType,&sale.Price,&sale.ClientName,&sale.ClientName,&sale.CreatedAt); err != nil {
			fmt.Println("error while getting list of sales ")
			return models.SaleResponse{}, err
		}
		sales = append(sales,sale)
	}
	return models.SaleResponse{
		Sales: sales,
		Count: count,
		
	} ,nil
}
func (s saleRepo) Update(sale models.UpdateSale) (string, error) {
	sales:=models.Sale{}
	sale.UpdatedAt = time.Now()
	if _, err := s.DB.Exec(`update sale set branch_id=$1,shop_assistant=$2,cashier=$3 ,payment_type=$4,price=$5,updated_at=$6 where id=$7`, 
	&sale.BranchID,&sale.ShopAssistantID,&sale.Cashier,&sale.PaymentType,&sale.Price,&sale.UpdatedAt,&sale.ID); err != nil {
		return "", err
	}
	if err := s.DB.QueryRow(`select * from sale where id=$1`, sales.ID).Scan(&sales.ID,&sales.BranchID,&sales.ShopAssistantID,&sales.Cashier,&sales.PaymentType,&sales.Price,&sales.Status,&sales.ClientName,&sales.UpdatedAt,&sales.CreatedAt); err != nil {
		fmt.Println("error while updating sales")
		return "", err
	}
	return sales.ID, nil
}
func (s saleRepo) Delete(pk models.PrimaryKey) error {
	if _, err := s.DB.Exec(`delete from sale where id=$1`, pk.ID); err != nil {
		return err
	}
	return nil

}
