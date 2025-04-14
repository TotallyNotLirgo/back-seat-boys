package models

type Role string

func (r Role) GreaterEqual(other Role) bool {
	return r.getRoleLevel() >= other.getRoleLevel()
}

func (r Role) getRoleLevel() int {
	switch r {
	case RoleNew:
		return 1
	case RoleUser:
		return 2
	case RoleAdmin:
		return 3
	}
	return -1
}

const (
	RoleNew   Role = "new"
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type EmailRequest struct {
	Email string `json:"email"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel struct {
	UserId uint
	Email  string
	Role   Role
}

type UserResponse struct {
	UserId uint   `json:"id"`
	Email  string `json:"email"`
	Role   Role   `json:"role"`
}
