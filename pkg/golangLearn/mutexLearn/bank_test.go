package mutexLearn

import (
	"golangpkg/pkg/golangLearn"
	"runtime"
	"testing"
)

func TestBank(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go golangLearn.Deposit(100)
	go func() {
		golangLearn.Deposit(200)          // A1
		t.Log("=", golangLearn.Balance()) // A2
	}()
	var x, y int
	go func() {
		x = 1 // A1
		t.Log("y:", y, " ") // A2
	}()
	go func() {
		y = 1                   // B1
		t.Log("x:", x, " ") // B2
	}()
}
