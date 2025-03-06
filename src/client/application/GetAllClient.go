package application

import (
	"github.com/lalo64/sgp/src/client/domain/entities"
	"github.com/lalo64/sgp/src/client/domain/ports"
)

type GetClientUseCase struct {
	ClientRepository ports.IClientRepository
}

func NewGetAllClientUseCase(clientRepository ports.IClientRepository) *GetClientUseCase {
	return &GetClientUseCase{ClientRepository: clientRepository}
}


func (c *GetClientUseCase) Run(limit, page int64, orderBy, orderDir string ) ([]entities.Client, error) {
	clients, err := c.ClientRepository.GetAll(limit, page, orderBy, orderDir)

	if err != nil {
		return []entities.Client{}, err
	}

	return clients, nil
}