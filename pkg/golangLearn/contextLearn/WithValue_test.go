package contextLearn

import (
	"context"
	"testing"
	"time"
)

func valueWork(ctx context.Context, t *testing.T) {
	valueKey := "fang"
	for {
		time.Sleep(1 * time.Second)
		select {
		// we received the signal of cancelation in this channel
		case <-ctx.Done():
			t.Log(ctx.Value(valueKey), "is cancel")
			t.Log("done")
			return
		default:
			t.Log(ctx.Value(valueKey), "int goroutine")
			t.Log("work")
		}
	}
}

func TestWithValue(t *testing.T) {
	valueKey := "fang"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx = context.WithValue(ctx, valueKey, "fangfang")
	//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
	go valueWork(ctx, t)
	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}
