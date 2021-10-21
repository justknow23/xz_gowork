package tools

import (
	"fmt"
)

// SliceRemove 删除Slice中的元素 []int64 []int []string []float64
func SliceRemove(s interface{}, index int) error {
	if ps, ok := s.(*[]string); ok {
		*ps = append((*ps)[:index], (*ps)[index+1:]...)
	} else if ps, ok := s.(*[]int); ok {
		*ps = append((*ps)[:index], (*ps)[index+1:]...)
	} else if ps, ok := s.(*[]int64); ok {
		*ps = append((*ps)[:index], (*ps)[index+1:]...)
	} else if ps, ok := s.(*[]float64); ok {
		*ps = append((*ps)[:index], (*ps)[index+1:]...)
	} else {
		return fmt.Errorf("<sliceRemove> Unsupported type: %T", s)
	}
	return nil
}

// SliceLast 取最后一个元素  返回 interface{}
func SliceLast(s interface{}) (interface{}, error) {
	if ps, ok := s.([]string); ok {
		last := (ps)[len(ps)-1]
		return last, nil
	}
	return nil, fmt.Errorf("<sliceLast> Unsupported type: %T", s)
}

// SliceUnique 切片排重 目前支持 []int64 []int []string
func SliceUnique(s interface{}) error {
	rs := make([]int, 0)
	switch ps := s.(type) {
	case *[]string:
		m := make(map[string]struct{}, len(*ps))
		for idx, v := range *ps {
			if _, ok := m[v]; ok {
				rs = append(rs, idx)
			} else {
				m[v] = struct{}{}
			}
		}
	case *[]int64:
		m := make(map[int64]struct{}, len(*ps))
		for idx, v := range *ps {
			if _, ok := m[v]; ok {
				rs = append(rs, idx)
			} else {
				m[v] = struct{}{}
			}
		}
	case *[]int:
		m := make(map[int]struct{}, len(*ps))
		for idx, v := range *ps {
			if _, ok := m[v]; ok {
				rs = append(rs, idx)
			} else {
				m[v] = struct{}{}
			}
		}
	default:
		return fmt.Errorf("<sliceUnique> Unsupported type: %T", s)
	}
	for idx := range rs {
		err := SliceRemove(s, idx)
		if err != nil {
			return err
		}
	}
	return nil
}

// SliceStringConvertInt64 -
func SliceStringConvertInt64(input []string) []int64 {
	output := make([]int64, len(input))
	for idx, s := range input {
		if i, e := IfInt64(s); e == nil {
			output[idx] = i
		} else {
			output[idx] = 0
		}
	}
	return output
}

// SplitSliceString -
func SplitSliceString(arr []string, num int64) [][]string {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]string{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]string, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

// SplitSliceInt64 -
func SplitSliceInt64(arr []int64, num int64) [][]int64 {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]int64{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]int64, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}
