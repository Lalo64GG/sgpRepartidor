package request

type CreateProductRequest struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Supplier_Id int `json:"supplier_id"`
}