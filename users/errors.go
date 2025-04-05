package users

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserConflict = errors.New("user with this email already exists")
)
