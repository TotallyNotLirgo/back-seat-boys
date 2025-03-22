package users

import (
	"errors"
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
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
	logger := log.GetLogger("GetUserByCredentials")
	logger.Info("Fetching user %v", email)
	var user User
	result := d.Connection.First(
		&user,
		"email = ? AND password = ?",
		email,
		password,
	)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logger.Warning("User not found")
		return nil
	}
	response := models.UserResponse{
		UserId:    user.Model.ID,
		Role:      user.Role,
		Email:     user.Email,
	}
	logger.Info("Returning user %v", user.Model.ID)
	return &response
}
func (d Database) GetUserByEmail(email string) *models.UserResponse {
	logger := log.GetLogger("GetUserByEmail")
	logger.Info("Fetching user %v", email)
	var user User
	result := d.Connection.First(&user, "email = ?", email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logger.Warning("User not found")
		return nil
	}
	response := models.UserResponse{
		UserId:    user.Model.ID,
		Role:      user.Role,
		Email:     user.Email,
	}
	logger.Info("Returning user %v", user.Model.ID)
	return &response
}

func (d *Database) CreateUser(
	user models.LoginRequest, role string,
) *models.UserResponse {
	logger := log.GetLogger("CreateUser")
	logger.Info("Creating user %v", user.Email)
	model := User{
		Email: user.Email,
		Password: user.Password,
		Role: role,
	}
	result := d.Connection.Create(&model)

	if result.Error != nil {
		fmt.Printf("Creation failed: %v", result.Error.Error())
		logger.Error(result.Error.Error())
		return nil
	}
	response := models.UserResponse{
		UserId:    model.Model.ID,
		Role:      model.Role,
		Email:     model.Email,
	}
	logger.Info("Returning user %v", model.Model.ID)
	return &response
}
