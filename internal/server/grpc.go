package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/spf13/viper"

	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/grpc/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/service"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
)

func NewGRPCServer(c *config.Server, sg *service.GRPCService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.GRPC.Network != "" {
		opts = append(opts, grpc.Network(c.GRPC.Network))
	}
	if c.GRPC.Addr != "" {
		opts = append(opts, grpc.Address(c.GRPC.Addr))
	}
	if c.GRPC.Timeout != "" {
		opts = append(opts, grpc.Timeout(viper.GetDuration(c.GRPC.Timeout)))
	}
	srv := grpc.NewServer(opts...)
	return registerGRPCServer(srv, sg)
}

// registerGRPCServer 注册grpc服务组
func registerGRPCServer(srv *grpc.Server, sg *service.GRPCService) *grpc.Server {
	v1.RegisterTestGRPCServer(srv, sg)
	v1.RegisterUserGRPCServer(srv, sg)

	return srv
}
