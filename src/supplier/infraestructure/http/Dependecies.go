package http

import (
	"log"

	encrypt "github.com/lalo64/sgp/src/shared/Encrypt"
	"github.com/lalo64/sgp/src/supplier/application"
	"github.com/lalo64/sgp/src/supplier/domain/ports"
	"github.com/lalo64/sgp/src/supplier/infraestructure/adapters"
	"github.com/lalo64/sgp/src/supplier/infraestructure/http/controllers"
)

var (
	supplierRepository ports.ISupplier
	encryptService encrypt.EncryptService 
)

func init() {
	var err error

	supplierRepository, err = adapters.NewSupplierRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing supplier repository: %v", err)
	}

	// Usamos el helper, que implementa la interfaz EncryptService
	encryptService, err = encrypt.NewEncryptHelper()
	if err != nil {
		log.Fatalf("Error initializing encrypt service: %v", err)
	}
}

func SetUpRegisterController() *controllers.CreateSupplierController {
	createService := application.NewCreateSupplierUseCase(supplierRepository, encryptService)
	return controllers.NewCreateSupplierController(createService)
}

func GetAllSupplierController() *controllers.GetAllSupplierController {
	getAllService := application.NewGetAllSupplierUseCase(supplierRepository)
	return controllers.NewGetAllSupplierUseCase(getAllService)
}

func GetByIdSupplierController() *controllers.GetSupplierByIdController {
	getByIdService := application.NewGetSupplierByIdUseCase(supplierRepository)
	return controllers.NewGetSupplierByIdController(getByIdService)
}

func CheckEmailSupplierController() *controllers.CheckEmailController {
	checkEmailService := application.NewCheckEmailUseCase(supplierRepository)
	return controllers.NewCheckEmailController(checkEmailService)
}

func AuthController() *controllers.AuthController {
	authService := application.NewAuthUseCase(supplierRepository)
	return controllers.NewAuthController(authService)
}
