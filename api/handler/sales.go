package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

func (h Handler) CreateSale(c *gin.Context) {
	createsale := models.CreateSale{}
	if err := c.ShouldBindJSON(&createsale); err != nil {
		handleResponse(c, "error while decoding sale from json", 400, err.Error())
		return
	}
	id, err := h.storage.Sale().Create(createsale)
	if err != nil {
		handleResponse(c, "error is while getting by id sale", 500, err.Error())
		return
	}
	res, err := h.storage.Sale().GetById(models.PrimaryKey{ID: id})
	if err != nil {

		handleResponse(c, "error while getting by id after creating sale", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, res)

}
func (h Handler) GetByIdSale(c *gin.Context) {
	var err error
	uid := c.Param("id")
	sale, err := h.storage.Sale().GetById(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by id", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, sale)

}
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
	sale, err := h.storage.Sale().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	handleResponse(c, "", 200, sale)

}

func (h Handler) UpdateSale(c *gin.Context) {
	updatesale := models.UpdateSale{}
	uid := c.Param("id")
	if err := c.ShouldBindJSON(&updatesale); err != nil {
		handleResponse(c, "erro while decoding sale", 500, err.Error())
		return
	}
	updatesale.ID = uid
	if _, err := h.storage.Sale().Update(updatesale); err != nil {
		handleResponse(c, "error while updating sale", 500, err.Error())
		return
	}
	ids := models.PrimaryKey{ID: uid}
	res, err := h.storage.Sale().GetById(ids)
	if err != nil {
		fmt.Println("error while getting by id")
		handleResponse(c, "error", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, res)

}
func (h Handler) DeleteSale(c *gin.Context) {
	uid := c.Param("id")
	saleid := models.PrimaryKey{ID: uid}
	if err := h.storage.Branch().Delete(saleid); err != nil {
		handleResponse(c, "error while deleting sale", 500, err)
		return
	}
	handleResponse(c, "", http.StatusOK, nil)
}
