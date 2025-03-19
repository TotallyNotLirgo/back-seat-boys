package users

import "github.com/TotallyNotLirgo/back-seat-boys/src/models"


type Parser interface {
    ReadJSON(payload any) error
    WriteJSON(status int, v any) error
    WriteString(status int, message string)
}

func Login(parser Parser) {
    request := models.LoginRequest{}
    e := parser.ReadJSON(&request)
    if e != nil {
        parser.WriteString(422, e.Error())
        return
    }
    e = parser.WriteJSON(200, request)
    if e != nil {
        parser.WriteString(422, e.Error())
        return
    }
}
