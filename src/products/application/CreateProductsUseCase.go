package application

import (
	"github.com/lalo64/sgp/src/products/domain/entities"
	"github.com/lalo64/sgp/src/products/domain/ports"
)

type CreateProductUseCase struct{
	ProductRepository ports.IProducts
}

func NewCreateProductsUseCase(productRepository ports.IProducts) *CreateProductUseCase{
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (p *CreateProductUseCase) Run(Name string, Price float64, Supplier_Id int) (entities.Products, error){

	product := entities.Products{
		Name: Name,
		Price: Price,
		Supplier_Id: Supplier_Id,
	}

	product, err := p.ProductRepository.Create(product)

	if err != nil {
		return entities.Products{}, err
	}

	return product, nil
}