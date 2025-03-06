package application

import "github.com/lalo64/sgp/src/client/domain/ports"

type DeleteClientUseCase struct {
	ClientRepository ports.IClientRepository
}


func NewDeleteClientUseCase(clientRepository ports.IClientRepository) *DeleteClientUseCase {
	return &DeleteClientUseCase{ClientRepository: clientRepository}
} 

func (c *DeleteClientUseCase) Run (id int) (bool, error) {
	_, err := c.ClientRepository.Delete(id)

	if err != nil {
		return false, err
	}

	return true, nil
}