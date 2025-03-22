package users

import "testing"

func TestGetUserByEmailReturnsNilWhenNotFound(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	result := database.GetUserByEmail("email")
	if result != nil {
		t.Fatalf("Expected: %v, got: %v", nil, result)
	}
}

func TestGetUserByEmailReturnsUserWhenFound(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	request := User{
		Email:    "email",
		Password: "password",
		Role:     "role",
	}
	conn.Create(&request)
	result := database.GetUserByEmail("email")
	if result == nil {
		t.Fatalf("Expected: result, got: %v", result)
	}
	if result.Email != "email" {
		t.Fatalf("Expected: email, got: %v", result.Email)
	}
	if result.Role != "role" {
		t.Fatalf("Expected: role, got: %v", result.Role)
	}
}
