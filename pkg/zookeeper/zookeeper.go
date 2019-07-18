package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

const zkServer1 = "192.168.137.128:2181"
const zkServer2 = "192.168.137.128:22181"
const zkServer3 = "192.168.137.128:32181"

var ZKClient *zk.Conn

func NewZookeeper() {
	var err error
	ZKClient, _, err = zk.Connect([]string{zkServer1, zkServer2, zkServer3}, 3*time.Second) //*10)
	if err != nil {
		panic(err)
	}
}
func NewZookeeperWithCallback(callback func(event zk.Event)) {
	option := zk.WithEventCallback(callback)
	var err error
	ZKClient, _, err = zk.Connect([]string{zkServer1, zkServer2, zkServer3}, 3*time.Second, option) //*10)
	if err != nil {
		panic(err)
	}
}
