package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-gin-grpc-layout/internal/pkg/ecode"
)

type Body struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      any    `json:"data"`
	RequestID string `json:"requestId"`
}

func Result(c *gin.Context, code int, data any, err any) {
	if data == nil {
		data = make(map[string]string, 0)
	}
	resp := Body{
		Code:      code,
		Data:      data,
		Msg:       ecode.Msg[code],
		RequestID: "3224khhk3-324532jkbk",
	}

	if e, ok := err.(error); ok {
		resp.Msg = resp.Msg + ": " + e.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func Success(c *gin.Context, data any) {
	Result(c, 0, data, nil)
}

func Fail(c *gin.Context, code int, err any) {
	Result(c, code, nil, err)
}
