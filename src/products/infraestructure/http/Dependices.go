package http

import (
	"log"

	"github.com/lalo64/sgp/src/products/application"
	"github.com/lalo64/sgp/src/products/domain/ports"
	"github.com/lalo64/sgp/src/products/infraestructure/adapters"
	"github.com/lalo64/sgp/src/products/infraestructure/http/controllers"
)

var productService ports.IProducts

func init() {
	var err error

	productService, err = adapters.NewProductRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing product repository: %v", err)
	}
}

func SetUpCreateProductController() *controllers.CreateProductController {
	createService := application.NewCreateProductsUseCase(productService)
	return controllers.NewCreateProductController(createService)
}

func SetUpGetAllProductsByIdSupplierController() *controllers.GetAllProductsByIdSupplierController {
	getAllProductsByIdSupplierService := application.NewGetAllProductsByIdSupplierUseCase(productService)
	return controllers.NewGetAllProductsByIdSupplierUseCase(getAllProductsByIdSupplierService)
}