package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/sgp/src/delivery/application"
	"github.com/lalo64/sgp/src/delivery/infraestructure/http/request"
	"github.com/lalo64/sgp/src/shared/responses"
)

type CreateDeliveryController struct {
	CreateDeliveryUseCase *application.CreateDeliveryUseCase
	Validator             *validator.Validate
}

func NewCreateDeliveryController(createDeliveryUseCase *application.CreateDeliveryUseCase) *CreateDeliveryController {
	return &CreateDeliveryController{
		CreateDeliveryUseCase: createDeliveryUseCase,
		Validator:             validator.New(),
	}
}

func (ctr *CreateDeliveryController) Run(ctx *gin.Context) {
	var req request.CreateDeliveryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "LLenar todos los campos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Validación adicional (si es necesario)
	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Convertir fecha de entrega desde el request
	deliveryDate, err := time.Parse("2006-01-02", req.DeliveryDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Fecha de entrega no válida",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Ejecutar el caso de uso
	delivery, err := ctr.CreateDeliveryUseCase.Run(req.DriverID, req.ClientID, req.SupplierID, req.ProductID, deliveryDate, req.Status)
	if err != nil {
		if strings.Contains(err.Error(), "el cliente es obligatorio") ||
			strings.Contains(err.Error(), "el proveedor es obligatorio") ||
			strings.Contains(err.Error(), "la fecha de entrega no puede ser en el pasado") {
			ctx.JSON(http.StatusBadRequest, responses.Response{
				Success: false,
				Message: "Datos incorrectos",
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

	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Pedido creado exitosamente",
		Data:    delivery,
		Error:   nil,
	})
}
