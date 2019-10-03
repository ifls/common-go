package msg

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"sync"
	"time"
)

const (
	NsqAddr = "35.201.225.232:4150"
	Topic   = "golang_sql_test"
)

func nsqPing() {
	nsq.Ping()
}

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

func wait() {

	//初始化配置
	config := nsq.NewConfig()
	//创造消费者，参数一时订阅的主题，参数二是使用的通道
	com, err := nsq.NewConsumer(Topic, "channel1", config)
	if err != nil {
		fmt.Println(err)
	}
	//添加处理回调
	com.AddHandler(&NsqHandler{nsqHandlerID: "One"})
	//连接对应的nsqd
	err = com.ConnectToNSQD(NsqAddr)
	if err != nil {
		fmt.Println(err)
	}

	//只是为了不结束此进程，这里没有意义
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	/*
		result:

		msg.Timestamp=2018-11-02 04:37:18, msg.nsqaddress=10.10.6.147:4150,msg.body=new data!
		98 One
		msg.Timestamp=2018-11-02 04:37:18, msg.nsqaddress=10.10.6.147:4150,msg.body=new data!
		99 One
		msg.Timestamp=2018-11-02 04:37:18, msg.nsqaddress=10.10.6.147:4150,msg.body=new data!
		100 One
		msg.Timestamp=2018-11-02 04:37:18, msg.nsqaddress=10.10.6.147:4150,msg.body=new data!

	*/

}

func publicTopicAndMessage() {
	//初始化配置
	config := nsq.NewConfig()
	for i := 0; i < 10; i++ {
		//创建100个生产者
		tPro, err := nsq.NewProducer(NsqAddr, config)
		if err != nil {
			fmt.Println(err)
		}
		//主题
		topic := Topic
		//主题内容
		tCommand := "new data!->" + string(i)
		//发布消息
		err = tPro.Publish(topic, []byte(tCommand))
		if err != nil {
			fmt.Println(err)
		}
	}
}
