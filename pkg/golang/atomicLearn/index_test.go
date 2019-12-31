package atomicLearn

import (
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

/**
https://golang.google.cn/pkg/sync/atomic/
These functions require great care to be used correctly.
Except for special, low-level applications, synchronization is better done with channels or the facilities of the sync package.
Share memory by communicating; don't communicate by sharing memory.
*/

func TestAtomic(t *testing.T) {
	var ops uint64 = 0

	for i := 0; i < 8; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	t.Log("ops:", opsFinal)
}
