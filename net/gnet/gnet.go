package main

import (
	"log"
	"time"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
)

var addr = "0.0.0.0:9000"

type echoServer struct {
	*gnet.EventServer
	pool *goroutine.Pool
}

func (es *echoServer) React(c gnet.Conn) (out []byte, action gnet.Action) {
	frame := c.Read()
	data := append([]byte{}, frame...)

	// Use ants pool to unblock the event-loop.
	err := es.pool.Submit(func() {
		time.Sleep(1 * time.Second)
		c.AsyncWrite(data)
	})
	if err != nil {
		log.Printf("react submit err = %s\n", err)
		return nil, gnet.Skip
	}

	return
}

func main() {
	p := goroutine.Default()
	defer p.Release()

	echo := &echoServer{pool: p}
	log.Fatal(gnet.Serve(echo, "tcp://"+addr, gnet.WithMulticore(true)))
}
