package users

import (
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"reflect"
)

type TestParser struct {
	request      any
	readError    error
	result       any
	status       int
	cookie       any
	pathKey      string
	pathParam    string
	access       models.UserResponse
	readJWTError bool
}

func (p *TestParser) ReadJSON(payload any) error {
	if p.readError != nil {
		return p.readError
	}
	reflect.ValueOf(payload).Elem().Set(reflect.ValueOf(p.request))
	return nil
}
func (p *TestParser) WriteJSON(status int, v any) {
	p.status = status
	p.result = v
}
func (p *TestParser) WriteString(status int, message string) {
	p.status = status
	p.result = message
}
func (p *TestParser) WriteJWTCookie(response models.UserResponse) {
	p.cookie = response
}
func (p *TestParser) ReadJWTCookie(request *models.UserResponse) {
	if p.readJWTError {
		return
	}
	request.Role = p.access.Role
	request.Email = p.access.Email
	request.UserId = p.access.UserId
}
func (p *TestParser) ReadPath(key string) string {
	if p.pathKey == key {
		return p.pathParam
	}
	return ""
}
