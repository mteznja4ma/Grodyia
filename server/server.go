package server

import (
	"github.com/google/uuid"
)

var (
	DefaultAddress        = ":0"
	DefaultName           = "grodyia"
	DefaultVersion        = "latest"
	DefaultId             = uuid.New().String()
	DefaultServer  Server = NewGRPCServer()
)

type Option func(*Options)

type Server interface {
	// Intialize Options
	Init(opts ...Option) error
	// Retrieve Options
	Options() Options
	// Start the Server
	Start() error
	// Stop the Server
	Stop() error
}
