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
	id       int
	email    string
	password string
	role     models.Role
}

type TestServiceAdapter struct {
	lastId int
	users  []*userModel
	errors map[string]bool
	tokens map[string]int
}

func NewServiceAdapter() TestServiceAdapter {
	return TestServiceAdapter{
		errors: make(map[string]bool),
		tokens: make(map[string]int),
	}
}

func (tsa *TestServiceAdapter) SetLogger(logger slog.Logger) {
}

func (tsa *TestServiceAdapter) insert(email, pass string, role models.Role) {
	tsa.lastId++
	pass = fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))
	tsa.users = append(tsa.users, &userModel{tsa.lastId, email, pass, role})
}

func (tsa *TestServiceAdapter) insert_token(id int, token string) {
	tsa.tokens[token] = id
}

func (tsa *TestServiceAdapter) GetUserById(
	id int,
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

func (tsa *TestServiceAdapter) DeleteUser(id int) error {
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
	id int, email, password string, role models.Role,
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
) (int, error) {
	if tsa.errors["InsertUser"] {
		return 0, ErrTestServer
	}
	tsa.lastId++
	tsa.users = append(tsa.users, &userModel{tsa.lastId, email, pass, role})
	return tsa.lastId, nil
}

func (tsa *TestServiceAdapter) SendEmail(id int, token string) error {
	if tsa.errors["SendEmail"] {
		return errors.New("Server error")
	}
	tsa.tokens[token] = id
	return nil
}

func (tsa *TestServiceAdapter) GetIdByToken(token string) (int, bool, error) {
	if tsa.errors["GetIdByToken"] {
		return 0, false, errors.New("Server error")
	}
	id, ok := tsa.tokens[token]
	return id, ok, nil
}

func (tsa *TestServiceAdapter) DeleteToken(token string) error {
	if tsa.errors["DeleteToken"] {
		return errors.New("Server error")
	}
	delete(tsa.tokens, token)
	return nil
}
