package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/salemzii/swing/service"
	"go.neonxp.dev/jsonrpc2/rpc"
	"go.neonxp.dev/jsonrpc2/transport"
)

func main() {

	server := rpc.New(
		rpc.WithLogger(rpc.StdLogger),
		rpc.WithTransport(&transport.HTTP{Bind: ":8080"}),
	)

	server.Register("hello", rpc.HS(Hello))
	server.Register("all", rpc.H(service.AllRecords))
	server.Register("create", rpc.H(service.CreateRecord))
	server.Register("lineno", rpc.H(service.GetRecordByNum))
	server.Register("function", rpc.H(service.GetRecordByFunction))
	server.Register("level", rpc.H(service.GetRecordByLevel))
	server.Register("bulkingest", rpc.H(service.CreateRecords))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func Hello(ctx context.Context) (string, error) {
	return "world", nil
}
