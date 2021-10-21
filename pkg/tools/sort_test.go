package tools_test

import (
	"xz_gowork/pkg/enum/enumint"
	"sort"
	"testing"
)

type MapSorter []Item
type Item struct {
	Key int
	Val string
}

func NewMapSorter(m map[int]string) MapSorter {
	ms := make(MapSorter, 0, len(m))
	for k, v := range m {
		ms = append(ms, Item{k, v})
	}
	return ms
}

func (ms MapSorter) Len() int {
	return len(ms)
}

func (ms MapSorter) Less(i, j int) bool {
	return ms[i].Val < ms[j].Val // 按值排序
	//return ms[i].Key < ms[j].Key // 按键排序

}

func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func TestSort(t *testing.T) {
	m := enumint.CrmVisitPurposeEnum.GetAll()
	ms := NewMapSorter(m)
	t.Logf("sort1: %+v", ms)
	sort.Sort(ms)
	t.Logf("sort: %+v", ms)
}
