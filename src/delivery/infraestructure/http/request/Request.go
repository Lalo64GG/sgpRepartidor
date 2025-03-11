package request

type CreateDeliveryRequest struct {
	ClientID     int    `json:"client_id" validate:"required"`
	DeliveryDate string `json:"delivery_date" validate:"required"`
	Status       string `json:"status" validate:"required,oneof=Pending InTransit Delivered Cancelled"`
	SupplierID   int    `json:"supplier_id" validate:"required"`
	ProductID	int    `json:"product_id" validate:"required"`
	DriverID     int   `json:"driver_id,omitempty"`
}

type AssignDriverRequest struct {
	DeliveryID int `json:"delivery_id" validate:"required"`
	DriverID   int `json:"driver_id" validate:"required"`
}