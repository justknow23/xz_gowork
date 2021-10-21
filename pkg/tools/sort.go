package tools

import (
	"sort"
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
	//return ms[i].Val < ms[j].Val // 按值排序
	return ms[i].Key < ms[j].Key // 按键排序

}

func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

// SortIntMap 排序
func SortIntMap(m map[int]string) MapSorter {
	ms := NewMapSorter(m)
	sort.Sort(ms)
	return ms
}
