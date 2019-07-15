package etcd

import (
	"context"
	"testing"
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
