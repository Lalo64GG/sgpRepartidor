package application

import (
	"github.com/lalo64/sgp/src/deliverydriver/domain/entities"
	"github.com/lalo64/sgp/src/deliverydriver/domain/ports"
)

type GetAllUseCase struct {
	DriverRepository ports.IDriver
}

func NewGetAllUseCase(driverRepository ports.IDriver) *GetAllUseCase {
	return &GetAllUseCase{DriverRepository: driverRepository}
}

func(d *GetAllUseCase) Run()([]entities.Driver, error){
	drivers, err := d.DriverRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return drivers, nil
}