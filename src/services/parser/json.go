package parser

import (
	"encoding/json"
	"fmt"
)

func (p Parser) WriteJSON(status int, v any) {
	p.Writer.Header().Add("Content-Type", "application/json")
	p.Writer.WriteHeader(status)
	e := json.NewEncoder(p.Writer).Encode(v)
	if e != nil {
		fmt.Printf("Could not write: \n%v\n", e)
	}
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
	p.WriteJSON(status, messageToJson(message))
}
