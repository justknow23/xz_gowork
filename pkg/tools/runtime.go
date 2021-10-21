package tools

import (
	"runtime"
	"strings"
)

func getCallerName(skips ...int) string {
	skip := 2
	if len(skips) > 0 {
		skip = skips[0]
	}
	pc, _, _, _ := runtime.Caller(skip)
	fullName := runtime.FuncForPC(pc).Name()
	name, err := SliceLast(strings.Split(fullName, "."))
	if err == nil {
		return name.(string)
	}
	return "getCallerName error:" + err.Error()
}
