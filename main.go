package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

type EndpointFacade struct {
	logger *slog.Logger
}

var appEnv = os.Getenv("APP_ENV")

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	r := gin.New()

	var logger *slog.Logger
	var closer func() error
	switch appEnv {
	case "PROD":
		logger, closer = getProdLogger()
	default:
		logger, closer = getDevLogger()
	}
	defer closer()

	f := EndpointFacade{logger}
	r.POST("/api/login", f.login)
	r.POST("/api/register", f.register)
	r.PATCH("/api/users/:id", f.update)
	r.DELETE("/api/users/:id", f.delete)
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.RunTLS(":8090", "cert/localhost.crt", "cert/localhost.key")
}
