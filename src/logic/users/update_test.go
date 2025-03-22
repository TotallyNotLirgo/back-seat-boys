package users

import (
	"fmt"
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"testing"
)

func TestUpdateInvalidBodyWrites422(t *testing.T) {
	parser := TestParser{
		request:   nil,
		readError: fmt.Errorf("Invalid body"),
	}
	database := TestDatabase{}
	Update(&parser, &database)

	if expected, got := 422, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := "Invalid body", parser.result; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestUpdateJWTErrorWrites401(t *testing.T) {
	parser := TestParser{
		request:      models.UserUpdateRequest{},
		pathKey:      "id",
		pathParam:    "12",
		readJWTError: true,
	}
	database := TestDatabase{
		email:    "admin",
		password: "admin",
		role:     models.Admin,
		userId:   12,
	}
	Update(&parser, &database)

	if expected, got := 401, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := "Unauthorized", parser.result; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestUpdateForbiddenWrites403(t *testing.T) {
	parser := TestParser{
		request: models.UserUpdateRequest{},
		access: models.UserResponse{
			UserId: 11,
			Email:  "user",
			Role:   models.User,
		},
		pathKey:   "id",
		pathParam: "12",
	}
	database := TestDatabase{
		email:    "admin",
		password: "admin",
		role:     models.Admin,
		userId:   12,
	}
	Update(&parser, &database)

	if expected, got := 403, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := "Insufficient permissions", parser.result; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestUpdateHighRoleWrites200UpgradingRole(t *testing.T) {
	parser := TestParser{
		request: models.UserUpdateRequest{
			Email:    "admin new",
			Password: "admin new",
			Role:     models.Admin,
		},
		access: models.UserResponse{
			UserId: 11,
			Email:  "user",
			Role:   models.Admin,
		},
		pathKey:   "id",
		pathParam: "12",
	}
	database := TestDatabase{
		email:    "user",
		password: "04f8996da763b7a969b1028ee3007569eaf3a635486ddab211d512c85b9df8fb",
		role:     models.User,
		userId:   12,
	}
	Update(&parser, &database)

	if expected, got := 200, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	result := parser.result.(*models.UserResponse)
	if expected, got := "user", result.Email; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := models.Admin, result.Role; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := uint(12), result.UserId; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	parser.request = models.LoginRequest{
		Email:    "user",
		Password: "user",
	}
	Login(&parser, &database)
	if expected, got := 200, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestUpdateYourselfWrites200UpdatingOtherFields(t *testing.T) {
	parser := TestParser{
		request: models.UserUpdateRequest{
			Email:    "user new",
			Password: "user new",
			Role:     models.Admin,
		},
		access: models.UserResponse{
			UserId: 12,
			Email:  "user",
			Role:   models.User,
		},
		pathKey:   "id",
		pathParam: "12",
	}
	database := TestDatabase{
		email:    "user",
		password: "04f8996da763b7a969b1028ee3007569eaf3a635486ddab211d512c85b9df8fb",
		role:     models.User,
		userId:   12,
	}
	Update(&parser, &database)

	if expected, got := 200, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	result := parser.result.(*models.UserResponse)
	if expected, got := "user new", result.Email; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := models.User, result.Role; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := uint(12), result.UserId; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	parser.request = models.LoginRequest{
		Email:    "user new",
		Password: "user new",
	}
	Login(&parser, &database)
	if expected, got := 200, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestUpdateHighRoleHighRoleWrites403(t *testing.T) {
	parser := TestParser{
		request: models.UserUpdateRequest{
			Email:    "admin new",
			Password: "admin new",
			Role:     models.Admin,
		},
		access: models.UserResponse{
			UserId: 11,
			Email:  "user",
			Role:   models.Admin,
		},
		pathKey:   "id",
		pathParam: "12",
	}
	database := TestDatabase{
		email:    "user",
		password: "04f8996da763b7a969b1028ee3007569eaf3a635486ddab211d512c85b9df8fb",
		role:     models.Admin,
		userId:   12,
	}
	Update(&parser, &database)
	if expected, got := 403, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	errMsg := "Cannot change permissions of another admin"
	if expected, got := errMsg, parser.result; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestUpdateHighInvalidRoleWrites422(t *testing.T) {
	parser := TestParser{
		request: models.UserUpdateRequest{
			Email:    "admin new",
			Password: "admin new",
			Role:     "invalid",
		},
		access: models.UserResponse{
			UserId: 11,
			Email:  "user",
			Role:   models.Admin,
		},
		pathKey:   "id",
		pathParam: "12",
	}
	database := TestDatabase{
		email:    "user",
		password: "04f8996da763b7a969b1028ee3007569eaf3a635486ddab211d512c85b9df8fb",
		role:     models.User,
		userId:   12,
	}
	Update(&parser, &database)

	if expected, got := 422, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := "Invalid role", parser.result; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestUpdateEmailExistsWrites409(t *testing.T) {
	parser := TestParser{
		request: models.UserUpdateRequest{
			Email:    "user",
			Password: "admin",
			Role:     models.Admin,
		},
		access: models.UserResponse{
			UserId: 12,
			Email:  "admin",
			Role:   models.Admin,
		},
		pathKey:   "id",
		pathParam: "12",
	}
	database := TestDatabase{
		email:    "user",
		password: "04f8996da763b7a969b1028ee3007569eaf3a635486ddab211d512c85b9df8fb",
		role:     models.User,
		userId:   12,
		getUserId: 11,
	}
	Update(&parser, &database)

	if expected, got := 409, parser.status; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	if expected, got := "Email taken", parser.result; expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}
