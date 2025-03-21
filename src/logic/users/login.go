package users

import (
	"crypto/sha256"
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
)

type LoginParser interface {
	ReadJSON(payload any) error
	WriteJSON(status int, v any)
	WriteString(status int, message string)
	WriteJWTCookie(response models.UserResponse)
}
type LoginDatabase interface {
	GetUserByCredentials(email, password string) *models.UserResponse
}

func Login(parser LoginParser, database LoginDatabase) {
	request := models.LoginRequest{}
	e := parser.ReadJSON(&request)
	if e != nil {
		parser.WriteString(422, e.Error())
		return
	}
	hash := sha256.Sum256([]byte(request.Password))
	request.Password = fmt.Sprintf("%x", hash)
	user := database.GetUserByCredentials(request.Email, request.Password)
	if user == nil {
		parser.WriteString(401, "Invalid credentials")
		return
	}
	parser.WriteJWTCookie(*user)
	parser.WriteJSON(200, user)
}
