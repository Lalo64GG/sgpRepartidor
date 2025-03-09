package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/products/application"
	"github.com/lalo64/sgp/src/shared/responses"
)

type GetAllProductsByIdSupplierController struct {
	ProductService *application.GetAllProductsByIdSupplierUseCase
}

func NewGetAllProductsByIdSupplierUseCase(productService *application.GetAllProductsByIdSupplierUseCase) *GetAllProductsByIdSupplierController {
	return &GetAllProductsByIdSupplierController{ProductService: productService}
}

func (ctr *GetAllProductsByIdSupplierController)Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid id",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	products, err := ctr.ProductService.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error getting products",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Products retrieved",
		Data: products,
		Error: nil,
	})
}