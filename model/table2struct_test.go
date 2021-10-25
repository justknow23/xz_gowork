package model_test

import (
	"testing"
	"xz_gowork/model"
)

func TestTable2Struct(t *testing.T) {
	t2t := model.Table2Struct{}
	err := t2t.GetTableStruct()
	if err != nil {
		t.Logf("err:%+v", err)
	}
	t.Logf("success")
}
