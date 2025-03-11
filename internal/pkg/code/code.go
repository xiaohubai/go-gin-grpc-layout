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

// 内部错误
const (
	Success      = 0    // 成功
	Failed       = 4000 // 系统内部错误
	ParamsFailed = 4001 // 参数校验错误
)

// 业务错误
const (
	TestFailed = 5001 // 测试失败
	SSEFailed  = 5002 // sse失败
)

var msgMap = map[int]string{
	// 内部错误
	Success:      "成功",
	Failed:       "系统内部错误",
	ParamsFailed: "参数校验错误",

	// 业务错误
	TestFailed: "测试失败",
	SSEFailed:  "sse失败",
}
