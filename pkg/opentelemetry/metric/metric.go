package metric

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type Labels struct {
	Env      string
	Service  string
	Protocol string
	Path     string
	Method   string
	Status   int
	Duration float64
}

func ReqInc(ctx context.Context, labels *Labels) {
	counter, _ := otel.Meter("request_counter").Int64Counter(
		fmt.Sprintf("%s_request_total", labels.Protocol),
		metric.WithDescription("请求总数"),
		metric.WithUnit("1"),
	)
	counter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("env", labels.Env),
			attribute.String("service", labels.Service),
			attribute.String("protocol", labels.Protocol),
			attribute.String("path", labels.Path),
			attribute.String("method", labels.Method),
			attribute.Int("status", labels.Status),
		),
	)
}

func Inc(ctx context.Context, key string) {
	counter, _ := otel.Meter("key_counter").Int64Counter(
		key,
		metric.WithDescription("请求总数"),
		metric.WithUnit("1"),
	)
	counter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("key", key),
		),
	)
}
