package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"testing"
)

func TestWatch(t *testing.T) {
	NewZookeeper()
	defer ZKClient.Close()
	zkLock := "/myZkLock"
	zkLockSub := zkLock + "/lock"
	zkValue := []byte("fangfang")
	acls := zk.WorldACL(zk.PermAll)
	//exist
	exist, _, err := ZKClient.Exists(zkLock)
	if err != nil {
		t.Error(err)
	}
	t.Log(zkLock, exist)
	if !exist {
		path, err := ZKClient.Create(zkLock, zkValue, 0, acls)
		if err != nil {
			t.Error(err)
		}
		t.Log(path)
	}
	// 监听子节点
	children, _, childCh, err := ZKClient.ChildrenW(zkLock)
	if err != nil {
		t.Error(err)
	}
	t.Log(children)
	// 创建临时的且有序的子节点
	for i := 0; i < 10; i++ {
		path, err := ZKClient.Create(zkLockSub, zkValue, 3, acls)
		if err != nil {
			t.Error(err)
		}
		t.Log(path)
		select {
		case event := <-childCh:
			t.Log("path:", event.Path)
			t.Log("type:", event.Type.String())
			t.Log("state:", event.State.String())
		}
	}
	children, _, err = ZKClient.Children(zkLock)
	if err != nil {
		t.Error(err)
	}
	t.Log(children)
}
