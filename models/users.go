package models

type Role string

const (
	RoleNew   Role = "new"
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel struct {
	Id    int
	Email string
	Role  Role
}

type UserResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}
