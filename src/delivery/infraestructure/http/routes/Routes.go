package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/delivery/infraestructure/http"
)

func DeliveryRoutes(router *gin.RouterGroup) {
	createDeliveryController := http.SetUpCreateDeliveryController()
	assingController := http.AssingDriverController()

	router.POST("/", createDeliveryController.Run)
	router.PATCH("/assing", assingController.Run)
}