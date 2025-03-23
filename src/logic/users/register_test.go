package users

import (
	"fmt"
	"testing"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
)

func prepareRegisterTest(t *testing.T) (TestParser, TestDatabase, TestCase) {
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

func TestRegisterInvalidBodyWrites422(t *testing.T) {
	parser, database, c := prepareRegisterTest(t)
	parser.readError = fmt.Errorf("Invalid body")
	Register(&parser, &database)

	c.AssertEquals(422, parser.status)
	c.AssertEquals("Invalid body", parser.result)
}

func TestRegisterUserEmailExistsWrites409(t *testing.T) {
	parser, database, c := prepareRegisterTest(t)
	parser.request = models.LoginRequest{Email: "admin", Password: "admin"}
	Register(&parser, &database)

	c.AssertEquals(409, parser.status)
	c.AssertEquals("User already exists", parser.result)
}

func TestRegisterUserCreatedWrites200(t *testing.T) {
	parser, database, c := prepareRegisterTest(t)
	parser.request = models.LoginRequest{Email: "user", Password: "user"}
	Register(&parser, &database)

	c.AssertEquals(200, parser.status)
	result := parser.result.(*models.UserResponse)
	c.AssertEquals("user", result.Email)
	c.AssertEquals(models.New, result.Role)
	c.AssertEquals(uint(13), result.UserId)
}

func TestRegisterDatabaseReturns500(t *testing.T) {
	parser, database, c := prepareRegisterTest(t)
	parser.request = models.LoginRequest{Email: "user", Password: "user"}
	database.requestFails = true
	Register(&parser, &database)

	c.AssertEquals(500, parser.status)
	c.AssertEquals("Could not create user", parser.result)
}

func TestLoginRegisteredUserWrites200(t *testing.T) {
	parser, database, c := prepareRegisterTest(t)
	parser.request = models.LoginRequest{Email: "user", Password: "user"}
	Register(&parser, &database)
	Login(&parser, &database)

	c.AssertEquals(200, parser.status)
	result := parser.result.(*models.UserResponse)
	c.AssertEquals("user", result.Email)
	c.AssertEquals(models.New, result.Role)
	c.AssertEquals(uint(13), result.UserId)
}
