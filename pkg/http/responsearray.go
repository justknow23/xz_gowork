package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type responsearray struct {
	ctx       *gin.Context
	code      int
	Status    int           `json:"status"`
	Content   []interface{} `json:"content"`
	ErrorMsg  string        `json:"errorMsg"`
	Timestamp int64         `json:"timestamp"`
}

func Define(c *gin.Context, data []interface{}) {
	newResponseArray(c, http.StatusOK, http.StatusOK, data).render()
}

func newResponseArray(ctx *gin.Context, httpCode, status int, data []interface{}) responsearray {
	var errMsg string
	return responsearray{
		ctx:       ctx,
		code:      httpCode,
		Status:    status,
		ErrorMsg:  errMsg,
		Content:   data,
		Timestamp: time.Now().Unix(),
	}
}

func (r responsearray) render() {
	r.ctx.JSON(r.code, r)
}
