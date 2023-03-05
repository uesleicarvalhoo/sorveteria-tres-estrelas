package auth

type SignInPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtToken struct {
	GrantType string `json:"grant_type"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiration"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token"`
}
