package entities

import "time"

type Delivery struct {
	DeliveryID   int      `json:"delivery_id"`
	DriverID     int     `json:"driver_id"`   
	ClientID     int      `json:"client_id"`     
	SupplierID   int      `json:"supplier_id"`   
	DeliveryDate time.Time  `json:"delivery_date"` 
	Status       string     `json:"status"`        
}
