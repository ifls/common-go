package socket

import "testing"

func TestServerMain(t *testing.T) {
	ServerMain("0.0.0.0:38888", TcpCallback{
		Dispatch:  nil,
		HbTimeout: nil,
	}, TcpParams{})
}

func TestClinetMain(*testing.T) {
	ClientMain("127.0.0.1:38888")
}
