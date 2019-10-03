package msg

import (
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"gocore/util"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

var topic string = "log_time"

type ConsumerHandler struct {
	t              *testing.T
	q              *nsq.Consumer
	messagesGood   int
	messagesFailed int
}

func (h *ConsumerHandler) LogFailedMessage(message *nsq.Message) {
	h.messagesFailed++
	h.q.Stop()
}

func (h *ConsumerHandler) HandleMessage(message *nsq.Message) error {
	msg := string(message.Body)
	if msg == "bad_test_case" {
		return errors.New("fail this message")
	}
	if msg != "multipublish_test_case" && msg != "publish_test_case" {
		h.t.Error("message 'action' was not correct:", msg)
	}
	h.messagesGood++
	return nil
}

func TestProducerConnection(t *testing.T) {
	config := nsq.NewConfig()
	host := "global.GetHost(global.ENVIR_DEVTES"

	//config.LocalAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:30001")

	w, _ := nsq.NewProducer(host+":4150", config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		t.Fatalf("should lazily connect - %s", err)
	}

	w.Stop()

	err = w.Publish("write_test", []byte("fail test"))
	if err != nsq.ErrStopped {
		t.Fatalf("should not be able to write after Stop()")
	}
}

func TestProducerPing(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(os.Stdout)

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w.Ping()

	if err != nil {
		t.Fatalf("should connect on ping")
	}

	w.Stop()

	err = w.Ping()
	if err != nsq.ErrStopped {
		t.Fatalf("should not be able to ping after Stop()")
	}
}

func TestProducerPublish(t *testing.T) {
	topicName := "publish" + strconv.Itoa(int(time.Now().Unix()))
	msgCount := 10

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	defer w.Stop()

	for i := 0; i < msgCount; i++ {
		err := w.Publish(topicName, []byte("publish_test_case"))
		if err != nil {
			t.Fatalf("error %s", err)
		}
	}

	err := w.Publish(topicName, []byte("bad_test_case"))
	if err != nil {
		t.Fatalf("error %s", err)
	}

	readMessages(topicName, t, msgCount)
}

func TestProducerMultiPublish(t *testing.T) {
	topicName := "multi_publish" + strconv.Itoa(int(time.Now().Unix()))
	msgCount := 10

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	defer w.Stop()

	var testData [][]byte
	for i := 0; i < msgCount; i++ {
		testData = append(testData, []byte("multipublish_test_case"))
	}

	err := w.MultiPublish(topicName, testData)
	if err != nil {
		t.Fatalf("error %s", err)
	}

	err = w.Publish(topicName, []byte("bad_test_case"))
	if err != nil {
		t.Fatalf("error %s", err)
	}

	readMessages(topicName, t, msgCount)
}

func TestProducerPublishAsync(t *testing.T) {
	topicName := "async_publish" + strconv.Itoa(int(time.Now().Unix()))
	msgCount := 10

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	defer w.Stop()

	responseChan := make(chan *nsq.ProducerTransaction, msgCount)
	for i := 0; i < msgCount; i++ {
		err := w.PublishAsync(topicName, []byte("publish_test_case"), responseChan, "test")
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

	for i := 0; i < msgCount; i++ {
		trans := <-responseChan
		if trans.Error != nil {
			t.Fatalf(trans.Error.Error())
		}
		if trans.Args[0].(string) != "test" {
			t.Fatalf(`proxied arg "%s" != "test"`, trans.Args[0].(string))
		}
	}

	err := w.Publish(topicName, []byte("bad_test_case"))
	if err != nil {
		t.Fatalf("error %s", err)
	}

	readMessages(topicName, t, msgCount)
}

func TestProducerMultiPublishAsync(t *testing.T) {
	topicName := "multi_publish" + strconv.Itoa(int(time.Now().Unix()))
	msgCount := 10

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	defer w.Stop()

	var testData [][]byte
	for i := 0; i < msgCount; i++ {
		testData = append(testData, []byte("multipublish_test_case"))
	}

	responseChan := make(chan *nsq.ProducerTransaction)
	err := w.MultiPublishAsync(topicName, testData, responseChan, "test0", 1)
	if err != nil {
		t.Fatalf(err.Error())
	}

	trans := <-responseChan
	if trans.Error != nil {
		t.Fatalf(trans.Error.Error())
	}
	if trans.Args[0].(string) != "test0" {
		t.Fatalf(`proxied arg "%s" != "test0"`, trans.Args[0].(string))
	}
	if trans.Args[1].(int) != 1 {
		t.Fatalf(`proxied arg %d != 1`, trans.Args[1].(int))
	}

	err = w.Publish(topicName, []byte("bad_test_case"))
	if err != nil {
		t.Fatalf("error %s", err)
	}

	readMessages(topicName, t, msgCount)
}

func TestProducerHeartbeat(t *testing.T) {
	topicName := "heartbeat" + strconv.Itoa(int(time.Now().Unix()))

	config := nsq.NewConfig()
	config.HeartbeatInterval = 100 * time.Millisecond
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	defer w.Stop()

	err := w.Publish(topicName, []byte("publish_test_case"))
	if err == nil {
		t.Fatalf("error should not be nil")
	}
	if identifyError, ok := err.(nsq.ErrIdentify); !ok ||
		identifyError.Reason != "E_BAD_BODY IDENTIFY heartbeat interval (100) is invalid" {
		t.Fatalf("wrong error - %s", err)
	}

	config = nsq.NewConfig()
	config.HeartbeatInterval = 1000 * time.Millisecond
	w, _ = nsq.NewProducer("127.0.0.1:4150", config)

	defer w.Stop()

	err = w.Publish(topicName, []byte("publish_test_case"))
	if err != nil {
		t.Fatalf(err.Error())
	}

	time.Sleep(1100 * time.Millisecond)

	msgCount := 10
	for i := 0; i < msgCount; i++ {
		err := w.Publish(topicName, []byte("publish_test_case"))
		if err != nil {
			t.Fatalf("error %s", err)
		}
	}

	err = w.Publish(topicName, []byte("bad_test_case"))
	if err != nil {
		t.Fatalf("error %s", err)
	}

	readMessages(topicName, t, msgCount+1)
}

func readMessages(topicName string, t *testing.T, msgCount int) {
	config := nsq.NewConfig()
	config.DefaultRequeueDelay = 0
	config.MaxBackoffDuration = 50 * time.Millisecond
	q, _ := nsq.NewConsumer(topicName, "ch", config)

	h := &ConsumerHandler{
		t: t,
		q: q,
	}
	q.AddHandler(h)

	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		t.Fatalf(err.Error())
	}
	<-q.StopChan

	if h.messagesGood != msgCount {
		t.Fatalf("end of test. should have handled a diff number of messages %d != %d", h.messagesGood, msgCount)
	}

	if h.messagesFailed != 1 {
		t.Fatal("failed message not done")
	}
}

func TestNsqProductor(t *testing.T) {
	strIP1 := nsqdtcp

	producer1, err := initProducer(strIP1)
	if err != nil {
		log.Fatal("init producer1 error:", err)
	}

	defer producer1.Stop()

	//读取控制台输入
	//reader := bufio.NewReader(os.Stdin)

	count := 0
	for {
		fmt.Print("please say:")
		//data, _, _ := reader.ReadLine()
		command := strconv.Itoa(count) + "->" + util.GetTime()

		err := producer1.public(topic, command)
		if err != nil {
			log.Fatal("producer1 public error:", err)
		}

		time.Sleep(100 * time.Millisecond)
		count++
	}
}

type nsqProducer struct {
	*nsq.Producer
}

//初始化生产者
func initProducer(addr string) (*nsqProducer, error) {
	util.DevInfo("init producer address:" + addr)
	producer, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		util.LogErr(err, zap.String("tag", "nsq.nsq.NewProducer"))
		return nil, err
	}
	return &nsqProducer{producer}, nil
}

