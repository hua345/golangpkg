package nsq

import (
	"github.com/nsqio/go-nsq"
	"sync"
	"testing"
)

func TestConsumer(t *testing.T) {
	//初始化配置
	config := nsq.NewConfig()
	//创造消费者，参数一时订阅的主题，参数二是使用的通道
	com, err := nsq.NewConsumer("fang", "fangChan", config)
	if err != nil {
		t.Error(err)
	}
	//添加处理回调
	com.AddHandler(&NsqHandler{nsqHandlerID: "fangChan"})
	//连接对应的nsqd
	err = com.ConnectToNSQD(tcpNsqdAddrr)
	if err != nil {
		t.Error(err)
	}

	//只是为了不结束此进程，这里没有意义
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
