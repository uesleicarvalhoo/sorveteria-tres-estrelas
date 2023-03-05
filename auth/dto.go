package auth

type SignInPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtToken struct {
	GrantType    string `json:"grant_type"`
	AcessToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expiration"`
}
