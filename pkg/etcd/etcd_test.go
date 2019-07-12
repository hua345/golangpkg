package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestEtcdPutGet(t *testing.T) {
	InitEtcd()
	testkey := "name"
	testValue := "fang"
	_, err := EtcdClient.Put(context.Background(), testkey, testValue)
	if err != nil {
		t.Log("Etcd Put Error")
		t.Error(err)
	}
	etcdResp, err := EtcdClient.Get(context.Background(), testkey)
	if err != nil {
		t.Error(err)
	}
	t.Log(etcdResp)
	t.Log(etcdResp.Kvs)
	if len(etcdResp.Kvs) != 0 {
		if testValue != string(etcdResp.Kvs[0].Value) {
			t.Error("Put/Get值不一致")
		}
	}
}

// https://github.com/etcd-io/etcd/blob/master/clientv3/example_lease_test.go
func TestEtcdPutWithLease(t *testing.T) {
	InitEtcd()
	testkey := "nameLease"
	testValue := "fang"

	// create a lease first,minimum lease TTL is 5-second
	resp, err := EtcdClient.Grant(context.Background(), int64(5))
	if err != nil {
		t.Log("Unable to create lease")
		t.Error(err)
	}
	// after 5 seconds, the key 'foo' will be removed
	_, err = EtcdClient.Put(context.Background(), testkey, testValue, clientv3.WithLease(resp.ID))
	if err != nil {
		t.Log("Etcd Put WithLease Error")
		t.Error(err)
	}
	getResp, err := EtcdClient.Get(context.Background(), testkey)
	if err != nil {
		t.Error(err)
	}
	if len(getResp.Kvs) == 0 || testValue != string(getResp.Kvs[0].Value) {
		t.Error("Etcd Put WithLease failed")
	}
	// revoking lease expires the key attached to its lease ID
	_, err = EtcdClient.Revoke(context.Background(), resp.ID)
	if err != nil {
		t.Error(err)
	}

	getResp, err = EtcdClient.Get(context.Background(), testkey)
	if err != nil {
		t.Error(err)
	}
	if len(getResp.Kvs) != 0 {
		t.Error("after 5 seconds Key not removed")
	}
	// Output: number of keys: 0
}

func TestEtcdWatch(t *testing.T) {
	InitEtcd()
	etcdWatchKey := "testWatchKey"
	etcdWatchValue := "liufang"
	watchChan := EtcdClient.Watch(context.Background(), etcdWatchKey)
	t.Log("set WATCH on " + etcdWatchKey)
	go func() {
		t.Log("started goroutine for Etcd PUT")
		for {
			time.Sleep(2 * time.Second)
			_, err := EtcdClient.Put(context.Background(), etcdWatchKey, etcdWatchValue)
			if err != nil {
				t.Error(err)
			}
		}
	}()
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			t.Log(event)
			if string(event.Kv.Key) != etcdWatchKey {
				t.Error()
			}
			if string(event.Kv.Value) != etcdWatchValue {
				t.Error()
			}
		}
		break
	}
}

func TestEtcdStatus(t *testing.T) {
	InitEtcd()
	endpoints := EtcdClient.Endpoints()
	for _, ep := range endpoints {
		resp, err := EtcdClient.Status(context.Background(), ep)
		if err != nil {
			t.Error(err)
		}
		t.Log(resp)
		if resp.Header.MemberId == resp.Leader {
			t.Log("endpoint: " + ep + " / Leader: true")
		} else {
			t.Log("endpoint: " + ep + " / Leader: false")
		}

	}
}

// 性能测试
//go test -bench=.
func BenchmarkETCDPut(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	InitEtcd()
	b.StartTimer()
	testkey := "name"
	testValue := "fang"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		_, err := EtcdClient.Put(context.Background(), testkey, testValue)
		if err != nil {
			b.Log("Etcd Put Error")
			b.Error(err)
		}
	}
}

// 性能测试
//go test -bench=.
func BenchmarkETCDGet(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	InitEtcd()
	b.StartTimer()
	testkey := "name"
	testValue := "fang"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		etcdResp, err := EtcdClient.Get(context.Background(), testkey)
		if err != nil {
			b.Error(err)
		}
		if len(etcdResp.Kvs) != 0 {
			if testValue != string(etcdResp.Kvs[0].Value) {
				b.Error("Put/Get值不一致")
			}
		}
	}
}
