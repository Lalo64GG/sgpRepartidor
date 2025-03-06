package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	encrypt "github.com/lalo64/sgp/src/shared/Encrypt"
	"github.com/lalo64/sgp/src/shared/middlewares"
	"github.com/lalo64/sgp/src/shared/responses"
	"github.com/lalo64/sgp/src/supplier/application"
	"github.com/lalo64/sgp/src/supplier/infraestructure/http/request"
)

type AuthController struct {
	AuthService *application.AuthUseCase
	EncryptHelper *encrypt.EncryptHelper
}

func NewAuthController(authService *application.AuthUseCase) *AuthController{
	return &AuthController{AuthService: authService}
}


func(ctr *AuthController) Run(ctx *gin.Context){
	var authRequest request.AuthRequest

	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Llene todos los campos",
			Error: err.Error(),
			Data: nil,
		})
        return
    }

	supplier, err := ctr.AuthService.Run(authRequest.Email)

	if err != nil {
		switch err.Error(){
		case "sql: no rows in result set":
		    ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false,
                Message: "El correo no existe",
                Error: err.Error(),
                Data: nil,
			})
		default:
			ctx.JSON(http.StatusInternalServerError, responses.Response{
				Success: false,
				Message: "Error al iniciar sesión",
				Error: err.Error(),
				Data: nil,
			})
		}
		return 
    }

	fmt.Println(supplier.Password)
	if err := ctr.EncryptHelper.Compare(supplier.Password, []byte(authRequest.Password)); err != nil{
		ctx.JSON(http.StatusUnauthorized, responses.Response{
			Success: false,
			Message: "Contraseña incorrecta",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	token, err := middlewares.GenerateJWT(int64(supplier.ID), supplier.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
            Success: false,
            Message: "Error al generar token",
            Error: err.Error(),
            Data: nil,
        })
        return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Sesión iniciada correctamente",
		Error: "",
        Data: map[string]interface{}{
			"token": token,
			"Id": supplier.ID,
			"Name": supplier.Name,
            "Email": supplier.Email,
			"Address": supplier.Address,
		},
	})
}