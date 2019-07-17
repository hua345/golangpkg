package redis

import (
	"testing"
	"time"
)

func TestSetNX(t *testing.T) {
	NewRedis()
	key := "tryLockName"
	value := "fangfang"
	result, err := RedisClient.SetNX(key, value, 5*time.Second).Result()
	if err != nil {
		t.Error(err)
	}
	if true == result {
		t.Log(result)
		t.Log("获取Redis锁成功")
	}
	result, err = RedisClient.SetNX(key, value, 5*time.Second).Result()
	if err != nil {
		t.Error(err)
	}
	if false == result {
		t.Log(result)
		t.Log("获取Redis锁失败")
	} else {
		t.Error("获取锁出现错误")
	}
}
