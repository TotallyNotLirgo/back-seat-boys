package users

import (
	"context"
	"errors"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	slogctx "github.com/veqryn/slog-context"
)

type AuthorizeServices interface {
	GetIdByToken(token, bucket string) (uint, bool, error)
	DeleteToken(token, bucket string) error
	UpdateUser(id uint, email, password string, role models.Role) error
	GetUserById(id uint) (*models.UserModel, error)
	SetLogger(logger slog.Logger)
}

func Authorize(
	ctx context.Context, s AuthorizeServices, token string,
) (response models.UserResponse, err error) {
	logger := slogctx.FromCtx(ctx)
	logger.Info("authorizing")
	s.SetLogger(*logger)
	id, ok, err := s.GetIdByToken(token, "Authorize")
	if err != nil {
		logger.Error(err.Error())
		return response, models.ErrServerError
	}
	if !ok {
		logger.Info("token not found")
		return response, errors.Join(models.ErrNotFound, ErrTokenNotFound)
	}
	err = s.DeleteToken(token, "Authorize")
	if err != nil {
		logger.Error(err.Error())
		return response, models.ErrServerError
	}
	user, err := s.GetUserById(id)
	if err != nil {
		logger.Error(err.Error())
		return response, models.ErrServerError
	}
	if user.Role.GreaterEqual(models.RoleUser) {
		response.Email = user.Email
		response.UserId = user.UserId
		response.Role = user.Role
		return response, nil
	}
	err = s.UpdateUser(id, "", "", models.RoleUser)
	if err != nil {
		logger.Error(err.Error())
		return response, models.ErrServerError
	}
	response.Email = user.Email
	response.UserId = user.UserId
	response.Role = models.RoleUser
	return response, nil
}
