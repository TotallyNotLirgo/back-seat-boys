package parser

import (
	"errors"
	"strconv"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidJSON = errors.Join(
		models.ErrBadRequest,
		errors.New("invalid json"),
	)
	ErrInvalidId = errors.Join(
		models.ErrBadRequest,
		errors.New("id not an int"),
	)
)

type Parser struct {
	*gin.Context
	Body gin.H
}

func (c *Parser) WriteJSONMessage(status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func (c *Parser) GetJSON() error {
	err := c.ShouldBindJSON(&c.Body)
	if err != nil {
		c.WriteErrorResponse(ErrInvalidJSON)
		return err
	}
	return nil
}

func (c *Parser) GetPathId() (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.WriteErrorResponse(ErrInvalidId)
		return 0, err
	}
	return uint(id), nil
}
