package mallhome

import (
	"time"

	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next MallHomeService) MallHomeService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	MallHomeService
}

func (mw logmw) GetMallHomeInfo(ctx context.Context, request interface{}) (output interface{}, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "uppercase",
			"input", ctx,
			"output", output,
			"err", error,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.MallHomeService.GetMallHomeInfo(s)
	return
}


