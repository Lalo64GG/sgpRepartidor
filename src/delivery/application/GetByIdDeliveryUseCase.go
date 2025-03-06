package application

import (
	"github.com/lalo64/sgp/src/delivery/domain/entities"
	"github.com/lalo64/sgp/src/delivery/domain/ports"
)

type GetByIdDeliveryUseCase struct {
	DeliveryRepository ports.IDelivery
}

func NewGetByIdDeliveryUseCase(deliveryRepository ports.IDelivery) *GetByIdDeliveryUseCase {
    return &GetByIdDeliveryUseCase{DeliveryRepository: deliveryRepository}
}

func (d *GetByIdDeliveryUseCase) Run(id int) (entities.Delivery, error) {
	delivery, err := d.DeliveryRepository.GetById(id)

	if err != nil {
		return entities.Delivery{}, err
	}

	return delivery, nil
}