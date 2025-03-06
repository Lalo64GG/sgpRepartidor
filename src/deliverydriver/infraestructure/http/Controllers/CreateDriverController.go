package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/sgp/src/deliverydriver/application"
	request "github.com/lalo64/sgp/src/deliverydriver/infraestructure/http/Request"
	"github.com/lalo64/sgp/src/shared/responses"
)

type CreateDriverController struct {
	DriverService *application.CreateDriverUseCase
	Validator *validator.Validate
}

func NewCreateDriverController(driverService *application.CreateDriverUseCase) *CreateDriverController {
	return &CreateDriverController{DriverService: driverService, Validator: validator.New()}
}

func (ctr *CreateDriverController) Run(ctx *gin.Context) {
	var req request.CreateDriverRequest
	
	if err := ctx.ShouldBindJSON(&req); err !=nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Llene todos los campos",
			Data: nil,
			Error: err.Error(),
		})

		return
	}

	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
            Success: false,
            Message: "Los datos ingresados no son válidos",
            Data: nil,
            Error: err.Error(),
        })
		return
	}

	driver, err := ctr.DriverService.Run(req.Name, req.Email, req.Password, req.FCM_TOKEN)

	if err != nil {

		if strings.Contains(err.Error(), "unique_client_email") {
			ctx.JSON(http.StatusConflict, responses.Response{
				Success: false,
				Message: "El email ya existe",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al crear el conductor",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Conductor creado correctamente",
		Data: driver,
		Error: nil,
	})
}