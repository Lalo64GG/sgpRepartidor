package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/delivery/application"
	"github.com/lalo64/sgp/src/shared/responses"
)

type GetAllSupplierIDController struct {
	GetAllSupplierIDUseCase *application.GetAllSupplierIDUseCase
}

func NewGetAllSupplierIDController(getAllSupplierIdUseCase *application.GetAllSupplierIDUseCase) *GetAllSupplierIDController {
	return &GetAllSupplierIDController{GetAllSupplierIDUseCase: getAllSupplierIdUseCase}
}

func (ctr *GetAllSupplierIDController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err  := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid ID",
			Error: err.Error(),
			Data: nil,
		})
		return 
	}
	deliverys, err := ctr.GetAllSupplierIDUseCase.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Internal Server Error",
			Error: err.Error(),
			Data:nil,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, responses.Response{
		Success: true,
		Message: "Deliveries Found",
		Error: nil,
		Data: deliverys,
	})

}