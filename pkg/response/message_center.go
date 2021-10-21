package response

import (
	"github.com/gin-gonic/gin"
	"gitlab.idc.xiaozhu.com/xz-go/common/util/app"
	"insurance/pkg/global"
)

// ErrorMessageCenter 错误响应 消息中心 拒绝并重试
func ErrorMessageCenter(c *gin.Context, err error, codes ...int) {
	httpCode := 200
	if len(codes) > 0 {
		httpCode = codes[0]
	}
	c.Header(global.ConsumeResultHeaderName, global.ConsumeResultHeaderReject)
	app.Error(c, httpCode, err)
}

// SuccessMessageCenter 消息中心成功返回
func SuccessMessageCenter(c *gin.Context, data interface{}) {
	c.Header(global.ConsumeResultHeaderName, global.ConsumeResultHeaderSuccess)
	app.Success(c, data)
}
