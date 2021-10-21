package errors

import (
	"fmt"
	"gitlab.idc.xiaozhu.com/xz-go/common/errors"
)

var (
	// ErrorServerInternal -
	ErrorServerInternal      = errors.ServerError
	NotFound                 = errors.NewError(1000001, "数据未找到")
	ErrorIdempotent          = errors.NewError(1001001, "幂等校验不同过")
	ErrorRepeat              = errors.NewError(1001002, "禁止重复消费")
	ErrorBeyondLimit         = errors.NewError(1002001, "超出限制")
	ErrorNotFoundPhoneRecord = errors.NewError(1002002, "通话信息不存在")
	ErrorBeyondLimitWarp     = func(str string) *errors.Error {
		return errors.Wrap(ErrorBeyondLimit, str)
	}
	NotFoundIdWarp = func(str string) *errors.Error {
		return errors.Wrap(NotFound, fmt.Sprintf("ID %s", str))
	}
	ErrorNotFoundWrap = func(str string) *errors.Error {
		return errors.Wrap(NotFound, str)
	}
	ErrorEmpty     = errors.NewError(1000050, "缺少必要参数")
	ErrorEmptyWrap = func(str string) *errors.Error {
		return errors.Wrap(ErrorEmpty, str)
	}

	// 业务级别
	ErrorFieldValue     = errors.NewError(3008, "错误的字段值")
	ErrorFieldValueWrap = func(str string) *errors.Error {
		return errors.Wrap(ErrorFieldValue, str)
	}
	ErrorNotLandlord     = errors.NewError(1001, "")
	ErrorNotLandlordWrap = func(str string) *errors.Error {
		return errors.Wrap(ErrorNotLandlord, str)
	}
	ErrorNotVisit     = errors.NewError(1002, "")
	ErrorNotVisitWrap = func(str string) *errors.Error {
		return errors.Wrap(ErrorNotVisit, str)
	}
)
