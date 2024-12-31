package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/spf13/viper"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/pkg/conf"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/service"
)

func NewHTTPServer(c *conf.Server, sh *service.HTTPService) *http.Server {
	var opts = []http.ServerOption{}
	if c.HTTP.Network != "" {
		opts = append(opts, http.Network(c.HTTP.Network))
	}
	if c.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.HTTP.Addr))
	}
	if c.HTTP.Timeout != "" {
		opts = append(opts, http.Timeout(viper.GetDuration(c.HTTP.Timeout)))
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", routers(sh))
	return srv
}
func routers(s *service.HTTPService) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	r1 := router.Group("")
	{
		r1.GET("/test", s.Test)       //登录
		r1.POST("/v1/login", s.Login) //登录
	}

	return router
}
