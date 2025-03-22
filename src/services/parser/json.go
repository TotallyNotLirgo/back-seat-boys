package parser

import (
	"encoding/json"
	"fmt"

	"github.com/TotallyNotLirgo/back-seat-boys/src/services/log"
)

func (p Parser) ReadPath(key string) string {
	return p.Request.PathValue(key)
}

func (p Parser) WriteJSON(status int, v any) {
	logger := log.GetLogger("WriteJSON")
	logger.Info("Writing JSON")
	p.Writer.Header().Add("Content-Type", "application/json")
	p.Writer.WriteHeader(status)
	e := json.NewEncoder(p.Writer).Encode(v)
	if e != nil {
		logger.Error(e.Error())
	}
}

func (p Parser) ReadJSON(payload any) error {
	logger := log.GetLogger("ReadJSON")
	logger.Info("Reading JSON")
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
