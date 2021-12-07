package gmp_test

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestOk(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			var buf [64]byte
			n := runtime.Stack(buf[:], false)
			idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
			id, err := strconv.Atoi(idField)
			if err != nil {
				panic(fmt.Sprintf("cannot get goroutine id: %v", err))
			}
			fmt.Println("go routine 1 i: ", i, id)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			var buf [64]byte
			n := runtime.Stack(buf[:], false)
			idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
			id, err := strconv.Atoi(idField)
			if err != nil {
				panic(fmt.Sprintf("cannot get goroutine id: %v", err))
			}
			fmt.Println("go routine 2 i: ", i, id)
			wg.Done()
		}(i)

	}
	wg.Wait()
}