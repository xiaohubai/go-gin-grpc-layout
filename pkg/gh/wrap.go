// gin http
package gh

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Wrap 泛型函数，封装请求处理逻辑，自动解析参数并处理错误
func Wrap[T any, R any](run func(c *gin.Context, req *T) (R, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqModel T
		var err error
		contentType := c.ContentType()

		switch c.ContentType() {
		// application/json
		case gin.MIMEJSON:
			err = c.ShouldBindJSON(&reqModel)
		// application/x-www-form-urlencoded
		case gin.MIMEPOSTForm:
			err = c.ShouldBind(&reqModel)
		// multipart/form-data
		case gin.MIMEMultipartPOSTForm:
			err = c.ShouldBind(&reqModel)
		default:
			err := fmt.Errorf("unsupported content type: %s", contentType)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// 参数解析/校验
		if err != nil || Validate(&reqModel) != nil {
			fail(c, CodeMsg{Code: 4001, Msg: fmt.Sprintf("invalid params: %v", err)})
			return
		}

		// 执行业务逻辑
		result, err := run(c, &reqModel)
		if err == nil {
			success(c, result)
			return
		}

		// 业务逻辑错误处理
		if v, ok := err.(*CodeMsg); ok {
			fail(c, CodeMsg{Code: v.Code, Msg: v.Msg})
			return
		}

		fail(c, CodeMsg{Code: 4000, Msg: fmt.Sprintf("err: %v", err)})
	}
}
