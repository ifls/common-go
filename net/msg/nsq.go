package msg

import (
	log2 "github.com/ifls/gocore/utils/log"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"log"
	"time"
)

var host = "192.168.8.101"
var nsqIp = host

//var nsqAdmin string = nsqIp + ":4171"       //nsq 4171
var nsqAdminTcp = nsqIp + ":4150" //nsqAdmin 4150
//var nsqHttp string = nsqIp + ":4151"       //nsqLookup 4151
//var nsqLookupTcp string = nsqIp + ":4160"  //nsq 4160
var nsqLookupHttp = nsqIp + ":4161" //nsqLookup 4161

var pAddr = nsqAdminTcp

var received uint64 = 0

var mhs []nsq.HandlerFunc

var producer *nsq.Producer
var consumer *nsq.Consumer

func init() {
	mhs = make([]nsq.HandlerFunc, 0)
	log.Println(consumer)
}

func createProducer(addr string) *nsq.Producer {
	producer, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		log2.LogErr(err, zap.String("reason", "nsq.createProducer"))
		return nil
	}
	return producer
}

func Consumer(topic string, channel string, addr string) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 100000 * time.Millisecond
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log2.LogErr(err, zap.String("reason", "nsq.NewConsumer"))
		return err
	}

	c.AddHandler(&H{c})

	err = c.ConnectToNSQLookupd(addr)
	if err != nil {
		log2.LogErr(err, zap.String("reason", "nsq.consumer.ConnectToNSQLookupd()"))
		return err
	}
	return nil
}

//func createNsqTopic() {
//
//}

//func AddMessageHandler(topic string, handle nsq.HandlerFunc) {
//	if topic != "" {
//
//		mhs = append(mhs, handle)
//	}
//}

type H struct {
	nsqConsumer *nsq.Consumer
}

//处理消息
func (h *H) HandleMessage(msg *nsq.Message) error {
	received++
	log2.DevInfo("receive ID:%s,addr:%s,message:%s", msg.ID, msg.NSQDAddress, string(msg.Body))

	for _, h := range mhs {
		if h != nil {
			_ = h(msg)
			break
		}
	}
	return nil
}

/*************** API *****************/
//topic:主题
//data:数据
func PublicMessage(topic string, data []byte) {
	msg := string(data)
	log2.DevInfo("PublicMessage message = " + msg)

	//
	if producer == nil {
		producer = createProducer(pAddr)
		if producer == nil {
			return
		}
	}

	//发布消息
	err := producer.Publish(topic, data)
	if err != nil {
		log2.LogErr(err, zap.String("reason", "nsq.PublicMessage"))
		return
	}
}

type NSQHandler struct {
	handler func(*nsq.Message) error
}

func (n *NSQHandler) SetHandler(h func(*nsq.Message) error) {
	n.handler = h
}

func (n *NSQHandler) HandleMessage(message *nsq.Message) error {
	return n.handler(message)
}

func ConsumeMessage(topic string, channel string, handler func(*nsq.Message) error) {
	consumer, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if nil != err {
		log2.LogErr(err)
		return
	}

	nh := NSQHandler{handler: handler}

	consumer.AddHandler(&nh)
	err = consumer.ConnectToNSQD(nsqAdminTcp)
	if nil != err {
		log2.LogErr(err)
		return
	}
}
