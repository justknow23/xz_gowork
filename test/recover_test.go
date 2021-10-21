package test

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	a:=30 * time.Second
	b:=a.Milliseconds()
	fmt.Println(b)
	//
	//e:= c()
	//fmt.Println(e)
}

func c() error {
	//defer func(err error) {
	//	if err := recover(); err != nil {
	//		fmt.Println(1111111111)
	//		err = errors.New("aaaaaaa")
	//	}
	//}(err)
	//panic("1")
	//fmt.Println(22222222222)
	//return err
	var err error
	Try(func() {
		panic("panic")
	}, func(e interface{}) {
		err = errors.New(fmt.Sprintf("gengtao:%+v",e))
	})
	return err
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}