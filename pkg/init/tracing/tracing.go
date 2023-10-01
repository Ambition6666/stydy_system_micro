package tracing

import (
	"io"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	con "github.com/uber/jaeger-client-go/config"
)

func InitTracer(service string, host string, port string) (opentracing.Tracer, io.Closer) {
	os.Setenv("JAEGER_AGENT_HOST", host)
	os.Setenv("JAEGER_AGENT_PORT", port)
	cfg, err := con.FromEnv()
	if err != nil {
		panic(err)
	}

	cfg.ServiceName = service
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter.LogSpans = true

	tracer, closer, err := cfg.NewTracer(con.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	return tracer, closer
}
