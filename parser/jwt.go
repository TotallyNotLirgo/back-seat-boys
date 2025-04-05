package parser

import (
	"errors"
	"fmt"
	"time"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/golang-jwt/jwt"
)

var secretKey = "mysecretkey"

func (c Parser) WriteJWTCookie(response models.UserResponse) error {
	value, err := generateJWT(response)
	if err != nil {
		return err
	}
	c.SetCookie(
		"JWT",
		value,
		int(time.Now().Add(365*24*time.Hour).Unix()),
		"/",
		"localhost",
		true,
		true,
	)
	return nil
}

func (c Parser) ReadJWTCookie(request *models.UserResponse) error {
	cookie, err := c.Cookie("JWT")
	if err != nil {
		return err
	}
	token, err := jwt.Parse(
		cookie,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("could not extract claims")
	}
	request.UserId = int(claims["userId"].(float64))
	role, ok := claims["role"].(string)
	if !ok {
		return errors.New("could not convert role")
	}
	request.Role = models.Role(role)
	request.Email, ok = claims["email"].(string)
	if !ok {
		return errors.New("could not convert email")
	}
	return nil
}

func generateJWT(response models.UserResponse) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["authorized"] = true
	claims["userId"] = response.UserId
	claims["email"] = response.Email
	claims["role"] = response.Role
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("could not encode: \n%v\n", err)
	}
	return tokenString, nil
}
