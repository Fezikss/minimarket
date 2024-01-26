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
	r.GET("/branch/:id", h.GetByIdBranch)
	r.GET("/branchs", h.GetListBranch)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/sale", h.CreateSale)
	r.GET("/sale/:id", h.GetByIdSale)
	r.GET("/sales", h.GetListSale)
	r.PUT("/sale/:id", h.UpdateSale)
	r.DELETE("/sale/:id", h.DeleteSale)

	return r
}
