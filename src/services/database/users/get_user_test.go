package users

import "testing"

func TestGetUserReturnsNilWhenNotFound(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	result := database.GetUser(1)
	if result != nil {
		t.Fatalf("Expected: %v, got: %v", nil, result)
	}
}

func TestGetUserReturnsUserWhenFound(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	request := User{
		Email:    "email",
		Password: "password",
		Role:     "role",
	}
	conn.Create(&request)
	result := database.GetUser(1)
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
