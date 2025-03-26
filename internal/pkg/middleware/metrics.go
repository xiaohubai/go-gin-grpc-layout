package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/opentelemetry/metric"
)

// Metrics returns a gin.HandlerFunc for exporting some Web metrics
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := config.GetConfig()
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		metric.ReqInc(c.Request.Context(), &metric.Labels{
			Env:      conf.App.Env,
			Service:  conf.App.Name,
			Protocol: conf.Server.HTTP.Name,
			Path:     path,
			Method:   method,
			Status:   c.Writer.Status(),
			Duration: time.Since(start).Seconds(),
		})
	}
}
