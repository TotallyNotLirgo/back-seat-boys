package users

import (
	"context"
	"errors"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	slogctx "github.com/veqryn/slog-context"
)

type DeleteServices interface {
	GetUserById(id int) (*models.UserModel, error)
	DeleteUser(id int) error
	SetLogger(logger slog.Logger)
}

func Delete(
	ctx context.Context, s DeleteServices, id int,
) (response models.UserResponse, err error) {
	logger := slogctx.FromCtx(ctx)
	logger.Info("Deleting", slog.Int("uid", id))
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
	err = s.DeleteUser(id)
	if err != nil {
		logger.Error("db error", slog.String("error", err.Error()))
		return response, models.ErrServerError
	}
	response.UserId = found.UserId
	response.Email = found.Email
	response.Role = found.Role
	return response, nil
}
