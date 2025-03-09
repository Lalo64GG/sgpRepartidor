package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/sgp/src/products/application"
	"github.com/lalo64/sgp/src/products/infraestructure/http/request"
	"github.com/lalo64/sgp/src/shared/responses"
)

type CreateProductController struct {
	ProductService *application.CreateProductUseCase
	Validator *validator.Validate
}

func NewCreateProductController(productService *application.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		ProductService: productService,
		Validator: validator.New(),
	}
}

func (ctr *CreateProductController) Run(ctx *gin.Context){
	var req request.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Llenar todos los campos",
			Data:    nil,
			Error:   err.Error(),
		})
		return 
	}

	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Los datos enviados no son válidos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	product, err := ctr.ProductService.Run(req.Name, req.Price, req.Supplier_Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al crear el producto",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Producto creado",
		Data:    product,
		Error:   nil,
	})

}