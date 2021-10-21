package errors

import (
	stdErrors "errors"
	pkgErrors "github.com/pkg/errors"
	"gitlab.idc.xiaozhu.com/xz-go/common/errors"
)

// Is stdErrors.Is
func Is(err, target error) bool { return stdErrors.Is(err, target) }

// As stdErrors.As
func As(err error, target interface{}) bool { return stdErrors.As(err, &target) }

// Unwrap stdErrors.Unwrap
func Unwrap(err error) error { return stdErrors.Unwrap(err) }

// New pkgErrors.New
func New(message string) error { return pkgErrors.New(message) }

// Trace common_errors.Trace
func Trace(err error) string { return errors.Trace(err) }

// Assert common_errors.Assert
func Assert(condition bool, err error) { errors.Assert(condition, err) }

// AssertFalse common_errors.AssertFalse
func AssertFalse(condition bool, err error) { errors.AssertFalse(condition, err) }

// AssertErrorNil common_errors.AssertErrorNil
func AssertErrorNil(err, throw error) { errors.AssertErrorNil(err, throw) }

// AssertNotEmpty common_errors.AssertNotEmpty
func AssertNotEmpty(obj interface{}, err error) { errors.AssertNotEmpty(obj, err) }
