package dto

type LoginRequest struct {
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password"`
}

type SignupResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Admin bool `json:"admin"`
}

type LoginResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Admin bool `json:"admin"`
}

type RefreshTokensRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokensResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
