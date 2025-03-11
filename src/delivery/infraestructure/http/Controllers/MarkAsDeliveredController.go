package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/sgp/src/delivery/application"
	"github.com/lalo64/sgp/src/shared/responses"
)

type MarkAsDeliveredController struct{
	DeliveryRepository *application.MarkAsDeliveredUseCase
	Validator *validator.Validate
}

func NewMarkAsDeliveredController(deliveryRepository *application.MarkAsDeliveredUseCase) *MarkAsDeliveredController{
	return &MarkAsDeliveredController{
		DeliveryRepository: deliveryRepository,
		Validator: validator.New(),
	}
}

func (ctr *MarkAsDeliveredController) Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	delivered, err := ctr.DeliveryRepository.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{

		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Pedido entregado correctamente",
		Data:    delivered,
		Error:   nil,
	})
}