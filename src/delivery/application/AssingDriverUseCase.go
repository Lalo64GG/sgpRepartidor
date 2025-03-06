package application

import (
	"errors"
	"fmt"

	"github.com/lalo64/sgp/src/delivery/application/services"
	"github.com/lalo64/sgp/src/delivery/domain/entities"
	"github.com/lalo64/sgp/src/delivery/domain/ports"
	driverPorts "github.com/lalo64/sgp/src/deliverydriver/domain/ports"
)

type AssignDriverUseCase struct {
	DeliveryRepository  ports.IDelivery
	DriverRepository    driverPorts.IDriver  
	NotificationService services.Notification
}

// NewAssignDriverUseCase crea una nueva instancia del use case
func NewAssignDriverUseCase(deliveryRepo ports.IDelivery, driverRepo driverPorts.IDriver, notificationService services.Notification) *AssignDriverUseCase {
	return &AssignDriverUseCase{
		DeliveryRepository:  deliveryRepo,
		DriverRepository:    driverRepo,
		NotificationService: notificationService,
	}
}

// Ejecuta la lógica del use case para asignar un repartidor a un pedido
func (u *AssignDriverUseCase) Run(deliveryID int, driverID int) (entities.Delivery, error) {
	delivery, err := u.DeliveryRepository.GetById(deliveryID)
	if err != nil {
		return entities.Delivery{}, err
	}
	if delivery.DeliveryID == 0 {
		return entities.Delivery{}, errors.New("delivery not found")
	}

	if delivery.Status != "Pending" {
		return entities.Delivery{}, errors.New("cannot assign driver, the delivery is not in pending state")
	}

	updatedDelivery, err := u.DeliveryRepository.AssignDriver(deliveryID, driverID)
	if err != nil {
		return entities.Delivery{}, err
	}

	driverId64 := int64(driverID)

	driver, err := u.DriverRepository.GetById(driverId64)
	if err != nil {
		return entities.Delivery{}, fmt.Errorf("error retrieving driver: %v", err)
	}
	if driver.FCM_TOKEN == "" {
		return entities.Delivery{}, errors.New("driver FCM token is missing")
	}

	fcmToken := string(driver.FCM_TOKEN)

	
	err = u.NotificationService.SendPushNotification(fcmToken, "Nuevo Pedido Asignado", "Se te ha asignado un nuevo producto para entregar.")
	if err != nil {
		return entities.Delivery{}, fmt.Errorf("error sending push notification: %v", err)
	}


	return updatedDelivery, nil
}
