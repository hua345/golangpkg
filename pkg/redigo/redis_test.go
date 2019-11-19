package redigo

import (
	"testing"
)

func TestRedisSet(t *testing.T) {
	testKey := "name"
	testValue := "fang"
	err := SetStr(testKey, testValue)
	if err != nil {
		t.Error(err)
	}
	value, err := GetStr(testKey)
	if err != nil {
		t.Error(err)
	}
	if value != testValue {
		t.Log("Redis Get " + testKey + ": " + value)
		t.Error("Redis Get " + testKey + " != " + testValue)
	}
}

func BenchmarkSetStr(b *testing.B) {
	testKey := "name"
	testValue := "fang"
	b.StopTimer() //停止压力测试的时间计数
	GetInstance()
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		err := SetStrWithExpire(testKey, testValue, 60*60)
		if err != nil {
			b.Error(err)
		}
	}
}
func BenchmarkGetStr(b *testing.B) {
	testKey := "name"
	b.StopTimer() //停止压力测试的时间计数
	GetInstance()
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		_, err := GetStr(testKey)
		if err != nil {
			b.Error(err)
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkRedisGetParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	GetInstance()
	b.StartTimer()
	testKey := "name"
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := GetStr(testKey)
			if err != nil {
				b.Error(err)
			}
		}
	})
}
