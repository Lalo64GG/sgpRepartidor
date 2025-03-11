package application

import (
	"github.com/lalo64/sgp/src/delivery/domain/entities"
	"github.com/lalo64/sgp/src/delivery/domain/ports"
)

type GetAllSupplierIDUseCase struct {
	DeliveryRepository ports.IDelivery
}

func NewGetAllSupplierIDUseCase(deliveryRepository ports.IDelivery) *GetAllSupplierIDUseCase{
	return &GetAllSupplierIDUseCase{DeliveryRepository: deliveryRepository}
}

func(d *GetAllSupplierIDUseCase) Run(id int64) ([]entities.Delivery, error){
	deliverys, err := d.DeliveryRepository.GetAllSupplierID(id)

	if err != nil{
		return []entities.Delivery{}, err
	}


	return deliverys, nil

}