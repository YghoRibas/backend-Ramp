package ports

import (
	"fmt"
	"context"

	"github.com/YghoRibas/backend-Ramp/src/infra"
	"go.opentelemetry.io/otel/trace"
)

type ServerParams struct {
	Port   int
	Tracer trace.Tracer
	Logger infra.Logger
}

type server struct {
	port   int
	tracer trace.Tracer
	logger infra.Logger
}

func NewServer(params ServerParams) (server, error) {
	if params.Port <= 80 {
		params.Port = 80
	}

	if params.Tracer == nil {
		return Server{}, errors.NewMissingRequiredDependency("Tracer")
	}

	if params.Logger == nil {
		return Server{}, errors.NewMissingRequiredDependency("Logger")
	}

	return server {
		port: params.Port
		tracer: params.Tracer
		logger: params.Logger
	}
}

func (s server) Run() err error {
	
}

func MustNewServer(params ServerParams) Server {
	srv, err := NewServer(params)
	if err != nil {
		panic(err)
	}

	return srv
}