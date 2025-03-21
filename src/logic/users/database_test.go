package users

import (
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
)

type TestDatabase struct {
	email       string
	password    string
	role        string
	userId      uint
	createFails bool
}

func (d TestDatabase) GetUserByCredentials(
	email, password string,
) *models.UserResponse {
	if d.email != email || d.password != password {
		return nil
	}
	response := models.UserResponse{
		UserId: d.userId,
		Role:   d.role,
		Email:  d.email,
	}
	return &response
}
func (d TestDatabase) GetUserByEmail(email string) *models.UserResponse {
	if d.email != email {
		return nil
	}
	response := models.UserResponse{
		UserId: d.userId,
		Role:   d.role,
		Email:  d.email,
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
		UserId: d.userId,
		Role:   d.role,
		Email:  d.email,
	}
	if d.createFails {
		return nil
	}
	return &response
}
