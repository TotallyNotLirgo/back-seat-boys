package users

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
)

type ExampleParser struct {
	request any
	error   error
	result  any
	status  int
	cookie  any
}

func (p *ExampleParser) ReadJSON(payload any) error {
	if p.error != nil {
		return p.error
	}
	reflect.ValueOf(payload).Elem().Set(reflect.ValueOf(p.request))
	return nil
}
func (p *ExampleParser) WriteJSON(status int, v any) {
	p.status = status
	p.result = v
}
func (p *ExampleParser) WriteString(status int, message string) {
	p.status = status
	p.result = message
}
func (p *ExampleParser) WriteJWTCookie(response models.AuthModel) {
	p.cookie = response
}

type ExampleDatabase struct {
	email     string
	password  string
	role      string
	lastLogin int64
	userId    int64
}

func (d ExampleDatabase) GetUserByCredentials(
	email, password string,
) *models.AuthModel {
	if d.email != email || d.password != password {
		return nil
	}
	response := models.AuthModel{
		Role:      d.role,
		LastLogin: d.lastLogin,
		UserId:    d.userId,
	}
	return &response
}

func TestLoginInvalidBodyWrites422(t *testing.T) {
	parser := ExampleParser{request: nil, error: fmt.Errorf("Invalid body")}
	database := ExampleDatabase{}
	Login(&parser, database)

	if expected, got := 422, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := "Invalid body", parser.result; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestLoginInvalidCredentialsWrites401(t *testing.T) {
	parser := ExampleParser{
		request: models.LoginRequest{
			Email:    "email",
			Password: "password",
		},
	}
	database := ExampleDatabase{
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
	parser := ExampleParser{
		request: models.LoginRequest{
			Email:    "admin",
			Password: "admin",
		},
	}
	database := ExampleDatabase{
		email:     "admin",
		password:  "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
		role:      "admin",
		lastLogin: 1234,
		userId: 12,
	}
	Login(&parser, database)

	if expected, got := 200, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	result := parser.result.(*models.AuthModel)
	if expected, got := "admin", result.Role; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := int64(1234), result.LastLogin; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := int64(12), result.UserId; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	cookie := parser.cookie.(models.AuthModel)
	if expected, got := "admin", cookie.Role; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := int64(1234), cookie.LastLogin; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := int64(12), cookie.UserId; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}
