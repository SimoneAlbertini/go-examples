package main

import (
	"errors"
	"log"
	"net/http"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
)

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("string is empty")
	}

	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

func main() {
	strService := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(strService),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(strService),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
