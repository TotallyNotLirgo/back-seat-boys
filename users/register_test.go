package users

import (
	"errors"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/go-cmp/cmp"
)

func TestRegisterInvalidPasswordReturnsBadRequest(t *testing.T) {
	expected := models.ErrBadRequest
	ctx, services := PrepareTest()
	_, result := Register(
		ctx,
		&services,
		models.UserRequest{Email: "email@email.com", Password: "pass1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestRegisterInvalidEmailReturnsBadRequest(t *testing.T) {
	expected := models.ErrBadRequest
	ctx, services := PrepareTest()
	_, result := Register(
		ctx,
		&services,
		models.UserRequest{Email: "emailemail.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestRegisterUserExistsReturnsConflict(t *testing.T) {
	expected := models.ErrConflict
	ctx, services := PrepareTest()
	services.insert("user@email.com", "pass", models.RoleUser)
	_, result := Register(
		ctx,
		&services,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestRegisterUserCorrect(t *testing.T) {
	expected := models.UserResponse{
		UserId: 2,
		Email:  "new@email.com",
		Role:   "new",
	}
	ctx, services := PrepareTest()
	services.insert("user@email.com", "pass", models.RoleUser)
	result, err := Register(
		ctx,
		&services,
		models.UserRequest{Email: "new@email.com", Password: "Password1!"},
	)
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}
	if diff := cmp.Diff(result, expected); diff != "" {
		t.Error(diff)
	}
	if len(services.users) != 2 {
		t.Fatal("user was not inserted")
	}
	if services.users[1].password == "Password1!" {
		t.Error("password was not encrypted")
	}
	for _, id := range services.tokens["Authorize"] {
		if id == 2 {
			return
		}
	}
	t.Error("email was not sent")
}

func TestRegisterGetServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.errors["GetUserByEmail"] = true
	_, result := Register(
		ctx,
		&services,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestRegisterInsertServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.errors["InsertUser"] = true
	_, result := Register(
		ctx,
		&services,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestRegisterSendEmailServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.errors["SendEmail"] = true
	_, result := Register(
		ctx,
		&services,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
