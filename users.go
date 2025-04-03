package main

import (
	"errors"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/context"
	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/TotallyNotLirgo/back-seat-boys/users"
	"github.com/gin-gonic/gin"
)

func getStatusCode(err error) int {
	switch {
	case errors.Is(err, models.ErrUnauthorized):
		return 401
	case errors.Is(err, models.ErrForbidden):
		return 403
	case errors.Is(err, models.ErrNotFound):
		return 404
	case errors.Is(err, models.ErrConflict):
		return 409
	case errors.Is(err, models.ErrBadRequest):
		return 422
	case errors.Is(err, models.ErrServerError):
		return 500
	}
	return 500
}

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
	response, err = users.Login(&f.services, request)
	if err != nil {
		code := getStatusCode(err)
		c.WriteJSONMessage(code, err.Error())
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
	response, err = users.Register(&f.services, request)
	if err != nil {
		code := getStatusCode(err)
		c.WriteJSONMessage(code, err.Error())
		return
	}

	l.Info("Processing request")
	c.JSON(200, response)
}
func (f EndpointFacade) update(ctx *gin.Context) {
	var err error
	var id int
	var request models.UserRequest
	var response models.UserResponse

	c := context.Context{Context: ctx}
	l := f.logger.With(
		slog.String("endpoint", "update"),
		slog.String("hash", getRandomHash()),
	)
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
