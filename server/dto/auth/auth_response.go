package authdto

type AuthRegisterResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token" form:"token"`
}
