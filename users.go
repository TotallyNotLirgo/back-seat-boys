package main

import (
	users "github.com/TotallyNotLirgo/back-seat-boys/src/logic/users"
	parser "github.com/TotallyNotLirgo/back-seat-boys/src/services/parser"
	http "net/http"
)

func (s Server) login(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Login")
	p := parser.Parser{Writer: w, Request: r}
	users.Login(p, s.database)
}

func (s Server) register(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Register")
	p := parser.Parser{Writer: w, Request: r}
	users.Register(p, s.database)
}

func (s Server) update(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Update")
	p := parser.Parser{Writer: w, Request: r}
	users.Update(p, s.database)
}

func (s Server) delete(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Delete")
	p := parser.Parser{Writer: w, Request: r}
	users.Delete(p, s.database)
}
