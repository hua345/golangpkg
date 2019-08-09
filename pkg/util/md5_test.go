package util

import (
	"testing"
)

func ExampleEncodeMD5() {
	EncodeMD5("hello")
}

// 单元测试
// go test -v
func TestEncodeMD5(t *testing.T) {
	t.Log(EncodeMD5("123456"))
	if EncodeMD5("hello") != "5d41402abc4b2a76b9719d911017c592" {
		t.Error(`EncodeMD5("hello") != "5d41402abc4b2a76b9719d911017c592"`)
	}
	if EncodeMD5("world") != "7d793037a0760186574b0282f2f435e7" {
		t.Error(`EncodeMD5("world") != "7d793037a0760186574b0282f2f435e7"`)
	}
}

// 压力测试
// go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench
// go test -test.bench=".*"
// go test -test.bench=".*" -count=5
func BenchmarkEncodeMD5(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		EncodeMD5("hello")
	}
}
