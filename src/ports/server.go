package ports

import (
	"github.com/YghoRibas/backend-Ramp/src/infra"
	"go.opentelemetry.io/otel/trace"
)

type ServerParams struct {
	Port int
	Tracer trace.Tracer
	Logger infra.Logger
}

type server struct {
	port int
	tracer trace.Tracer
	logger infra.Logger
}

func NewServer() {
	return 
}

func (s server) Run() err error {

}