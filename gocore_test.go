package gocore

import (
	"flag"
	"fmt"
	"github.com/ifls/gocore/net/socket"
	"log"
	"os"
	"testing"
	"time"

	tp "github.com/henrylee2cn/teleport"
)

func test() {
	// proto level
	tp.SetLoggerLevel("ERROR")

	cli := tp.NewPeer(tp.PeerConfig{})
	defer log.Println(cli.Close())

	cli.RoutePush(new(Push))

	sess, err := cli.Dial(":9090")
	if err != nil {
		tp.Fatalf("%v", err)
	}

	//var result int
	//rerr := sess.Call("/math/add?author=henrylee2cn",
	//	[]int{1, 2, 3, 4, 5},
	//	&result,
	//).Rerror()
	//if rerr != nil {
	//	tp.Fatalf("%v", rerr)
	//}
	//tp.Printf("result: %d", result)

	tp.Printf("wait for 10s... %v", sess)
	time.Sleep(time.Second * 10)
}

func Test2(t *testing.T) {
	fmt.Println(os.Args)

	cmd := flag.String("type", "main", "")

	flag.Parse()

	fmt.Println("type:", *cmd)

	if *cmd == "client" {
		fmt.Println("client")
		test()
	} else if *cmd == "server" {
		fmt.Println("server")
		_ = socket.TcpMain()
	}
}

// Push push handler
type Push struct {
	tp.PushCtx
}

// Push handles '/push/status' message
func (p *Push) Status(arg *string) *tp.Rerror {
	tp.Printf("%s", *arg)
	return nil
}

func TestNsq(t *testing.T) {
	//msg.nsqPing()
	//msg.publicTopicAndMessage()
	//
	//msg.wait()
}

func TestNsqHandler_HandleMessage(t *testing.T) {
	for i := 0; i < 5; i++ {
		log.Println(i)
	}
}
