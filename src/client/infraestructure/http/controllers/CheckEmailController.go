package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/client/application"
	"github.com/lalo64/sgp/src/shared/responses"
)

type CheckEmailController struct {
	CheckEmailService *application.CheckEmailUseCase
}

func NewCheckEmailController(checkEmailService *application.CheckEmailUseCase) *CheckEmailController{
	return &CheckEmailController{CheckEmailService: checkEmailService}
}


func (ctr *CheckEmailController)Run(ctx *gin.Context){

	email := ctx.Param("email")

	status, err := ctr.CheckEmailService.Run(email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
            Message: "Error al verificar el correo",
            Data:    nil,
            Error:   err.Error(),
		})
		return 
	}

	log.Println(status)

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