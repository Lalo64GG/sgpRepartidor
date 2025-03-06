package application

import (
	"github.com/lalo64/sgp/src/supplier/domain/entities"
	"github.com/lalo64/sgp/src/supplier/domain/ports"
)

type GetAllSupplierUseCase struct {
	SupplierRepository ports.ISupplier
}


func NewGetAllSupplierUseCase(supplierRepository ports.ISupplier) *GetAllSupplierUseCase {
	return &GetAllSupplierUseCase{SupplierRepository: supplierRepository}
}

func (s *GetAllSupplierUseCase) Run(limit, page int64, orderBy, orderDir string) ([]entities.Supplier, error){
	suppliers, err := s.SupplierRepository.GetAll(limit, page, orderBy, orderDir)

	if err != nil {
		return []entities.Supplier{}, err
	}

	return suppliers, nil
}