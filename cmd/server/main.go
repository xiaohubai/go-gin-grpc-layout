package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/data"

	"github.com/xiaohubai/go-gin-grpc-layout/internal/server"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/service"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/log"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/opentelemetry"

	_ "go.uber.org/automaxprocs"
)

func newApp(c *config.Conf) *kratos.App {
	return kratos.New(
		kratos.ID(c.App.ID),
		kratos.Name(c.App.Name),
		kratos.Version(c.App.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			server.NewHTTPServer(&c.Server, service.NewHTTPService()),
			server.NewGRPCServer(&c.Server, service.NewGRPCService()),
		),
	)
}

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	conf := config.GetConfig()
	if err := log.Init(&conf.Log); err != nil {
		panic(err)
	}

	if err := opentelemetry.Init(); err != nil {
		panic(err)
	}

	if err := data.Init(&conf.Data); err != nil {
		panic(err)
	}

	if err := newApp(conf).Run(); err != nil {
		panic(err)
	}

}
