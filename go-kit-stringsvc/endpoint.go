package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// Create endpoint (look for type Endpoint in go-kit!)
func makeUppercaseEndpoint(service StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := service.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{V: v, Err: err.Error()}, nil
		}
		return uppercaseResponse{V: v, Err: ""}, nil
	}
}

func makeCountEndpoint(service StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := service.Count(req.S)
		return countResponse{V: v}, nil
	}
}
