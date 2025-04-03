package users

import (
	"errors"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/go-cmp/cmp"
)

func TestLoginUserNotFoundReturnsUnauthorized(t *testing.T) {
	expected := models.ErrUnauthorized
	services := NewServiceAdapter()
	_, result := Login(
		&services,
		models.UserRequest{Email: "email@email.com", Password: "pass1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestLoginUserCorrect(t *testing.T) {
	expected := models.UserResponse{
		Id:    1,
		Email: "user@email.com",
		Role:  models.RoleUser,
	}
	services := NewServiceAdapter()
	services.insert("user@email.com", "Password1!", models.RoleUser)
	result, err := Login(
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
	services := NewServiceAdapter()
	services.errors["GetUserByCredentials"] = true
	_, result := Login(
		&services,
		models.UserRequest{Email: "email@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
