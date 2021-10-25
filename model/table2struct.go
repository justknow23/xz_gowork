package model

import "github.com/gohouse/converter"

type Table2Struct struct {
}

func (t *Table2Struct) GetTableStruct() error {
	t2t := converter.NewTable2Struct()
	t2t.Config(&converter.T2tConfig{
		RmTagIfUcFirsted: false,
		TagToLower:       false,
		UcFirstOnly:      false,
		SeperatFile:      false,
	})

	err := t2t.Table("pre_insurance").
		EnableJsonTag(true).
		PackageName("model").
		TagKey("orm").
		RealNameMethod("TableName").
		SavePath("preinsurance.go").
		Dsn("xzsys:f50d63b0063b6f62b48b1fc9af754f1a@tcp(mysql-test-master-n0.idc.xiaozhu.com:3306)/xz_insurance?charset=utf8mb4&parseTime=True&loc=Local").
		Run()
	if err != nil {
		return err
	}
	return nil
}
