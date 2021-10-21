package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"gitlab.idc.xiaozhu.com/xz-go/common/server"
)

const (
	// nameField nameField
	nameField = "name"
)

// newLogger -
func newLogger(name ...string) *logrus.Entry {
	if len(name) == 0 {
		return logrus.New().WithFields(nil)
	}
	return logrus.New().WithField(nameField, name[0])
}

// GetLogger -
func GetLogger(ctx context.Context, name ...string) *logrus.Entry {
	if srv := server.FromContext(ctx); srv != nil {
		if len(name) == 0 {
			return srv.Log().WithFields(nil)
		}
		return srv.Log().WithField(nameField, name[0])
	}
	return newLogger(name...)
}
