package socket

import (
	"fmt"
	net2 "github.com/ifls/gocore/net"
	"github.com/ifls/gocore/util"
	"go.uber.org/zap"
	"log"
	"net"
	"time"
)

const (
	ProtocolTcp  = "tcp"
	PacketLength = 4
)

type TcpCallback struct {
	Dispatch    func(conn net.Conn, cmd int32, packet []byte)
	HbTimeout   func(conn net.Conn)
	Connected   func(conn net.Conn)
	Closed      func(conn net.Conn)
	ReConnected func(conn net.Conn)
}

type TcpParams struct {
	BufSize int
}

//DONE：
// 服务端在本机的任一端口建立TCP监听
// 为接入的每一个客户端开辟一条独立的协程
// 循环接收客户端消息，不管客户端说什么，都自动回复“已阅xxx”
// 连接处理器不断接收客户端信息，知道客户端申请断开连接，服务端把连接关闭。
func ServerMain(addr string, callback TcpCallback, params TcpParams) {
	//启动监听器
	listener, err := net.Listen(ProtocolTcp, addr)
	util.LogErr(err, zap.String(util.LogTagReason, fmt.Sprintf("listen on %s->%s", ProtocolTcp, addr)))
	log.Println(params)
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	//tcp服务端应不断处于等待连接状态
	for {
		//阻塞，直到建立连接
		conn, err := listener.Accept()

		if err != nil {
			util.LogErr(err, zap.String(util.LogTagReason, "listener.Accept()"))
		} else {
			//每个连接, 一个子协程进行对话
			go handleConnect(conn, callback)
		}
	}
}

func handleConnect(conn net.Conn, callback TcpCallback) {
	defer func() {
		err := conn.Close()
		log.Println(err)
	}()

	connBuff := make([]byte, 0)
	readBuff := make([]byte, 1024)

	for {
		n, err := conn.Read(readBuff)
		if err != nil {
			util.LogErr(err, zap.String(util.LogTagReason, "Conn.Read() err"))
			return
		}

		if n <= 0 {
			util.LogErr(err, zap.String(util.LogTagReason, "Conn.Read() <=0 eof"))
			return
		}

		connBuff = append(connBuff, readBuff[:n]...)
		if len(connBuff) > PacketLength {
			connBuff = readPackets(conn, connBuff, callback)
		}
	}
}

func readPackets(conn net.Conn, buff []byte, callback TcpCallback) []byte {
	log.Println(conn)
	log.Println(callback)
	var packet []byte
	for {
		buff, packet = Unpacket(buff) //对缓冲区进行分包处理

		if packet != nil {
			//handle(conn, packet, callback)
		} else {
			break
		}
	}
	return buff
}

func Unpacket(buff []byte) ([]byte, []byte) {
	length := len(buff)

	//如果包长小于header 就直接返回 因为接收的数据不完整
	if length < PacketLength {
		return buff, nil
	}

	packetLength := int(net2.BytesToUInt32(buff)) + 4

	if length < packetLength {
		return buff, nil
	}

	//划分数组
	return buff[packetLength:], buff[PacketLength:packetLength]
}

//func handle(conn net.Conn, packet []byte, callback TcpCallback) {
//sample := &pb.HeaderContainer{}
//err := proto.Unmarshal(packet, sample)
//if err != nil {
//	util.LogErr(err, zap.String(util.LOGTAG_REASON, "proto.Unmarshal()"))
//	return
//}
//
//callback.Dispatch(conn, sample.GetHeader().Cmd, packet)
//}

//----------------------- client -------------------//

