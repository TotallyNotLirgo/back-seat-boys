package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthModel struct {
	Role      string
	LastLogin int64
	UserId    int64
}
