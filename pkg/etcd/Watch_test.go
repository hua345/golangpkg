package etcd

import (
	"context"
	"testing"
	"time"
)

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
