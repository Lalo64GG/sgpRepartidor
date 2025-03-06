package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/sgp/src/client/application"
	"github.com/lalo64/sgp/src/client/infraestructure/http/request"
	"github.com/lalo64/sgp/src/shared/responses"
)

type CreateClientController struct {
	ClientService *application.CreateClientUseCase
	Validator     *validator.Validate
}

func NewCreateClientController(clientService *application.CreateClientUseCase) *CreateClientController {
	// Inicializamos el validador aquí
	return &CreateClientController{
		ClientService: clientService,
		Validator:     validator.New(), // Inicializando el validador
	}
}

func (ctr *CreateClientController) Run(ctx *gin.Context) {
	var req request.CreateClientRequest

	// Validamos si hay errores en la solicitud
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Llene todos los campos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Validamos los datos usando el validador
	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Los datos enviados no son válidos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Procesamos la creación del cliente
	client, err := ctr.ClientService.Run(req.Name, req.Email, req.Password, req.Address)

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

		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error al crear el cliente",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Si todo está bien, devolvemos la respuesta exitosa
	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Cliente creado correctamente",
		Data:    client,
		Error:   nil,
	})
}
