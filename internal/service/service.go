package service

import (
	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/grpc/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/biz"
)

type HTTPService struct {
	biz *biz.Usecase
}

type GRPCService struct {
	biz *biz.Usecase
	v1.UnimplementedGRPCServer
}

func NewHTTPService() *HTTPService {
	return &HTTPService{
		biz: biz.NewUsecase(),
	}
}

func NewGRPCService() *GRPCService {
	return &GRPCService{
		biz: biz.NewUsecase(),
	}
}
