package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/TotallyNotLirgo/back-seat-boys/services"
	"github.com/gin-gonic/gin"
)

type EndpointFacade struct {
	services *services.TestServiceAdapter
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

	r.Use(loggerMiddleware(*logger))
	r.Use(authMiddleware())
	r.Use(CORSMiddleware())
	f := EndpointFacade{services.NewServiceAdapter(*logger)}
	r.POST("/api/login", f.login)
	r.POST("/api/logout", f.logout)
	r.POST("/api/register", f.register)
	r.PATCH("/api/users/:id", f.update)
	r.DELETE("/api/users/:id", f.delete)
	r.POST("/api/authorize/:token", f.authorize)
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.RunTLS(":8090", "cert/bake-roll.crt", "cert/bake-roll.key")
}
