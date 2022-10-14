package dto

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserPayload struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
