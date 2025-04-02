package models

type Role string

const (
	New   Role = "new"
	User  Role = "user"
	Admin Role = "admin"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
    Id    int    `json:"id"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}
