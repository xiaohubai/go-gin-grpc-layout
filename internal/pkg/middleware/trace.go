package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
	"go.opentelemetry.io/contrib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
	"go.opentelemetry.io/otel/trace"
)

type traceConf struct {
	TracerProvider trace.TracerProvider
	Propagators    propagation.TextMapPropagator
}

// Option specifies instrumentation configuration options.
type Option func(*traceConf)

const (
	tracerKey  = "otel-tracer"
	tracerName = "otelgin"
)

func Trace(opts ...Option) gin.HandlerFunc {
	cfg := traceConf{}
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.TracerProvider == nil {
		cfg.TracerProvider = otel.GetTracerProvider()
	}
	tracer := cfg.TracerProvider.Tracer(
		tracerName,
		trace.WithInstrumentationVersion(contrib.Version()),
	)
	if cfg.Propagators == nil {
		cfg.Propagators = otel.GetTextMapPropagator()
	}
	return func(c *gin.Context) {
		c.Set(tracerKey, tracer)
		savedCtx := c.Request.Context()
		defer func() {
			c.Request = c.Request.WithContext(savedCtx)
		}()
		ctx := cfg.Propagators.Extract(savedCtx, propagation.HeaderCarrier(c.Request.Header))
		route := c.FullPath()
		opts := []trace.SpanStartOption{
			trace.WithAttributes(
				semconv.HostType(config.GetConfig().Server.HTTP.Name),
				semconv.HostIP(c.ClientIP()),
				semconv.HTTPRequestMethodOriginal(c.Request.Method),
				semconv.HTTPRoute(route),
			),
			trace.WithSpanKind(trace.SpanKindServer),
		}
		spanName := route
		if spanName == "" {
			spanName = fmt.Sprintf("HTTP %s route not found", c.Request.Method)
		}
		ctx, span := tracer.Start(ctx, spanName, opts...)
		defer span.End()

		c.Request = c.Request.WithContext(ctx)

		c.Next()

		status := c.Writer.Status()
		if status > 0 {
			span.SetAttributes(
				semconv.HTTPResponseStatusCode(status),
			)
		}
		if len(c.Errors) > 0 {
			span.SetAttributes(attribute.String("gin.errors", c.Errors.String()))
		}
	}
}
