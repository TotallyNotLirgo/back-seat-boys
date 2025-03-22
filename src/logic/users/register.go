package users

import (
	"crypto/sha256"
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
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
	logger := log.GetLogger("Register")
	logger.Info("Initializing")
	request := models.LoginRequest{}
	e := parser.ReadJSON(&request)
	if e != nil {
		logger.Warning(e.Error())
		parser.WriteString(422, e.Error())
		return
	}
	logger.Info("Fetching by email")
	user := database.GetUserByEmail(request.Email)
	if user != nil {
		logger.Warning("User already exists")
		parser.WriteString(409, "User already exists")
		return
	}
	logger.Info("Hashing password")
	hash := sha256.Sum256([]byte(request.Password))
	request.Password = fmt.Sprintf("%x", hash)
	logger.Info("Creating user")
	user = database.CreateUser(request, models.New)
	if user == nil {
		logger.Error("Could not create user")
		parser.WriteString(500, "Could not create user")
		return
	}
	parser.WriteJSON(200, user)
	logger.Info("Responding")
}
