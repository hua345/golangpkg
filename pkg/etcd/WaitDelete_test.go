package etcd

import (
	"context"
	"testing"
)

// etcdctl get testKey --order="DESCEND" --sort-by="CREATE" --limit=0
func TestWaitDeletes(t *testing.T) {
	InitEtcd()

	testDeleteKey := "testKey"
	_, err := waitDeletes(context.TODO(), EtcdClient, testDeleteKey, int64(100000))
	if err != nil {
		t.Error(err)
	}
}

// etcdctl get testKey --order="DESCEND" --sort-by="CREATE" --limit=0
func TestWaitDelete(t *testing.T) {
	InitEtcd()

	testDeleteKey := "testKey"
	err := waitDelete(context.TODO(), EtcdClient, testDeleteKey, int64(100000))
	if err != nil {
		t.Error(err)
	}
}
