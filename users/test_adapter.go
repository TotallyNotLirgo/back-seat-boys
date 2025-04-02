//go:build test

package users

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
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
}

func (tsa *TestServiceAdapter) insert(email, pass string, role models.Role) {
	tsa.lastId++
	pass = fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))
	tsa.users = append(tsa.users, &userModel{tsa.lastId, email, pass, role})
}

func (tsa *TestServiceAdapter) GetUserByEmail(
	email string,
) (*models.UserModel, error) {
	if _, ok := tsa.errors["GetUserByEmail"]; ok {
		return nil, errors.New("Server error")
	}
	for _, user := range tsa.users {
		if user.email != email {
			continue
		}
		return &models.UserModel{
			Id:    user.id,
			Email: user.email,
			Role:  user.role,
		}, nil
	}
	return nil, nil
}

func (tsa *TestServiceAdapter) InsertUser(
	email, pass string, role models.Role,
) (int, error) {
	if _, ok := tsa.errors["InsertUser"]; ok {
		return 0, errors.New("Server error")
	}
	tsa.lastId++
	tsa.users = append(tsa.users, &userModel{tsa.lastId, email, pass, role})
	return tsa.lastId, nil
}
