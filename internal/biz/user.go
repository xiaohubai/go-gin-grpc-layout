package biz

import (
	"fmt"

	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/log"
)

func (uc *Usecase) Login(ctx *gin.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	user, err := uc.db.NewUserDB().FindByConds(ctx, map[string]any{
		"user_name": req.UserName,
	}, nil)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	resp := &v1.LoginResponse{
		UserName: user.UserName,
	}
	log.Info("login success", log.AddField("id", user.ID))
	return resp, nil
}
