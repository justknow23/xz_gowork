package errors

import (
	"gitlab.idc.xiaozhu.com/xz-go/common/errors"
)

var (
	// ErrorServerInternal -
	ErrorServerInternal = errors.ServerError
	// ErrorData 数据错误
	ErrorData = errors.NewError(50001, "数据错误")
	// ErrorDataNotFound -
	ErrorDataNotFound = errors.NewError(50404, "数据不存在")
)
