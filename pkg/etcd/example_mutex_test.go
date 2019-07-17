package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3/concurrency"
	"sync"
	"testing"
	"time"
)

func TestExampleMutex(t *testing.T) {
	InitEtcd()
	// create two separate sessions for lock competition
	sission1, err := concurrency.NewSession(EtcdClient)
	if err != nil {
		t.Error(err)
	}
	defer sission1.Close()
	mutex1 := concurrency.NewMutex(sission1, "/my-lock/")

	session2, err := concurrency.NewSession(EtcdClient)
	if err != nil {
		t.Error(err)
	}
	defer session2.Close()
	mutex2 := concurrency.NewMutex(session2, "/my-lock/")

	// acquire lock for sission1
	if err := mutex1.Lock(context.TODO()); err != nil {
		t.Error(err)
	}
	t.Log("acquired lock for sission1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// wait until sission1 is locks /my-lock/
		if err := mutex2.Lock(context.TODO()); err != nil {
			t.Error(err)
		}
	}()

	if err := mutex1.Unlock(context.TODO()); err != nil {
		t.Log(err)
	}
	t.Log("released lock for sission1")
	<-m2Locked
	t.Log("acquired lock for session2")
}
func mutexWork(index int, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	sission1, err := concurrency.NewSession(EtcdClient, concurrency.WithTTL(5))
	if err != nil {
		t.Error(err)
	}
	defer sission1.Close()
	mutex1 := concurrency.NewMutex(sission1, "/my-lock/")
	// acquire lock for sission1
	if err := mutex1.Lock(context.TODO()); err != nil {
		t.Error(err)
	}
	t.Log("acquired lock for sission", index)
	time.Sleep(1000 * time.Millisecond)
	if err := mutex1.Unlock(context.TODO()); err != nil {
		t.Error(err)
	}
	t.Log("released lock for sission", index)
}
func TestExampleMutex2(t *testing.T) {
	InitEtcd()
	wg := sync.WaitGroup{}
	wg.Add(10)
	// create two separate sessions for lock competition
	for i := 0; i < 10; i++ {
		go mutexWork(i, &wg, t)
	}
	wg.Wait()
}

// 性能测试
//go test -bench=.
func BenchmarkEtcdLock(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	InitEtcd()
	b.StartTimer()
	lockKey := "myLock"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		sission1, err := concurrency.NewSession(EtcdClient, concurrency.WithTTL(5))
		if err != nil {
			b.Error(err)
		}
		defer sission1.Close()
		mutex1 := concurrency.NewMutex(sission1, lockKey)
		// acquire lock for sission1
		if err := mutex1.Lock(context.TODO()); err != nil {
			b.Error(err)
		}
		if err := mutex1.Unlock(context.TODO()); err != nil {
			b.Error(err)
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkRedisLockParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	InitEtcd()
	b.StartTimer()
	lockKey := "myLock"
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sission1, err := concurrency.NewSession(EtcdClient, concurrency.WithTTL(5))
			if err != nil {
				b.Error(err)
			}
			defer sission1.Close()
			mutex1 := concurrency.NewMutex(sission1, lockKey)
			// acquire lock for sission1
			if err := mutex1.Lock(context.TODO()); err != nil {
				b.Error(err)
			}
			if err := mutex1.Unlock(context.TODO()); err != nil {
				b.Error(err)
			}
		}
	})
}
