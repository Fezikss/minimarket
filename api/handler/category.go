package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

// CreateCategory godoc
// @Router /category [POST]
// @Summary      Creates a new category
// @Description  create a new category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        category body models.CreateCategory false "category"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateCategory(c *gin.Context) {
	createcategory := models.CreateCategory{}
	if err := c.ShouldBindJSON(&createcategory); err != nil {
		handleResponse(c, "error while decoding category from json", 400, err.Error())
		return
	}
	id, err := h.storage.Category().Create(context.Background(), createcategory)
	if err != nil {
		handleResponse(c, "error is while creating category", 500, err.Error())
		return
	}
	res, err := h.storage.Category().GetById(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {

		handleResponse(c, "error while getting by id after creating category", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, res)

}

// GetCategory godoc
// @Router       /category/{id} [GET]
// @Summary      Gets category
// @Description  get category by ID
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        id path string true "category"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetByIdCategory(c *gin.Context) {
	var err error
	uid := c.Param("id")
	category, err := h.storage.Category().GetById(context.Background(), models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by id", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, category)

}

// GetCategoryList godoc
// @Router       /categories [GET]
// @Summary      Get category list
// @Description  get category list
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.CategoryResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetListCategory(c *gin.Context) {
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
	category, err := h.storage.Category().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	handleResponse(c, "", 200, category)

}

// UpdateCategory godoc
// @Router       /category/{id} [PUT]
// @Summary      Update category
// @Description  update category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param 		 id path string true "category_id"
// @Param        category body models.UpdateCategory true "category"
// @Success      200  {object}  models.Category
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateCategory(c *gin.Context) {
	updatecategory := models.UpdateCategory{}
	uid := c.Param("id")
	if err := c.ShouldBindJSON(&updatecategory); err != nil {
		handleResponse(c, "error while decoding category", 500, err.Error())
		return
	}
	updatecategory.ID = uid
	if _, err := h.storage.Category().Update(context.Background(), updatecategory); err != nil {
		handleResponse(c, "error while updating branch", 500, err.Error())
		return
	}
	ids := models.PrimaryKey{ID: uid}
	res, err := h.storage.Category().GetById(context.Background(), ids)
	if err != nil {
		handleResponse(c, "error while getting category by ID", 500, err.Error())
		return
	}
	handleResponse(c, "category updated successfully", 200, res)
}

// DeleteCategory godoc
// @Router       /category/{id} [DELETE]
// @Summary      Delete category
// @Description  delete category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param 		 id path string true "category_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteCategory(c *gin.Context) {
	uid := c.Param("id")
	categoryid := models.PrimaryKey{ID: uid}
	if err := h.storage.Category().Delete(context.Background(), categoryid); err != nil {
		handleResponse(c, "error while deleting category", 500, err)
		return
	}
	handleResponse(c, "", http.StatusOK, nil)
}
