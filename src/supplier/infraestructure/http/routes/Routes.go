package routes

import (
	"github.com/gin-gonic/gin"

)

func SupplierRoutes(router *gin.RouterGroup) {
	AuthRoutes(router)
	ProtectedRoutes(router)
}
