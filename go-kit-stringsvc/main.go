package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	var strService StringService
	strService = stringService{}
	strService = loggingMiddleware{logger: logger, next: strService}

	uppercaseEndpoint := makeUppercaseEndpoint(strService)
	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		decodeUppercaseRequest,
		encodeResponse,
	)

	countEndpoint := makeCountEndpoint(strService)
	countHandler := httptransport.NewServer(
		countEndpoint,
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8000", nil)
}
