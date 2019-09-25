package redis

import (
	"errors"
	"github.com/hua345/golangpkg/pkg/util"
	"time"

	"github.com/go-redis/redis"
)

var (
	luaRelease = redis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("del", KEYS[1]) else return 0 end`)
)

var (
	// ErrTryLockFailed is returned when a lock cannot be tryLock.
	ErrTryLockFailed = errors.New("redisLock: tryLock Failed")

	// ErrLockNotHeld is returned when trying to release an inactive lock.
	ErrUnLockFailed = errors.New("redisLock: unLock delete key failed")
)

// RedisSession wraps a redis session.
type RedisSession struct {
	// The number of time the acquisition of a lock will be retried.
	// Default: 0 = do not retry
	RetryCount int `json:"retryCount"`

	// RetryTimeout is the amount of time to wait between retries.
	// Default: 100ms
	RetryTimeout time.Duration `json:"retryTimeout"`
	client       *redis.Client `json:"client"`
}

const DefaultRetryCount = 5
const DefaultRetryTimeout = 100 * time.Millisecond

// NewRedisSession creates a new RedisSession instance with a custom namespace.
func NewRedisSession(client *redis.Client) *RedisSession {
	redisSession := &RedisSession{client: client}
	redisSession.RetryCount = DefaultRetryCount
	redisSession.RetryTimeout = DefaultRetryTimeout
	return redisSession
}

// TryLock tries to tryLock a new lock using a key with the given TTL.
// May return ErrNotObtained if not successful.
func (redisSession *RedisSession) TryLock(key string, ttl time.Duration) (*Lock, error) {
	if redisSession == nil || redisSession.client == nil {
		panic("RedisClient Need Init")
	}
	value := util.GetUUID32()
	for {
		ok, err := redisSession.tryLock(key, value, ttl)
		if err != nil {
			return nil, err
		} else if ok {
			return &Lock{session: redisSession, key: key, value: value}, nil
		}
		time.Sleep(redisSession.RetryTimeout)
	}
	return nil, ErrTryLockFailed
}

func (redisSession *RedisSession) tryLock(key, value string, ttl time.Duration) (bool, error) {
	ok, err := redisSession.client.SetNX(key, value, ttl).Result()
	if err == redis.Nil {
		err = nil
	}
	return ok, err
}

// Lock represents an obtained, distributed lock.
type Lock struct {
	session *RedisSession `json:"session"`
	key     string        `json:"key"`
	value   string        `json:"value"`
}

// Key returns the redis key used by the lock.
func (l *Lock) Key() string {
	return l.key
}

// Release manually releases the lock.
// May return ErrLockNotHeld.
func (l *Lock) Release() error {
	res, err := luaRelease.Run(l.session.client, []string{l.key}, l.value).Result()
	if err == redis.Nil {
		return ErrUnLockFailed
	} else if err != nil {
		return err
	}
	if i, ok := res.(int64); !ok || i != 1 {
		return ErrUnLockFailed
	}
	return nil
}
