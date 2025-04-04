//go:build test

package users

import (
	"crypto/sha256"
	"errors"
	"fmt"

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
	tokens map[string]string
}

func NewServiceAdapter() TestServiceAdapter {
	return TestServiceAdapter{
		errors: make(map[string]bool),
		tokens: make(map[string]string),
	}
}

func (tsa *TestServiceAdapter) insert(email, pass string, role models.Role) {
	tsa.lastId++
	pass = fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))
	tsa.users = append(tsa.users, &userModel{tsa.lastId, email, pass, role})
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
			UserId:    user.id,
			Email: user.email,
			Role:  user.role,
		}, nil
	}
	return nil, nil
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
			UserId:    user.id,
			Email: user.email,
			Role:  user.role,
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

func (tsa *TestServiceAdapter) SendEmail(email, token string) error {
	if tsa.errors["SendEmail"] {
		return errors.New("Server error")
	}
	tsa.tokens[email] = token
	return nil
}
