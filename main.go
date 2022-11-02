package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/salemzii/swing/app"
	"github.com/salemzii/swing/service"
	"go.neonxp.dev/jsonrpc2/rpc"
	"go.neonxp.dev/jsonrpc2/transport"
)

func main() {
	port := os.Getenv("PORT")
	server := rpc.New(
		rpc.WithLogger(rpc.StdLogger),
		rpc.WithTransport(&transport.HTTP{Bind: fmt.Sprintf(":%s", port)}),
		rpc.WithMiddleware(app.TokenMiddleware(context.Background())),
	)

	server.Register("hello", rpc.HS(Hello))
	server.Register("records.all", rpc.H(service.AllRecords))
	server.Register("records.create.one", rpc.H(service.CreateRecord))
	server.Register("records.lineno", rpc.H(service.GetRecordByNum))
	server.Register("records.function", rpc.H(service.GetRecordByFunction))
	server.Register("records.level", rpc.H(service.GetRecordByLevel))
	server.Register("records.create.many", rpc.H(service.CreateRecords))
	server.Register("records.duration.15", rpc.H(service.GetLast15MinutesRecords))
	server.Register("records.duration.x", rpc.H(service.GetLastXMinutesRecords))
	server.Register("records.delete.one", rpc.H(service.DeleteRecordF))
	server.Register("records.delete.many", rpc.H(service.DeleteRecordsF))
	server.Register("users.create", rpc.H(service.CreateUserAccount))
	server.Register("users.login", rpc.H(service.LoginUserAccount))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func Hello(ctx context.Context) (string, error) {
	return "world", nil
}
