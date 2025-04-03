package users

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/uuid"
)

var (
	ErrUserConflict = errors.New("user with this email already exists")
)

type RegisterServices interface {
	GetUserByEmail(email string) (*models.UserModel, error)
	InsertUser(email, pass string, role models.Role) (int, error)
	SendEmail(email, token string) error
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
		return response, errors.Join(models.ErrConflict, ErrUserConflict)
	}
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(request.Password)))
	id, err := s.InsertUser(request.Email, password, models.RoleNew)
	if err != nil {
		return response, errors.Join(models.ErrServerError, err)
	}
	err = s.SendEmail(request.Email, uuid.New().String())
	if err != nil {
		return response, errors.Join(models.ErrBadRequest, err)
	}
	response.Id = id
	response.Email = request.Email
	response.Role = models.RoleNew
	return response, nil
}
