package telemetry

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

type TracerConfig struct {
	Enable        bool    `env:"TRACER_ENABLE" env-default:"true"`
	AgentEndpoint string  `env:"TRACER_AGENT_ENDPOINT" env-default:"localhost"`
	AgentPort     string  `env:"TRACER_AGENT_PORT" env-default:"6831"`
	SamplingRatio float64 `env:"TRACER_SAMPLING_RATIO" env-default:"1"` // configures the fraction of tracers sampled
}

func InitTracer(ctx context.Context, cfg TracerConfig, serviceName string) (trace.Tracer, error) {
	if !cfg.Enable {
		return otel.Tracer(serviceName), nil // defaults to no-op tracer
	}
	
	// Creates OTel Exporter implementation that exports the collected spans to Jaeger
	exp, err := createJaegerExporter(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create jaeger exporter: %w", err)
	}

	// Create a new tracer provider that exports to the configured exporter
	tracerProvider := createTracerProvider(serviceName, exp, cfg.SamplingRatio)

	// Registers the tracer provider as the global tracer provider
	otel.SetTracerProvider(tracerProvider)

	return otel.Tracer(serviceName), nil
}

func createJaegerExporter(c TracerConfig) (*jaeger.Exporter, error) {
	return jaeger.New(
		jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost(c.AgentEndpoint),
			jaeger.WithAgentPort(c.AgentPort),
		),
	)
}

func createTracerProvider(serviceName string, exp *jaeger.Exporter, samplingRatio float64) trace.TracerProvider {
	return tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(serviceName))),
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(samplingRatio))),
	)
}
