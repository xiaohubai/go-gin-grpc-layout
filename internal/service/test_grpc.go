package service

import (
	"context"

	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/grpc/v1"
	httpReq "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
)

func (s *GRPCService) Test(ctx context.Context, req *v1.TestRequest) (*v1.TestResponse, error) {
	resp, err := s.biz.Test(ctx, &httpReq.TestRequest{})
	if err != nil {
		return nil, err
	}

	return &v1.TestResponse{
		Id:      resp.Id,
		Message: resp.Message,
	}, nil

}
