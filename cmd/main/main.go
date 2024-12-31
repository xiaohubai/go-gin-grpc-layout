package main

import (
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/pkg/conf"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/server"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/service"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/log"

	_ "go.uber.org/automaxprocs"
)

var cfg = flag.String("app_conf", "./configs/config.yaml", "app config file")

func main() {
	c, err := config.Read[conf.Conf](*cfg)
	if err != nil {
		panic(err)
	}

	if err := log.New(c); err != nil {
		panic(err)
	}

	/* if err := data.Init(&c.Data); err != nil {
		panic(err)
	} */

	app := kratos.New(
		kratos.ID(c.ID),
		kratos.Name(c.Name),
		kratos.Version(c.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			server.NewHTTPServer(&c.Server, service.NewHTTPService()),
			server.NewGRPCServer(&c.Server, service.NewGRPCService()),
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}

}
