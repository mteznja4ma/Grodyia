package registry

type Registry interface {
	Init(...Option) error
	Options() Options
}

type Service struct {
	Id        string
	Name      string
	Version   string
	Metadata  map[string]string
	Endpoints []*Endpoint
}

type Endpoint struct {
	Name     string
	Metadata map[string]string
}

type Option func(*Options)
