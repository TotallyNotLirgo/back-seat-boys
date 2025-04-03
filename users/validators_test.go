package users

import (
	"errors"
	"strings"
	"testing"
)

func TestIsPasswordValidCorrectReturnsNil(t *testing.T) {
	expected := (error)(nil)
	result := IsPasswordValid("Password1!")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestIsPasswordValidTooShortReturnsError(t *testing.T) {
	expected := ErrPassTooShort
	result := IsPasswordValid("Pass1!")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestIsPasswordValidTooLongReturnsError(t *testing.T) {
	expected := ErrPassTooLong
	result := IsPasswordValid(strings.Repeat("Pas1!", 100))
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestIsPasswordValidUpperMissingReturnsError(t *testing.T) {
	expected := ErrPassUpperMissing
	result := IsPasswordValid("password1!")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestIsPasswordValidLowerMissingReturnsError(t *testing.T) {
	expected := ErrPassLowerMissing
	result := IsPasswordValid("PASSWORD1!")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestIsPasswordValidNumberMissingReturnsError(t *testing.T) {
	expected := ErrPassDigitMissing
	result := IsPasswordValid("Password!")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestIsPasswordValidSpecialMissingReturnsError(t *testing.T) {
	expected := ErrPassOtherMissing
	result := IsPasswordValid("Password1")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestIsEmailValidValidReturnsNil(t *testing.T) {
	expected := (error)(nil)
	result := IsEmailValid("email@email.com")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestIsEmailValidInvalidReturnsError(t *testing.T) {
	expected := ErrEmailInvalid
	result := IsEmailValid("email@")
	if !errors.Is(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
