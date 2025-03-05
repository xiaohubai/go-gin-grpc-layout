package cron

import (
	"context"
	"flag"

	"github.com/xiaohubai/go-gin-grpc-layout/internal/service/cron"
)

var config = flag.String("conf", "./conf/app.toml", "app config file")

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cron.Start(ctx)

	select {}
}
