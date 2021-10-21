package lormg

import (
	"context"
	"strings"

	"insurance/pkg/global"
	"insurance/pkg/log"

	"gitlab.idc.xiaozhu.com/xz-go/common/orm"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// FromContext 获取orm.DB
func FromContext(ctx context.Context) *gorm.DB {
	tx := orm.FromContext(ctx)
	tx.Logger = NewOrmLogger(ctx)
	return tx
}

// NewOrmLogger -
func NewOrmLogger(ctx context.Context) logger.Interface {
	olw := &ormLoggerWriter{
		ctx: ctx,
	}
	return logger.New(olw, logger.Config{}).LogMode(parseSQLLogLevel(global.Settings.LogLevel))
}

// parseSQLLogLevel takes a string level and returns the Logrus log level constant.
func parseSQLLogLevel(lvl string)  logger.LogLevel {
	switch strings.ToLower(lvl) {
	case "panic":
		return logger.Silent
	case "fatal":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn", "warning":
		return logger.Warn
	case "info":
		return logger.Warn
	case "debug":
		return logger.Info
	default:
		return logger.Info
	}
}

type ormLoggerWriter struct {
	ctx context.Context
}

func (olw *ormLoggerWriter) Printf(format string, args ...interface{}) {
	log.GetLogger(olw.ctx, "gorm").Infof(format, args...)
}
