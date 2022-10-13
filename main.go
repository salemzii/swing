package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"go.neonxp.dev/jsonrpc2/rpc"
	"go.neonxp.dev/jsonrpc2/transport"
)

func main() {
	server := rpc.New(
		rpc.WithLogger(rpc.StdLogger),
		rpc.WithTransport(&transport.HTTP{Bind: ":8080"}),
	)

	server.Register("hello", rpc.HS(Hello))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func Hello(ctx context.Context) (string, error) {
	return "world", nil
}
