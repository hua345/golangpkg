package zookeeper

import (
	"testing"
)

func TestZkExistW(t *testing.T) {
	NewZookeeper()
	defer ZKClient.Close()
	zkPath := "/zkGolangTest"
	//zkValue := []byte("fangfang")
	exist, stat, err := ZKClient.Exists(zkPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(zkPath, exist)
	t.Log(ZkStateStringFormat(stat))
}
