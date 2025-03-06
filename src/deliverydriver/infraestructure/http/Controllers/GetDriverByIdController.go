package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/deliverydriver/application"
)

type GetDriverByIdController struct{
	DriverService *application.GetDriverByIdUseCase
}

func NewGetDriverByIdController(driverService *application.GetDriverByIdUseCase) *GetDriverByIdController {
	return &GetDriverByIdController{
		DriverService: driverService,
	}
}


func (ctr *GetDriverByIdController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return
    }

    driver, err := ctr.DriverService.Run(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, driver)
}