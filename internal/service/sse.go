package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/pkg/code"
)

func (s *HTTPService) SSE(ctx *gin.Context, req *v1.SSERequest) (*v1.SSEResponse, error) {
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")

	msgChan := make(chan string)
	defer close(msgChan)

	_, err := s.biz.SSE(ctx, req, &msgChan)
	if err != nil {
		return nil, code.WithError(code.SSEFailed, nil)
	}

	resp := &v1.SSEResponse{}

	for {
		select {
		case msg := <-msgChan:
			ctx.SSEvent("message", msg)
			ctx.Writer.Flush()
		case <-ctx.Request.Context().Done():
			return resp, nil
		}
	}
}
