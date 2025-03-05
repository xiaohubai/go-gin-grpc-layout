package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/pkg/code"
)

func (s *HTTPService) Login(ctx *gin.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	resp, err := s.biz.Login(ctx, req)
	if err != nil {
		return nil, code.WithError(code.LoginFailed, nil)
	}
	return resp, err

}
