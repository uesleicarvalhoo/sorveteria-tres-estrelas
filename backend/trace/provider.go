package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

var serviceName string //nolint: gochecknoglobals

type ProviderConfig struct {
	Endpoint       string
	ServiceName    string
	ServiceVersion string
	Environment    string
	Disabled       bool
}

type Provider struct {
	Provider trace.TracerProvider
}

func NewProvider(config ProviderConfig) (Provider, error) {
	serviceName = config.ServiceName

	if config.Disabled {
		return Provider{Provider: trace.NewNoopTracerProvider()}, nil
	}

	exp, err := zipkin.New(config.Endpoint)
	if err != nil {
		return Provider{}, err
	}

	prv := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(config.ServiceName),
				semconv.ServiceVersionKey.String(config.ServiceVersion),
				attribute.String("environment", config.Environment),
			),
		),
	)

	otel.SetTracerProvider(prv)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return Provider{Provider: prv}, nil
}

func (p *Provider) Close(ctx context.Context) error {
	if prov, ok := p.Provider.(*tracesdk.TracerProvider); ok {
		return prov.Shutdown(ctx)
	}

	return nil
}
