package users

import (
	"crypto/sha256"
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
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
	logger := log.GetLogger("Login")
	logger.Info("Initializing")
	request := models.LoginRequest{}
	e := parser.ReadJSON(&request)
	if e != nil {
		logger.Warning(e.Error())
		parser.WriteString(422, e.Error())
		return
	}
	logger.Info("Hashing password")
	hash := sha256.Sum256([]byte(request.Password))
	request.Password = fmt.Sprintf("%x", hash)
	logger.Info("Fetching by credentials")
	user := database.GetUserByCredentials(request.Email, request.Password)
	if user == nil {
		logger.Warning("Invalid credentials")
		parser.WriteString(401, "Invalid credentials")
		return
	}
	logger.Info("Writing cookie")
	parser.WriteJWTCookie(*user)
	parser.WriteJSON(200, user)
	logger.Info("Responding")
}
