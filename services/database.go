package services

import (
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
	tokens map[string]string
}

func NewServiceAdapter() *TestServiceAdapter {
	return &TestServiceAdapter{
		errors: make(map[string]bool),
		tokens: make(map[string]string),
	}
}

func (tsa *TestServiceAdapter) GetUserByEmail(
	email string,
) (*models.UserModel, error) {
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

func (tsa *TestServiceAdapter) GetUserByCredentials(
	email, password string,
) (*models.UserModel, error) {
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
	tsa.lastId++
	tsa.users = append(tsa.users, &userModel{tsa.lastId, email, pass, role})
	return tsa.lastId, nil
}

func (tsa *TestServiceAdapter) SendEmail(email, token string) error {
	tsa.tokens[email] = token
	return nil
}
