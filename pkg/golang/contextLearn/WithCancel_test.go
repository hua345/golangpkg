package contextLearn

import (
	"context"
	"testing"
	"time"
)

func work(ctx context.Context, t *testing.T) {
	for {
		time.Sleep(1 * time.Second)
		select {
		// we received the signal of cancelation in this channel
		case <-ctx.Done():
			t.Log("done")
			return
		default:
			t.Log("work")
		}
	}
}
func TestWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
	go work(ctx, t)
	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}

func TestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
	go work(ctx, t)
	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}

func TestWithDeadline(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
	go work(ctx, t)
	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}
