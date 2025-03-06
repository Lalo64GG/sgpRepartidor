package ports

import "github.com/lalo64/sgp/src/delivery/domain/entities"

type IDelivery interface {
	Create(entities.Delivery) (entities.Delivery, error)
	UpdateStatus(status string) (bool, error)
	AssignDriver(deliveryID, driverID int) (entities.Delivery, error)
	GetById(id int) (entities.Delivery, error)
}