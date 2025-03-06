package entities

type Supplier struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	ContactInfo string `json:"contact_info"`
}