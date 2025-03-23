package users

import (
	"errors"

	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
	"gorm.io/gorm"
)

func (d Database) DeleteUser(id uint) {
	logger := log.GetLogger("DeleteUser")
	logger.Info("Deleting user %v", id)
	result := d.Connection.Delete(&User{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logger.Warning("User not found")
		return
	}
	return
}
