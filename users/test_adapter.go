//go:build test

package users

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
)

var (
	ErrTestServer = errors.New("Server error")
)

type userModel struct {
	id       uint
	email    string
	password string
	role     models.Role
}

type TestServiceAdapter struct {
	lastId uint
	users  []*userModel
	errors map[string]bool
	tokens map[string]map[string]uint
}

func NewServiceAdapter() TestServiceAdapter {
	tokens := make(map[string]map[string]uint)
	tokens["TemporaryLogin"] = make(map[string]uint)
	tokens["Authorize"] = make(map[string]uint)
	return TestServiceAdapter{
		errors: make(map[string]bool),
		tokens: tokens,
	}
}

func (tsa *TestServiceAdapter) SetLogger(logger slog.Logger) {
}

func (tsa *TestServiceAdapter) insert(email, pass string, role models.Role) {
	tsa.lastId++
	pass = fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))
	tsa.users = append(tsa.users, &userModel{tsa.lastId, email, pass, role})
}

func (tsa *TestServiceAdapter) insert_token(id uint, token, bucket string) {
	tsa.tokens[bucket][token] = id
}

func (tsa *TestServiceAdapter) GetUserById(
	id uint,
) (*models.UserModel, error) {
	if tsa.errors["GetUserById"] {
		return nil, ErrTestServer
	}
	for _, user := range tsa.users {
		if user.id != id {
			continue
		}
		return &models.UserModel{
			UserId: user.id,
			Email:  user.email,
			Role:   user.role,
		}, nil
	}
	return nil, nil
}

func (tsa *TestServiceAdapter) GetUserByEmail(
	email string,
) (*models.UserModel, error) {
	if tsa.errors["GetUserByEmail"] {
		return nil, ErrTestServer
	}
	for _, user := range tsa.users {
		if user.email != email {
			continue
		}
		return &models.UserModel{
			UserId: user.id,
			Email:  user.email,
			Role:   user.role,
		}, nil
	}
	return nil, nil
}

func (tsa *TestServiceAdapter) DeleteUser(id uint) error {
	if tsa.errors["DeleteUser"] {
		return ErrTestServer
	}
	var user *userModel
	var i int
	for i, user = range tsa.users {
		if user.id != id {
			continue
		}
		break
	}
	tsa.users = append(tsa.users[:i], tsa.users[i+1:]...)
	return nil
}

func (tsa *TestServiceAdapter) UpdateUser(
	id uint, email, password string, role models.Role,
) error {
	if tsa.errors["UpdateUser"] {
		return ErrTestServer
	}
	for _, user := range tsa.users {
		if user.id != id {
			continue
		}
		if email != "" {
			user.email = email
		}
		if password != "" {
			user.password = password
		}
		if role != "" {
			user.role = role
		}
		break
	}
	return nil
}

func (tsa *TestServiceAdapter) GetUserByCredentials(
	email, password string,
) (*models.UserModel, error) {
	if tsa.errors["GetUserByCredentials"] {
		return nil, ErrTestServer
	}
	for _, user := range tsa.users {
		if user.email != email {
			continue
		}
		if user.password != password {
			continue
		}
		return &models.UserModel{
			UserId: user.id,
			Email:  user.email,
			Role:   user.role,
		}, nil
	}
	return nil, nil
}

func (tsa *TestServiceAdapter) InsertUser(
	email, pass string, role models.Role,
) (uint, error) {
	if tsa.errors["InsertUser"] {
		return 0, ErrTestServer
	}
	tsa.lastId++
	tsa.users = append(tsa.users, &userModel{tsa.lastId, email, pass, role})
	return tsa.lastId, nil
}

func (tsa *TestServiceAdapter) SendEmail(id uint, token, bucket string) error {
	if tsa.errors["SendEmail"] {
		return errors.New("Server error")
	}
	tsa.tokens[bucket][token] = id
	return nil
}

func (tsa *TestServiceAdapter) GetIdByToken(token, bucket string) (uint, bool, error) {
	if tsa.errors["GetIdByToken"] {
		return 0, false, errors.New("Server error")
	}
	id, ok := tsa.tokens[bucket][token]
	return id, ok, nil
}

func (tsa *TestServiceAdapter) DeleteToken(token, bucket string) error {
	if tsa.errors["DeleteToken"] {
		return errors.New("Server error")
	}
	delete(tsa.tokens[bucket], token)
	return nil
}
