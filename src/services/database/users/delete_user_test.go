package users

import "testing"

func TestDeleteUserDeletesNothingWhenNotFound(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	database.DeleteUser(1)
}

func TestDeleteUserDeletesUserWhenFound(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	request := User{
		Email:    "email",
		Password: "password",
		Role:     "role",
	}
	conn.Create(&request)
	database.DeleteUser(1)
	result := database.GetUser(1)
	if result != nil {
		t.Fatalf("Expected: nil, got: %v", result)
	}
}
