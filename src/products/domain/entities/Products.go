package entities

import()

type Products struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Supplier_Id int `json:"supplier_id"`
}