package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

var (
	//nsqd的地址，使用了tcp监听的端口
	tcpNsqdAddrr = "192.168.137.128:4150"
)

//声明一个结构体，实现HandleMessage接口方法（根据文档的要求）
type NsqHandler struct {
	//消息数
	msqCount int64
	//标识ID
	nsqHandlerID string
}

//实现HandleMessage方法
//message是接收到的消息
func (s *NsqHandler) HandleMessage(message *nsq.Message) error {
	//没收到一条消息+1
	s.msqCount++
	//打印输出信息和ID
	fmt.Println(s.msqCount, s.nsqHandlerID)
	//打印消息的一些基本信息
	fmt.Printf("msg.Timestamp=%v, msg.nsqaddress=%s,msg.body=%s \n", time.Unix(0, message.Timestamp).Format("2006-01-02 03:04:05"), message.NSQDAddress, string(message.Body))
	return nil
}
