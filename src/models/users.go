package models

const (
	Admin = "admin"
	User  = "user"
	New   = "new"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserId    int64 `json:"userId"`
	Email    string `json:"email"`
	Role      string `json:"role"`
}
