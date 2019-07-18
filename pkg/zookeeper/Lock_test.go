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
	m2Locked := make(chan struct{})
	zkLock2 := zk.NewLock(ZKClient, lockPath, acls)
	go func() {
		defer close(m2Locked)
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
	<-m2Locked
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
	time.Sleep(1000 * time.Millisecond)
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
	wg.Add(10)
	// create two separate sessions for lock competition
	for i := 0; i < 10; i++ {
		go zkLockWork(i, &wg, t)
	}
	wg.Wait()
}
