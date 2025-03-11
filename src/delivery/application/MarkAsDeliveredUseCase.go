package application

import (
	"github.com/lalo64/sgp/src/delivery/domain/entities"
	"github.com/lalo64/sgp/src/delivery/domain/ports"
)

type MarkAsDeliveredUseCase struct {
	DeliveryRepository ports.IDelivery
}

func NewMarkAsDeliveredUseCase(deliveryRepository ports.IDelivery) *MarkAsDeliveredUseCase{
	return &MarkAsDeliveredUseCase{DeliveryRepository: deliveryRepository}
}

func (d *MarkAsDeliveredUseCase) Run(id int)(entities.Delivery,error){
	delivery, err := d.DeliveryRepository.MarkAsDelivered(id)
	if err != nil{
		return entities.Delivery{}, err
	}

	return delivery, nil
}