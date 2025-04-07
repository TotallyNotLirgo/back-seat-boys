package users

import (
	"errors"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/go-cmp/cmp"
)

func TestDeleteInvalidIdReturnsNotFound(t *testing.T) {
	expected := models.ErrNotFound
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleUser)
	_, result := Delete(ctx, &services, 2)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestDeleteValidIdReturnsUserData(t *testing.T) {
	expected := models.UserResponse{
		UserId: 2,
		Email:  "user2@email.com",
		Role:   models.RoleUser,
	}
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleUser)
	services.insert("user2@email.com", "Password1!", models.RoleUser)
	result, err := Delete(ctx, &services, 2)
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}
	if diff := cmp.Diff(result, expected); diff != "" {
		t.Error(diff)
	}
	if len(services.users) != 1 {
		t.Fatal("user was not deleted")
	}
}

func TestDeleteDeleteServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleUser)
	services.insert("user2@email.com", "Password1!", models.RoleUser)
	services.errors["DeleteUser"] = true
	_, result := Delete(ctx, &services, 2)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
