package application

import (
	"github.com/lalo64/sgp/src/deliverydriver/domain/entities"
	"github.com/lalo64/sgp/src/deliverydriver/domain/ports"
)

type AuthUseCase struct {
	DriverRepository ports.IDriver
}

func NewAuthUseCase(driverRepository ports.IDriver) *AuthUseCase {
    return &AuthUseCase{DriverRepository: driverRepository}
}


func (a *AuthUseCase) Run(email string) (entities.Driver, error) {
	driver, err := a.DriverRepository.GetByEmail(email)

	if err != nil {
		return entities.Driver{}, err
	}

	return driver, nil
}