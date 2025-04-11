package users

import (
	"errors"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/go-cmp/cmp"
)

func TestTemporaryLoginInvalidTokenReturnsNotFound(t *testing.T) {
	expected := models.ErrNotFound
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	_, result := TemporaryLogin(ctx, &services, "abcxyz")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestTemporaryLoginValidTokenReturnsNilChangingRole(t *testing.T) {
	expected := models.UserResponse{
		UserId: 1,
		Email:  "user@email.com",
		Role:   models.RoleNew,
	}
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.insert_token(1, "abcxyz", "TemporaryLogin")
	result, err := TemporaryLogin(ctx, &services, "abcxyz")
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}
	if diff := cmp.Diff(result, expected); diff != "" {
		t.Error(diff)
	}
	if _, ok := services.tokens["abcxyz"]; ok {
		t.Error("token was not removed")
	}
}

func TestTemporaryLoginGetByTokenServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.insert_token(1, "abcxyz", "TemporaryLogin")
	services.errors["GetIdByToken"] = true
	_, result := TemporaryLogin(ctx, &services, "abcxyz")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestTemporaryLoginDeleteTokenServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.insert_token(1, "abcxyz", "TemporaryLogin")
	services.errors["DeleteToken"] = true
	_, result := TemporaryLogin(ctx, &services, "abcxyz")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestTemporaryLoginGetUserServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.insert_token(1, "abcxyz", "TemporaryLogin")
	services.errors["GetUserById"] = true
	_, result := TemporaryLogin(ctx, &services, "abcxyz")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
