package users

import (
	"context"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/uuid"
	slogctx "github.com/veqryn/slog-context"
)

type ForgotPasswordServices interface {
	SetLogger(logger slog.Logger)
	SendEmail(id int, token, bucket string) error
	GetUserByEmail(email string) (*models.UserModel, error)
}

func ForgotPassword(
	ctx context.Context, s ForgotPasswordServices, email string,
) error {
	logger := slogctx.FromCtx(ctx)
	logger.Info("forgot password")
	s.SetLogger(*logger)
	user, err := s.GetUserByEmail(email)
	if err != nil {
		logger.Error(err.Error())
		return models.ErrServerError
	}
	if user == nil {
		logger.Warn("forgot password not found", slog.String("email", email))
		return nil
	}
	err = s.SendEmail(user.UserId, uuid.New().String(), "TemporaryLogin")
	if err != nil {
		logger.Error(err.Error())
		return models.ErrServerError
	}
	return nil
}
