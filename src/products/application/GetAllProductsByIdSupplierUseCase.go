package application

import(
	"github.com/lalo64/sgp/src/products/domain/ports"
	"github.com/lalo64/sgp/src/products/domain/entities"
)

type GetAllProductsByIdSupplierUseCase struct {
	ProductsRepository ports.IProducts
}

func NewGetAllProductsByIdSupplierUseCase(productsRepository ports.IProducts) *GetAllProductsByIdSupplierUseCase {
	return &GetAllProductsByIdSupplierUseCase{ProductsRepository: productsRepository}
}

func (s GetAllProductsByIdSupplierUseCase) Run(id int64) ([]entities.Products, error) {
	products, err := s.ProductsRepository.GetAllByIdSupplier(id)

	if err != nil {
		return []entities.Products{}, err
	}

	return products, nil
}