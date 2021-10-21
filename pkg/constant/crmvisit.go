package constant

import "time"

const (
	GoroutineTimeout       = 90 * time.Second
	GoroutineTimeoutExport = 1800 * time.Second
	UploadFileName         = "visitrecord"
	MaxCrmVisitExport      = 5000
)
