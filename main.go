package main

import (
	http "net/http"

	"github.com/TotallyNotLirgo/back-seat-boys/src/services/database/users"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Server struct {
	database *users.Database
	logger   log.Logger
}

func main() {
	logger := log.GetLogger("main")
	logger.Info("Starting server")
	conn, err := gorm.Open(
		sqlite.Open("main.db"),
		&gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Silent),
		},
	)
	if err != nil {
		panic("failed to connect database")
	}
	logger.Info("Connected to the DB")
	conn.AutoMigrate(&users.User{})
	server := Server{
		database: &users.Database{Connection: conn},
		logger:   logger,
	}
	http.HandleFunc("POST /api/login", server.login)
	http.HandleFunc("POST /api/register", server.register)
	http.HandleFunc("PATCH /api/users/{id}", server.update)
	http.HandleFunc("DELETE /api/users/{id}", server.delete)
	http.ListenAndServeTLS(
		":8090",
		"cert/localhost.crt",
		"cert/localhost.key",
		nil,
	)
}
