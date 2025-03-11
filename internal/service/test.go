package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/pkg/code"
)

func (s *HTTPService) Test(ctx *gin.Context, req *v1.TestRequest) (*v1.TestResponse, error) {
	resp, err := s.biz.Test(ctx, req)
	if err != nil {
		return nil, code.WithError(code.TestFailed, nil)
	}
	return resp, err
}


