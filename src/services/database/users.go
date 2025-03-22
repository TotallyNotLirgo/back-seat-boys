package users

import (
	"errors"
	"fmt"
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string
	Password  string
	Role      string
}

type Database struct {
	Connection *gorm.DB
}

func (d Database) GetUserByCredentials(
	email, password string,
) *models.UserResponse {
	var user User
	result := d.Connection.First(
		&user,
		"email = ? AND password = ?",
		email,
		password,
	)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	response := models.UserResponse{
		UserId:    user.Model.ID,
		Role:      user.Role,
		Email:     user.Email,
	}
	return &response
}
func (d Database) GetUserByEmail(email string) *models.UserResponse {
	var user User
	result := d.Connection.First(&user, "email = ?", email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	response := models.UserResponse{
		UserId:    user.Model.ID,
		Role:      user.Role,
		Email:     user.Email,
	}
	return &response
}

func (d *Database) CreateUser(
	user models.LoginRequest, role string,
) *models.UserResponse {
	model := User{
		Email: user.Email,
		Password: user.Password,
		Role: role,
	}
	result := d.Connection.Create(&model)

	if result.Error != nil {
		fmt.Printf("Creation failed: %v", result.Error.Error())
		return nil
	}
	response := models.UserResponse{
		UserId:    model.Model.ID,
		Role:      model.Role,
		Email:     model.Email,
	}
	return &response
}
