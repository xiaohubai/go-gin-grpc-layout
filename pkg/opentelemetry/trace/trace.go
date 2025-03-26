package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// NewSpan creates a span using default tracer.
func NewSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return otel.Tracer("").Start(ctx, spanName, opts...)
}

func TraceID(ctx context.Context) string {
	if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
		return span.TraceID().String()
	}
	return ""
}

func SpanID(ctx context.Context) string {
	if span := trace.SpanContextFromContext(ctx); span.HasSpanID() {
		return span.SpanID().String()
	}
	return ""
}
