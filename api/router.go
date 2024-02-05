package api

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main.go/api/handler"
	"main.go/storage"
	_"main.go/api/docs"

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

	r.POST("/category", h.CreateCategory)
	r.GET("/category/:id", h.GetByIdCategory)
	r.GET("/categories", h.GetListCategory)
	r.PUT("/category/:id", h.UpdateCategory)
	r.DELETE("/category/:id", h.DeleteCategory)

	r.POST("/product", h.CreateProduct)
	r.GET("/product/:id", h.GetByIdProduct)
	r.GET("/products", h.GetListProduct)
	r.PUT("/product/:id", h.UpdateProduct)
	r.DELETE("/product/:id", h.DeleteProduct)

	r.POST("/storage", h.CreateStorage)
	r.GET("/storage/:id", h.GetByIdStorage)
	r.GET("/storages", h.GetListStorage)
	r.PUT("/storage/:id", h.UpdateStorage)
	r.DELETE("/storage/:id", h.DeleteStorage)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
