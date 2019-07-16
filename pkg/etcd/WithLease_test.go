package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"testing"
)

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
		t.Error("after 5 seconds myKey not removed")
	}
	// Output: number of keys: 0
}
