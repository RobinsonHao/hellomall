package mallhome

import (
	"errors"
	"strings"
	"context"
	"net/url"
	"time"
	jujuratelimit "github.com/juju/ratelimit"
	"github.com/sony/gobreaker"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	//"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MallHomeService provides operations on strings.
type MallHomeService interface {
	GetMallHomeInfo(ctx context.Context, request interface{}) (interface{}, error)
}

type mallHomeService struct{}

func (mallHomeService) GetMallHomeInfo(ctx context.Context, request interface{}) (error) {
	go func() {
		<-getUserBaseInfo(ctx, request)
		<-getHotGoods(ctx, request)
	}()
	return  nil
}


func getUserBaseInfo(ctx context.Context, request interface{}) (interface{}, error) {
	var (
		qps         = 100                    // beyond which we will return an error
		maxAttempts = 3                      // per request, before giving up
		maxTime     = 250 * time.Millisecond // wallclock time, before giving up
	)


	var (
		sbUserBase   sd.FixedSubscriber
	)

	var eUserBaseInfo endpoint.Endpoint
	userBaseInstance := "localhost:8001"
	eUserBaseInfo = makeGetUserInfoProxy(ctx, userBaseInstance)
	eUserBaseInfo = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(eUserBaseInfo)
	eUserBaseInfo = ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(float64(qps), int64(qps)))(eUserBaseInfo)
    sbUserBase = append(sbUserBase, eUserBaseInfo)
	balancer := lb.NewRoundRobin(sbUserBase)
	retry := lb.Retry(maxAttempts, maxTime, balancer)
	return eUserBaseInfo
}




func makeGetUserInfoProxy(ctx context.Context, instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	if u.Path == "" {
		u.Path = "/getUseInfo"
	}
	return httptransport.NewClient(
		"GET",
		u,
		encodeRequest,
		decodeGetUserInfoResponse,
	).Endpoint()
}



func getHotGoods(ctx context.Context, request interface{}) (interface{}, error) {
	var (
		qps         = 100                    // beyond which we will return an error
		maxAttempts = 3                      // per request, before giving up
		maxTime     = 250 * time.Millisecond // wallclock time, before giving up
	)

	var (
		sbHotGoods   sd.FixedSubscriber
	)

	var eHotGoods endpoint.Endpoint
	hotGoodsInstance := "localhost:8001"
	eHotGoods = makeGetUserInfoProxy(ctx, hotGoodsInstance)
	eHotGoods = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(eHotGoods)
	eHotGoods = ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(float64(qps), int64(qps)))(eHotGoods)
    sbHotGoods = append(sbUserBase, eHotGoods)
	balancer := lb.NewRoundRobin(sbHotGoods)
	retry := lb.Retry(maxAttempts, maxTime, balancer)
	return eHotGoods
}




func makeGetHotGoodsProxy(ctx context.Context, instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	if u.Path == "" {
		u.Path = "/getHotGoods"
	}
	return httptransport.NewClient(
		"GET",
		u,
		encodeRequest,
		decodeGetHotGoodsResponse,
	).Endpoint()
}


// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for StringService.
type ServiceMiddleware func(MallHomeService) MallHomeService
