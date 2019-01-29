package main

import (
	"go-examples/addressbook/storage"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)

	var addressBook AddressbookService
	addressBook = addressbookService{
		logger:  logger,
		storage: storage.NewRedisStorage("tcp", "redis", "6379"),
	}
	addressBook = loggingMiddleware{logger: logger, next: addressBook}

	lookForHandler := httptransport.NewServer(
		makeLookForEndpoint(addressBook),
		decodeLookForRequest,
		encodeLookForRespose,
	)

	http.Handle("/lookfor", lookForHandler)
	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}
