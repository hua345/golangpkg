package etcd

import (
	"context"
	"testing"
)

func TestWaitDelete(t *testing.T) {
	InitEtcd()

	testDeleteKey := "testKey"
	err := waitDelete(context.TODO(), EtcdClient, testDeleteKey)
	if err != nil {
		t.Error(err)
	}
}
