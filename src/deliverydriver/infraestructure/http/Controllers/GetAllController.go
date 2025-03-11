package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/deliverydriver/application"
	"github.com/lalo64/sgp/src/shared/responses"

	"net/http"
)

type GetAllController struct {
	GetAllUseCase *application.GetAllUseCase
}

func NewGetAllController(getAllUseCase *application.GetAllUseCase) *GetAllController {
	return &GetAllController{GetAllUseCase: getAllUseCase}
}

func (ctr *GetAllController) Run(ctx *gin.Context) {
	drivers, err := ctr.GetAllUseCase.Run()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Internal Server Error",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Drivers Found",
		Error: nil,
		Data: drivers,
	})
}