package users

import (
	"errors"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/google/go-cmp/cmp"
)

func TestUpdateInvalidPasswordReturnsBadRequest(t *testing.T) {
	expected := models.ErrBadRequest
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleUser)
	_, result := Update(
		ctx, &services, 1, models.UserRequest{Password: "pass1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestUpdateInvalidEmailReturnsBadRequest(t *testing.T) {
	expected := models.ErrBadRequest
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleUser)
	_, result := Update(
		ctx, &services, 1, models.UserRequest{Email: "email123"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestUpdateAlreadyExistsReturnsConflict(t *testing.T) {
	expected := models.ErrConflict
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleUser)
	services.insert("user2@email.com", "Password1!", models.RoleUser)
	_, result := Update(
		ctx, &services, 1, models.UserRequest{Email: "user2@email.com"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestUpdateAlreadyExistsTheSameReturnsNoConflict(t *testing.T) {
	expected := (error)(nil)
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleUser)
	services.insert("user2@email.com", "Password1!", models.RoleUser)
	_, result := Update(
		ctx, &services, 1, models.UserRequest{Email: "user@email.com"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestUpdateNotFoundReturnsNotFound(t *testing.T) {
	expected := models.ErrNotFound
	ctx, services := PrepareTest()
	services.insert("user@email.com", "Password1!", models.RoleUser)
	_, result := Update(
		ctx, &services, 2, models.UserRequest{Email: "new@email.com"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestUpdateUserBothCorrect(t *testing.T) {
	expected := models.UserResponse{
		UserId: 1,
		Email:  "new@email.com",
		Role:   "new",
	}
	ctx, services := PrepareTest()
	services.insert("user@email.com", "pass", models.RoleUser)
	result, err := Update(
		ctx,
		&services,
		1,
		models.UserRequest{Email: "new@email.com", Password: "Password1!"},
	)
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}
	if diff := cmp.Diff(result, expected); diff != "" {
		t.Error(diff)
	}
}

func TestUpdateUserEmailCorrect(t *testing.T) {
	ctx, services := PrepareTest()
	services.insert("user2@email.com", "pass2", models.RoleUser)
	services.insert("user@email.com", "pass", models.RoleUser)
	oldPassword := services.users[1].password
	Update(
		ctx,
		&services,
		2,
		models.UserRequest{Email: "new@email.com"},
	)
	if services.users[1].password != oldPassword {
		t.Error("password was changed")
	}
	if services.users[1].email != "new@email.com" {
		t.Error("email was not changed")
	}
	if services.users[1].role != models.RoleNew {
		t.Error("role was not changed")
	}
	for _, id := range services.tokens["Authorize"] {
		if id == 2 {
			return
		}
	}
	t.Error("email was not sent")
}

func TestUpdateUserPasswordCorrect(t *testing.T) {
	ctx, services := PrepareTest()
	services.insert("user@email.com", "pass", models.RoleUser)
	oldPassword := services.users[0].password
	Update(
		ctx,
		&services,
		1,
		models.UserRequest{Password: "Password1!"},
	)
	if services.users[0].password == "Password1!" {
		t.Error("password was not encrypted")
	}
	if services.users[0].password == oldPassword {
		t.Error("password was not changed")
	}
	if services.users[0].email != "user@email.com" {
		t.Error("email was changed")
	}
	if services.users[0].role != models.RoleUser {
		t.Error("role was changed")
	}
	if _, ok := services.tokens["new@email.com"]; ok {
		t.Error("email was sent")
	}
}

func TestUpdateGetByEmailServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.errors["GetUserByEmail"] = true
	services.insert("user@email.com", "Password1!", models.RoleUser)
	_, result := Update(
		ctx,
		&services,
		1,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestUpdateUpdateServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.errors["UpdateUser"] = true
	services.insert("user@email.com", "Password1!", models.RoleUser)
	_, result := Update(
		ctx,
		&services,
		1,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestUpdateGetByIdServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.errors["GetUserById"] = true
	services.insert("user@email.com", "Password1!", models.RoleUser)
	_, result := Update(
		ctx,
		&services,
		1,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestUpdateSendEmailServerError(t *testing.T) {
	expected := models.ErrServerError
	ctx, services := PrepareTest()
	services.errors["SendEmail"] = true
	services.insert("user@email.com", "Password1!", models.RoleUser)
	_, result := Update(
		ctx,
		&services,
		1,
		models.UserRequest{Email: "user@email.com", Password: "Password1!"},
	)
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
