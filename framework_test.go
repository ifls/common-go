package gocore

import (
	"flag"
	"fmt"
	"gocore/net"
	"os"
	"testing"
	"time"

	tp "github.com/henrylee2cn/teleport"
)

func test() {
	// proto level
	tp.SetLoggerLevel("ERROR")

	cli := tp.NewPeer(tp.PeerConfig{})
	defer cli.Close()

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

func main() {
	fmt.Println(os.Args)

	cmd := flag.String("type", "main", "")

	flag.Parse()

	fmt.Println("type:", *cmd)

	if *cmd == "client" {
		fmt.Println("client")
		test()
	} else if *cmd == "server" {
		fmt.Println("server")
		net.SocketMain()
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
	//msg.wait()
}
