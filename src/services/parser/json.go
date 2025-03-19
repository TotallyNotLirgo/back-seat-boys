package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Parser struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (p Parser) WriteJSON(status int, v any) error {
	p.Writer.Header().Add("Content-Type", "application/json")
	p.Writer.WriteHeader(status)
	return json.NewEncoder(p.Writer).Encode(v)
}

func (p Parser) ReadJSON(payload any) error {
	if p.Request.Body == nil {
		return fmt.Errorf("Missing body request")
	}
	return json.NewDecoder(p.Request.Body).Decode(payload)
}

func messageToJson(message string) map[string]string {
	return map[string]string{"message": message}
}

func (p Parser) WriteString(status int, message string) {
	e := p.WriteJSON(status, messageToJson(message))
	if e != nil {
		fmt.Printf("Could not write error: \n%v\n", e)
	}
}

