package users

import (
	"errors"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/go-cmp/cmp"
)

func TestLoginUserNotFoundReturnsUnauthorized(t *testing.T) {
	expected := models.ErrUnauthorized
	ctx, services := PrepareTest()
	services.insert("email@email.com", "Password1!", models.RoleUser)
	_, result := Login(
		ctx,
		&services,
		models.UserRequest{Email: "email@email.com", Password: "pass1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestLoginUserCorrect(t *testing.T) {
	expected := models.UserResponse{
		UserId: 2,
		Email:  "user@email.com",
		Role:   models.RoleUser,
	}
	ctx, services := PrepareTest()
	services.insert("admin@email.com", "Password1!", models.RoleUser)
	services.insert("user@email.com", "Password1!", models.RoleUser)
	result, err := Login(
		ctx,
		&services,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}
	if diff := cmp.Diff(result, expected); diff != "" {
		t.Error(diff)
	}
}

func TestLoginGetReturnsServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.errors["GetUserByCredentials"] = true
	_, result := Login(
		ctx,
		&services,
		models.UserRequest{Email: "email@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
