package handler

import (
	"context"
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

// CreateStorageTransaction godoc
// @Router       /storagetransaction [POST]
// @Summary      Create a new stransaction
// @Description  create a new stransaction
// @Tags         stransaction
// @Accept       json
// @Produce      json
// @Param 		 stransaction body models.CreateStorageTransaction false "stransaction"
// @Success      200  {object}  models.StorageTransaction
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateStorageTransaction(c *gin.Context) {
	rtransaction := models.CreateStorageTransaction{}

	if err := c.ShouldBindJSON(&rtransaction); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.StorageTransaction().Create(context.Background(), rtransaction)
	if err != nil {
		handleResponse(c, "error while creating repository transaction", http.StatusInternalServerError, err.Error())
		return
	}

	createdRTransaction, err := h.storage.StorageTransaction().GetById(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, createdRTransaction)
}

// GetByIdStorageTransaction godoc
// @Router       /storagetransaction/{id} [GET]
// @Summary      Get stransaction by id
// @Description  get stransaction by id
// @Tags         stransaction
// @Accept       json
// @Produce      json
// @Param 		 id path string true "stransaction_id"
// @Success      200  {object}  models.StorageTransaction
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetByIdStorageTransaction(c *gin.Context) {
	uid := c.Param("id")

	fmt.Println("eeeeeee")
	repository, err := h.storage.StorageTransaction().GetById(context.Background(), models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting storage transaction by ID", http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("rrrrrrrrrrr")
	handleResponse(c, "", http.StatusOK, repository)
}

// GetStorageTransactionList godoc
// @Router       /storagetransactions [GET]
// @Summary      Get stransaction list
// @Description  get stransaction list
// @Tags         stransaction
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.StorageTransactionsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetStorageTransactionList(c *gin.Context) {
	var (
		page, limit int
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	search := c.Query("search")

	response, err := h.storage.StorageTransaction().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting repository list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, response)
}

// UpdateStorageTransaction godoc
// @Router       /storagetransaction/{id} [PUT]
// @Summary      Update stransaction
// @Description  get stransaction
// @Tags         stransaction
// @Accept       json
// @Produce      json
// @Param 		 id path string true "stransaction_id"
// @Param 		 rtransaction body models.UpdateStorageTransaction false "stransaction"
// @Success      200  {object}  models.Storage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateStorageTransaction(c *gin.Context) {
	uid := c.Param("id")

	rTransaction := models.UpdateStorageTransaction{}
	if err := c.ShouldBindJSON(&rTransaction); err != nil {
		handleResponse(c, "error while reading from body", http.StatusBadRequest, err.Error())
		return
	}

	rTransaction.ID = uid
	if _, err := h.storage.StorageTransaction().Update(context.Background(), rTransaction); err != nil {
		handleResponse(c, "error while updating repository transaction ", http.StatusInternalServerError, err.Error())
		return
	}

	updatedRTransaction, err := h.storage.StorageTransaction().GetById(context.Background(), models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, updatedRTransaction)
}

// DeleteStorageTransaction godoc
// @Router       /storagetransaction/{id} [DELETE]
// @Summary      Delete stransaction
// @Description  delete stransaction
// @Tags         stransaction
// @Accept       json
// @Produce      json
// @Param 		 id path string true "stransaction_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteStorageTransaction(c *gin.Context) {
	uid := c.Param("id")
	uuid:=models.PrimaryKey{ID: uid}

	if err := h.storage.StorageTransaction().Delete(context.Background(), uuid); err != nil {
		handleResponse(c, "error while deleting repository transaction ", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "storage transaction deleted")
}