//go:build test

package users

import (
	"testing"
)

type TestCase struct {
	t *testing.T
}

func (c *TestCase) AssertEquals(expected, got any) {
	if expected != got {
		c.t.Fatalf("Expected %v, got %v", expected, got)
	}
}
