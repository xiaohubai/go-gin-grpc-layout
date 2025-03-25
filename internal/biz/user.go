package biz

import (
	"context"

	v1 "github.com/xiaohubai/go-gin-grpc-layout/api/http/v1"
)

func (uc *Usecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	resp := &v1.LoginResponse{}

	return resp, nil
}

func (uc *Usecase) UserInfo(ctx context.Context, req *v1.UserInfoRequest) (*v1.UserInfoResponse, error) {
	user, err := uc.db.NewUserDB().FindByConds(ctx, map[string]any{
		"username": req.UserName,
	}, []string{"Role"})
	if err != nil {
		return nil, err
	}
	resp := &v1.UserInfoResponse{
		UserName: user.Username,
		RoleId:   user.RoleID,
		RoleName: user.Role.RoleName,
	}
	return resp, nil
}
