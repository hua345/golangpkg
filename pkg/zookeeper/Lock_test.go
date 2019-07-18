package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"testing"
	"time"
)

func TestZkLock(t *testing.T) {
	NewZookeeper()
	defer ZKClient.Close()
	lockPath := "/lockPath"
	acls := zk.WorldACL(zk.PermAll)

	zkLock := zk.NewLock(ZKClient, lockPath, acls)
	err := zkLock.Lock()
	if err != nil {
		t.Error(err)
	}
	t.Log("acquired lock for sission1")
	waitChan := make(chan struct{})
	zkLock2 := zk.NewLock(ZKClient, lockPath, acls)
	go func() {
		defer close(waitChan)
		err = zkLock2.Lock()
		if err != nil {
			t.Error(err)
		}
		t.Log("acquired lock for session2")
		err = zkLock2.Unlock()
		if err != nil {
			t.Error(err)
		}
		t.Log("released lock for sission2")
	}()
	time.Sleep(time.Millisecond * 100)
	err = zkLock.Unlock()
	if err != nil {
		t.Error(err)
	}
	t.Log("released lock for sission1")
	<-waitChan
}
func zkLockWork(index int, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	lockPath := "/lockPath"
	acls := zk.WorldACL(zk.PermAll)
	zkLock := zk.NewLock(ZKClient, lockPath, acls)
	err := zkLock.Lock()
	if err != nil {
		t.Error(err)
	}
	t.Log("acquired lock for sission", index)
	time.Sleep(100 * time.Millisecond)
	err = zkLock.Unlock()
	if err != nil {
		t.Error(err)
	}
	t.Log("released lock for sission", index)
}
func TestZkLock2(t *testing.T) {
	NewZookeeper()
	defer ZKClient.Close()
	wg := sync.WaitGroup{}
	wg.Add(100)
	// create two separate sessions for lock competition
	for i := 0; i < 100; i++ {
		go zkLockWork(i, &wg, t)
	}
	wg.Wait()
}

// 性能测试
//go test -bench=.
func BenchmarkZkLock(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewZookeeper()
	defer ZKClient.Close()
	retryTime := 2 * time.Second
	retryTimer := time.NewTimer(retryTime)
	select {
	case <-retryTimer.C:
		b.Log("2s timer")
		retryTimer.Stop()
	}
	b.StartTimer()
	lockPath := "/lockPath"
	acls := zk.WorldACL(zk.PermAll)
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		zkLock := zk.NewLock(ZKClient, lockPath, acls)
		err := zkLock.Lock()
		if err != nil {
			b.Error(err)
		}
		err = zkLock.Unlock()
		if err != nil {
			b.Error(err)
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkZkLockParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewZookeeper()
	defer ZKClient.Close()
	retryTime := 2 * time.Second
	retryTimer := time.NewTimer(retryTime)
	select {
	case <-retryTimer.C:
		b.Log("2s timer")
		retryTimer.Stop()
	}
	b.StartTimer()
	lockPath := "/lockPath"
	acls := zk.WorldACL(zk.PermAll)
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			zkLock := zk.NewLock(ZKClient, lockPath, acls)
			err := zkLock.Lock()
			if err != nil {
				b.Error(err)
			}
			err = zkLock.Unlock()
			if err != nil {
				b.Error(err)
			}
		}
	})
}
