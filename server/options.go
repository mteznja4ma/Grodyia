package server

import "grodyia/logger"

type Options struct {
	Logger logger.Logger

	Name    string
	Id      string
	Version string
	Address string
}
