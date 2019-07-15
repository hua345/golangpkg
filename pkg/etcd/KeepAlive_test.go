package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestEtcdKeepAlive(t *testing.T) {
	InitEtcd()
	ctx := context.Background()
	testkey := "nameLease"
	testValue := "fang"
	LeaseTTL := 5
	// create a lease first,minimum lease TTL is 5-second
	leaseResp, err := EtcdClient.Grant(ctx, int64(LeaseTTL))
	if err != nil {
		t.Log("Unable to create lease")
		t.Error(err)
	}
	// after 5 seconds, the key 'foo' will be removed
	_, err = EtcdClient.Put(ctx, testkey, testValue, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		t.Log("Etcd Put WithLease Error")
		t.Error(err)
	}
	// the key 'foo' will be kept forever
	leaseKeepAliveChan, err := EtcdClient.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		t.Log("Etcd KeepAlive Error")
		t.Error(err)
	}
	var keepAliveInfo *clientv3.LeaseKeepAliveResponse
	timeStart := time.Now()
	for i := 0; i < 2; i++ {
		keepAliveInfo = <-leaseKeepAliveChan
		t.Log(keepAliveInfo)
		t.Log("ttl:", keepAliveInfo.TTL)
	}
	duration := time.Since(timeStart).Seconds()
	t.Log(duration, " Seconds")
}
