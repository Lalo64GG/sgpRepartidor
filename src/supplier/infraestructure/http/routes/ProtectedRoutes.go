package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/supplier/infraestructure/http"
)

func ProtectedRoutes(router *gin.RouterGroup) {
	getAllSupplier := http.GetAllSupplierController()
	getByIdSupplier := http.GetByIdSupplierController()


	router.GET("/", getAllSupplier.Run)
	router.GET("/:id", getByIdSupplier.Run)
}