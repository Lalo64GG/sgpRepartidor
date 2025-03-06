package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/supplier/infraestructure/http"
)


func AuthRoutes(router *gin.RouterGroup){
	createController := http.SetUpRegisterController()
	authController := http.AuthController()
	checkEmailSupplier := http.CheckEmailSupplierController()

	router.POST("/", createController.Run)
	router.POST("/auth", authController.Run)
	router.GET("/check-email", checkEmailSupplier.Run)
}


