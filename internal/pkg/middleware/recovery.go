package middleware

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/log"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]
				bufs := string(buf)

				log.Error(c, "panic", log.AddField("stack", bufs))
				c.Abort()
			}
		}()
		c.Next()
	}
}
