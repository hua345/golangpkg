package goroutineLearn

import (
	"testing"
	"time"
)

func TestLeak(t *testing.T) {
	go t.Log("你好, 并发!") // 干活的

	// 10000: 80M左右
	for i := 0; i < 10000; i++ {
		go func() { <-make(chan int) } () // 滥竽充数的, Goroutine 泄露
	}
	go func() {} () // 滥竽充数的, 但不是 Goroutine 泄露

	time.Sleep(30 * time.Second)
	t.Log("Done")
}