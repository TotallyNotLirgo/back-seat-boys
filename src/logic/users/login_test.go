package users

import (
	"fmt"
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"testing"
)

func TestLoginInvalidBodyWrites422(t *testing.T) {
	parser := TestParser{
		request: nil,
		error:   fmt.Errorf("Invalid body"),
	}
	database := TestDatabase{}
	Login(&parser, database)

	if expected, got := 422, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := "Invalid body", parser.result; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestLoginInvalidCredentialsWrites401(t *testing.T) {
	parser := TestParser{
		request: models.LoginRequest{
			Email:    "email",
			Password: "password",
		},
	}
	database := TestDatabase{
		email:    "admin",
		password: "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
	}
	Login(&parser, database)

	if expected, got := 401, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := "Invalid credentials", parser.result; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestLoginValidCredentialsWrites200(t *testing.T) {
	parser := TestParser{
		request: models.LoginRequest{
			Email:    "admin",
			Password: "admin",
		},
	}
	database := TestDatabase{
		email:     "admin",
		password:  "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
		role:      models.Admin,
		userId:    12,
	}
	Login(&parser, database)

	if expected, got := 200, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	result := parser.result.(*models.UserResponse)
	if expected, got := "admin", result.Email; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := models.Admin, result.Role; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := int64(12), result.UserId; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	cookie := parser.cookie.(models.UserResponse)
	if expected, got := models.Admin, cookie.Role; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := int64(12), cookie.UserId; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}
