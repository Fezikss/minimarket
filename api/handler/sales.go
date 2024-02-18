package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

// StartSale godoc
// @Router       /startsale [POST]
// @Summary      Creates a new sale
// @Description  create a new sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        sale body models.CreateSale false "sale"
// @Success      201  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) StartSale(c *gin.Context) {
	startsale := models.CreateSale{}
	if err := c.ShouldBindJSON(&startsale); err != nil {
		handleResponse(c, "error while decoding sale from json", 400, err.Error())
		return
	}
	id, err := h.storage.Sale().Create(context.Background(), startsale)
	if err != nil {
		handleResponse(c, "error while creating sale", 500, err.Error())
		return
	}
	res, err := h.storage.Sale().GetById(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error while getting sale by id after creation", 500, err.Error())
		return
	}
	handleResponse(c, "succesfully created", 200, res)

}

// UpdateEndSale godoc
// @Router       /endsale/{id} [PUT]
// @Summary      Update endsale
// @Description  update endsale
// @Tags         endsale
// @Accept       json
// @Produce      json
// @Param 		 id path string true "end_sale_id"
// @Param        status body models.SaleStatus
// @Success      200  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) EndSale(c *gin.Context) {
	uid := c.Param("id")
	salestatus := models.SaleStatus{}
	if err := c.ShouldBindJSON(&salestatus); err != nil {
		handleResponse(c, "error while reading status from body", 500, err)
		return
	}

	basket, err := h.storage.Basket().GetList(context.Background(), models.GetListRequest{
		Page:   1,
		Limit:  10,
		Search: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting baskets list", 500, err.Error())
		log.Fatal("error", err.Error())
		return
	}
	totalPrice := 0
	selectedbaskets := map[string]models.Basket{}

	for _, v := range basket.Basket {
		totalPrice += int(v.Price)
		selectedbaskets[v.ProductID] = v
	}

	idprice, err := h.storage.Sale().UpdatePriceSale(context.Background(), models.EndSaleUpdate{
		ID:    uid,
		Price: float64(totalPrice),
	})
	fmt.Println("id", idprice)

	response, err := h.storage.Sale().GetById(context.Background(), models.PrimaryKey{ID: idprice})
	if err != nil {
		handleResponse(c, "error", 500, "err")
	}
	fmt.Println("resp", response)

	storages, err := h.storage.Storage().GetList(context.Background(), models.GetListRequest{Page: 1, Limit: 100000})
	if err != nil {
		handleResponse(c, "error in line 93 sale logic", 500, err.Error())
		return
	}

	mapp := map[string]models.Storage{}
	for _, s := range storages.Storages {
		mapp[s.ID] = s
	}

	for i, v := range mapp {
		if v.ProductID == selectedbaskets[v.ProductID].ProductID {

			up, err := h.storage.Storage().Update(context.Background(), models.UpdateStorage{
				ID:        i,
				ProductID: v.ProductID,
				BranchID:  v.BranchID,
				Count:     v.Count - selectedbaskets[v.ProductID].Quantity,
			})
			if err != nil {
				handleResponse(c, "error in line 108 handler", 500, "error")
			}
			fmt.Println("up", up)
			st, err := h.storage.StorageTransaction().Create(context.Background(), models.CreateStorageTransaction{
				StaffID:                response.Cashier,
				ProductID:              v.ProductID,
				StorageTransactionType: "minus",
				Price:                  int(selectedbaskets[v.ProductID].Price),
				Quantity:               selectedbaskets[v.ProductID].Quantity,
				BranchID:               v.BranchID,
			})
			fmt.Println("create st", st)
			if err != nil {
				fmt.Println(err.Error())
			}

		}
	}

	staffs, err := h.storage.Staff().GetList(context.Background(), models.GetListRequest{
		Page:   1,
		Limit:  10,
		Search: response.Cashier,
	})
	if err != nil {
		handleResponse(c, "error while ggeting staff by id in endsale", 500, err.Error())
		return
	}

	for _, v := range staffs.Staffs {
		staftarifid, err := h.storage.StaffTariff().GetById(context.Background(), models.PrimaryKey{ID: v.TariffID})
		if err != nil {
			handleResponse(c, "error wile getting bu id of stafftarif", 500, err.Error())
			return
		}
		if staftarifid.TariffType == "percent" && response.PaymentType == "card" {
			_, err := h.storage.Staff().Update(context.Background(), models.UpdateStaff{
				ID:        v.ID,
				BranchID:  v.BranchID,
				TariffID:  v.TariffID,
				StaffType: v.StaffType,
				Name:      v.Name,
				Login:     v.Login,
				Balance:   (uint(response.Price) * 10) / 100,
			})
			if err != nil {
				handleResponse(c, "err", 500, "cannot calculate staff balance")
				return
			}

		}
		if staftarifid.TariffType == "fixed" && response.PaymentType == "cash" {
			h.storage.Staff().Update(context.Background(), models.UpdateStaff{
				ID:        v.ID,
				BranchID:  v.BranchID,
				TariffID:  v.TariffID,
				StaffType: v.StaffType,
				Name:      v.Name,
				Login:     v.Login,
				Balance:   uint(response.Price),
			})
		}

	}

	var transactiontype string
	switch salestatus.Status {
	case "cancel":
		transactiontype = "withdraw"
	case "success":
		transactiontype = "topup"
	}
	createdTransaction, err := h.storage.Transaction().Create(c.Request.Context(), models.CreateTransaction{
		SaleID:          response.ID,
		StaffID:         response.Cashier,
		TransactionType: transactiontype,
		SourceType:      "sales",
		Amount:          response.Price,
		Description:     "desciption:'created' ",
	})
	if err != nil {
		handleResponse(c, "error", 500, err)
		return
	}
	idt, err := h.storage.Transaction().GetById(context.Background(), models.PrimaryKey{ID: createdTransaction})
	if err != nil {
		handleResponse(c, "error", 500, err)
		return
	}

	handleResponse(c, "", 200, idt)

}

