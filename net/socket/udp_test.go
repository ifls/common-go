package socket

import (
	"log"
	"net"
	"testing"
)

func TestUdpServer(t *testing.T) {
	udpAddr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:9998")
	if err != nil {
		t.Error(err)
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err := udpConn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	log.Printf("udp listening on %v\n", udpAddr.String())

	for true {
		t.Error(handleUdpConn(udpConn))
	}
}

func TestUdpCli(t *testing.T) {
	udpAddr, _ := net.ResolveUDPAddr("udp4", "127.0.0.1:9998")

	//连接udpAddr，返回 udpConn
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("udp dial ok ")

	// 发送数据
	lstr := "0123456789"
	str := ""
	for i := 0; i < 1; i++ {
		str = str + lstr
	}
	//最大一次性能写9216B, 可以根据写缓冲区变化
	err = udpConn.SetWriteBuffer(10250)
	if err != nil {
		t.Error(err)
	}
	length, err := udpConn.Write([]byte(str))
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("client write len = %d\n", length)

	//读取数据
	buf := make([]byte, 1024)
	length, err = udpConn.Read(buf)
	if err != nil {
		t.Error(err)
	}

	log.Printf("client read len = %d\n", length)
	log.Printf("client len = %d, read data = %v\n", len(buf), string(buf))
}
