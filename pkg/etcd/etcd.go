package etcd

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var EtcdClient *clientv3.Client

/**
https://etcd.io/
https://github.com/etcd-io/etcd/tree/master/clientv3
etcd/clientv3 is the official Go etcd client for v3.
etcd v3 uses gRPC for remote procedure calls. And clientv3 uses grpc-go to connect to etcd.
Make sure to close the client after using it. If the client is not closed, the connection will have leaky goroutines.
To specify client request timeout, pass context.WithTimeout to APIs:
*/
func InitEtcd() {
	var err error
	EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.137.128:2379", "192.168.137.128:22379", "192.168.137.128:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(fmt.Sprintf("Unable to create etcd client: %v", err))
	}
}
