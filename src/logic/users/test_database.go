//go:build test

package users

import (
	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
)

type TestUser struct {
	email    string
	password string
	role     string
	userId   uint
}

type TestDatabase struct {
	users        []*TestUser
	requestFails bool
}

func (d TestDatabase) GetUserByCredentials(
	email, password string,
) *models.UserResponse {
	for _, user := range d.users {
		if user.email != email || user.password != password {
			continue
		}
		response := models.UserResponse{
			UserId: user.userId,
			Role:   user.role,
			Email:  user.email,
		}
		return &response
	}
	return nil
}
func (d TestDatabase) GetUserByEmail(email string) *models.UserResponse {
	for _, user := range d.users {
		if user.email != email {
			continue
		}
		response := models.UserResponse{
			UserId: user.userId,
			Role:   user.role,
			Email:  user.email,
		}
		return &response
	}
	return nil
}
func (d TestDatabase) GetUser(id uint) *models.UserResponse {
	for _, user := range d.users {
		if user.userId != id {
			continue
		}
		response := models.UserResponse{
			UserId: user.userId,
			Role:   user.role,
			Email:  user.email,
		}
		return &response
	}
	return nil
}

func (d *TestDatabase) CreateUser(
	user models.LoginRequest, role string,
) *models.UserResponse {
	if d.requestFails {
		return nil
	}
	lastUserId := d.users[len(d.users)-1].userId
	newUser := TestUser{
		email:    user.Email,
		password: user.Password,
		role:     role,
		userId:   lastUserId + 1,
	}
	d.users = append(d.users, &newUser)
	response := models.UserResponse{
		UserId: newUser.userId,
		Role:   newUser.role,
		Email:  newUser.email,
	}
	return &response
}

func (d *TestDatabase) UpdateCredentials(
	id uint, email, password string,
) *models.UserResponse {
	for _, user := range d.users {
		if user.userId != id {
			continue
		}
		if email != "" {
			user.email = email
		}
		if password != "" {
			user.password = password
		}
		response := models.UserResponse{
			UserId: user.userId,
			Role:   user.role,
			Email:  user.email,
		}
		return &response
	}
	return nil
}

func (d *TestDatabase) UpdateRole(
	id uint, role string,
) *models.UserResponse {
	for _, user := range d.users {
		if user.userId != id {
			continue
		}
		if role != "" {
			user.role = role
		}
		response := models.UserResponse{
			UserId: user.userId,
			Role:   user.role,
			Email:  user.email,
		}
		return &response
	}
	return nil
}

func (d *TestDatabase) DeleteUser(id uint) {
	index := 0
	for i, user := range d.users {
		if user.userId == id {
			index = i
			break
		}
	}
	d.users = append(d.users[:index], d.users[index+1:]...)
}
