package services

import (
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
)

type userModel struct {
	id       int
	email    string
	password string
	role     models.Role
}

type TestServiceAdapter struct {
	logger slog.Logger
	lastId int
	users  []*userModel
	errors map[string]bool
	tokens map[string]map[string]int
}

func NewServiceAdapter(logger slog.Logger) *TestServiceAdapter {
	tokens := make(map[string]map[string]int)
	tokens["TemporaryLogin"] = make(map[string]int)
	tokens["Authorize"] = make(map[string]int)
	return &TestServiceAdapter{
		errors: make(map[string]bool),
		tokens: tokens,
	}
}

func (tsa *TestServiceAdapter) SetLogger(logger slog.Logger) {
	tsa.logger = logger
}

func (tsa *TestServiceAdapter) GetUserById(
	id int,
) (*models.UserModel, error) {
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

func (tsa *TestServiceAdapter) UpdateUser(
	id int, email, password string, role models.Role,
) error {
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
		return nil
	}
	return nil
}

func (tsa *TestServiceAdapter) DeleteUser(id int) error {
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

func (tsa *TestServiceAdapter) InsertUser(
	email, pass string, role models.Role,
) (int, error) {
	tsa.lastId++
	tsa.users = append(tsa.users, &userModel{tsa.lastId, email, pass, role})
	return tsa.lastId, nil
}

func (tsa *TestServiceAdapter) SendEmail(id int, token, bucket string) error {
	tsa.tokens[bucket][token] = id
	return nil
}

func (tsa *TestServiceAdapter) GetIdByToken(
	token, bucket string,
) (int, bool, error) {
	id, ok := tsa.tokens[bucket][token]
	return id, ok, nil
}

func (tsa *TestServiceAdapter) DeleteToken(token, bucket string) error {
	delete(tsa.tokens[bucket], token)
	return nil
}
