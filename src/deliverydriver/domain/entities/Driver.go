package entities


type Driver struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	FCM_TOKEN string   `json:"fcm_token"`
}