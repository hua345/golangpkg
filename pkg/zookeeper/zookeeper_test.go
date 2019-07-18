package zookeeper

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"testing"
)

func ZkStateStringFormat(s *zk.Stat) string {
	return fmt.Sprintf("Czxid:%d\nMzxid: %d\nCtime: %d\nMtime: %d\nVersion: %d\nCversion: %d\nAversion: %d\nEphemeralOwner: %d\nDataLength: %d\nNumChildren: %d\nPzxid: %d\n",
		s.Czxid, s.Mzxid, s.Ctime, s.Mtime, s.Version, s.Cversion, s.Aversion, s.EphemeralOwner, s.DataLength, s.NumChildren, s.Pzxid)
}
func TestZooKeeper(t *testing.T) {
	NewZookeeper()
	defer ZKClient.Close()
	zkPath := "/zkGolangTest"
	zkValue := []byte("fangfang")

	err := ZKClient.Delete(zkPath, -1)
	if err != nil && err != zk.ErrNoNode {
		t.Fatalf("Delete returned error: %+v", err)
	}
	var flags int32 = 0
	//flags有4种取值：
	//0:永久，除非手动删除
	//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	//zk.FlagSequence  = 2:会自动在节点后面添加序号
	//3:Ephemeral和Sequence，即，短暂且自动添加序号
	var acls = zk.WorldACL(zk.PermAll) //控制访问权限模式
	createdPath, err := ZKClient.Create(zkPath, zkValue, flags, acls)
	if err != nil {
		t.Fatalf("Create returned error: %+v", err)
	}
	t.Log(createdPath)
	if createdPath != zkPath {
		t.Fatalf("Create returned different zkPath '%s' != '%s'", createdPath, zkPath)
	}
	data, stat, err := ZKClient.Get(zkPath)
	if err != nil {
		t.Fatalf("Get returned error: %+v", err)
	}
	t.Log(zkPath, string(data))
	t.Log(ZkStateStringFormat(stat))
}
