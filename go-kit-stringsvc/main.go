package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "calling endpoint")
			return next(ctx, request)
		}
	}
}

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	strService := stringService{}

	uppercaseEndpoint := makeUppercaseEndpoint(strService)
	uppercaseEndpoint = loggingMiddleware(log.With(logger, "method", "uppercase"))(uppercaseEndpoint)
	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		decodeUppercaseRequest,
		encodeResponse,
	)

	countEndpoint := makeCountEndpoint(strService)
	countEndpoint = loggingMiddleware(log.With(logger, "method", "count"))(countEndpoint)
	countHandler := httptransport.NewServer(
		countEndpoint,
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8000", nil)
}
