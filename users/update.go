package users

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/uuid"
	slogctx "github.com/veqryn/slog-context"
)

type UpdateServices interface {
	GetUserById(id uint) (*models.UserModel, error)
	GetUserByEmail(email string) (*models.UserModel, error)
	UpdateUser(id uint, email, password string, role models.Role) error
	SendEmail(id uint, token, bucket string) error
	SetLogger(logger slog.Logger)
}

func Update(
	ctx context.Context, s UpdateServices, id uint, request models.UserRequest,
) (response models.UserResponse, err error) {
	var email string
	var password string
	var role models.Role
	logger := slogctx.FromCtx(ctx)
	logger.Info("updating", slog.Int("uid", int(id)))
	s.SetLogger(*logger)
	found, err := s.GetUserById(id)
	if err != nil {
		logger.Error("db error", slog.String("error", err.Error()))
		return response, models.ErrServerError
	}
	if found == nil {
		logger.Info("user not found")
		return response, errors.Join(models.ErrNotFound, ErrUserNotFound)
	}
	if request.Password != "" {
		if err = IsPasswordValid(request.Password); err != nil {
			logger.Info("invalid password")
			return response, errors.Join(models.ErrBadRequest, err)
		}
		password = fmt.Sprintf("%x", sha256.Sum256([]byte(request.Password)))
	}
	if request.Email != "" {
		if err = IsEmailValid(request.Email); err != nil {
			logger.Info("invalid email")
			return response, errors.Join(models.ErrBadRequest, err)
		}
		conflict, err := s.GetUserByEmail(request.Email)
		if err != nil {
			logger.Error("db error", slog.String("error", err.Error()))
			return response, models.ErrServerError
		}
		if conflict != nil && conflict.UserId != id {
			logger.Info("user already exists")
			return response, errors.Join(models.ErrConflict, ErrUserConflict)
		}
		err = s.SendEmail(id, uuid.New().String(), "Authorize")
		if err != nil {
			logger.Error(
				"email engine error",
				slog.String("error", err.Error()),
			)
			return response, models.ErrServerError
		}
		email = request.Email
		role = models.RoleNew
	}
	err = s.UpdateUser(id, email, password, role)
	if err != nil {
		logger.Error("db error", slog.String("error", err.Error()))
		return response, models.ErrServerError
	}
	logger.Info("user updated, returning")
	response.UserId = id
	response.Email = email
	response.Role = role
	return response, nil
}
