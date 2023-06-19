package authdto

type AuthRegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
