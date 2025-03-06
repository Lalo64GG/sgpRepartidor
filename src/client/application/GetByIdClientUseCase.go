package application

import (
	"github.com/lalo64/sgp/src/client/domain/entities"
	"github.com/lalo64/sgp/src/client/domain/ports"
)

type GetByIdClientUseCase struct {
	ClientRepository ports.IClientRepository
}
 
func NewGetByIdClientUseCase(clientRepository ports.IClientRepository) *GetByIdClientUseCase{
	return &GetByIdClientUseCase{ClientRepository: clientRepository}
}


func (c *GetByIdClientUseCase) Run(id int) (entities.Client, error){
	client, err := c.ClientRepository.GetById(id)

	if err != nil {
		return entities.Client{}, err
	}

	return client, nil
}