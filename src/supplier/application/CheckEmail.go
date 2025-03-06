package application

import "github.com/lalo64/sgp/src/supplier/domain/ports"

type CheckEmailUseCase struct {
	SupplierRepository ports.ISupplier
}

func NewCheckEmailUseCase(supplierRepository ports.ISupplier) *CheckEmailUseCase {
	return &CheckEmailUseCase{SupplierRepository: supplierRepository}
}

func (s *CheckEmailUseCase) Run(email string) (bool, error) {
	status, err := s.SupplierRepository.CheckEmail(email)

	if err != nil {
		return false, err
	}

	return status, nil
}