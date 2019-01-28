package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)

	var addrBookSrv AddressbookService
	addrBookSrv = addressbookService{
		logger: logger,
	}
	addrBookSrv = loggingMiddleware{logger: logger, next: addrBookSrv}

	lookForHandler := httptransport.NewServer(
		makeLookForEndpoint(addrBookSrv),
		decodeLookForRequest,
		encodeLookForRespose,
	)

	http.Handle("/lookfor", lookForHandler)
	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}
