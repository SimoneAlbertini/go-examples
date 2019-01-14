package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeLookForEndpoint(service AddressbookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(lookForRequest)
		contact, err := service.LookFor(req.Name, req.LastName)
		response := lookForRespose{
			Name:        contact.name,
			LastName:    contact.lastName,
			Address:     contact.address,
			PhoneNumber: contact.phoneNumber,
		}

		if err != nil {
			response.Error = err.Error()
		}

		return response, err
	}
}
