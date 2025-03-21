package users

import (
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"time"
)

type ExampleDatabase struct {
	userId   int64
	email    string
	password string
	role     string
}

func (d ExampleDatabase) GetUserByCredentials(
	email, password string,
) *models.UserResponse {
	if d.email != email || d.password != password {
		return nil
	}
	response := models.UserResponse{
		UserId:    d.userId,
		Role:      d.role,
		Email:     d.email,
		LastLogin: time.Now().Unix(),
	}
	return &response
}
func (d ExampleDatabase) GetUserByEmail(email string) *models.UserResponse {
	if d.email != email {
		return nil
	}
	response := models.UserResponse{
		UserId:    d.userId,
		Role:      d.role,
		Email:     d.email,
		LastLogin: time.Now().Unix(),
	}
	return &response
}

func (d *ExampleDatabase) CreateUser(
	user models.LoginRequest, role string,
) *models.UserResponse {
	d.email = user.Email
	d.password = user.Password
	d.role = role
	response := models.UserResponse{
		UserId:    d.userId,
		Role:      d.role,
		Email:     d.email,
		LastLogin: time.Now().Unix(),
	}
	return &response
}
