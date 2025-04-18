package users

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	slogctx "github.com/veqryn/slog-context"
)

type LoginServices interface {
	GetUserByCredentials(email, password string) (*models.UserModel, error)
	SetLogger(logger slog.Logger)
}

func Login(
	ctx context.Context, s LoginServices, request models.UserRequest,
) (response models.UserResponse, err error) {
	logger := slogctx.FromCtx(ctx)
	logger.Info("login", slog.String("email", request.Email))
	s.SetLogger(*logger)
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(request.Password)))
	user, err := s.GetUserByCredentials(request.Email, password)
	if err != nil {
		logger.Error("db error", slog.String("error", err.Error()))
		return response, models.ErrServerError
	}
	if user == nil {
		logger.Info("user not found")
		return response, errors.Join(models.ErrUnauthorized, ErrUserNotFound)
	}
	logger.Info("user found, returning")
	response.UserId = user.UserId
	response.Email = user.Email
	response.Role = user.Role
	return response, nil
}
