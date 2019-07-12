package etcd

import (
	"fmt"
	"testing"
	"time"
)

func TestEtcdMutex(t *testing.T) {
	InitEtcd()
	for i := 0; i < 10; i++ {
		etcdMutex := &EtcdMutex{
			TTL: 10,
			Key: "lock",
		}
		//groutine1
		go func() {
			err := etcdMutex.Lock()
			if err != nil {
				fmt.Println("groutine" + string(i) + "抢锁失败")
				fmt.Println(err)
				return
			}
			fmt.Println("groutine" + string(i) + "抢锁成功")
			time.Sleep(10 * time.Second)
			defer etcdMutex.UnLock()
		}()
	}
	time.Sleep(10 * time.Second)
}
