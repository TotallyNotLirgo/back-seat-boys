package parser

import "net/http"


type Parser struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

