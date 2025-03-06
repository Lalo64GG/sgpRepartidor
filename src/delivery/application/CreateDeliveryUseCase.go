package application

import (
	"errors"
	"time"

	"github.com/lalo64/sgp/src/delivery/domain/entities"
	"github.com/lalo64/sgp/src/delivery/domain/ports"
)

type CreateDeliveryUseCase struct {
	DeliveryRepository ports.IDelivery
}

func NewCreateDeliveryUseCase(deliveryRepository ports.IDelivery) *CreateDeliveryUseCase {
	return &CreateDeliveryUseCase{DeliveryRepository: deliveryRepository}
}

func (c *CreateDeliveryUseCase) Run(driverID, clientID, supplierID int, deliveryDate time.Time, status string) (entities.Delivery, error) {
	if clientID == 0 {
		return entities.Delivery{}, errors.New("el cliente es obligatorio")
	}
	if supplierID == 0 {
		return entities.Delivery{}, errors.New("el proveedor es obligatorio")
	}
	if deliveryDate.Before(time.Now()) {
		return entities.Delivery{}, errors.New("la fecha de entrega no puede ser en el pasado")
	}

	if status == "" {
		status = "Pending"
	}

	delivery := entities.Delivery{
		DriverID:     driverID,
		ClientID:     clientID,
		SupplierID:   supplierID,
		DeliveryDate: deliveryDate,
		Status:       status,
	}

	newDelivery, err := c.DeliveryRepository.Create(delivery)
	if err != nil {
		return entities.Delivery{}, err
	}

	return newDelivery, nil
}
