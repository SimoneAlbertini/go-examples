package main

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeLookForEndpoint(service AddressbookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(lookForRequest)
		cnt, err := service.LookFor(req.Name, req.LastName)

		response := lookForRespose{
			Name:        cnt.Name,
			LastName:    cnt.LastName,
			Address:     cnt.Address,
			PhoneNumber: cnt.PhoneNumber,
		}

		if cnt.IsEmpty() {
			response.Error = errors.New("Contact Not Found").Error()
		}

		if err != nil {
			response.Error = err.Error()
		}

		return response, err
	}
}
