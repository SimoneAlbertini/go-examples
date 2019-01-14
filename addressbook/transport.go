package main

import (
	"context"
	"encoding/json"
	"net/http"
)

//Request and response for LookFor request
type lookForRequest struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

type lookForRespose struct {
	Name        string `json:"name"`
	LastName    string `json:"lastname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
	Error       string `json:"error,omitempty"`
}

func decodeLookForRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request lookForRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeLookForRespose(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}
