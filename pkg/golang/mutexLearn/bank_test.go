package mutexLearn

import (
	"runtime"
	"testing"
	"time"
)

func TestBank(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go Deposit(100)
	go func() {
		Deposit(200)          // A1
		t.Log("=", Balance()) // A2
	}()
	var x, y int
	go func() {
		x = 1               // A1
		t.Log("y:", y, " ") // A2
	}()
	go func() {
		y = 1               // B1
		t.Log("x:", x, " ") // B2
	}()
	time.Sleep(1 * time.Second)
}
