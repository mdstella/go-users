package middleware

import (
	"time"

	"github.com/go-kit/kit/log"

	stringService "github.com/mdstella/go-users/service"
)

//LoggingMiddleware - the logging middleware class
type LoggingMiddleware struct {
	Logger log.Logger
	Next   stringService.StringService
}

//Uppercase - this logs uppercase method data
func (mw LoggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method:", "uppercase",
			"input:", s,
			"output:", output,
			"err:", err,
			"took:", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Uppercase(s)
	return
}

//Count - this logs count method data
func (mw LoggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method:", "count",
			"input:", s,
			"n:", n,
			"took:", time.Since(begin),
		)
	}(time.Now())

	n = mw.Next.Count(s)
	return
}
