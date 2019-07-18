package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"testing"
)

func TestWatch(t *testing.T) {
	NewZookeeper()
	defer ZKClient.Close()
	zkPath := "/zkWatchTest"
	zkValue := []byte("fangfang")
	acls := zk.WorldACL(zk.PermAll)
	//exist
	exist, _, err := ZKClient.Exists(zkPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(zkPath, exist)

	if exist {
		// delete
		err = ZKClient.Delete(zkPath, -1)
		if err != nil {
			t.Error(err)
		}
	}
	children, _, childCh, err := ZKClient.ChildrenW("/")
	if err != nil {
		t.Error(err)
	}
	t.Log(children)

	path, err := ZKClient.Create(zkPath, zkValue, 0, acls)
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
	children, _, childCh, err = ZKClient.ChildrenW(zkPath)
	if err != nil {
		t.Error(err)
	}

	// delete
	err = ZKClient.Delete(zkPath, -1)
	if err != nil {
		t.Error(err)
	}

	select {
	case event := <-childCh:
		t.Log("path:", event.Path)
		t.Log("type:", event.Type.String())
		t.Log("state:", event.State.String())
	}
}
