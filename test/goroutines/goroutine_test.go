package goroutines

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"testing"
)

var wg = sync.WaitGroup{}

// channel 与 sync 同步组合方式实现控制 goroutine
func TestGoroutineNum(t *testing.T) {
	//模拟用户需求go业务的数量
	task_cnt := math.MaxInt64

	ch := make(chan bool, 3)

	for i := 0; i < task_cnt; i++ {
		wg.Add(1)
		ch <- true
		go doBusiness(ch, i)
	}

	wg.Wait()
}

// 模拟执行业务的 goroutine
func doBusiness(ch chan bool, i int) {
	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch
	wg.Done()
}
