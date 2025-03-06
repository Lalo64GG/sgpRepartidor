package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/sgp/src/shared/responses"
	"github.com/lalo64/sgp/src/delivery/application"
	"github.com/lalo64/sgp/src/delivery/infraestructure/http/request"
)

type AssignDriverController struct {
	AssignDriverUseCase *application.AssignDriverUseCase
	Validator           *validator.Validate
}

func NewAssignDriverController(assignDriverUseCase *application.AssignDriverUseCase) *AssignDriverController {
	return &AssignDriverController{
		AssignDriverUseCase: assignDriverUseCase,
		Validator:           validator.New(),
	}
}

func (ctr *AssignDriverController) Run(ctx *gin.Context) {
	var req request.AssignDriverRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Validar la solicitud
	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Llamar al UseCase para asignar el repartidor
	delivery, err := ctr.AssignDriverUseCase.Run(req.DeliveryID, req.DriverID)
	if err != nil {
		// Si ocurre un error, manejarlo adecuadamente
		if strings.Contains(err.Error(), "delivery_not_found") {
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false,
				Message: "Pedido no encontrado",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}

		if strings.Contains(err.Error(), "driver_not_found") {
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false,
				Message: "Repartidor no encontrado",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Repartidor asignado correctamente",
		Data:    delivery,
		Error:   nil,
	})
}
