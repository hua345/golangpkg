package redis

import (
	"testing"
)

func TestTime(t *testing.T) {
	redisTime, err := GetInstance().Time().Result()
	if err != nil {
		t.Error(err)
	}
	// 2006/01/02 15:04:05
	t.Log(redisTime.Format("200601021504"))
	t.Log(redisTime)
}
