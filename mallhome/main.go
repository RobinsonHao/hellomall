package mallhome

import (
	"context"
	//"flag"
	"net/http"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "hellomall",
		Subsystem: "mallhome_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "hellomaill",
		Subsystem: "mallhome_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "hellomaill",
		Subsystem: "mallhome_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	var svc MallHomeService
	svc = mallHomeService{}
	svc = proxyingMiddleware(context.Background(), logger)(svc)
	svc = loggingMiddleware(logger)(svc)
	svc = instrumentingMiddleware(requestCount, requestLatency, countResult)(svc)

	mallHomeHandler := httptransport.NewServer(
		makeMallHomeEndpoint(svc),
		decodeMallHomeRequest,
		encodeResponse,
	)
	

	http.Handle("/uppercase", mallHomeHandler)
	http.Handle("/metrics", stdprometheus.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080" ,nil))
}
