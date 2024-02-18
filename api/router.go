package api

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "main.go/api/docs"
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
	r.POST("/startsale", h.CreateSale)
	r.PUT("/endsale/:id", h.EndSale)

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

	r.POST("/basket", h.CreateBasket)
	r.GET("/basket/:id", h.GetByIdBasket)
	r.GET("/baskets", h.GetListBasket)
	r.PUT("/basket/:id", h.UpdateBasket)
	r.DELETE("/basket/:id", h.DeleteBasket)

	r.POST("/storagetransaction", h.CreateStorageTransaction)
	r.GET("/storagetransaction/:id", h.GetByIdStorageTransaction)
	r.GET("/storagetransactions", h.GetStorageTransactionList)
	r.PUT("/storagetransaction/:id", h.UpdateStorageTransaction)
	r.DELETE("/storagetransaction/:id", h.DeleteStorageTransaction)

	r.POST("/staff", h.CreateStaff)
	r.GET("/staff/:id", h.GetByIdStaff)
	r.GET("/staffs", h.GetListStaff)
	r.PUT("/staff/:id", h.UpdateStaff)
	r.PATCH("/staff/:id", h.UpdateStaffPassword)
	r.DELETE("/staff/:id", h.DeleteStaff)

	r.POST("/transaction", h.CreateTransaction)
	r.GET("/transaction/:id", h.GetTransaction)
	r.GET("/transactions", h.GetTransactionList)
	r.PUT("/transaction/:id", h.UpdateTransaction)
	r.DELETE("/transaction/:id", h.DeleteTransaction)

	r.POST("/staff-tariff", h.CreateStaffTariff)
	r.GET("/staff-tariff/:id", h.GetStaffTariff)
	r.GET("/staff-tariffs", h.GetStaffTariffList)
	r.PUT("/staff-tariff/:id", h.UpdateStaffTariff)
	r.DELETE("/staff-tariff/:id", h.DeleteStaffTariff)

	r.POST("/staff_tarif")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
