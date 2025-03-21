package users

import (
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
)

type TestDatabase struct {
	email     string
	password  string
	role      string
	lastLogin int64
	userId    int64
}

func (d TestDatabase) GetUserByCredentials(
	email, password string,
) *models.UserResponse {
	if d.email != email || d.password != password {
		return nil
	}
	response := models.UserResponse{
		UserId:    d.userId,
		Role:      d.role,
		Email:     d.email,
		LastLogin: d.lastLogin,
	}
	return &response
}
func (d TestDatabase) GetUserByEmail(email string) *models.UserResponse {
	if d.email != email {
		return nil
	}
	response := models.UserResponse{
		UserId:    d.userId,
		Role:      d.role,
		Email:     d.email,
		LastLogin: d.lastLogin,
	}
	return &response
}

func (d *TestDatabase) CreateUser(
	user models.LoginRequest, role string,
) *models.UserResponse {
	d.email = user.Email
	d.password = user.Password
	d.role = role
	response := models.UserResponse{
		UserId:    d.userId,
		Role:      d.role,
		Email:     d.email,
		LastLogin: d.lastLogin,
	}
	return &response
}
