package users

import (
	"errors"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/go-cmp/cmp"
)

func TestAuthorizeInvalidTokenReturnsNotFound(t *testing.T) {
	expected := models.ErrNotFound
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	_, result := Authorize(ctx, &services, "abcxyz")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestAuthorizeValidTokenReturnsNilChangingRole(t *testing.T) {
	expected := models.UserResponse{
		UserId: 1,
		Email:  "user@email.com",
		Role:   models.RoleUser,
	}
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.insert_token(1, "abcxyz")
	result, err := Authorize(ctx, &services, "abcxyz")
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}
	if diff := cmp.Diff(result, expected); diff != "" {
		t.Error(diff)
	}
	if services.users[0].role != models.RoleUser {
		t.Error("role was not changed")
	}
	if _, ok := services.tokens["abcxyz"]; ok {
		t.Error("token was not removed")
	}
}

func TestAuthorizeWillNotRevertRole(t *testing.T) {
	expected := models.UserResponse{
		UserId: 1,
		Email:  "user@email.com",
		Role:   models.RoleAdmin,
	}
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleAdmin)
	services.insert_token(1, "abcxyz")
	result, err := Authorize(ctx, &services, "abcxyz")
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}
	if diff := cmp.Diff(result, expected); diff != "" {
		t.Error(diff)
	}
	if services.users[0].role != models.RoleAdmin {
		t.Error("role was changed")
	}
	if _, ok := services.tokens["abcxyz"]; ok {
		t.Error("token was not removed")
	}
}

func TestAuthorizeGetByTokenServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.insert_token(1, "abcxyz")
	services.errors["GetIdByToken"] = true
	_, result := Authorize(ctx, &services, "abcxyz")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestAuthorizeDeleteTokenServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.insert_token(1, "abcxyz")
	services.errors["DeleteToken"] = true
	_, result := Authorize(ctx, &services, "abcxyz")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestAuthorizeUpdateServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.insert_token(1, "abcxyz")
	services.errors["UpdateUser"] = true
	_, result := Authorize(ctx, &services, "abcxyz")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestAuthorizeGetUserServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleNew)
	services.insert_token(1, "abcxyz")
	services.errors["GetUserById"] = true
	_, result := Authorize(ctx, &services, "abcxyz")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
