package users

import (
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"testing"
)

func TestCreateUserInsertsDataIntoDB(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	request := models.LoginRequest{Email: "email", Password: "password"}
	result := database.CreateUser(request, "role")
	if result == nil {
		t.Fatalf("Expected: result, got: %v", result)
	}
	var user User
	conn.First(&user, result.UserId)
	if user.Email != "email" {
		t.Fatalf("Expected: email, got: %v", user.Email)
	}
	if user.Password != "password" {
		t.Fatalf("Expected: password, got: %v", user.Password)
	}
	if user.Role != "role" {
		t.Fatalf("Expected: role, got: %v", user.Role)
	}
}
