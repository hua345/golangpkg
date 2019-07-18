package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"testing"
)

func TestChildren(t *testing.T) {
	NewZookeeper()
	defer ZKClient.Close()
	zkPath := "/zkWatchTest"
	zkLock := zkPath + "/lock"
	zkValue := []byte("fangfang")
	acls := zk.WorldACL(zk.PermAll)
	// 判断分布式锁的永久节点是否存在
	exist, _, err := ZKClient.Exists(zkPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(zkPath, exist)
	if !exist {
		path, err := ZKClient.Create(zkPath, zkValue, 0, acls)
		if err != nil {
			t.Error(err)
		}
		t.Log(path)
	}
	//flags有4种取值：
	//0:永久，除非手动删除
	//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	//zk.FlagSequence  = 2:会自动在节点后面添加序号
	//3:Ephemeral和Sequence，即，短暂且自动添加序号
	for i := 0; i < 10; i++ {
		path, err := ZKClient.Create(zkLock, zkValue, 3, acls)
		if err != nil {
			t.Error(err)
		}
		t.Log(path)
	}

	childrenList, _, err := ZKClient.Children(zkPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(childrenList)
}
