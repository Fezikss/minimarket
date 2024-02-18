package handler

import (
	"context"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

// CreateBasket   godoc
// @Router               /basket [POST]
// @Summary              Creates a new basket
// @Description          Creates a new basket
// @Tags                 basket
// @Accept               json
// @Produce              json
// @Param                basket body models.CreateBasket true "basket"
// @Success              201  {object}  models.Response
// @Failure              400  {object}  models.Response
// @Failure              404  {object}  models.Response
// @Failure              500  {object}  models.Response
func (h Handler) CreateBasket(c *gin.Context) {

	createbasket := models.CreateBasket{}
	if err := c.ShouldBindJSON(&createbasket); err != nil {
		handleResponse(c, "error while decoding basket from json", 400, err.Error())
		return
	}

	storage, err := h.storage.Storage().GetList(context.Background(), models.GetListRequest{
		Page:   1,
		Limit:  10,
		Search: createbasket.ProductID,
	})
	if err != nil {
		handleResponse(c, "error in line 39", 500, err.Error())
	}

	
	price, err := h.storage.Product().GetById(context.Background(), models.PrimaryKey{ID: createbasket.ProductID})

	totalsum := float64(createbasket.Quantity) * price.Price

	IsTrue := false
	baskets, err := h.storage.Basket().GetList(context.Background(), models.GetListRequest{Page: 1, Limit: 10, Search: createbasket.SaleID})
	for _, v := range baskets.Basket {

		if v.ProductID == createbasket.ProductID {
			for _, v1 := range storage.Storages {
				if v1.Count < createbasket.Quantity + v.Quantity{

					handleResponse(c, "not enough product in storage", 302, "")
					return

				}
			}
			IsTrue = true
			_, err := h.storage.Basket().Update(context.Background(), models.UpdateBasket{
				ID:        v.ID,
				ProductID: v.ProductID,
				Price:     v.Price + totalsum,
				Quantity:  v.Quantity + createbasket.Quantity,
			})
			if err != nil {
				fmt.Println(err.Error())
				handleResponse(c, "error creating basket", 500, "")
			}
		}
	}
	// basket create -> prodID
	// basket create -> basketGetlist.prodID = basket.prodID -> update qlamz

	
	if !IsTrue {
		for _, v := range storage.Storages {
			if v.Count < createbasket.Quantity {
	
				handleResponse(c, "not enough product in storage", 300, "")
				return
	
			}
		}
		createbasket.Price = totalsum
		_, err := h.storage.Basket().Create(context.Background(), createbasket)
		if err != nil {
			handleResponse(c, "error is while getting by id basket", 500, err.Error())
			return
		}
	}

	handleResponse(c, "", 200, "success")

}

// GetByIdBasket  godoc
// @Router            /basket{id} [GET]
// @Summary           Get basket by id
// @Description       get basket by id
// @Tags              basket
// @Accept            json
// @Produce           json
// @Param             id path string true "basket"
// @Success           200  {object}  models.Basket
// @Failure           400  {object}  models.Response
// @Failure           404  {object}  models.Response
// @Failure           500  {object}  models.Response
func (h Handler) GetByIdBasket(c *gin.Context) {
	var err error
	uid := c.Param("id")
	basket, err := h.storage.Basket().GetById(context.Background(), models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by id", 500, err.Error())
		return
	}
	handleResponse(c, "", 200, basket)

}

// GetBasketList            godoc
// @Router                 /baskets [GET]
// @Summary                 Get baskets list
// @Description             get baskets list
// @Tags                    basket
// @Accept                  json
// @Produce                 json
// @Param                   page query string false "page"
// @Param 		            limit query string false "limit"
// @Param 		            search query string false "search"
// @Success                 200  {object}  models.BasketResponse
// @Failure                 400  {object}  models.Response
// @Failure                 404  {object}  models.Response
// @Failure                 500  {object}  models.Response
func (h Handler) GetListBasket(c *gin.Context) {
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
	basket, err := h.storage.Basket().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	handleResponse(c, "", 200, basket)

}

// UpdateBasket godoc
// @Router       /basket/{id} [PUT]
// @Summary      Update basket
// @Description  update basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param         id path string true "basket_id"
// @Param        basket body models.UpdateBasket true "basket"
// @Success      200  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBasket(c *gin.Context) {
	update := models.UpdateBasket{}
	uid := c.Param("id")
	if err := c.ShouldBindJSON(&update); err != nil {
		handleResponse(c, "error while decoding basket", 500, err.Error())
		return
	}
	update.ID = uid
	if _, err := h.storage.Basket().Update(context.Background(), update); err != nil {
		handleResponse(c, "error while updating basket", 500, err.Error())
		return
	}
	ids := models.PrimaryKey{ID: uid}
	res, err := h.storage.Basket().GetById(context.Background(), ids)
	if err != nil {
		handleResponse(c, "error while getting basket by ID", 500, err.Error())
		return
	}
	handleResponse(c, "basket updated successfully", 200, res)
}

// DeleteBasket godoc
// @Router       /basket/{id} [DELETE]
// @Summary      Delete basket
// @Description  delete basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param 		 id path string true "basket_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBasket(c *gin.Context) {
	uid := c.Param("id")
	basketid := models.PrimaryKey{ID: uid}
	if err := h.storage.Branch().Delete(context.Background(), basketid); err != nil {
		handleResponse(c, "error while deleting basket", 500, err)
		return
	}
	handleResponse(c, "", http.StatusOK, nil)
}
