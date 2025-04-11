package users

import (
	"errors"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
)

func TestForgotPasswordCreatesTemporaryLoginSession(t *testing.T) {
	expected := (error)(nil)
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	result := ForgotPassword(ctx, &services, "user@email.com")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
	for _, id := range services.tokens["TemporaryLogin"] {
		if id == 1 {
			return
		}
	}
	t.Error("email was not sent")
}

func TestForgotPasswordReturnsNoErrorWhenEmailNotFound(t *testing.T) {
	expected := (error)(nil)
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	result := ForgotPassword(ctx, &services, "missing@email.com")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
	if len(services.tokens["TemporaryLogin"]) != 0 {
		t.Errorf("expected len: %v, got: %v", 0, len(services.tokens))
	}
}

func TestForgotPasswordGetUserServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.errors["GetUserByEmail"] = true
	result := ForgotPassword(ctx, &services, "user@email.com")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestForgotPasswordSendEmailServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.errors["SendEmail"] = true
	result := ForgotPassword(ctx, &services, "user@email.com")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
