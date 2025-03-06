package application

import (
	"github.com/lalo64/sgp/src/deliverydriver/domain/entities"
	"github.com/lalo64/sgp/src/deliverydriver/domain/ports"
)

type GetDriverByIdUseCase struct{
	driverRepo ports.IDriver
}

func NewGetDriverByIdUseCase(driverRepo ports.IDriver) *GetDriverByIdUseCase {
	return &GetDriverByIdUseCase{
		driverRepo: driverRepo,
	}
}


func (g *GetDriverByIdUseCase) Run(id int64) (entities.Driver, error) {
	driver, err := g.driverRepo.GetById(id)
	if err != nil {
		return entities.Driver{}, err
	}
	return driver, nil
}