package main

import (
	http "net/http"

	database "github.com/TotallyNotLirgo/back-seat-boys/src/services/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Server struct {
	database *database.Database
}

func main() {
	conn, err := gorm.Open(
		sqlite.Open("main.db"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		panic("failed to connect database")
	}
	conn.AutoMigrate(&database.User{})
	server := Server{
		database: &database.Database{Connection: conn},
	}
	http.HandleFunc("POST /api/login", server.login)
	http.HandleFunc("POST /api/register", server.register)
	http.ListenAndServe(":8090", nil)
}