// CreateSale godoc
// @Router       /sale [POST]
// @Summary      Creates a new sale
// @Description  create a new sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        sale body models.CreateSale false "sale"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateSale(c *gin.Context) {
	createsale := models.CreateSale{}
	if err := c.ShouldBindJSON(&createsale); err != nil {
		handleResponse(c, "error while decoding sale from json", 400, err.Error())
		return
	}
	id, err := h.storage.Sale().Create(context.Background(), createsale)
	if err != nil {
		handleResponse(c, "error while creating sale", 500, err.Error())
		return
	}
	res, err := h.storage.Sale().GetById(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error while getting sale by id after creation", 500, err.Error())
		return
	}
	handleResponse(c, "sale created successfully", 200, res)
}

// GetSale godoc
// @Router       /sale/{id} [GET]
// @Summary      Gets sale
// @Description  get sale by ID
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        id path string true "sale"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetByIdSale(c *gin.Context) {
	var err error
	uid := c.Param("id")
	sale, err := h.storage.Sale().GetById(context.Background(), models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by id", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, sale)

}

// GetSaleList godoc
// @Router       /sales [GET]
// @Summary      Get user list
// @Description  get user list
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.SaleResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetListSale(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)
	pagestring := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pagestring)
	if err != nil {
		handleResponse(c, "error while converting pagestr", 400, err)
		return
	}
	limitstr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitstr)
	if err != nil {
		handleResponse(c, "error while converting limit ", 400, err.Error())
		return

	}
	search = c.Query("search")
	sale, err := h.storage.Sale().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	handleResponse(c, "", 200, sale)

}

// UpdateSale godoc
// @Router       /sale/{id} [PUT]
// @Summary      Update sale
// @Description  update sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param 		 id path string true "sale_id"
// @Param        sale body models.UpdateSale true "sale"
// @Success      200  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateSale(c *gin.Context) {
	updatesale := models.UpdateSale{}
	uid := c.Param("id")
	if err := c.ShouldBindJSON(&updatesale); err != nil {
		handleResponse(c, "erro while decoding sale", 500, err.Error())
		return
	}
	updatesale.ID = uid
	if _, err := h.storage.Sale().Update(context.Background(), updatesale); err != nil {
		handleResponse(c, "error while updating sale", 500, err.Error())
		return
	}
	ids := models.PrimaryKey{ID: uid}
	res, err := h.storage.Sale().GetById(context.Background(), ids)
	if err != nil {
		fmt.Println("error while getting by id")
		handleResponse(c, "error", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, res)

}

// DeleteSale godoc
// @Router       /sale/{id} [DELETE]
// @Summary      Delete sale
// @Description  delete sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param 		 id path string true "sale_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteSale(c *gin.Context) {
	uid := c.Param("id")
	saleid := models.PrimaryKey{ID: uid}
	if err := h.storage.Sale().Delete(context.Background(), saleid); err != nil {
		handleResponse(c, "error while deleting sale", 500, err)
		return
	}
	handleResponse(c, "", http.StatusOK, nil)
}
