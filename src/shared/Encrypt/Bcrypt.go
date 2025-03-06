package encrypt

import "golang.org/x/crypto/bcrypt"

type EncryptHelper struct{}

func NewEncryptHelper() (EncryptService, error) {
	return &EncryptHelper{}, nil
}

func (h *EncryptHelper) Encrypt(pwd []byte) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (h *EncryptHelper) Compare(hashPwd string, plaindPwd[]byte) error{
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), plaindPwd)
	if err != nil {
		return err
	}
	return nil
}