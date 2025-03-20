package users

import (
	"time"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
)

type ExampleDatabase struct {
	email     string
	password  string
	role      string
	lastLogin int
}

func (d ExampleDatabase) GetUserByCredentials(
	email, password string,
) *models.AuthModel {
	response := models.AuthModel{
		Role:      "admin",
		LastLogin: time.Now().Unix(),
		UserId: 12,
	}
	return &response
}
