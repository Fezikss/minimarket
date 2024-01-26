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

	r.POST("/sale", h.)
	r.GET("/sale", h.)
	r.GET("/sales", h.)
	r.PUT("/sale", h.)
	r.DELETE("/sale", h.DeleteBranch)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
