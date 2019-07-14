package mutexLearn

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func incrementor(s string, t *testing.T) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		counter++
		t.Log(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

func TestMutex(t *testing.T) {
	wg.Add(2)
	go incrementor("Foo:", t)
	go incrementor("Bar:", t)
	wg.Wait()
	t.Log("Final Counter:", counter)
}
