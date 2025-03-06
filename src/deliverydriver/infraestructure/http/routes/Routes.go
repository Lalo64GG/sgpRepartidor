package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/deliverydriver/infraestructure/http"
)

func DriverRoutes(router *gin.RouterGroup) {
	CreateDriverController := http.SetUpCreateDriverController()
	authController := http.AuthController()

	router.POST("/", CreateDriverController.Run)
	router.POST("/auth", authController.Run)
}
