package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/shared/responses"
	"github.com/lalo64/sgp/src/supplier/application"
)

type CheckEmailController struct {
	SupplierService *application.CheckEmailUseCase
}

func NewCheckEmailController(supplierService *application.CheckEmailUseCase) *CheckEmailController {
	return &CheckEmailController{SupplierService: supplierService}
}

func (ctr *CheckEmailController) Run(ctx *gin.Context){
	email := ctx.Param("email")

	status, err := ctr.SupplierService.Run(email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error al verificar el correo",
			Data:    nil,
			Error:   err.Error(),
		})
		return 
	}

	message := ""
	if status {
		message = "El correo ya se encuentra registrado"
	} else {
		message = "Correo válido"
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Error:   nil,
		Message: message,
		Data:    status,
	})
}