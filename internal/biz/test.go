package biz

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/log"
)

func (uc *Usecase) Test(ctx *gin.Context, req *v1.TestRequest) (*v1.TestResponse, error) {
	resp := &v1.TestResponse{
		ID:      req.ID,
		Message: "test",
	}
	log.Info(ctx, "test success", log.AddField("id", resp.ID), log.AddField("message", resp.Message))
	return resp, nil
}
