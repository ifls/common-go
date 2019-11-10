package socket

import (
	"log"
	"net"
)

func init() {

}

func handleUdpConn(conn *net.UDPConn) error {
	// 读取数据
	buf := make([]byte, 10240)
	// 设置buf[20] = 45， 读取的数据小于20字节， 则不会被覆盖
	log.Println("before read block")
	//超过读缓冲区大小的数据会直接丢弃
	length, udpAddr, err := conn.ReadFromUDP(buf)
	log.Println("after read block")
	if err != nil {
		return err
	}

	//logContent := strings.Replace(string(buf),"\n","",1)
	log.Printf("server read len = %d\n", length)
	log.Printf("server datat len = %d, read data: %v\n", len(buf), buf)

	// 发送数据
	length, err = conn.WriteToUDP([]byte("ok\r\n"), udpAddr)
	if err != nil {
		return err
	}

	log.Printf("server write len = %d\n", length)
	return nil
}
