package cache_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"insurance/pkg/cache"
	"insurance/test"
)

//--- PASS: TestNewDistributedLock (0.56s)
//=== RUN   TestNewDistributedLock/test_lock
//dl.Lock(4) success
//dl.Lock(1) failed
//dl.Lock(0) failed
//dl.Lock(2) failed
//dl.Lock(3) failed
//dl.Unlock(4) success
//    --- PASS: TestNewDistributedLock/test_lock (0.56s)
//PASS
func TestNewDistributedLock(t *testing.T) {
	type args struct {
		opt []cache.Option
	}
	tests := []struct {
		name string
		args args
		want *cache.DistributedLock
	}{
		{
			name: "test lock",
			args: args{opt: []cache.Option{
				cache.WithContext(test.GetCtx()),
				cache.Timeout(1 * time.Second),
				cache.Keys(map[string]interface{}{
					"scene":  "calc",
					"userID": "1",
				}),
				cache.SetData("1"),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl, _ := cache.NewDistributedLock(tt.args.opt...)

			wg := sync.WaitGroup{}

			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func(idx int) {
					defer wg.Done()
					if dl.Lock() {
						fmt.Printf("dl.Lock(%d) success\n", idx)
						time.Sleep(500 * time.Millisecond)
						if err := dl.Unlock(); err != nil {
							fmt.Printf("dl.Unlock(%d) error:%v\n", idx, err)
						} else {
							fmt.Printf("dl.Unlock(%d) success\n", idx)
						}
					} else {
						time.Sleep(50 * time.Millisecond)
						fmt.Printf("dl.Lock(%d) failed\n", idx)
					}
				}(i)
			}
			wg.Wait()
		})
	}
}
