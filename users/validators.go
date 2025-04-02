package users

import (
	"errors"
	"net/mail"
	"strings"
)

var (
	ErrPassTooShort     = errors.New("password too short")
	ErrPassTooLong      = errors.New("password too long")
	ErrPassUpperMissing = errors.New("password must contain uppercase letters")
	ErrPassLowerMissing = errors.New("password must contain lowercase letters")
	ErrPassDigitMissing = errors.New("password must contain digits")
	ErrPassOtherMissing = errors.New("password must contain special characters")
	ErrEmailInvalid     = errors.New("email address is invalid")
)

const (
	minPassLen = 8
	maxPassLen = 255
	digits     = "1234567890"
	special    = " !\"#$&'()+,./:;<=?@[\\]^_`{|}~"
)

func IsPasswordValid(p string) error {
	switch {
	case len(p) < minPassLen:
		return ErrPassTooShort
	case len(p) > maxPassLen:
		return ErrPassTooLong
	case strings.ToLower(p) == p:
		return ErrPassUpperMissing
	case strings.ToUpper(p) == p:
		return ErrPassLowerMissing
	case !strings.ContainsAny(p, digits):
		return ErrPassDigitMissing
	case !strings.ContainsAny(p, special):
		return ErrPassOtherMissing
	}
	return nil
}

func IsEmailValid(a string) error {
	_, err := mail.ParseAddress(a)
	if err != nil {
		return ErrEmailInvalid
	}
	return nil
}
