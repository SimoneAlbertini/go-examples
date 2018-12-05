package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

// ** middleware function, do not log request params, better tu use a middleware struct
// ** In order to use it, decorate the handler init in main func like this:
// ** uppercaseEndpoint = loggingMiddleware(log.With(logger, "method", "uppercase"))(uppercaseEndpoint)
// ** countEndpoint = loggingMiddleware(log.With(logger, "method", "count"))(countEndpoint)

// func loggingMiddleware(logger log.Logger) endpoint.Middleware {
// 	return func(next endpoint.Endpoint) endpoint.Endpoint {
// 		return func(ctx context.Context, request interface{}) (interface{}, error) {
// 			logger.Log("msg", "calling endpoint")
// 			defer logger.Log("msg", "calling endpoint")
// 			return next(ctx, request)
// 		}
// 	}
// }

// It implements the StringService interface!
type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.next.Count(s)
	return
}
