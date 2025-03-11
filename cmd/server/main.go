package main

import (
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/data"

	"github.com/xiaohubai/go-gin-grpc-layout/internal/server"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/service"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/log"

	_ "go.uber.org/automaxprocs"
)

var conf = flag.String("app_conf", "./configs/config.yaml", "app config file")

func newApp(c *config.Conf) *kratos.App {
	return kratos.New(
		kratos.ID(c.ID),
		kratos.Name(c.Name),
		kratos.Version(c.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			server.NewHTTPServer(&c.Server, service.NewHTTPService()),
			server.NewGRPCServer(&c.Server, service.NewGRPCService()),
		),
	)
}

func main() {
	cf, err := config.Read[config.Conf](*conf)
	if err != nil {
		panic(err)
	}

	if err := log.Init(&cf.Log); err != nil {
		panic(err)
	}

	if err := data.Init(&cf.Data); err != nil {
		panic(err)
	}

	if err := newApp(cf).Run(); err != nil {
		panic(err)
	}

}
