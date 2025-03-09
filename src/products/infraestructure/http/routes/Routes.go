package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/products/infraestructure/http"
)

func ProductsRoutes(router *gin.RouterGroup) {
	CreateProductController := http.SetUpCreateProductController()
	GetAllByIdSupplierController := http.SetUpGetAllProductsByIdSupplierController()


	router.POST("/", CreateProductController.Run)
	router.GET("/:id", GetAllByIdSupplierController.Run)
}