package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/shared/responses"
	"github.com/lalo64/sgp/src/supplier/application"
)

type GetSupplierByIdController struct {
	SupplierService *application.GetSupplierByIdUseCase
}

func NewGetSupplierByIdController(supplierService *application.GetSupplierByIdUseCase) *GetSupplierByIdController{
	return &GetSupplierByIdController{SupplierService: supplierService}
}

func (ctr *GetSupplierByIdController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid id",
			Data: nil,
			Error: err.Error(),
		})
	}

	supplier, err := ctr.SupplierService.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
            Success: false,
            Message: "Error getting supplier",
            Data: nil,
            Error: err.Error(),
        })
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Supplier retrieved",
		Data: supplier,
		Error: nil,
	})


	
}