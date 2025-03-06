package http

import (
	"log"

	"github.com/lalo64/sgp/src/client/application"
	"github.com/lalo64/sgp/src/client/domain/ports"
	"github.com/lalo64/sgp/src/client/infraestructure/adapters"
	"github.com/lalo64/sgp/src/client/infraestructure/http/controllers"
	encrypt "github.com/lalo64/sgp/src/shared/Encrypt"
)

var (
	clientRepository  ports.IClientRepository
	encryptService  encrypt.EncryptService
)

func init(){
	var err error

	clientRepository, err = adapters.NewClientRepositoryMysql()
	if err != nil {
		log.Fatalf("Error creating client repository: %v", err)
	}

	encryptService, err = encrypt.NewEncryptHelper()
	if err != nil {
        log.Fatalf("Error creating encrypt service: %v", err)
    }
}


func SetUpRegisterController() *controllers.CreateClientController{
	createUserUseCase := application.NewCreateClientUseCase(clientRepository, encryptService)
	return controllers.NewCreateClientController(createUserUseCase)
}

func CheckEmailController() *controllers.CheckEmailController{
	checkEmailUseCase := application.NewCheckEmailUseCase(clientRepository)
	return controllers.NewCheckEmailController(checkEmailUseCase)
}

func AuthController() *controllers.AuthController{
	authUseCase := application.NewAuthUseCase(clientRepository)
    return controllers.NewAuthController(authUseCase)
}
