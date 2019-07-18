package zookeeper

//import (
//	"github.com/samuel/go-zookeeper/zk"
//	"testing"
//	"time"
//)
//
//func TestWatch(t *testing.T){
//	NewZookeeper()
//	defer ZKClient.Close()
//	zkPath := "/zkGolangTest"
//	zkValue := []byte("fangfang")
//
//	children, stat, childCh, err := zk.("/")
//	if err != nil {
//		t.Fatalf("Children returned error: %+v", err)
//	} else if stat == nil {
//		t.Fatal("Children returned nil stat")
//	} else if len(children) < 1 {
//		t.Fatal("Children should return at least 1 child")
//	}
//
//}
