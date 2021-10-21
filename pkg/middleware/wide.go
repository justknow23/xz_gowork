package middleware

import (
	"github.com/gin-gonic/gin"
)

// WideMiddleware 全局通用中间件
var WideMiddleware gin.HandlerFunc

// initWide
func initWide() {
	WideMiddleware = CheckWide()
}

// Setup 初始化
func Setup() {
	initWide()
}

func CheckWide() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO write your logic
		// c is gin context
		// 获取header中的数据，可通过：c.GetHeader(key string)
		// 获取query中的数据，可通过：c.GetQuery(key string)
		// 获取form中的数据，可通过：c.GetPostForm(key string)
		// ....
		// gin.Context提供很多种方式去获取请求中的参数，在使用中可参考其方法签名

		// before handler
		// 触发下一个中间件函数，一直到实际的请求处理函数处理完成，因为路由请求handler其实也是一个遵循中间键定义的HandlerFunc
		c.Next()
		// after handler
	}
}
