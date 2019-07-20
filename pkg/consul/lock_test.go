package consul

import (
	"github.com/hashicorp/consul/api"
	"testing"
)

func createTestLock(t *testing.T, c *api.Client, key string) (*api.Lock, *api.Session) {
	session := c.Session()

	se := &api.SessionEntry{
		Name:     api.DefaultLockSessionName,
		TTL:      api.DefaultLockSessionTTL,
		Behavior: api.SessionBehaviorDelete,
	}
	id, _, err := session.CreateNoChecks(se, nil)
	t.Log(id)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	opts := &api.LockOptions{
		Key:         key,
		Session:     id,
		SessionName: se.Name,
		SessionTTL:  se.TTL,
	}
	lock, err := c.LockOpts(opts)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	return lock, session
}

func TestLockUnlock(t *testing.T) {
	NewConsul()
	lock, _ := createTestLock(t, ConsulClient, "test/lock")
	defer lock.Destroy()
	// Initial unlock should fail
	err := lock.Unlock()
	if err != api.ErrLockNotHeld {
		t.Fatalf("err: %v", err)
	}

	// Should work
	leaderCh, err := lock.Lock(nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if leaderCh == nil {
		t.Fatalf("not leader")
	}

	// Double lock should fail
	_, err = lock.Lock(nil)
	if err != api.ErrLockHeld {
		t.Fatalf("err: %v", err)
	}

	// Should be leader
	select {
	case <-leaderCh:
		t.Fatalf("should be leader")
	default:
	}

	// Initial unlock should work
	err = lock.Unlock()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
}
