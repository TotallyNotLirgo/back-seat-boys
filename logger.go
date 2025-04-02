package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/MatusOllah/slogcolor"
	"log/slog"
	"os"
	"time"
)

func getDevLogger() (*slog.Logger, func() error) {
	options := slogcolor.DefaultOptions
	options.Level = slog.LevelDebug
	logger := slog.New(
		slogcolor.NewHandler(os.Stderr, options),
	)
	return logger, func() error { return nil }
}

func getProdLogger() (*slog.Logger, func() error) {
	f, err := os.OpenFile("log.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{}))
	return logger, f.Close
}

func getRandomHash() string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(time.Now().UnixNano()))
	hash := sha256.Sum256(b)
	return fmt.Sprintf("%x", hash[:4])
}
