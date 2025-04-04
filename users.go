package main

import (
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/context"
	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/TotallyNotLirgo/back-seat-boys/users"
	"github.com/gin-gonic/gin"
)

func (f EndpointFacade) login(ctx *gin.Context) {
	var err error
	var request models.UserRequest
	var response models.UserResponse

	c := context.Context{Context: ctx}
	l := f.logger.With(
		slog.String("endpoint", "login"),
		slog.String("hash", getRandomHash()),
	)
	if err = c.Unmarshal(&request); err != nil {
		l.Info(err.Error())
		return
	}
	l.Info("Processing the request")
	response, err = users.Login(f.services, request)
	if err != nil {
		c.WriteErrorResponse(err)
		return
	}
	err = c.WriteJWTCookie(response)
	if err != nil {
		c.WriteJSONMessage(500, err.Error())
		return
	}

	c.JSON(200, response)
}
func (f EndpointFacade) register(ctx *gin.Context) {
	var err error
	var request models.UserRequest
	var response models.UserResponse

	c := context.Context{Context: ctx}
	l := f.logger.With(
		slog.String("endpoint", "register"),
		slog.String("hash", getRandomHash()),
	)
	if err = c.Unmarshal(&request); err != nil {
		l.Info(err.Error())
		return
	}
	response, err = users.Register(f.services, request)
	if err != nil {
		c.WriteErrorResponse(err)
		return
	}

	l.Info("Processing request")
	c.JSON(200, response)
}
func (f EndpointFacade) update(ctx *gin.Context) {
	var err error
	var id int
	var request models.UserRequest
	var permissions models.UserResponse
	var response models.UserResponse

	c := context.Context{Context: ctx}
	l := f.logger.With(
		slog.String("endpoint", "update"),
		slog.String("hash", getRandomHash()),
	)
	err = c.ReadJWTCookie(&permissions)
	if err != nil {
		c.WriteJSONMessage(422, err.Error())
		return
	}
	l = l.With(slog.Int("userId", permissions.UserId))
	if id, err = c.GetPathId(); err != nil {
		l.Info(err.Error())
		return
	}
	if err = c.Unmarshal(&request); err != nil {
		l.Info(err.Error())
		return
	}

	l.Info("Processing request")
	c.JSON(200, gin.H{"id": id, "response": response})
}
func (f EndpointFacade) delete(ctx *gin.Context) {
	var err error
	var id int
	var response models.UserResponse

	c := context.Context{Context: ctx}
	l := f.logger.With(
		slog.String("endpoint", "delete"),
		slog.String("hash", getRandomHash()),
	)
	if id, err = c.GetPathId(); err != nil {
		l.Info(err.Error())
		return
	}

	l.Info("Processing request")
	c.JSON(200, gin.H{"id": id, "response": response})
}
