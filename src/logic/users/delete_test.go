package users

import (
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"testing"
)

func prepareDeleteTest(t *testing.T) (TestParser, TestDatabase, TestCase) {
	parser := TestParser{
		pathKey:   "id",
		pathParam: "12",
		access: models.UserResponse{
			UserId: 11,
			Email:  "user",
			Role:   models.User,
		},
	}
	database := TestDatabase{
		users: []*TestUser{{
			email:    "user",
			password: "04f8996da763b7a969b1028ee3007569eaf3a635486ddab211d512c85b9df8fb",
			role:     models.User,
			userId:   12,
		}},
	}
	c := TestCase{t}
	return parser, database, c
}

func TestDeleteInvalidBodyWrites422(t *testing.T) {
	parser, database, c := prepareDeleteTest(t)
	parser.pathParam = "hi"
	Delete(&parser, &database)

	c.AssertEquals(422, parser.status)
	c.AssertEquals("Id should be an int", parser.result)
}

func TestDeleteJWTErrorWrites401(t *testing.T) {
	parser, database, c := prepareDeleteTest(t)
	parser.readJWTError = true
	Delete(&parser, &database)

	c.AssertEquals(401, parser.status)
	c.AssertEquals("Unauthorized", parser.result)
}

func TestDeleteOtherWrites403(t *testing.T) {
	parser, database, c := prepareDeleteTest(t)
	Delete(&parser, &database)

	c.AssertEquals(403, parser.status)
	c.AssertEquals("Insufficient permissions", parser.result)
}

func TestDeleteAsAdminWrites200(t *testing.T) {
	parser, database, c := prepareDeleteTest(t)
	parser.access.Role = models.Admin
	Delete(&parser, &database)

	c.AssertEquals(200, parser.status)
	result := parser.result.(*models.UserResponse)
	c.AssertEquals("user", result.Email)
	c.AssertEquals(models.User, result.Role)
	c.AssertEquals(uint(12), result.UserId)

	c.AssertEquals((*models.UserResponse)(nil), database.GetUser(12))
}

func TestDeleteAsYourselfWrites200(t *testing.T) {
	parser, database, c := prepareDeleteTest(t)
	parser.access.UserId = 12
	Delete(&parser, &database)

	c.AssertEquals(200, parser.status)
	result := parser.result.(*models.UserResponse)
	c.AssertEquals("user", result.Email)
	c.AssertEquals(models.User, result.Role)
	c.AssertEquals(uint(12), result.UserId)

	c.AssertEquals((*models.UserResponse)(nil), database.GetUser(12))
}

func TestDeleteAdminAsAdminWrites403(t *testing.T) {
	parser, database, c := prepareDeleteTest(t)
	parser.access.Role = models.Admin
	database.users[0].role = models.Admin
	Delete(&parser, &database)

	c.AssertEquals(403, parser.status)
	c.AssertEquals("Insufficient permissions", parser.result)
}

func TestDeleteUserNotFoundWrites404(t *testing.T) {
	parser, database, c := prepareDeleteTest(t)
	parser.access.Role = models.Admin
	parser.pathParam = "999"
	Delete(&parser, &database)

	c.AssertEquals(404, parser.status)
	c.AssertEquals("User not found", parser.result)
}
