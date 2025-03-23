package users

import (
	"fmt"
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"testing"
)

func prepareLoginTest(t *testing.T) (TestParser, TestDatabase, TestCase) {
	parser := TestParser{}
	database := TestDatabase{
		users: []*TestUser{{
			email:    "admin",
			password: "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
			role:     models.Admin,
			userId:   12,
		}},
	}
	c := TestCase{t}
	return parser, database, c
}

func TestLoginInvalidBodyWrites422(t *testing.T) {
	parser, database, c := prepareLoginTest(t)
	parser.readError = fmt.Errorf("Invalid body")
	Login(&parser, database)

	c.AssertEquals(422, parser.status)
	c.AssertEquals("Invalid body", parser.result)
}

func TestLoginInvalidCredentialsWrites401(t *testing.T) {
	parser, database, c := prepareLoginTest(t)
	parser.request = models.LoginRequest{Email: "email", Password: "password"}
	Login(&parser, database)

	c.AssertEquals(401, parser.status)
	c.AssertEquals("Invalid credentials", parser.result)
}

func TestLoginValidCredentialsWrites200(t *testing.T) {
	parser, database, c := prepareLoginTest(t)
	parser.request = models.LoginRequest{Email: "admin", Password: "admin"}
	Login(&parser, database)

	c.AssertEquals(200, parser.status)
	result := parser.result.(*models.UserResponse)
	c.AssertEquals("admin", result.Email)
	c.AssertEquals(models.Admin, result.Role)
	c.AssertEquals(uint(12), result.UserId)
	cookie := parser.cookie.(models.UserResponse)
	c.AssertEquals("admin", cookie.Email)
	c.AssertEquals(models.Admin, cookie.Role)
	c.AssertEquals(uint(12), cookie.UserId)
}
