package http
import (
	"github.com/gin-gonic/gin"
	"gitlab.idc.xiaozhu.com/xz-go/common/errors"
	"net/http"
	"time"
)

const (
	defaultErrCode = 4000
	defaultErrMsg  = "server internal error"
)

type response struct {
	ctx       *gin.Context
	code      int
	Status    int         `json:"status"`
	Content   interface{} `json:"content"`
	ErrorMsg  string      `json:"errorMsg"`
	Timestamp int64       `json:"timestamp"`
}

func Success(c *gin.Context, data interface{}) {
	newResponse(c, http.StatusOK, http.StatusOK, data).render()
}

func Error(c *gin.Context, httpCode int, err error) {
	if ce, ok := err.(errors.CustomError); ok {
		errCode := ce.Code()
		errMsg := ce.Error()
		newResponse(c, httpCode, errCode, errMsg).render()
		return
	}

	newResponse(c, httpCode, defaultErrCode, defaultErrMsg).render()
}

func newResponse(ctx *gin.Context, httpCode, status int, data interface{}) response {
	var errMsg string
	if status != http.StatusOK {
		errMsg = data.(string)
		data = nil
	}
	return response{
		ctx:       ctx,
		code:      httpCode,
		Status:    status,
		ErrorMsg:  errMsg,
		Content: data,
		Timestamp: time.Now().Unix(),
	}
}

func (r response) render() {
	r.ctx.JSON(r.code, r)
}
