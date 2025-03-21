package main

import (
	database "github.com/TotallyNotLirgo/back-seat-boys/src/services/database"
	http "net/http"
)

type Server struct {
	database *database.ExampleDatabase
}

func main() {
	server := Server{
		database: &database.ExampleDatabase{},
	}
	http.HandleFunc("POST /api/login", server.login)
	http.HandleFunc("POST /api/register", server.register)
	http.ListenAndServe(":8090", nil)
}
