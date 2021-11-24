package main

import (
	"gitlab.idc.xiaozhu.com/xz-go/common/orm"
	"gitlab.idc.xiaozhu.com/xz-go/common/redis"
	"net/http"
	"xz_gowork/pkg/global"
	"xz_gowork/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	global.Setup()
	orm.Setup()
	redis.Setup()
	middleware.Setup()

	// 3.监听端口，默认在8081
	// Run("里面不指定端口号默认为8081")
	r.Run(":8001")
}