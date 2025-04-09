package users

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrTokenNotFound = errors.New("token not found")
	ErrUserConflict  = errors.New("user with this email already exists")
)
