package application

import (
	"github.com/lalo64/sgp/src/supplier/domain/entities"
	"github.com/lalo64/sgp/src/supplier/domain/ports"
)

type AuthUseCase struct {
	SupplierRepository ports.ISupplier
}

func NewAuthUseCase(supplierRepository ports.ISupplier) *AuthUseCase {
	return &AuthUseCase{SupplierRepository: supplierRepository}  
}


func (s AuthUseCase) Run(email string)(entities.Supplier, error){
	user, err := s.SupplierRepository.GetByEmail(email)

	if err != nil {
        return entities.Supplier{}, err
    }

	return user, nil
}