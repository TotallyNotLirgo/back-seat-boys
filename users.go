package main

import (
	users "github.com/TotallyNotLirgo/back-seat-boys/src/logic/users"
	database "github.com/TotallyNotLirgo/back-seat-boys/src/services/database"
	parser "github.com/TotallyNotLirgo/back-seat-boys/src/services/parser"
	http "net/http"
)

func (s Server) login(w http.ResponseWriter, r *http.Request) {
	p := parser.Parser{Writer: w, Request: r}
	d := database.ExampleDatabase{}
	users.Login(p, d)
}
