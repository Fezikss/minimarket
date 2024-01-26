package api

import (
	"github.com/gin-gonic/gin"

	"main.go/api/handler"
	"main.go/storage"
)

func New(store storage.IStorage) *gin.Engine {
	h := handler.New(store)
	r := gin.New()
	r.POST("/branch", h.CreateBranch)
	r.GET("/branch", h.GetByIdBranch)
	r.GET("/branchs", h.GetListBranch)
	r.PUT("/branch", h.UpdateBranch)
	r.DELETE("/branch", h.DeleteBranch)

	r.POST("/sale", h.CreateSale)
	r.GET("/sale", h.GetByIdSale)
	r.GET("/sales", h.GetListSale)
	r.PUT("/sale", h.UpdateSale)
	r.DELETE("/sale", h.DeleteSale)

	return r
}
