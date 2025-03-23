package users

import (
	"strconv"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
)

type DeleteParser interface {
	WriteJSON(status int, v any)
	ReadPath(key string) string
	WriteString(status int, message string)
	ReadJWTCookie(request *models.UserResponse)
}
type DeleteDatabase interface {
	GetUser(id uint) *models.UserResponse
	DeleteUser(id uint)
}

func Delete(parser DeleteParser, database DeleteDatabase) {
	logger := log.GetLogger("Delete")
	logger.Info("Initializing")
	userIdInt, err := strconv.Atoi(parser.ReadPath("id"))
	if err != nil {
		logger.Warning(err.Error())
		parser.WriteString(422, "Id should be an int")
		return
	}
	userId := uint(userIdInt)
	var permissions models.UserResponse
	parser.ReadJWTCookie(&permissions)
	if permissions.Email == "" {
		logger.Warning("Unauthorized")
		parser.WriteString(401, "Unauthorized")
		return
	}
	logger.Warning("Fetching user")
	user := database.GetUser(userId)
	if user == nil {
		logger.Warning("User not found")
		parser.WriteString(404, "User not found")
		return
	}
	if permissions.Role == models.Admin && user.Role != models.Admin {
		logger.Info("Deleting other")
		database.DeleteUser(userId)
		parser.WriteJSON(200, user)
		return
	}
	if user.UserId == permissions.UserId {
		logger.Info("Deleting yourself")
		database.DeleteUser(userId)
		parser.WriteJSON(200, user)
		return
	}
	logger.Warning("Insufficient permissions")
	parser.WriteString(403, "Insufficient permissions")

}
