package users

import (
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
)

func (d *Database) CreateUser(
	user models.LoginRequest, role string,
) *models.UserResponse {
	logger := log.GetLogger("CreateUser")
	logger.Info("Creating user %v", user.Email)
	model := User{
		Email:    user.Email,
		Password: user.Password,
		Role:     role,
	}
	result := d.Connection.Create(&model)

	if result.Error != nil {
		fmt.Printf("Creation failed: %v", result.Error.Error())
		logger.Error(result.Error.Error())
		return nil
	}
	response := models.UserResponse{
		UserId: model.Model.ID,
		Role:   model.Role,
		Email:  model.Email,
	}
	logger.Info("Returning user %v", model.Model.ID)
	return &response
}
