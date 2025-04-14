package users

import (
	"context"
	"errors"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	slogctx "github.com/veqryn/slog-context"
)

type TemporaryLoginServices interface {
	GetIdByToken(token, bucket string) (uint, bool, error)
	DeleteToken(token, bucket string) error
	GetUserById(id uint) (*models.UserModel, error)
	SetLogger(logger slog.Logger)
}

func TemporaryLogin(
	ctx context.Context, s TemporaryLoginServices, token string,
) (response models.UserResponse, err error) {
	logger := slogctx.FromCtx(ctx)
	logger.Info("temporary login")
	s.SetLogger(*logger)
	id, ok, err := s.GetIdByToken(token, "TemporaryLogin")
	if err != nil {
		logger.Error(err.Error())
		return response, models.ErrServerError
	}
	if !ok {
		logger.Info("token not found")
		return response, errors.Join(models.ErrNotFound, ErrTokenNotFound)
	}
	err = s.DeleteToken(token, "TemporaryLogin")
	if err != nil {
		logger.Error(err.Error())
		return response, models.ErrServerError
	}
	user, err := s.GetUserById(id)
	if err != nil {
		logger.Error(err.Error())
		return response, models.ErrServerError
	}
	response.Email = user.Email
	response.Role = user.Role
	response.UserId = user.UserId
	return response, nil
}
