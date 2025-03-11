package biz

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/log"
)

func (uc *Usecase) SSE(ctx *gin.Context, req *v1.SSERequest, msgChan *chan string) (*v1.SSEResponse, error) {
	resp := &v1.SSEResponse{}

	log.Info(ctx, "sse success")
	return resp, nil
}
