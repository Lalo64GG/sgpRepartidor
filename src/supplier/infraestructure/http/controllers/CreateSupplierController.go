package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/sgp/src/shared/responses"
	"github.com/lalo64/sgp/src/supplier/application"
	"github.com/lalo64/sgp/src/supplier/infraestructure/http/request"
)

type CreateSupplierController struct {
	SupplierService *application.CreateSupplierUseCase
	Validator *validator.Validate
}

func NewCreateSupplierController(supplierService *application.CreateSupplierUseCase) *CreateSupplierController {
	return &CreateSupplierController{SupplierService: supplierService, Validator: validator.New()}
}

func (ctr *CreateSupplierController) Run(ctx *gin.Context) {
	var req request.CreateSupplierRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
			Error: err.Error(),
		})
        return
    }

	if err := ctr.Validator.Struct(req); err != nil{
		ctx.JSON(http.StatusBadRequest, responses.Response{
            Success: false,
            Message: err.Error(),
            Data: nil,
            Error: err.Error(),
        })
        return
    }

	if err := ctr.Validator.Struct(req); err != nil{
		ctx.JSON(http.StatusBadRequest, responses.Response{
            Success: false,
            Message: err.Error(),
            Data: nil,
            Error: err.Error(),
        })
        return
    }

	supplier, err := ctr.SupplierService.Run(req.Name, req.Email, req.Password, req.Address, req.ContactInfo)

	if err != nil {

		if strings.Contains(err.Error(), "unique_client_email") {
			ctx.JSON(http.StatusConflict, responses.Response{
				Success: false,
				Message: "El email ya existe",
				Data: nil,
				Error: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, responses.Response{
            Success: false,
            Message: err.Error(),
            Data: nil,
            Error: err.Error(),
        })
        return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Proveedor creado correctamente",
		Data: supplier,
        Error: nil,
	})

	

}

