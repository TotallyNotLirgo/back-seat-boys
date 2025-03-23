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
	GetUserByEmail(email string) *models.UserResponse
	UpdateRole(id uint, role string) *models.UserResponse
}

type UpdateData struct {
	parser      UpdateParser
	database    UpdateDatabase
	logger      log.Logger
	request     models.UserUpdateRequest
	userId      uint
	permissions models.UserResponse
	user        *models.UserResponse
}

func Update(parser UpdateParser, database UpdateDatabase) {
	data := UpdateData{
		parser:   parser,
		database: database,
		logger:   log.GetLogger("Update"),
	}
	data.logger.Info("Initializing")
	err := data.gatherRequest()
	if err != nil {
		data.logger.Warning(err.Error())
		parser.WriteString(422, err.Error())
		return
	}
	if data.permissions.Email == "" {
		data.logger.Warning("Unauthorized")
		parser.WriteString(401, "Unauthorized")
		return
	}
	data.user = data.database.GetUser(data.userId)
	if data.user == nil {
		data.logger.Info("User not found")
		parser.WriteString(404, "User not found")
		return
	}
	if data.permissions.UserId == data.userId {
		data.logger.Info("Updating as user")
		data.updateCredentials()
		return
	} else if data.permissions.Role == models.Admin {
		data.logger.Info("Updating as admin")
		data.updateRole()
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
	switch d.request.Role {
	case "":
	case models.Admin:
	case models.New:
	case models.User:
	default:
		return fmt.Errorf("Invalid role")
	}
	userId, err := strconv.Atoi(d.parser.ReadPath("id"))
	if err != nil {
		return err
	}
	d.userId = uint(userId)
	d.parser.ReadJWTCookie(&d.permissions)
	return nil
}

func (d *UpdateData) updateRole() {
	if d.user.Role == models.Admin {
		d.logger.Warning("Admin vs Admin")
		d.parser.WriteString(403, "Cannot change permissions of another admin")
		return
	}
	result := d.database.UpdateRole(d.userId, d.request.Role)
	d.logger.Warning("Updated role")
	d.parser.WriteJSON(200, result)
}

func (d *UpdateData) updateCredentials() {
	if d.request.Email != "" {
		d.logger.Warning("Checking email")
		user := d.database.GetUserByEmail(d.request.Email)
		if user != nil && user.UserId != d.userId {
			d.logger.Warning("Email taken")
			d.parser.WriteString(409, "Email taken")
			return
		}
	}
	if d.request.Password != "" {
		d.logger.Info("Hashing password")
		hash := sha256.Sum256([]byte(d.request.Password))
		d.request.Password = fmt.Sprintf("%x", hash)
	}
	result := d.database.UpdateCredentials(
		d.userId, d.request.Email, d.request.Password,
	)
	d.logger.Warning("Updated credentials")
	d.parser.WriteJSON(200, result)
}
