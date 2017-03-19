package hotgoods

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   HotGoodsService
}

func (mw loggingMiddleware) GetHotGoodsInfo(s string) (out interface{}, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", out,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	out, err = mw.next.GetHotGoodsInfo(s)
	return
}

