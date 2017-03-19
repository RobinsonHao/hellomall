package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           UserBaseInfoService
}

func (mw instrumentingMiddleware) GetUserBaseInfo(s string) (output interface{}, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetUserBaseInfo", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.GetUserBaseInfo(s)
	return
}


