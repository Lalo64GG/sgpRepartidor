package application

import (
	"github.com/lalo64/sgp/src/delivery/domain/ports"
	"github.com/lalo64/sgp/src/delivery/domain/entities"
)

type GetAllDriverIDUseCase struct {
	DriverRepository ports.IDelivery
}

func NewGetAllDriverIDUseCase(driverRepository ports.IDelivery)*GetAllDriverIDUseCase{
	return &GetAllDriverIDUseCase{DriverRepository: driverRepository}
}

func(uc *GetAllDriverIDUseCase) Run(id int64) ([]entities.Delivery, error){
	deliveries, err := uc.DriverRepository.GetAllDriverID(id)

	if err != nil {
		return nil, err
	}

	return deliveries, nil
}