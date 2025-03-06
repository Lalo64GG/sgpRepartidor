package request

type CreateDriverRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FCM_TOKEN string `json:"fcm_token" binding:"required"`
}


type AuthRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