//发布消息
func (np *nsqProducer) public(topic, message string) error {
	err := np.Publish(topic, []byte(message))
	if err != nil {
		log.Println("nsq public error:", err)
		return err
	}
	return nil
}

func TestNsqComsumer(t *testing.T) {
	err := initConsumer("test1", "test-channel1", nsqlookupdhttp)
	if err != nil {
		log.Fatal("init Consumer error")
	}
	err = initConsumer("test2", "test-channel2", nsqlookupdhttp)
	if err != nil {
		log.Fatal("init Consumer error")
	}
	select {}
}

type nsqHandler struct {
	nsqConsumer      *nsq.Consumer
	messagesReceived int
}

//处理消息
func (nh *nsqHandler) HandleMessage(msg *nsq.Message) error {
	nh.messagesReceived++
	fmt.Printf("receive ID:%s,addr:%s,message:%s", msg.ID, msg.NSQDAddress, string(msg.Body))
	fmt.Println()
	return nil
}

func initConsumer(topic, channel, addr string) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 3 * time.Second
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log.Println("init Consumer nsq.NewConsumer error:", err)
		return err
	}

	handler := &nsqHandler{nsqConsumer: c}
	c.AddHandler(handler)

	err = c.ConnectToNSQLookupd(addr)
	if err != nil {
		log.Println("init Consumer ConnectToNSQLookupd error:", err)
		return err
	}
	return nil
}

func TestNsqPublishMessage(t *testing.T) {
	i := 0
	for i < 100000000 {
		time.Sleep(2000 * time.Millisecond)
		PublicMessage(topic, []byte(util.GetTime()+";"+strconv.Itoa(i)))
		i++
	}

}

func TestNsqCsm(t *testing.T) {
	Consumer(topic, "t1", nsqlookupdhttp)
	//AddMessagehandler(topic, func(msg *nsq.Message) error{
	//	util.DevInfo("msg=%v", msg)
	//	return nil
	//})
	select {}
}

func TestNsq(t *testing.T) {
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()

		consumer, err := nsq.NewConsumer(topic, topic+"_testConsumer", nsq.NewConfig())
		if nil != err {
			util.LogErr(err)
			return
		}

		consumer.AddHandler(&NSQHandler{})
		err = consumer.ConnectToNSQD(nsqdtcp)
		if nil != err {
			util.LogErr(err)
			return
		}

		select {}
	}()

	waiter.Wait()
}

func TestConsumer(t *testing.T) {
	waiter := sync.WaitGroup{}
	waiter.Add(1)
	ConsumeMessage(topic, topic+"_testConsumer2", func(message *nsq.Message) error {
		util.DevInfo("%s [%+v] %+v %+v\n", message.ID, string(message.Body), message.Timestamp, message.NSQDAddress)
		message.Finish()
		return nil
	})
	waiter.Wait()
}
