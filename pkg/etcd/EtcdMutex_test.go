package etcd

import (
	"fmt"
	"testing"
	"time"
)

func TestEtcdMutex(t *testing.T) {
	InitEtcd()
	myLockKey := "myLock"
	for i := 0; i < 10; i++ {
		//groutine1
		go func() {
			etcdMutex, err := NewLock(myLockKey, 5)
			if err != nil {
				t.Error("NewLock error")
			}
			err = etcdMutex.Lock()
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
	time.Sleep(5 * time.Second)
}
