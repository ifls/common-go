package msg

import (
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"gocore/util"
	"time"
)

var host string = "192.168.8.101"
var nsqip = host
var nsqAdmin string = nsqip + ":4171"       //nsqd 4171
var nsqdtcp string = nsqip + ":4150"        //nsqadmin 4150
var nsqdhttp string = nsqip + ":4151"       //nsqloopup 4151
var nsqlookupdtcp string = nsqip + ":4160"  //nsqd 4160
var nsqlookupdhttp string = nsqip + ":4161" //nsqlookup 4161

var paddr string = nsqdtcp

var received uint64 = 0

var mhs []nsq.HandlerFunc

func init() {
	mhs = make([]nsq.HandlerFunc, 0)
}

var producer *nsq.Producer
var consumer *nsq.Consumer

func createProducer(addr string) *nsq.Producer {
	producer, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		util.LogErr(err, zap.String("reason", "nsq.createProducer"))
		return nil
	}
	return producer
}

func Consumer(topic string, channel string, addr string) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 100000 * time.Millisecond
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		util.LogErr(err, zap.String("reason", "nsq.NewConsumer"))
		return err
	}

	c.AddHandler(&H{c})

	err = c.ConnectToNSQLookupd(addr)
	if err != nil {
		util.LogErr(err, zap.String("reason", "nsq.consumer.ConnectToNSQLookupd()"))
		return err
	}
	return nil
}

func createNsqTopic() {

}

func AddMessageHandler(topic string, handle nsq.HandlerFunc) {
	if topic != "" {

		mhs = append(mhs, handle)
	}
}

type H struct {
	nsqConsumer *nsq.Consumer
}

//处理消息
func (h *H) HandleMessage(msg *nsq.Message) error {
	received++
	util.DevInfo("receive ID:%s,addr:%s,message:%s", msg.ID, msg.NSQDAddress, string(msg.Body))

	for _, h := range mhs {
		if h != nil {
			h(msg)
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
	util.DevInfo("PublicMessage message = " + msg)

	//
	if producer == nil {
		producer = createProducer(paddr)
		if producer == nil {
			return
		}
	}

	//发布消息
	err := producer.Publish(topic, data)
	if err != nil {
		util.LogErr(err, zap.String("reason", "nsq.PublicMessage"))
		return
	}
}

type NSQHandler struct {
	handler func(*nsq.Message) error
}

func (this *NSQHandler) SetHandler(h func(*nsq.Message) error) {
	this.handler = h
}

func (this *NSQHandler) HandleMessage(message *nsq.Message) error {
	return this.handler(message)
}

func ConsumeMessage(topic string, channel string, handler func(*nsq.Message) error) {
	consumer, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if nil != err {
		util.LogErr(err)
		return
	}

	nh := NSQHandler{handler: handler}

	consumer.AddHandler(&nh)
	err = consumer.ConnectToNSQD(nsqdtcp)
	if nil != err {
		util.LogErr(err)
		return
	}
}
