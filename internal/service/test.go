package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/utils/response"
)

func (s *HTTPService) Test(c *gin.Context) {
	data := "test success"
	response.Success(c, data)
}
