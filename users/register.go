package users

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
)

type RegisterServices interface {
	GetUserByEmail(email string) (*models.UserModel, error)
	InsertUser(email, pass string, role models.Role) (int, error)
}

func Register(
	s RegisterServices, request models.UserRequest,
) (response models.UserResponse, err error) {
	err = IsPasswordValid(request.Password)
	if err != nil {
		return response, errors.Join(models.ErrBadRequest, err)
	}
	err = IsEmailValid(request.Email)
	if err != nil {
		return response, errors.Join(models.ErrBadRequest, err)
	}
	conflict, err := s.GetUserByEmail(request.Email)
	if err != nil {
		return response, errors.Join(models.ErrServerError, err)
	}
	if conflict != nil {
		return response, errors.Join(
			models.ErrConflict,
			errors.New("user with this email already exists"),
		)
	}
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(request.Password)))
	id, err := s.InsertUser(request.Email, password, models.RoleNew)
	if err != nil {
		return response, errors.Join(models.ErrServerError, err)
	}
	response.Id = id
	response.Email = request.Email
	response.Role = models.RoleNew
	return response, nil
}
