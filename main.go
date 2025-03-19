package main

import "net/http"

type Server struct{}

func main() {
    server := Server{}
    http.HandleFunc("POST /api/login", server.login)
    http.ListenAndServe(":8090", nil)
}
