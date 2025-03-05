package code

import (
	"fmt"

	"github.com/xiaohubai/go-gin-grpc-layout/pkg/gh"
)

func WithError(code int, err error) error {
	msg, ok := msgMap[code]
	if !ok {
		msg = "未知错误"
	}
	if err != nil {
		return &gh.CodeMsg{Code: code, Msg: fmt.Sprintf("%s: %s", msg, err.Error())}
	}
	return &gh.CodeMsg{Code: code, Msg: msg}
}

const (
	Success      = 0    // 成功
	Failed       = 4000 // 系统内部错误
	ParamsFailed = 4002 // 参数校验错误
	LoginFailed  = 4003 // 登录失败
)

var msgMap = map[int]string{
	Success:      "成功",
	Failed:       "系统内部错误",
	ParamsFailed: "参数校验错误",
	LoginFailed:  "登录失败",
}
