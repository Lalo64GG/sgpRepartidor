package application

import (
	"github.com/lalo64/sgp/src/client/domain/entities"
	"github.com/lalo64/sgp/src/client/domain/ports"
	encrypt "github.com/lalo64/sgp/src/shared/Encrypt"
)

type CreateClientUseCase struct {
	ClientRepository ports.IClientRepository
	EncryptService encrypt.EncryptService
}

func NewCreateClientUseCase(clientRepository ports.IClientRepository, encryptService encrypt.EncryptService) *CreateClientUseCase {
    return &CreateClientUseCase{ClientRepository: clientRepository, EncryptService: encryptService}
}

func (c *CreateClientUseCase) Run(Name, Email, Password, Address string ) (entities.Client, error) {

	encryptedPass, err := c.EncryptService.Encrypt([]byte(Password))

	if err != nil {
		return entities.Client{}, err    
	}

	client := entities.Client{
		Name: Name,
		Email: Email,
		Password: encryptedPass,
		Address: Address,
	}


	newClient, err := c.ClientRepository.Create(client)

	if err != nil {
		return entities.Client{}, err
	}

	return newClient, nil
}