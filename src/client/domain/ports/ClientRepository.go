package ports

import (
	"github.com/lalo64/sgp/src/client/domain/entities"
)

type IClientRepository interface {
	Create(client entities.Client) (entities.Client, error)
	GetAll(limit, page int64, orderBy, orderDir string ) ([]entities.Client, error)
	GetById(id int) (entities.Client, error)
	Delete(id int) (bool, error)
	GetByEmail(email string) (entities.Client, error)
	CheckEmail(email string) (bool, error)
}