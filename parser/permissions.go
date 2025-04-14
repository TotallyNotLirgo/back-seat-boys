package parser

import (
	"errors"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
)

var (
	ErrPermissionForbidden = errors.Join(
		models.ErrForbidden,
		errors.New("you don't have sufficient permissions"),
	)
	ErrPermissionUnauthorized = errors.Join(
		models.ErrUnauthorized,
		errors.New("you need to be logged in"),
	)
)

func (c *Parser) CheckAccessAcceptOnlySelf(id uint) error {
	ctx := c.Request.Context()
	perm, ok := ctx.Value("permissions").(models.UserResponse)
	if !ok {
		c.WriteErrorResponse(ErrPermissionUnauthorized)
		return ErrPermissionUnauthorized
	}
	if perm.UserId == id {
		return nil
	}
	c.WriteErrorResponse(ErrPermissionForbidden)
	return ErrPermissionForbidden
}

func (c *Parser) CheckAccessAcceptSelf(id uint, role models.Role) error {
	ctx := c.Request.Context()
	perm, ok := ctx.Value("permissions").(models.UserResponse)
	if !ok {
		c.WriteErrorResponse(ErrPermissionUnauthorized)
		return ErrPermissionUnauthorized
	}
	if perm.UserId == id || perm.Role.GreaterEqual(role) {
		return nil
	}
	c.WriteErrorResponse(ErrPermissionForbidden)
	return ErrPermissionForbidden
}

func (c *Parser) CheckAccess(role models.Role) error {
	ctx := c.Request.Context()
	perm, ok := ctx.Value("permissions").(models.UserResponse)
	if !ok {
		c.WriteErrorResponse(ErrPermissionUnauthorized)
		return ErrPermissionUnauthorized
	}
	if perm.Role.GreaterEqual(role) {
		return nil
	}
	c.WriteErrorResponse(ErrPermissionForbidden)
	return ErrPermissionForbidden
}
