package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/spf13/viper"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/pkg/conf"
)

func NewHTTPServer(c *conf.Server) *http.Server {
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
	return srv
}
