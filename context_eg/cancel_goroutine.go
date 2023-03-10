// Package context_eg
// Title       : cancel_goroutine.go
// Author      : Tuffy  2023/3/2 15:22
// Description :
package context_eg

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func myTask1() {
	fmt.Println("myTask1 -- 1")
	runtime.Gosched()
	time.Sleep(3)
	fmt.Println("myTask1 -- 2")

	time.Sleep(3)
	fmt.Println("myTask1 -- 3")

}

func CancelGoroutine() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(1)
			}
		}
	}(ctx)
	time.Sleep(1)
	cancel()
	time.Sleep(10)
}
