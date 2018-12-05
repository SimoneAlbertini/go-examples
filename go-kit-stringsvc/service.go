package main

// StringService service features
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// Service implementation
type stringService struct{}
