package dto

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token"`
}
