package registry

import (
	"time"

	"grodyia/logger"
)

type Options struct {
	Logger  logger.Logger
	Timeout time.Duration
}
