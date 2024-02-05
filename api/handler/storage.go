package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

// CreateStorage godoc
// @Router       /storage [POST]
// @Summary      Creates a new storage
// @Description  create a new storage
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param        storage body models.CreateStorage false "storage"
// @Success      201  {object}  models.Storage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateStorage(c *gin.Context) {
	createstorage := models.CreateStorage{}
	if err := c.ShouldBindJSON(&createstorage); err != nil {
		handleResponse(c, "error while decoding storage from json", 400, err.Error())
		return
	}
	id, err := h.storage.Storage().Create(context.Background(),createstorage)
	if err != nil {
		handleResponse(c, "error is while getting by id storage", 500, err.Error())
		return
	}
	res, err := h.storage.Storage().GetById(context.Background(),models.PrimaryKey{ID: id})
	if err != nil {

		handleResponse(c, "error while getting by id after creating storage", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, res)

}

// GetByIdStorage godoc
// @Router       /storage/{id} [GET]
// @Summary      Get storage by id
// @Description  get storage by ID
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param        id path string true "storage"
// @Success      200  {object}  models.Storage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetByIdStorage(c *gin.Context) {
	var err error
	uid := c.Param("id")
	storage, err := h.storage.Storage().GetById(context.Background(),models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by id", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, storage)

}

// GetStorageList godoc
// @Router       /storages [GET]
// @Summary      Get storage list
// @Description  get storage list
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.StorageResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetListStorage(c *gin.Context) {
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
	storage, err := h.storage.Storage().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	handleResponse(c, "", 200, storage)

}

// UpdateStorage godoc
// @Router       /storage/{id} [PUT]
// @Summary      Update storage
// @Description  update storage
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param 		 id path string true "storage_id"
// @Param        storage body models.UpdateStorage true "storage"
// @Success      200  {object}  models.Storage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateStorage(c *gin.Context) {
	updatestorage := models.UpdateStorage{}
	uid := c.Param("id")
	if err := c.ShouldBindJSON(&updatestorage); err != nil {
		handleResponse(c, "error while decoding storage update", 500, err.Error())
		return
	}
	fmt.Println(updatestorage)
	updatestorage.ID = uid
	if _, err := h.storage.Storage().Update(context.Background(),updatestorage); err != nil {
		handleResponse(c, "error while updating storage", 500, err.Error())
		return
	}
	ids := models.PrimaryKey{ID: uid}
	res, err := h.storage.Storage().GetById(context.Background(),ids)
	if err != nil {
		handleResponse(c, "error while getting storage by ID", 500, err.Error())
		return
	}
	handleResponse(c, "storage updated successfully", 200, res)
}

// DeleteStorage godoc
// @Router       /storage/{id} [DELETE]
// @Summary      Delete storage
// @Description  delete storage
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param 		 id path string true "storage_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteStorage(c *gin.Context) {
	uid := c.Param("id")
	storageid := models.PrimaryKey{ID: uid}
	if err := h.storage.Storage().Delete(context.Background(),storageid); err != nil {
		handleResponse(c, "error while deleting storage", 500, err)
		return
	}
	handleResponse(c, "", http.StatusOK, nil)
}
