package opentelemetry

import (
	"context"

	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func Init() error {
	conf := config.GetConfig()
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(conf.App.Name),
		semconv.ServiceVersion(conf.App.Version),
		semconv.DeploymentEnvironmentName(conf.App.Env),
	)

	ctx := context.Background()

	// Trace
	jaeger, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(conf.OpenTelemetry.Jaeger.Endpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return err
	}
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(jaeger),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Metric
	prometheus, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(conf.OpenTelemetry.Prometheus.Endpoint),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	metricProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(prometheus)),
		metric.WithResource(res),
	)
	otel.SetMeterProvider(metricProvider)

	defer func() {
		_ = tracerProvider.Shutdown(ctx)
		_ = metricProvider.Shutdown(ctx)
	}()

	return nil
}
