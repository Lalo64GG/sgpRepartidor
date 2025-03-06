package http

import (
	"log"

	"github.com/lalo64/sgp/src/deliverydriver/application"
	"github.com/lalo64/sgp/src/deliverydriver/domain/ports"
	"github.com/lalo64/sgp/src/deliverydriver/infraestructure/adapters"
	controllers "github.com/lalo64/sgp/src/deliverydriver/infraestructure/http/Controllers"
	encrypt "github.com/lalo64/sgp/src/shared/Encrypt"
)

var (
	driverRepository ports.IDriver
	encryptService encrypt.EncryptService
)

func init(){
	var err error
    driverRepository, err = adapters.NewDriverRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing driver repository: %v", err)
	}

	encryptService, err = encrypt.NewEncryptHelper()
	if err != nil {
		log.Fatalf("Error initializing encrypt service: %v", err)
	}
}

func SetUpCreateDriverController() *controllers.CreateDriverController{
	createService := application.NewCreateDriverUseCase(driverRepository, encryptService)
    return controllers.NewCreateDriverController(createService)
}

func AuthController() *controllers.AuthController{
	authService := application.NewAuthUseCase(driverRepository)
	return controllers.NewAuthController(authService)
}