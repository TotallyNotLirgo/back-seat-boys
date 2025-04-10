package main

import (
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/TotallyNotLirgo/back-seat-boys/parser"
	"github.com/TotallyNotLirgo/back-seat-boys/users"
	"github.com/gin-gonic/gin"
	slogctx "github.com/veqryn/slog-context"
)

func (f EndpointFacade) login(c *gin.Context) {
	var err error
	var request models.UserRequest
	var response models.UserResponse

	ctx := c.Request.Context()
	logger := slogctx.FromCtx(ctx)
	logger.Info("login")
	p := parser.Parser{Context: c}
	if err = p.Unmarshal(&request); err != nil {
		logger.Info(err.Error())
		return
	}
	response, err = users.Login(ctx, f.services, request)
	if err != nil {
		p.WriteErrorResponse(err)
		return
	}
	err = p.SetJWTCookie(response)
	if err != nil {
		logger.Error("could not encode JWT", slog.String("error", err.Error()))
		p.WriteJSONMessage(500, "could not encode JWT")
		return
	}

	p.JSON(200, response)
}

func (f EndpointFacade) logout(c *gin.Context) {
	ctx := c.Request.Context()
	logger := slogctx.FromCtx(ctx)
	logger.Info("logout")
	p := parser.Parser{Context: c}
	p.ResetJWTCookie()
	p.JSON(200, map[string]string{"message": "ok"})
}

func (f EndpointFacade) register(c *gin.Context) {
	var err error
	var request models.UserRequest
	var response models.UserResponse

	ctx := c.Request.Context()
	logger := slogctx.FromCtx(ctx)
	logger.Info("register")
	p := parser.Parser{Context: c}

	if err = p.Unmarshal(&request); err != nil {
		logger.Info(err.Error())
		return
	}

	response, err = users.Register(ctx, f.services, request)
	if err != nil {
		p.WriteErrorResponse(err)
		return
	}
	err = p.SetJWTCookie(response)
	if err != nil {
		logger.Error("could not encode JWT", slog.String("error", err.Error()))
		p.WriteJSONMessage(500, "could not encode JWT")
		return
	}

	p.JSON(200, response)
}
func (f EndpointFacade) update(c *gin.Context) {
	var err error
	var id int
	var request models.UserRequest
	var response models.UserResponse

	ctx := c.Request.Context()
	logger := slogctx.FromCtx(ctx)
	logger.Info("update")
	p := parser.Parser{Context: c}
	if id, err = p.GetPathId(); err != nil {
		logger.Info(err.Error())
		return
	}
	if err = p.CheckAccessAcceptOnlySelf(id); err != nil {
		logger.Info(err.Error())
		return
	}
	if err = p.Unmarshal(&request); err != nil {
		logger.Info(err.Error())
		return
	}
	response, err = users.Update(ctx, f.services, id, request)
	if err != nil {
		p.WriteErrorResponse(err)
		return
	}

	p.JSON(200, response)
}
func (f EndpointFacade) delete(c *gin.Context) {
	var err error
	var id int
	var response models.UserResponse

	p := parser.Parser{Context: c}
	ctx := c.Request.Context()
	logger := slogctx.FromCtx(ctx)
	logger.Info("delete")
	if id, err = p.GetPathId(); err != nil {
		logger.Info(err.Error())
		return
	}
	if err = p.CheckAccessAcceptSelf(id, models.RoleAdmin); err != nil {
		logger.Info(err.Error())
		return
	}
	response, err = users.Delete(ctx, f.services, id)
	if err != nil {
		p.WriteErrorResponse(err)
		return
	}

	p.JSON(200, response)
}
func (f EndpointFacade) authorize(c *gin.Context) {
	var err error
	var token string
	var response models.UserResponse

	p := parser.Parser{Context: c}
	ctx := c.Request.Context()
	logger := slogctx.FromCtx(ctx)
	logger.Info("authorize")
	token = c.Param("token")
	response, err = users.Authorize(ctx, f.services, token)
	if err != nil {
		p.WriteErrorResponse(err)
		return
	}

	p.JSON(200, response)
}
