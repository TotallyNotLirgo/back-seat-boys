package main

import (
	"github.com/TotallyNotLirgo/back-seat-boys/src/logic/users"
	"github.com/TotallyNotLirgo/back-seat-boys/src/services/parser"
	"net/http"
)

func (s Server) login(w http.ResponseWriter, r *http.Request) {
	p := parser.Parser{Writer: w, Request: r}
	users.Login(p)
}
