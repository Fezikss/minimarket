package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

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
	id, err := h.storage.Sale().Create(context.Background(),createsale)
	if err != nil {
		handleResponse(c, "error while creating sale", 500, err.Error())
		return
	}
	res, err := h.storage.Sale().GetById(context.Background(),models.PrimaryKey{ID: id})
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
	sale, err := h.storage.Sale().GetById(context.Background(),models.PrimaryKey{ID: uid})
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
	sale, err := h.storage.Sale().GetList(context.Background(),models.GetListRequest{
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
	if _, err := h.storage.Sale().Update(context.Background(),updatesale); err != nil {
		handleResponse(c, "error while updating sale", 500, err.Error())
		return
	}
	ids := models.PrimaryKey{ID: uid}
	res, err := h.storage.Sale().GetById(context.Background(),ids)
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
	if err := h.storage.Sale().Delete(context.Background(),saleid); err != nil {
		handleResponse(c, "error while deleting sale", 500, err)
		return
	}
	handleResponse(c, "", http.StatusOK, nil)
}
