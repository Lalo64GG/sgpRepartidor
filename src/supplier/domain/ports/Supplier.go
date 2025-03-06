package ports

import "github.com/lalo64/sgp/src/supplier/domain/entities"

type ISupplier interface{
	Create(supplier entities.Supplier) (entities.Supplier, error)
	GetAll(limit, page int64, orderBy, orderDir string ) ([]entities.Supplier, error)
	GetById(id int64) (entities.Supplier, error)
	GetByEmail(email string) (entities.Supplier, error)
	CheckEmail(email string) (bool, error)
}