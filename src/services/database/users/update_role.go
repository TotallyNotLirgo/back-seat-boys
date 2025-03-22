package users

import (
	"errors"
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
	"gorm.io/gorm"
)

func (d *Database) UpdateRole(
	id uint, role string,
) *models.UserResponse {
	logger := log.GetLogger("UpdateRole")
	logger.Info("Updating user %v", id)
	var user User
	result := d.Connection.First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logger.Warning("User not found")
		return nil
	}
	if role != "" {
		user.Role = role
	}
	d.Connection.Save(&user)
	response := models.UserResponse{
		UserId: user.Model.ID,
		Role:   user.Role,
		Email:  user.Email,
	}
	logger.Info("Returning user %v", user.Model.ID)
	return &response
}
