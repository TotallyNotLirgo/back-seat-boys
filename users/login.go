package users

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type LoginServices interface {
	GetUserByCredentials(email, password string) (*models.UserModel, error)
}

func Login(
	s LoginServices, request models.UserRequest,
) (response models.UserResponse, err error) {
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(request.Password)))
	user, err := s.GetUserByCredentials(request.Email, password)
	if err != nil {
		return response, errors.Join(models.ErrServerError, err)
	}
	if user == nil {
		return response, errors.Join(models.ErrUnauthorized, ErrUserNotFound)
	}
	response.UserId = user.UserId
	response.Email = user.Email
	response.Role = user.Role
	return response, nil
}
