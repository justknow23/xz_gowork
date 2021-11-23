package defers_test

import (
	"fmt"
	_ "github.com/gogo/protobuf/test"
	"testing"
)

func TestDeferFunc(t *testing.T) {
	//deferFuncParameter()
	deferFuncParameter1()
}

func printArray(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}
func printArray1(a int) {
	fmt.Println(a)
}
func deferFuncParameter() {
	var aArray = [3]int{1, 2, 3}

	defer printArray(&aArray)

	aArray[0] = 10
	return
}

func deferFuncParameter1() {
	var aArray = 1

	defer printArray1(aArray)

	aArray = 10
	return
}