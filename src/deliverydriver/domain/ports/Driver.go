package ports

import "github.com/lalo64/sgp/src/deliverydriver/domain/entities"

type IDriver interface {
	Create(entities.Driver) (entities.Driver, error)
	GetByEmail(email string) (entities.Driver, error)
	GetById(id int64) (entities.Driver, error)
	GetAll()([]entities.Driver, error)
}