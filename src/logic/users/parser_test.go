package users

import (
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"reflect"
)

type TestParser struct {
	request any
	error   error
	result  any
	status  int
	cookie  any
}

func (p *TestParser) ReadJSON(payload any) error {
	if p.error != nil {
		return p.error
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
