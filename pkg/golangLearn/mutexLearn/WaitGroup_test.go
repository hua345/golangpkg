package mutexLearn

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func incrementor(s string, t *testing.T) {
	for i := 0; i < 20; i++ {
		//time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		counter++
		t.Log(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

func TestWaitGroup(t *testing.T) {
	wg.Add(3)
	go incrementor("AA:", t)
	go incrementor("BB:", t)
	go incrementor("CC:", t)
	wg.Wait()
	t.Log("Final Counter:", counter)
}
