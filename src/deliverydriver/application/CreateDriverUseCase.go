package application

import (
	"github.com/lalo64/sgp/src/deliverydriver/domain/entities"
	"github.com/lalo64/sgp/src/deliverydriver/domain/ports"
	encrypt "github.com/lalo64/sgp/src/shared/Encrypt"
)

type CreateDriverUseCase struct {
	DriverRepository ports.IDriver
	EncryptService encrypt.EncryptService
}

func NewCreateDriverUseCase(driverRepository ports.IDriver, encryptService encrypt.EncryptService) *CreateDriverUseCase {
    return &CreateDriverUseCase{DriverRepository: driverRepository, EncryptService: encryptService}
}


func (c *CreateDriverUseCase) Run(Name, Email, Password, FCM_TOKEN string) (entities.Driver, error) {
	hashedPassword, err := c.EncryptService.Encrypt([]byte(Password))

	if err != nil {
		return entities.Driver{}, err
	}
	
	driver := entities.Driver{
		Name: Name,
		Email: Email,
		Password: hashedPassword,
		FCM_TOKEN: FCM_TOKEN,
	}


	newDriver, err := c.DriverRepository.Create(driver)

	if err != nil {
		return entities.Driver{}, err
	}

	return newDriver, nil
}