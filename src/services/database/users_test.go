package users

import (
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func prepareConnection() (conn *gorm.DB) {
	conn, err := gorm.Open(
		sqlite.Open(":memory:"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		panic("failed to connect database")
	}
	conn.AutoMigrate(&User{})
	return
}

func TestGetUserByCredentialsReturnsNilWhenNotFound(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	result := database.GetUserByCredentials("email", "password")
	if result != nil {
		t.Fatalf("Expected: %v, got: %v", nil, result)
	}
}

func TestGetUserByCredentialsReturnsUserWhenFound(t *testing.T) {
	conn := prepareConnection()
	database := Database{conn}
	request := User{
		Email:    "email",
		Password: "password",
		Role:     "role",
	}
	conn.Create(&request)
	result := database.GetUserByCredentials("email", "password")
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
