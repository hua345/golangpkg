package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/clientv3util"
	"log"
	"testing"
)

// https://github.com/etcd-io/etcd/blob/master/clientv3/clientv3util/example_key_test.go
func TestKeyMissing(t *testing.T) {
	InitEtcd()
	kvc := clientv3.NewKV(EtcdClient)
	// perform a put only if key is missing
	// It is useful to do the check atomically to avoid overwriting
	// the existing key which would generate potentially unwanted events,
	// unless of course you wanted to do an overwrite no matter what.
	_, err := kvc.Txn(context.Background()).
		If(clientv3util.KeyMissing("hello")).
		Then(clientv3.OpPut("hello", "hello world")).
		Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func TestKeyExists(t *testing.T) {
	InitEtcd()
	kvc := clientv3.NewKV(EtcdClient)
	// perform a delete only if key already exists
	_, err := kvc.Txn(context.Background()).
		If(clientv3util.KeyExists("hello")).
		Then(clientv3.OpDelete("hello")).
		Commit()
	if err != nil {
		log.Fatal(err)
	}
}
