package main

import (
	"context"
	"flag"

	"github.com/TeamFat/go-way/code/rpcx/common"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args common.Args, reply *common.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	//https://github.com/rpcx-ecosystem/rpcx-examples3/tree/master/102basic
	flag.Parse()

	s := server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
