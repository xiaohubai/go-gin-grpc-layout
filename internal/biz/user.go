package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/sagikazarmark/slog-shim"
	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
)

func (uc *Usecase) Login(c *gin.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {

	slog.Info("Login")
	return nil, nil
}