func ClientMain(addr string) {
	//客户端拨号连接服务端
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	defer func() {
		err := conn.Close()
		log.Println(err)
	}()
	util.LogErr(err, zap.String("reason", "net.Dial"))

	//创建一个承载信息的桶
	//bytearr := make([]byte, 1024)
	count := 0
	//准备一个终端输入读取器
	//reader := bufio.NewReader(os.Stdin)
	for {
		//读取器不断从终端接收读取，每收到一个终端信息就发送给服务器
		//lineByte, _, err := reader.ReadLine()
		//SErrorFunc(err, "reader.Read")

		//lineByte := logSvrPacket(count)
		lineByte := []byte("fwefwe")
		//for i:=0;i<10;i++ {
		//	lineByte = append(lineByte, getPacket()...)
		//}
		//fmt.Printf("%s\n", string([]byte{0x00, 0x00, 0x00, 0x41}))
		util.DevInfo("send to server:%v, len=%v\n", lineByte, len(lineByte))
		count, err = conn.Write(lineByte)
		if err != nil {
			log.Println(err)
		}
		count += len(lineByte)
		//for i := 0; i < len(lineByte); i++ {
		//
		//	time.Sleep(200 * time.Millisecond)
		//	fmt.Printf("send to server:%d %v\n", len(lineByte), lineByte[i:i+1])
		//	conn.Write(lineByte[i : i+1])
		//}
		time.Sleep(2000 * time.Millisecond)
		//接收服务端的返回信息
		//n, err := conn.Read(bytearr)
		//
		//serverMsg := string([]byte(bytearr[:n]))
		//fmt.Println("get from server:：", serverMsg)
		//if string(lineByte) != "bye" {
		//	fmt.Println("服务端信息：", serverMsg)
		//} else {
		//	fmt.Println("客户端程序结束，GAME OVER!")
		//	os.Exit(0)
		//
		//}
		//break
	}

}

//func logSvrPacket(c int) []byte {
//	packets := make([]byte, 0)
//	n := 1
//	for i := c; i < c+n; i++ {
//		logReq := &pb.LogReq{
//			Header: &pb.Header{
//				Sign:                 0x77ff3b8d,
//				Cmd:                  0x3001,
//				Type:                 0x01,
//				Length:               0x04,
//				XXX_NoUnkeyedLiteral: struct{}{},
//				XXX_unrecognized:     nil,
//				XXX_sizecache:        0,
//			},
//			Log: []byte("log [" + strconv.Itoa(i) + "]" + time.Now().String()),
//		}
//
//		bt1, err := proto.Marshal(logReq)
//		//loginReq.Header.Length = uint32(len(bt1))
//		//bt1, err = proto.Marshal(loginReq)
//		if err != nil {
//			util.DevInfo("err = %v\n", err)
//		}
//
//		packets = append(packets, UInt32ToBytes(uint32(len(bt1)))...)
//		packets = append(packets, bt1...)
//
//	}
//
//	return packets
//}

//func testPacket() []byte {
//	packets := make([]byte, 0)
//	n := 100
//	for i := 0; i < n; i++ {
//		loginReq := &pb.LoginReq{
//			Header: &pb.Header{
//				Sign:                 0x77ff3b8d,
//				Cmd:                  int32(pb.Protocol_CLI_LOGIN),
//				Type:                 0x01,
//				Length:               0x04,
//				XXX_NoUnkeyedLiteral: struct{}{},
//				XXX_unrecognized:     nil,
//				XXX_sizecache:        0,
//			},
//			Uid: uint32(i),
//		}
//
//		bt1, err := proto.Marshal(loginReq)
//		//loginReq.Header.Length = uint32(len(bt1))
//		//bt1, err = proto.Marshal(loginReq)
//		if err != nil {
//			log.Printf("err = %v\n", err)
//		}
//
//		packets = append(packets, UInt32ToBytes(uint32(len(bt1)))...)
//		packets = append(packets, bt1...)
//
//	}
//
//	for i := 0; i < n; i++ {
//		logoutReq := &pb.LogoutReq{
//			Header: &pb.Header{
//				Sign:                 0x77ff3b8d,
//				Cmd:                  int32(pb.Protocol_CLI_LOGOUT),
//				Type:                 0x01,
//				Length:               0x04,
//				XXX_NoUnkeyedLiteral: struct{}{},
//				XXX_unrecognized:     nil,
//				XXX_sizecache:        0,
//			},
//			Uid: uint32(i),
//		}
//
//		bt2, err := proto.Marshal(logoutReq)
//		//logoutReq.Header.Length = uint32(len(bt2))
//		//bt2, err = proto.Marshal(logoutReq)
//		if err != nil {
//			log.Printf("err = %v\n", err)
//		}
//
//		packets = append(packets, UInt32ToBytes(uint32(len(bt2)))...)
//		packets = append(packets, bt2...)
//	}
//	return packets
//}
