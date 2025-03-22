package main

import (
	http "net/http"

	database "github.com/TotallyNotLirgo/back-seat-boys/src/services/database"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Server struct {
	database *database.Database
	logger log.Logger
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
	conn.AutoMigrate(&database.User{})
	server := Server{
		database: &database.Database{Connection: conn},
		logger: logger,
	}
	http.HandleFunc("POST /api/login", server.login)
	http.HandleFunc("POST /api/register", server.register)
	http.ListenAndServe(":8090", nil)
}
