package ports

import "github.com/lalo64/sgp/src/products/domain/entities"

type IProducts interface {
	Create(entities.Products) (entities.Products, error)
	GetAllByIdSupplier(id int64) ([]entities.Products, error)
}