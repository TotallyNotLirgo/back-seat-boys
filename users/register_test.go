package users

import (
	"errors"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/go-cmp/cmp"
)

func TestRegisterInvalidPasswordReturnsBadRequest(t *testing.T) {
	expected := models.ErrBadRequest
	services := TestServiceAdapter{}
	_, result := Register(
		&services,
		models.UserRequest{Email: "email@email.com", Password: "pass1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestRegisterInvalidEmailReturnsBadRequest(t *testing.T) {
	expected := models.ErrBadRequest
	services := TestServiceAdapter{}
	_, result := Register(
		&services,
		models.UserRequest{Email: "emailemail.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestRegisterUserExistsReturnsConflict(t *testing.T) {
	expected := models.ErrConflict
	services := TestServiceAdapter{}
	services.insert("user@email.com", "pass", models.RoleUser)
	_, result := Register(
		&services,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestRegisterUserCorrect(t *testing.T) {
	expected := models.UserResponse{Id: 2, Email: "new@email.com", Role: "new"}
	services := TestServiceAdapter{}
	services.insert("user@email.com", "pass", models.RoleUser)
	result, err := Register(
		&services,
		models.UserRequest{Email: "new@email.com", Password: "Password1!"},
	)
	if err != nil {
		t.Errorf("Expected: %v, got: %v", nil, err)
	}
	if diff := cmp.Diff(result, expected); diff != "" {
		t.Error(diff)
	}
	if len(services.users) != 2 {
		t.Fatal("User was not inserted")
	}
	if services.users[1].password == "Password1!" {
		t.Error("Password was not encrypted")
	}
}

func TestRegisterGetServerError(t *testing.T) {
	expected := models.ErrServerError
	services := TestServiceAdapter{}
	services.errors = make(map[string]bool, 1)
	services.errors["GetUserByEmail"] = true
	_, result := Register(
		&services,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestRegisterInsertServerError(t *testing.T) {
	expected := models.ErrServerError
	services := TestServiceAdapter{}
	services.errors = make(map[string]bool, 1)
	services.errors["InsertUser"] = true
	_, result := Register(
		&services,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
