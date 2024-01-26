package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

func (h Handler) CreateBranch(c *gin.Context) {
	createbranch := models.CreateBranch{}
	if err := c.ShouldBindJSON(&createbranch); err != nil {
		handleResponse(c, "error while decoding branch from json", 400, err.Error())
		return
	}
	id, err := h.storage.Branch().Create(createbranch)
	if err != nil {
		handleResponse(c, "error is while getting by id branch", 500, err.Error())
		return
	}
	res, err := h.storage.Branch().GetById(models.PrimaryKey{ID: id})
	if err != nil {

		handleResponse(c, "error while getting by id after creating branch", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, res)

}
func (h Handler) GetByIdBranch(c *gin.Context) {
	var err error
	uid := c.Param("id")
	branch, err := h.storage.Branch().GetById(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by id", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, branch)

}
func (h Handler) GetListBranch(c *gin.Context) {
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
	branch, err := h.storage.Branch().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	handleResponse(c, "", 200, branch)

}

func (h Handler) UpdateBranch(c *gin.Context) {
	updatebranch := models.UpdateBranch{}
	uid := c.Param("id")
	if err := c.ShouldBindJSON(&updatebranch); err != nil {
		handleResponse(c, "error while decoding branch", 500, err.Error())
		return
	}
	updatebranch.ID = uid
	if _, err := h.storage.Branch().Update(updatebranch); err != nil {
		handleResponse(c, "error while updating branch", 500, err.Error())
		return
	}
	ids := models.PrimaryKey{ID: uid}
	res, err := h.storage.Branch().GetById(ids)
	if err != nil {
		handleResponse(c, "error while getting branch by ID", 500, err.Error())
		return
	}
	handleResponse(c, "branch updated successfully", 200, res)
}

func (h Handler) DeleteBranch(c *gin.Context) {
	uid := c.Param("id")
	branchid := models.PrimaryKey{ID: uid}
	if err := h.storage.Branch().Delete(branchid); err != nil {
		handleResponse(c, "error while deleting branch", 500, err)
		return
	}
	handleResponse(c, "", http.StatusOK, nil)
}
