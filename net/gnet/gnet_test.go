package main

import (
	"encoding/binary"
	"github.com/smallnest/goframe"
	"log"
	"net"
	"testing"
)

var connection net.Conn
var fc goframe.FrameConn

func init() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	connection = conn
	encoderConfig := goframe.EncoderConfig{
		ByteOrder:                       binary.BigEndian,
		LengthFieldLength:               4,
		LengthAdjustment:                3,
		LengthIncludesLengthFieldLength: false,
	}

	decoderConfig := goframe.DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldOffset:   1,
		LengthFieldLength:   4,
		LengthAdjustment:    2,
		InitialBytesToStrip: 3,
	}

	fc = goframe.NewLengthFieldBasedFrameConn(encoderConfig, decoderConfig, conn)
}

func TestGnet(t *testing.T) {
	defer connection.Close()

	var bs []byte
	for i := 0; i < 222; i++ {
		bs = append(bs, byte(i))
	}

	err := fc.WriteFrame(bs)
	if err != nil {
		t.Fatal(err)
	}

	buf, err := fc.ReadFrame()
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("received: %+v \n", buf)
}
