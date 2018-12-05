package main

import (
	"errors"
	"strings"
)

// StringService service features
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// Service implementation
type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("string is empty")
	}

	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
