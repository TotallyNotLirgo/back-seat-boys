package parser

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TotallyNotLirgo/back-seat-boys/src/models"
	"github.com/golang-jwt/jwt"
)

var secretKey = "mysecretkey"

func (p Parser) WriteJWTCookie(response models.UserResponse) {
	cookie := http.Cookie{}
	cookie.Name = "JWT"
	cookie.Value = generateJWT(response)
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"
	http.SetCookie(p.Writer, &cookie)
}

func (p Parser) ReadJWTCookie(request *models.UserResponse) error {
	cookie, err := p.Request.Cookie("JWT")
	if err != nil {
		return err
	}
	token, err := jwt.Parse(
		cookie.Value,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("unable to extract claims")
	}
	request.UserId = int64(claims["userId"].(float64))
	request.Role = claims["role"].(string)
	request.LastLogin = int64(claims["lastLogin"].(float64))
	fmt.Printf(
		"%v, %v, %v",
		request.UserId,
		request.Role,
		request.LastLogin,
		)
	return nil
}

func generateJWT(response models.UserResponse) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["authorized"] = true
	claims["userId"] = response.UserId
	claims["role"] = response.Role
	claims["lastLogin"] = response.LastLogin
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Printf("could not encode: \n%v\n", err)
		return ""
	}
	return tokenString
}
