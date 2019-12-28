package tcp

import (
	"github.com/ifls/gocore/util"
	"go.uber.org/zap"
	"log"
	"net"
	"testing"
	"time"
)

func TestTcpListen(t *testing.T) {
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			var total []byte
			for {
				// Echo all incoming data.
				//io.Copy(os.Stdout, c)
				buffer := make([]byte, 1024)
				n, err := c.Read(buffer)
				if err != nil {
					t.Fatal(err)
				}
				total = append(total, buffer[:n]...)
				log.Printf("received %d %d %#v\n", n, len(total), buffer)
				time.Sleep(1 * time.Millisecond)
			}
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

func TestTcpClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		log.Printf("err:%s\n", err)
	}
	defer func() {
		err := conn.Close()
		log.Println(err)
	}()
	util.LogErr(err, zap.String("reason", "net.Dial"))

	//创建一个承载信息的桶
	count := 0

	for {
		lineByte := []byte("123")

		log.Printf("send to server:%#v, len=%v\n", lineByte, len(lineByte))
		c, err := conn.Write(lineByte)
		if err != nil {
			t.Fatal(err)
		}
		count += c
		log.Printf("write to len = %d %#v %d\n", count, lineByte, len(lineByte))
		//time.Sleep(20 * time.Millisecond)
		if count > 13333 {
			t.Fatalf("count = %d\n", count)
			return
		}
	}
}

func TestLookupIP(t *testing.T) {
	ips, err := net.LookupIP("macbookpro")
	if err != nil {
		t.Fatal(err)
	}

	log.Println(ips)
}
