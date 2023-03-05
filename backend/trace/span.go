package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func NewSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return otel.Tracer(serviceName).Start(ctx, name)
}

func AddSpanTags(span trace.Span, tags map[string]string) {
	list := []attribute.KeyValue{}

	for k, v := range tags {
		list = append(list, attribute.Key(k).String(v))
	}

	span.SetAttributes(list...)
}

func AddSpanEvents(span trace.Span, name string, events map[string]string) {
	list := []trace.EventOption{}

	for k, v := range events {
		list = append(list, trace.WithAttributes(attribute.Key(k).String(v)))
	}

	span.AddEvent(name, list...)
}

func AddSpanError(span trace.Span, err error) {
	span.RecordError(err)
}

func FailSpan(span trace.Span, message string) {
	span.SetStatus(codes.Error, message)
}

func SpanFromContext(ctx context.Context) trace.Span {
	return trace.SpanFromContext(ctx)
}
