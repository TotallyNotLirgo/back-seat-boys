package users

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
)

type UpdateParser interface {
	ReadPath(key string) string
	ReadJSON(payload any) error
	WriteJSON(status int, v any)
	WriteString(status int, message string)
	ReadJWTCookie(request *models.UserResponse)
}
type UpdateDatabase interface {
	GetUser(id uint) *models.UserResponse
	UpdateCredentials(id uint, email, password string) *models.UserResponse
	UpdateRole(id uint, role string) *models.UserResponse
}

type UpdateData struct {
	parser   UpdateParser
	database UpdateDatabase
	logger   log.Logger
	request  models.UserUpdateRequest
	userId   uint
}

func Update(parser UpdateParser, database UpdateDatabase) {
	data := UpdateData{
		parser: parser,
		database: database,
		logger: log.GetLogger("Update"),
	}
	data.logger.Info("Initializing")
	err := data.gatherRequest()
	if err != nil {
		data.logger.Warning(err.Error())
		parser.WriteString(422, err.Error())
		return
	}
	var permissions models.UserResponse
	parser.ReadJWTCookie(&permissions)
	if permissions.Email == "" {
		data.logger.Warning("Unauthorized")
		parser.WriteString(401, "Unauthorized")
		return
	} else if permissions.Role == models.Admin {
		data.logger.Info("Updating as admin")
		data.updateRole()
		return
	} else if permissions.UserId == data.userId {
		data.logger.Info("Updating as user")
		data.updateCredentials()
		return
	}
	data.logger.Info("Responding")
	parser.WriteString(403, "Insufficient permissions")
}

func (d *UpdateData) gatherRequest() error {
	d.request = models.UserUpdateRequest{}
	err := d.parser.ReadJSON(&d.request)
	if err != nil {
		return err
	}
	userId, err := strconv.Atoi(d.parser.ReadPath("id"))
	if err != nil {
		return err
	}
	d.userId = uint(userId)
	return nil
}

func (d *UpdateData) updateRole() {
	user := d.database.GetUser(d.userId)
	if user.Role == models.Admin {
		d.logger.Warning("Admin vs Admin")
		d.parser.WriteString(403, "Cannot change permissions of another admin")
		return
	}
	result := d.database.UpdateRole(d.userId, d.request.Role)
	d.logger.Warning("Updated role")
	d.parser.WriteJSON(200, result)
}

func (d *UpdateData) updateCredentials() {
	d.logger.Info("Hashing password")
	hash := sha256.Sum256([]byte(d.request.Password))
	d.request.Password = fmt.Sprintf("%x", hash)
	result := d.database.UpdateCredentials(
		d.userId, d.request.Email, d.request.Password,
	)
	d.logger.Warning("Updated credentials")
	d.parser.WriteJSON(200, result)
}
