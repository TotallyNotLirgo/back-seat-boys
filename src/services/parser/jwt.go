package parser

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
	"github.com/golang-jwt/jwt"
)

var secretKey = "mysecretkey"

func (p Parser) WriteJWTCookie(response models.UserResponse) {
	logger := log.GetLogger("WriteJWTCookie")
	logger.Info("Writing JWT")
	cookie := http.Cookie{}
	cookie.Name = "JWT"
	logger.Info("Generating JWT")
	value, err := generateJWT(response)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	cookie.Value = value
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"
	logger.Info("Setting cookie")
	http.SetCookie(p.Writer, &cookie)
}

func (p Parser) ReadJWTCookie(request *models.UserResponse) {
	logger := log.GetLogger("ReadJWTCookie")
	logger.Info("Reading JWT")
	cookie, err := p.Request.Cookie("JWT")
	if err != nil {
		logger.Error(err.Error())
		return
	}
	token, err := jwt.Parse(
		cookie.Value,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logger.Error("Unable to extract claims")
		return
	}
	request.UserId = uint(claims["userId"].(float64))
	request.Role = claims["role"].(string)
	logger.Info("Extracted user %v (%v)", request.UserId, request.Role)
}

func generateJWT(response models.UserResponse) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["authorized"] = true
	claims["userId"] = response.UserId
	claims["role"] = response.Role
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("could not encode: \n%v\n", err)
	}
	return tokenString, nil
}
