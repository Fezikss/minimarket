package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

// CreateBranch   godoc
// @Router               /branch [POST]
// @Summary              Creates a new branch
// @Description          Creates a new branch
// @Tags                 branch
// @Accept               json
// @Produce              json
// @Param                branch body models.CreateBranch true "branch"
// @Success              201  {object}  models.Response
// @Failure              400  {object}  models.Response
// @Failure              404  {object}  models.Response
// @Failure              500  {object}  models.Response
func (h Handler) CreateBranch(c *gin.Context) {
	createbranch := models.CreateBranch{}
	if err := c.ShouldBindJSON(&createbranch); err != nil {
		handleResponse(c, "error while decoding branch from json", 400, err.Error())
		return
	}
	id, err := h.storage.Branch().Create(context.Background(),createbranch)
	if err != nil {
		handleResponse(c, "error is while getting by id branch", 500, err.Error())
		return
	}
	res, err := h.storage.Branch().GetById(context.Background(),models.PrimaryKey{ID: id})
	if err != nil {

		handleResponse(c, "error while getting by id after creating branch", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, res)

}

// GetByIdBranch   godoc
// @Router            /branch/{id} [GET]
// @Summary           Get branch by id
// @Description       get branch by id
// @Tags              branch
// @Accept            json
// @Produce           json
// @Param             id path string true "branch"
// @Success           200  {object}  models.Branch
// @Failure           400  {object}  models.Response
// @Failure           404  {object}  models.Response
// @Failure           500  {object}  models.Response
func (h Handler) GetByIdBranch(c *gin.Context) {
	var err error
	uid := c.Param("id")
	branch, err := h.storage.Branch().GetById(context.Background(),models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by id", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, branch)

}


// GetBranchList           godoc
// @Router                 /branchs [GET]
// @Summary                 Get branchs list
// @Description             get branchs list
// @Tags                    branch
// @Accept                  json
// @Produce                 json
// @Param                   page query string false "page"
// @Param 		            limit query string false "limit"
// @Param 		            search query string false "search"
// @Success                 200  {object}  models.BranchResponse
// @Failure                 400  {object}  models.Response
// @Failure                 404  {object}  models.Response
// @Failure                 500  {object}  models.Response
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
	branch, err := h.storage.Branch().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	handleResponse(c, "", 200, branch)

}

// UpdateBranch godoc
// @Router       /branch/{id} [PUT]
// @Summary      Update branch
// @Description  update branch
// @Tags         branch
// @Accept       json
// @Produce      json
//@Param         id path string true "branch_id"
// @Param        branch body models.UpdateBranch true "branch"
// @Success      200  {object}  models.Branch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBranch(c *gin.Context) {
	updatebranch := models.UpdateBranch{}
	uid := c.Param("id")
	if err := c.ShouldBindJSON(&updatebranch); err != nil {
		handleResponse(c, "error while decoding branch", 500, err.Error())
		return
	}
	updatebranch.ID = uid
	if _, err := h.storage.Branch().Update(context.Background(),updatebranch); err != nil {
		handleResponse(c, "error while updating branch", 500, err.Error())
		return
	}
	ids := models.PrimaryKey{ID: uid}
	res, err := h.storage.Branch().GetById(context.Background(),ids)
	if err != nil {
		handleResponse(c, "error while getting branch by ID", 500, err.Error())
		return
	}
	handleResponse(c, "branch updated successfully", 200, res)
}


// DeleteBranch godoc
// @Router       /branch/{id} [DELETE]
// @Summary      Delete branch
// @Description  delete branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param 		 id path string true "branch_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBranch(c *gin.Context) {
	uid := c.Param("id")
	branchid := models.PrimaryKey{ID: uid}
	if err := h.storage.Branch().Delete(context.Background(),branchid); err != nil {
		handleResponse(c, "error while deleting branch", 500, err)
		return
	}
	handleResponse(c, "", http.StatusOK, nil)
}
