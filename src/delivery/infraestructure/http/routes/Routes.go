package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/delivery/infraestructure/http"
)

func DeliveryRoutes(router *gin.RouterGroup) {
	createDeliveryController := http.SetUpCreateDeliveryController()
	assingController := http.AssingDriverController()
	getAllSupplierId := http.GetAllSupplierID()
	getAllDriverId := http.GetAllDriverID()
	markAsDelivered := http.MarkAsDelivered()

	router.POST("/", createDeliveryController.Run)
	router.PATCH("/assing", assingController.Run)
	router.PATCH("/delivered/:id", markAsDelivered.Run)
	router.GET("/:id", getAllSupplierId.Run)
	router.GET("/driver/:id", getAllDriverId.Run)
}