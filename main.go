package main

import (
	"io"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/general"
	"github.com/TotallyNotLirgo/back-seat-boys/services/database"
	"github.com/gin-gonic/gin"
)

type EndpointFacade struct {
	services *services.TestServiceAdapter
}

func main() {
	config := general.GetConfig()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	r := gin.New()

	var logger *slog.Logger
	var closer func() error
	switch config.APP_ENV {
	case "PROD":
		logger, closer = general.GetProdLogger()
	default:
		logger, closer = general.GetDevLogger()
	}
	defer closer()

	r.Use(loggerMiddleware(*logger))
	r.Use(authMiddleware())
	r.Use(CORSMiddleware())
	f := EndpointFacade{services.NewServiceAdapter(*logger)}
	r.POST("/api/login", f.login)
	r.POST("/api/logout", f.logout)
	r.POST("/api/forgot_password", f.forgotPassword)
	r.POST("/api/temporary_login", f.temporaryLogin)
	r.POST("/api/register", f.register)
	r.PATCH("/api/users/:id", f.update)
	r.DELETE("/api/users/:id", f.delete)
	r.POST("/api/authorize/:token", f.authorize)
	r.SetTrustedProxies([]string{config.TRUSTED_PROXIES})
	r.RunTLS(config.PORT, config.CERT_FILE, config.KEY_FILE)
}
