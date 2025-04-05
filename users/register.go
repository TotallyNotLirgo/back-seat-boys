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

type RegisterServices interface {
	GetUserByEmail(email string) (*models.UserModel, error)
	InsertUser(email, pass string, role models.Role) (int, error)
	SendEmail(email, token string) error
}

func Register(
	ctx context.Context, s RegisterServices, request models.UserRequest,
) (response models.UserResponse, err error) {
	logger := slogctx.FromCtx(ctx)
	logger = logger.With(slog.String("email", request.Email))

	if err = IsPasswordValid(request.Password); err != nil {
		logger.Info("invalid password")
		return response, errors.Join(models.ErrBadRequest, err)
	}
	if err = IsEmailValid(request.Email); err != nil {
		logger.Info("invalid email")
		return response, errors.Join(models.ErrBadRequest, err)
	}
	conflict, err := s.GetUserByEmail(request.Email)
	if err != nil {
		logger.Error("db error", slog.String("error", err.Error()))
		return response, models.ErrServerError
	}
	if conflict != nil {
		logger.Info("user already exists")
		return response, errors.Join(models.ErrConflict, ErrUserConflict)
	}
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(request.Password)))
	id, err := s.InsertUser(request.Email, password, models.RoleNew)
	if err != nil {
		logger.Error("db error", slog.String("error", err.Error()))
		return response, models.ErrServerError
	}
	err = s.SendEmail(request.Email, uuid.New().String())
	if err != nil {
		logger.Error("email engine error", slog.String("error", err.Error()))
		return response, models.ErrServerError
	}
	logger.Info("user created, returning")
	response.UserId = id
	response.Email = request.Email
	response.Role = models.RoleNew
	return response, nil
}
