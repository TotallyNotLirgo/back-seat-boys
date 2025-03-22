package users

import (
	"errors"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
	"gorm.io/gorm"
)

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
		UserId: user.Model.ID,
		Role:   user.Role,
		Email:  user.Email,
	}
	logger.Info("Returning user %v", user.Model.ID)
	return &response
}
