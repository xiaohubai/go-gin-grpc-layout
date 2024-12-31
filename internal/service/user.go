package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/pkg/ecode"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/utils/response"
)

func (s *HTTPService) Login(c *gin.Context) {
	req := &v1.LoginRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}

	data, err := s.biz.Login(c, req)
	if err != nil {
		response.Fail(c, ecode.LoginFailed, err)
		return
	}
	response.Success(c, data)
}
