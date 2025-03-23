package users

import (
	"fmt"
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"testing"
)

func prepareUpdateTest(t *testing.T) (TestParser, TestDatabase, TestCase) {
	parser := TestParser{
		pathKey:   "id",
		pathParam: "12",
		request: models.UserUpdateRequest{
			Email:    "new",
			Password: "new",
			Role:     models.Admin,
		},
		access: models.UserResponse{
			UserId: 11,
			Email:  "user",
			Role:   models.User,
		},
	}
	users := []*TestUser{
		{
			email:    "admin",
			password: "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
			role:     models.Admin,
			userId:   5,
		},
		{
			email:    "user",
			password: "04f8996da763b7a969b1028ee3007569eaf3a635486ddab211d512c85b9df8fb",
			role:     models.Admin,
			userId:   12,
		},
	}
	database := TestDatabase{users: users}
	c := TestCase{t}
	return parser, database, c
}

func TestUpdateInvalidBodyWrites422(t *testing.T) {
	parser, database, c := prepareUpdateTest(t)
	parser.readError = fmt.Errorf("Invalid body")
	Update(&parser, &database)

	c.AssertEquals(422, parser.status)
	c.AssertEquals("Invalid body", parser.result)
}

func TestUpdateJWTErrorWrites401(t *testing.T) {
	parser, database, c := prepareUpdateTest(t)
	parser.request = models.UserUpdateRequest{}
	parser.readJWTError = true
	Update(&parser, &database)

	c.AssertEquals(401, parser.status)
	c.AssertEquals("Unauthorized", parser.result)
}

func TestUpdateForbiddenWrites403(t *testing.T) {
	parser, database, c := prepareUpdateTest(t)
	parser.request = models.UserUpdateRequest{}
	Update(&parser, &database)

	c.AssertEquals(403, parser.status)
	c.AssertEquals("Insufficient permissions", parser.result)
}

func TestUpdateHighRoleWrites200UpgradingRole(t *testing.T) {
	parser, database, c := prepareUpdateTest(t)
	parser.access.Role = models.Admin
	database.users[1].role = models.User
	Update(&parser, &database)

	c.AssertEquals(200, parser.status)
	result := parser.result.(*models.UserResponse)
	c.AssertEquals("user", result.Email)
	c.AssertEquals(models.Admin, result.Role)
	c.AssertEquals(uint(12), result.UserId)

	parser.request = models.LoginRequest{Email: "user", Password: "user"}
	Login(&parser, &database)
	c.AssertEquals(200, parser.status)
}

func TestUpdateYourselfWrites200UpdatingOtherFields(t *testing.T) {
	parser, database, c := prepareUpdateTest(t)
	parser.access.UserId = 12
	database.users[1].role = models.User
	Update(&parser, &database)

	c.AssertEquals(200, parser.status)
	result := parser.result.(*models.UserResponse)
	c.AssertEquals("new", result.Email)
	c.AssertEquals(models.User, result.Role)
	c.AssertEquals(uint(12), result.UserId)

	parser.request = models.LoginRequest{Email: "new", Password: "new"}
	Login(&parser, &database)
	c.AssertEquals(200, parser.status)
}

func TestUpdateHighRoleHighRoleWrites403(t *testing.T) {
	parser, database, c := prepareUpdateTest(t)
	parser.access.Role = models.Admin
	Update(&parser, &database)
	c.AssertEquals(403, parser.status)
	c.AssertEquals("Cannot change permissions of another admin", parser.result)
}

func TestUpdateHighInvalidRoleWrites422(t *testing.T) {
	parser, database, c := prepareUpdateTest(t)
	parser.request = models.UserUpdateRequest{Role: "Invalid"}
	parser.access.Role = models.Admin
	database.users[1].role = models.User
	Update(&parser, &database)

	c.AssertEquals(422, parser.status)
	c.AssertEquals("Invalid role", parser.result)
}

func TestUpdateEmailExistsWrites409(t *testing.T) {
	parser, database, c := prepareUpdateTest(t)
	parser.request = models.UserUpdateRequest{Email: "admin"}
	parser.access.UserId = 12
	Update(&parser, &database)

	c.AssertEquals(409, parser.status)
	c.AssertEquals("Email taken", parser.result)
}

func TestUpdateUserNotFoundWrites404(t *testing.T) {
	parser, database, c := prepareUpdateTest(t)
	parser.access.Role = models.Admin
	parser.pathParam = "999"
	Update(&parser, &database)

	c.AssertEquals(404, parser.status)
	c.AssertEquals("User not found", parser.result)
}
