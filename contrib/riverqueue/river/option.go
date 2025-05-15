package river

import (
	"github.com/DataDog/dd-trace-go/v2/ddtrace/tracer"
	"github.com/DataDog/dd-trace-go/v2/instrumentation"
)

type config struct {
	service                  string
	insertSpanName           string
	workSpanName             string
	measured                 bool
	spanOpts                 []tracer.StartSpanOption
	enableDistributedTracing bool
}

func defaultConfig() *config {
	return &config{
		service:                  instr.ServiceName(instrumentation.ComponentConsumer, nil),
		insertSpanName:           instr.OperationName(instrumentation.ComponentProducer, nil),
		workSpanName:             instr.OperationName(instrumentation.ComponentConsumer, nil),
		measured:                 false,
		enableDistributedTracing: false,
	}
}

// Option is used to customize spans started by InsertMiddleware or WorkerMiddleware.
type Option func(cfg *config)

// WithService sets the service name tag for traces started by InsertMiddleware or WorkerMiddleware.
func WithService(service string) Option {
	return func(cfg *config) {
		cfg.service = service
	}
}

// WithDistributedTracing enables distributed tracing if set to true. Default is false.
func WithDistributedTracing(enable bool) Option {
	return func(cfg *config) {
		cfg.enableDistributedTracing = enable
	}
}

// WithMeasured sets the measured tag for traces started by InsertMiddleware or WorkerMiddleware.
func WithMeasured() Option {
	return func(cfg *config) {
		cfg.measured = true
	}
}
