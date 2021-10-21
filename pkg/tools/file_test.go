package tools

import (
	"testing"
	"time"

	"insurance/pkg/global"
)

func Test_CreateCSVFile(t *testing.T) {

	f, e := CreateCSVFile([]string{"111"}, "aaaa-%s", time.Now().Format(global.DateFmtYMD))

	t.Logf("%#v\n", f)
	t.Logf("%#v\n", e)
}
