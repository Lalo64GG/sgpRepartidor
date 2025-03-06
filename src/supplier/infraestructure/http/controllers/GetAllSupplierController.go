package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/shared/responses"
	"github.com/lalo64/sgp/src/supplier/application"
)

type GetAllSupplierController struct {
	SupplierService *application.GetAllSupplierUseCase
}


func NewGetAllSupplierUseCase(supplierService *application.GetAllSupplierUseCase) *GetAllSupplierController {
	return &GetAllSupplierController{SupplierService: supplierService}
}


func (ctr *GetAllSupplierController) Run(ctx *gin.Context){
	limit := parseQueryParam(ctx, "limit", 5)
	page := parseQueryParam(ctx, "page", 1)
	orderBy := ctx.Query("orderBy")
	orderDir := ctx.Query("orderDir")

	if orderDir != "asc" && orderDir != "desc" {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid order direction",
		})
	}

	suppliers, err := ctr.SupplierService.Run(limit, page, orderBy, orderDir)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
            Success: false,
            Message: "Error getting suppliers",
			Data: nil,
			Error: err.Error(),
        })
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Suppliers retrieved",
		Data: suppliers,
		Error: nil,
	})
}

func parseQueryParam(ctx *gin.Context, key string, defaultValue int64) int64 {
	queryValue := ctx.Query(key)
	if queryValue == "" {
		return defaultValue
	}
	value, err := strconv.ParseInt(queryValue, 10, 64)
	if err != nil || value <= 0 {
		return defaultValue
	}
	return value
}
