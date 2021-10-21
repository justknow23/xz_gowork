package tools

import (
	"testing"
)

func Test_CreateCSVFile(t *testing.T) {

	f, e := CreateCSVFile([]string{"111", "111", "111"}, "a.csv", "/tmp")

	t.Logf("%#v\n", f)
	t.Logf("%#v\n", e)
}

func Test_CreateCSV(t *testing.T) {
	var rows [][]string
	rows = append(rows, []string{"a", "b", "c"},[]string{"111", "111", "111"})
	f, e := CreateCsv(rows, "b.csv", "/tmp")
	t.Logf("%#v\n", f)
	t.Logf("%#v\n", e)
}

func Test_ReadCSVFile(t *testing.T) {

	f, e := ReadCsv("/tmp/b.csv")

	t.Logf("%#v\n", f)
	t.Logf("%#v\n", e)
}