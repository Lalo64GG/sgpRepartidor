package application

import "github.com/lalo64/sgp/src/delivery/domain/ports"

type UpdateStatusDeliveryUseCase struct {
	DeliveryRepository ports.IDelivery
}

func NewUpdateStatusDeliveryUseCase(deliveryRepository ports.IDelivery) *UpdateStatusDeliveryUseCase {
    return &UpdateStatusDeliveryUseCase{DeliveryRepository: deliveryRepository}
}

func (d *UpdateStatusDeliveryUseCase) Run(status string) (bool, error) {
	updated, err := d.DeliveryRepository.UpdateStatus(status)

	if err != nil {
		return false, err
	}

	return updated, nil
}