//go:build test

package users

import (
	"context"
	"log/slog"
	"os"

	slogctx "github.com/veqryn/slog-context"
)

func PrepareTest() (context.Context, TestServiceAdapter) {
	adapter := NewServiceAdapter()
	hOpt := slog.HandlerOptions(slog.HandlerOptions{Level: slog.LevelDebug})
	h := slogctx.NewHandler(slog.NewTextHandler(os.Stdout, &hOpt), nil)
	ctx := slogctx.NewCtx(context.Background(), slog.New(h))
	return ctx, adapter
}
