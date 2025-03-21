package users

import (
	"crypto/sha256"
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
)

type RegisterParser interface {
	ReadJSON(payload any) error
	WriteJSON(status int, v any)
	WriteString(status int, message string)
}
type RegisterDatabase interface {
	GetUserByCredentials(email, password string) *models.UserResponse
	GetUserByEmail(email string) *models.UserResponse
	CreateUser(user models.LoginRequest, role string) *models.UserResponse
}

func Register(parser RegisterParser, database RegisterDatabase) {
	request := models.LoginRequest{}
	e := parser.ReadJSON(&request)
	if e != nil {
		parser.WriteString(422, e.Error())
		return
	}
	user := database.GetUserByEmail(request.Email)
	if user != nil {
		parser.WriteString(409, "User already exists")
		return
	}
	hash := sha256.Sum256([]byte(request.Password))
	request.Password = fmt.Sprintf("%x", hash)
	user = database.CreateUser(request, models.New)
	parser.WriteJSON(200, user)
}
