package ecode

const (
	Success      = 0    // 成功
	Failed       = 4000 // 系统内部错误
	ParamsFailed = 4002 // 参数校验错误
	LoginFailed  = 4003 // 登录失败
)

var Msg = map[int]string{
	Success:      "成功",
	Failed:       "系统内部错误",
	ParamsFailed: "参数校验错误",
	LoginFailed:  "登录失败",
}
