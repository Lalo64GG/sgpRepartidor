package encrypt

type EncryptService interface {
	Encrypt(pwd []byte) (string, error)
	Compare(hashedPwd string, plainPwd []byte) error
}