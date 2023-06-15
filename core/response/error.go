package response

import (
	"github.com/gin-gonic/gin"
)

// NoRoute 无路由的响应
func NoRoute(c *gin.Context) {
	Fail(c, RespType{code: 400, msg: "请求接口不存在"})
}

// NoMethod 无方法的响应
func NoMethod(c *gin.Context) {
	Fail(c, RespType{code: 400, msg: "请求方法不允许"})
}
