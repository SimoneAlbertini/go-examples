package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request and Response for Uppercase(string)
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` //errors don't JSON mashall, serialize in string format
}

// Request and Response for Count(string)
type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

// decoder for Uppercase Request
func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//  decoder for Count Request
func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// encoder for http response
func encodeResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}
