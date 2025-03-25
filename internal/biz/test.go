package biz

import (
	"context"

	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/log"
)

func (uc *Usecase) Test(ctx context.Context, req *v1.TestRequest) (*v1.TestResponse, error) {
	resp := &v1.TestResponse{
		ID:      req.ID,
		Message: "test",
	}
	log.Info(ctx, "test success", log.AddField("message", resp.Message))
	return resp, nil
}
