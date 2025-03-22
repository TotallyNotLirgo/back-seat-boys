package users

import (
	"testing"
)

func TestUpdateRoleReturnsNilWhenNotFound(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	result := database.UpdateCredentials(1, "new email", "new password")
	if result != nil {
		t.Fatalf("Expected: %v, got: %v", nil, result)
	}
}

func TestUpdateRoleUpdatesDataInDB(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	request := User{
		Email:    "email",
		Password: "password",
		Role:     "role",
	}
	conn.Create(&request)
	result := database.UpdateRole(1, "admin")
	if result == nil {
		t.Fatalf("Expected: result, got: %v", result)
	}
	var user User
	conn.First(&user, result.UserId)
	if user.Role != "admin" {
		t.Fatalf("Expected: password, got: %v", user.Password)
	}
}

func TestUpdateRolePatchesDataInDB(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	request := User{
		Email:    "email",
		Password: "password",
		Role:     "role",
	}
	conn.Create(&request)
	result := database.UpdateRole(1, "")
	if result == nil {
		t.Fatalf("Expected: result, got: %v", result)
	}
	var user User
	conn.First(&user, result.UserId)
	if user.Role != "role" {
		t.Fatalf("Expected: password, got: %v", user.Password)
	}
}
