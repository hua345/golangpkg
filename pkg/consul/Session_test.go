package consul

import (
	"github.com/hashicorp/consul/api"
	"testing"
)

func TestSessionCreateDestroy(t *testing.T) {
	NewConsul()
	session := ConsulClient.Session()
	se := &api.SessionEntry{
		Name:     api.DefaultLockSessionName,
		TTL:      api.DefaultLockSessionTTL,
		Behavior: api.SessionBehaviorDelete,
	}
	id, meta, err := session.Create(se, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
	t.Log(meta.RequestTime)
	meta, err = session.Destroy(id, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(meta.RequestTime)
}
