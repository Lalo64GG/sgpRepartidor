package application

import (
	"github.com/lalo64/sgp/src/client/domain/ports"
)

type CheckEmailUseCase struct {
	UserRepository ports.IClientRepository
}


func NewCheckEmailUseCase(userRepository ports.IClientRepository) *CheckEmailUseCase {
	return &CheckEmailUseCase{UserRepository: userRepository}
}

func (s CheckEmailUseCase) Run(email string) (bool, error) {
	status, err := s.UserRepository.CheckEmail(email)

	if err != nil {
		return false, err
	}

	return status, nil
}