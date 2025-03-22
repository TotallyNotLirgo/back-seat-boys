package users

import (
	"testing"
)

func TestUpdateCredentialsReturnsNilWhenNotFound(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	result := database.UpdateCredentials(1, "new email", "new password")
	if result != nil {
		t.Fatalf("Expected: %v, got: %v", nil, result)
	}
}

func TestUpdateCredentialsUpdatesDataInDB(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	request := User{
		Email:    "email",
		Password: "password",
		Role:     "role",
	}
	conn.Create(&request)
	result := database.UpdateCredentials(1, "new email", "new password")
	if result == nil {
		t.Fatalf("Expected: result, got: %v", result)
	}
	var user User
	conn.First(&user, result.UserId)
	if user.Email != "new email" {
		t.Fatalf("Expected: email, got: %v", user.Email)
	}
	if user.Password != "new password" {
		t.Fatalf("Expected: password, got: %v", user.Password)
	}
}

func TestUpdateCredentialsPatchesDataInDB(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	request := User{
		Email:    "email",
		Password: "password",
		Role:     "role",
	}
	conn.Create(&request)
	result := database.UpdateCredentials(1, "", "")
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
}
