package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
)

func TestTxn(t *testing.T) {
	InitEtcd()
	lockPrefix := "mylock"
	// create a lease first,minimum lease TTL is 5-second
	leaseResp, err := EtcdClient.Grant(context.TODO(), int64(50))
	if err != nil {
		t.Log("Unable to create lease")
		t.Error(err)
	}
	lockKey := fmt.Sprintf("%s%x", lockPrefix+"/", leaseResp.ID)
	cmp := clientv3.Compare(clientv3.CreateRevision(lockKey), "=", 0)
	// put self in lock waiters via myKey; oldest waiter holds lock
	put := clientv3.OpPut(lockKey, "", clientv3.WithLease(leaseResp.ID))
	// reuse key in case this session already holds the lock
	get := clientv3.OpGet(lockKey)
	// fetch current holder to complete uncontended path with only one RPC
	getOwner := clientv3.OpGet(lockPrefix, clientv3.WithFirstCreate()...)
	resp, err := EtcdClient.Txn(context.TODO()).If(cmp).Then(put, getOwner).Else(get, getOwner).Commit()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
