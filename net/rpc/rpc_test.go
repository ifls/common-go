package rpc

import (
	"context"
	"testing"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith struct {
}

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B

	return nil
}

func TestRpc(t *testing.T) {
	//s := server.NewServer()
	//s.RegisterName("Arith", new(Arith), "")
	//s.Serve("tcp", ":8972")
	err := Arith{}.Mul(nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}
