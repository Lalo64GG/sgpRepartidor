package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/client/infraestructure/http"
)

func ClientRoutes(router *gin.RouterGroup){
	createClientController := http.SetUpRegisterController()
	checkEmailController := http.CheckEmailController()
	authController := http.AuthController()


	router.POST("/", createClientController.Run)
	router.GET("/check-email/:email", checkEmailController.Run)
	router.POST("/auth", authController.Run)
}