package application

import (
	encrypt "github.com/lalo64/sgp/src/shared/Encrypt"
	"github.com/lalo64/sgp/src/supplier/domain/entities"
	"github.com/lalo64/sgp/src/supplier/domain/ports"
)

type CreateSupplierUseCase struct {
	SupplierRepository ports.ISupplier
	Encryptor encrypt.EncryptService
}

func NewCreateSupplierUseCase(supplierRepository ports.ISupplier, encryptService encrypt.EncryptService) *CreateSupplierUseCase {
    return &CreateSupplierUseCase{SupplierRepository: supplierRepository, Encryptor: encryptService}
}

func (s *CreateSupplierUseCase) Run(Name, Email, Password, Address, ContactInfo string ) (entities.Supplier, error) {

	hashPassword, error :=s.Encryptor.Encrypt([]byte(Password))

	if error != nil {
		return entities.Supplier{}, error
	}

	supplier := entities.Supplier{
		Name: Name,
		Email: Email,
		Password: hashPassword,
		Address: Address,
		ContactInfo: ContactInfo,
	}

	newSupplier, err := s.SupplierRepository.Create(supplier)

	if err != nil {
		return entities.Supplier{}, err
	}

	return newSupplier, nil

}