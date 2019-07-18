package zookeeper

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"testing"
	"time"
)

// zk事件类型
//var (
//	eventNames = map[EventType]string{
//		EventNodeCreated:         "EventNodeCreated",
//		EventNodeDeleted:         "EventNodeDeleted",
//		EventNodeDataChanged:     "EventNodeDataChanged",
//		EventNodeChildrenChanged: "EventNodeChildrenChanged",
//		EventSession:             "EventSession",
//		EventNotWatching:         "EventNotWatching",
//	}
//)
func zkCallback(event zk.Event) {
	fmt.Println(">>>>>>>>>>>>>>>>>>>")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("<<<<<<<<<<<<<<<<<<<")
}

func TestEvent(t *testing.T) {
	NewZookeeperWithCallback(zkCallback)
	defer ZKClient.Close()
	zkPath := "/zkWatchTest"
	zkValue := []byte("fangfang")
	//exist
	exist, _, err := ZKClient.Exists(zkPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(zkPath, exist)
	// try create
	//flags有4种取值：
	//0:永久，除非手动删除
	//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	//zk.FlagSequence  = 2:会自动在节点后面添加序号
	//3:Ephemeral和Sequence，即，短暂且自动添加序号
	var acls = zk.WorldACL(zk.PermAll)
	createdPath, err := ZKClient.Create(zkPath, zkValue, zk.FlagEphemeral, acls)
	if err != nil {
		t.Error(err)
	}
	t.Log(createdPath)
	time.Sleep(time.Second * 2)
	//exist
	exist, _, err = ZKClient.Exists(zkPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(zkPath, exist)
	// delete
	err = ZKClient.Delete(zkPath, -1)
	if err != nil {
		t.Error(err)
	}
}
