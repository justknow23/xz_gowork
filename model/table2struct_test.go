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

func TestTable2Struct_GetTableStruct(t1 *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &model.Table2Struct{}
			if err := t.GetTableStruct(); (err != nil) != tt.wantErr {
				t1.Errorf("GetTableStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}