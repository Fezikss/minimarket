package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

// CreateProduct godoc
// @Router       /product [POST]
// @Summary      Creates a new product
// @Description  create a new product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        product body models.CreateProduct false "product"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateProduct(c *gin.Context) {
	createproduct := models.CreateProduct{}
	if err := c.ShouldBindJSON(&createproduct); err != nil {
		handleResponse(c, "error while decoding product from json", 400, err.Error())
		return
	}
	id, err := h.storage.Product().Create(context.Background(), createproduct)
	if err != nil {
		handleResponse(c, "error is while getting by id branch", 500, err.Error())
		return
	}
	res, err := h.storage.Product().GetById(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {

		handleResponse(c, "error while getting by id after creating product", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, res)

}

// GetProduct godoc
// @Router       /product/{id} [GET]
// @Summary      Gets product
// @Description  get product by ID
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id path string true "product"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetByIdProduct(c *gin.Context) {
	var err error
	uid := c.Param("id")
	product, err := h.storage.Product().GetById(context.Background(), models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by id", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, product)

}

// GetProductList godoc
// @Router       /products [GET]
// @Summary      Get product list
// @Description  get product list
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Param        barcode query int false "barcode"
// @Success      200  {object}  models.ProductResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetListProduct(c *gin.Context) {
	var (
		page, limit int
		search      string
		barcode     int
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
	barcodestr := c.Query("barcode")
	barcode,err=strconv.Atoi(barcodestr)
	if err!=nil{
		handleResponse(c,"error while converting barcode",400,err.Error())
	}
	product, err := h.storage.Product().GetList(context.Background(), models.GetListRequestProduct{
		Page:    page,
		Limit:   limit,
		Search:  search,
		Barcode: barcode,
	})
	handleResponse(c, "", 200, product)

}

// UpdateProduct godoc
// @Router       /product/{id} [PUT]
// @Summary      Update product
// @Description  update product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path string true "product_id"
// @Param        product body models.UpdateProduct true "product"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateProduct(c *gin.Context) {
	updateproduct := models.UpdateProduct{}
	uid := c.Param("id")
	if err := c.ShouldBindJSON(&updateproduct); err != nil {
		handleResponse(c, "error while decoding product", 500, err.Error())
		return
	}
	updateproduct.ID = uid
	fmt.Println(updateproduct)
	if _, err := h.storage.Product().Update(context.Background(), updateproduct); err != nil {
		handleResponse(c, "error while updating product", 500, err.Error())
		return
	}
	ids := models.PrimaryKey{ID: uid}
	res, err := h.storage.Product().GetById(context.Background(), ids)
	if err != nil {
		handleResponse(c, "error while getting product by ID", 500, err.Error())
		return
	}
	handleResponse(c, "product updated successfully", 200, res)
}

// DeleteProduct godoc
// @Router       /product/{id} [DELETE]
// @Summary      Delete product
// @Description  delete product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path string true "product_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteProduct(c *gin.Context) {
	uid := c.Param("id")
	productid := models.PrimaryKey{ID: uid}
	if err := h.storage.Product().Delete(context.Background(), productid); err != nil {
		handleResponse(c, "error while deleting product", 500, err)
		return
	}
	handleResponse(c, "", http.StatusOK, nil)
}
