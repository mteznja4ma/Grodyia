package server

import "google.golang.org/grpc"

type grpcServer struct {
	server *grpc.Server
}

func NewGRPCServer(opts ...Option) Server {
	options := Options{
		Name: DefaultName,
	}

	for _, o := range opts {
		o(&options)
	}

	return &grpcServer{
		server: grpc.NewServer(),
	}
}

func (s *grpcServer) Init(opts ...Option) error {
	return nil
}

func (s *grpcServer) Options() Options {
	return Options{}
}

func (s *grpcServer) Start() error {
	return nil
}

func (s *grpcServer) Stop() error {
	return nil
}
