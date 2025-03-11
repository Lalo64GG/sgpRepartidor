package http

import (
	"log"

	"github.com/lalo64/sgp/src/delivery/application"
	"github.com/lalo64/sgp/src/delivery/application/services"
	"github.com/lalo64/sgp/src/delivery/domain/ports"
	driverPorts	"github.com/lalo64/sgp/src/deliverydriver/domain/ports"
	"github.com/lalo64/sgp/src/delivery/infraestructure/adapters"
	driverAdapters "github.com/lalo64/sgp/src/deliverydriver/infraestructure/adapters"
	controllers "github.com/lalo64/sgp/src/delivery/infraestructure/http/Controllers"
	"github.com/lalo64/sgp/src/delivery/infraestructure/http/Controllers/helpers"
)

var (
	deliveryRepository ports.IDelivery
	notificationService services.Notification
	driverRepository driverPorts.IDriver
)

func init() {
	var err error

	deliveryRepository, err = adapters.NewDeliveryRepositoryMysql()
	if err != nil {
        log.Fatalf("Error initializing delivery repository: %v", err)
    }

	notificationService, err = helpers.NewPushNotificationService()
	if err != nil {
		log.Fatalf("Error initializing notification service: %v", err)
	}

	driverRepository, err = driverAdapters.NewDriverRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing driver repository: %v", err)
	}



}

func SetUpCreateDeliveryController() *controllers.CreateDeliveryController {
	createService := application.NewCreateDeliveryUseCase(deliveryRepository)
	return controllers.NewCreateDeliveryController(createService)
}

func AssingDriverController() *controllers.AssignDriverController {
    assingController := application.NewAssignDriverUseCase(deliveryRepository, driverRepository, notificationService)
    return controllers.NewAssignDriverController(assingController)
}

func GetAllSupplierID() *controllers.GetAllSupplierIDController{
	getAllSupplierId := application.NewGetAllSupplierIDUseCase(deliveryRepository)
	return controllers.NewGetAllSupplierIDController(getAllSupplierId)
}

func GetAllDriverID() *controllers.GetAllDriverIDController{
	getAllDriverId := application.NewGetAllDriverIDUseCase(deliveryRepository)
	return controllers.NewGetAllDriverIDController(getAllDriverId)
}