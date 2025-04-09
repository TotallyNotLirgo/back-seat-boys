package users

import (
	"context"
	"errors"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	slogctx "github.com/veqryn/slog-context"
)

type AuthorizeServices interface {
	GetIdByToken(string) (int, bool, error)
	DeleteToken(string) error
	UpdateUser(id int, email, password string, role models.Role) error
	GetUserById(id int) (*models.UserModel, error)
	SetLogger(logger slog.Logger)
}

func Authorize(
	ctx context.Context, s AuthorizeServices, token string,
) (response models.UserResponse, err error) {
	logger := slogctx.FromCtx(ctx)
	logger.Info("Authorizing")
    s.SetLogger(*logger)
	id, ok, err := s.GetIdByToken(token)
	if err != nil {
		logger.Error(err.Error())
		return response, models.ErrServerError
	}
	if !ok {
		logger.Info("token not found")
		return response, errors.Join(models.ErrNotFound, ErrTokenNotFound)
	}
	err = s.DeleteToken(token)
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
