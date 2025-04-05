package main

import (
	"context"
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"github.com/TotallyNotLirgo/back-seat-boys/parser"
	"github.com/gin-gonic/gin"
	slogctx "github.com/veqryn/slog-context"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var perm models.UserResponse
		p := parser.Parser{Context: c}
		ctx := c.Request.Context()
		logger := slogctx.FromCtx(ctx)

		if err := p.ReadJWTCookie(&perm); err == nil {
			logger = logger.With(
				slog.Group(
					"permissions",
					slog.Int("userId", perm.UserId),
					slog.String("role", string(perm.Role)),
				),
			)
			ctx = slogctx.NewCtx(ctx, logger)
			c.Request = c.Request.WithContext(
				context.WithValue(ctx, "permissions", perm),
			)
		}

		c.Next()
	}
}

func loggerMiddleware(logger slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := logger.With("hash", getRandomHash())
		ctx := slogctx.NewCtx(c.Request.Context(), l)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
