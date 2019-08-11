package nsq

import (
	"github.com/nsqio/go-nsq"
	"testing"
)

func TestProducerConnection(t *testing.T) {
	producer, err := nsq.NewProducer(tcpNsqdAddrr, nsq.NewConfig())
	if err != nil {
		t.Error(err)
	}
	err = producer.Publish("fang", []byte("hello world"))
	if err != nil {
		t.Error(err)
	}
	producer.Stop()
}

// 并发性能测试
//go test -bench=.
func BenchmarkProducer(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	producer, err := nsq.NewProducer(tcpNsqdAddrr, nsq.NewConfig())
	if err != nil {
		b.Error(err)
	}
	defer producer.Stop()
	b.StartTimer()
	// 测试一个对象或者函数在多线程的场景下面是否安全
	for i := 0; i < b.N; i++ {
		err = producer.Publish("fang", []byte("fang love you"))
		if err != nil {
			b.Error(err)
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkProducerParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	producer, err := nsq.NewProducer(tcpNsqdAddrr, nsq.NewConfig())
	if err != nil {
		b.Error(err)
	}
	defer producer.Stop()
	b.StartTimer()
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err = producer.Publish("fang", []byte("fang love you"))
			if err != nil {
				b.Error(err)
			}
		}
	})
}
