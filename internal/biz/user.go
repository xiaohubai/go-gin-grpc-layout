package biz

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/log"
)

func (uc *Usecase) Login(c *gin.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {

	log.Info("Login")
	return nil, nil
}
