package parser

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Parser struct {
	*gin.Context
	Body gin.H
}

func (c *Parser) WriteJSONMessage(status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func (c *Parser) WriteError(status int, message string, args ...any) error {
	err := fmt.Errorf(message, args...)
	c.WriteJSONMessage(status, err.Error())
	return err
}

func (c *Parser) GetJSON() error {
	err := c.ShouldBindJSON(&c.Body)
	if err != nil {
		c.WriteJSONMessage(422, "invalid json")
		return err
	}
	return nil
}

func (c *Parser) GetPathId() (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.WriteJSONMessage(422, "id not an int")
		return 0, err
	}
	return id, nil
}
