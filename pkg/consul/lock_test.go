package consul

import (
	"github.com/hashicorp/consul/api"
	"sync"
	"testing"
	"time"
)

func createTestLock(t *testing.T, c *api.Client, key string) *api.Lock {
	session := c.Session()

	se := &api.SessionEntry{
		Name:     api.DefaultLockSessionName,
		TTL:      api.DefaultLockSessionTTL,
		Behavior: api.SessionBehaviorDelete,
	}
	id, _, err := session.CreateNoChecks(se, nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	opts := &api.LockOptions{
		Key:         key,
		Session:     id,
		SessionName: se.Name,
		SessionTTL:  se.TTL,
	}
	lock, err := c.LockOpts(opts)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	return lock
}

func TestLockUnlock(t *testing.T) {
	NewConsul()
	lockKey := "test/lock"
	lock := createTestLock(t, ConsulClient, lockKey)
	// Initial unlock should fail
	err := lock.Unlock()
	if err != api.ErrLockNotHeld {
		t.Fatalf("err: %v", err)
	}

	// Should work
	leaderCh, err := lock.Lock(nil)
	if err != nil {
		t.Error(err)
	}
	// Double lock should fail
	_, err = lock.Lock(nil)
	if err != api.ErrLockHeld {
		t.Error(err)
	}

	// Should be leader
	select {
	case <-leaderCh:
		t.Fatalf("should be leader")
	default:
	}

	// Initial unlock should work
	err = lock.Unlock()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
}
func lockWork(index int, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	lockKey := "test/lock"
	lock := createTestLock(t, ConsulClient, lockKey)
	// acquire lock for sission1
	_, err := lock.Lock(nil)
	if err != nil {
		t.Error(err)
	}
	t.Log("acquired lock for sission", index)
	time.Sleep(100 * time.Millisecond)
	err = lock.Unlock()
	if err != nil {
		t.Error(err)
	}
	t.Log("released lock for sission", index)
}
func TestConsulLock(t *testing.T) {
	NewConsul()
	wg := sync.WaitGroup{}
	wg.Add(100)
	// create two separate sessions for lock competition
	for i := 0; i < 100; i++ {
		go lockWork(i, &wg, t)
	}
	wg.Wait()
}

// 性能测试
//go test -bench=.
func BenchmarkEtcdLock(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewConsul()
	b.StartTimer()
	lockKey := "test/lock"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		session := ConsulClient.Session()
		se := &api.SessionEntry{
			Name:     api.DefaultLockSessionName,
			TTL:      api.DefaultLockSessionTTL,
			Behavior: api.SessionBehaviorDelete,
		}
		id, _, err := session.CreateNoChecks(se, nil)
		if err != nil {
			b.Error(err)
		}
		opts := &api.LockOptions{
			Key:         lockKey,
			Session:     id,
			SessionName: se.Name,
			SessionTTL:  se.TTL,
		}
		lock, err := ConsulClient.LockOpts(opts)
		if err != nil {
			b.Error(err)
		}
		// Should work
		_, err = lock.Lock(nil)
		if err != nil {
			b.Error(err)
		}
		// Initial unlock should work
		err = lock.Unlock()
		if err != nil {
			b.Fatalf("err: %v", err)
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkRedisLockParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewConsul()
	b.StartTimer()
	lockKey := "test/lock"
	b.Log("Hello")
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			session := ConsulClient.Session()
			se := &api.SessionEntry{
				Name:     api.DefaultLockSessionName,
				TTL:      api.DefaultLockSessionTTL,
				Behavior: api.SessionBehaviorDelete,
			}
			id, _, err := session.CreateNoChecks(se, nil)
			if err != nil {
				b.Error(err)
			}

			opts := &api.LockOptions{
				Key:         lockKey,
				Session:     id,
				SessionName: se.Name,
				SessionTTL:  se.TTL,
			}
			lock, err := ConsulClient.LockOpts(opts)
			if err != nil {
				b.Error(err)
			}
			// Should work
			_, err = lock.Lock(nil)
			if err != nil {
				b.Error(err)
			}
			// Initial unlock should work
			err = lock.Unlock()
			if err != nil {
				b.Fatalf("err: %v", err)
			}
		}
	})
}
