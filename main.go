package main

import (
	"net/http"
	"os"

	"golang.org/x/net/context"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	decoder "github.com/mdstella/go-users/decoder"
	encoder "github.com/mdstella/go-users/encoder"
	stringEndpoint "github.com/mdstella/go-users/endpoint"
	middleware "github.com/mdstella/go-users/middleware"
	StringService "github.com/mdstella/go-users/service"
	stringService "github.com/mdstella/go-users/service/impl"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var svc StringService.StringService
	svc = stringService.StringServiceImpl{}
	svc = middleware.LoggingMiddleware{Logger: logger, Next: svc}
	svc = middleware.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, CountResult: countResult, Next: svc}

	ctx := context.Background()

	uppercaseHandler := httptransport.NewServer(
		ctx,
		stringEndpoint.MakeUppercaseEndpoint(svc),
		decoder.DecodeUppercaseRequest,
		encoder.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		ctx,
		stringEndpoint.MakeCountEndpoint(svc),
		decoder.DecodeCountRequest,
		encoder.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))

}
