package gh

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/opentelemetry/trace"
)

type CodeMsg struct {
	Code int
	Msg  string
}

func (w *CodeMsg) Error() string {
	return fmt.Sprintf("%#+v", w)
}

type Body struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
	TraceID string `json:"traceId"`
}

func result(c *gin.Context, code int, data any, msg string) {
	if data == nil {
		data = make(map[string]string, 0)
	}
	resp := Body{
		Code:    code,
		Data:    data,
		Msg:     msg,
		TraceID: trace.TraceID(c.Request.Context()),
	}
	c.JSON(http.StatusOK, resp)
}

func success(ctx *gin.Context, data any) {
	result(ctx, 0, data, "success")
}

func fail(ctx *gin.Context, codeMsg CodeMsg) {
	result(ctx, codeMsg.Code, nil, codeMsg.Msg)
}
