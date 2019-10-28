package timeLearn

import (
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	retryTime := 2 * time.Second
	retryTimer := time.NewTimer(retryTime)
	select {
	case <-retryTimer.C:
		println("2s timer")
		retryTimer.Stop()
	}
}

func TestTime(t *testing.T) {
	retryTime := 1 * time.Second
	retryTimer := time.NewTimer(retryTime)
	defer retryTimer.Stop()
	for i := 0; i < 10; i++ {
		select {
		case <-retryTimer.C:
			println("1s timer")
			retryTimer.Reset(retryTime)
		}
	}
}

func TestSleep(t *testing.T) {
	retryTime := 1 * time.Second
	for i := 0; i < 10; i++ {
		time.Sleep(retryTime)
		println("after 1s")
	}
}
