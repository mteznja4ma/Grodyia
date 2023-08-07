package registry

import "time"

type Watcher interface {
	Next() (*Result, error)
	Stop()
}

type Result struct {
	Service *Service
	Action  string
}

type EventType int

const (
	Create EventType = iota
	Delete
	Update
)

func (e EventType) String() string {
	switch e {
	case Create:
		return "create"
	case Delete:
		return "delete"
	case Update:
		return "update"
	default:
		return "unknown"
	}
}

type Event struct {
	// Timestamp of the event
	Timestamp time.Time
	// Type of event
	Type EventType
	// Service that was created, deleted, updated
	Service *Service
	// Id is registry node id
	Id string
}
