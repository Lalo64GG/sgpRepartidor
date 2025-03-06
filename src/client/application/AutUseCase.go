package application

import (
	"github.com/lalo64/sgp/src/client/domain/entities"
	"github.com/lalo64/sgp/src/client/domain/ports"
)

type AuthUseCase struct {
	UserRepository ports.IClientRepository
}

func NewAuthUseCase(userRepository ports.IClientRepository) *AuthUseCase {
	return &AuthUseCase{UserRepository: userRepository}  
}


func (s AuthUseCase) Run(email string)(entities.Client, error){
	user, err := s.UserRepository.GetByEmail(email)

	if err != nil {
        return entities.Client{}, err
    }

	return user, nil
}