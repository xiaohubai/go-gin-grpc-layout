package main

import (
	"context"

	"github.com/xiaohubai/go-gin-grpc-layout/internal/conf"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c, err := config.InitConfig[conf.Config](ctx)
	if err != nil {
		panic(err)
	}

}
