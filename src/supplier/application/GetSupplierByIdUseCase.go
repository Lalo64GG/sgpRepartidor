package application

import (
	"github.com/lalo64/sgp/src/supplier/domain/entities"
	"github.com/lalo64/sgp/src/supplier/domain/ports"
)

type GetSupplierByIdUseCase struct {
	SupplierRepository ports.ISupplier
}

func NewGetSupplierByIdUseCase(supplierRepository ports.ISupplier) *GetSupplierByIdUseCase{
	return &GetSupplierByIdUseCase{SupplierRepository: supplierRepository}
}

func (s *GetSupplierByIdUseCase) Run(id int64) (entities.Supplier, error){
	supplier, err := s.SupplierRepository.GetById(id)

    if err != nil {
        return entities.Supplier{}, err
    }

    return supplier, nil
}