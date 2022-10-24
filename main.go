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
		rpc.WithTransport(&transport.HTTP{Bind: ":8081"}),
		//rpc.WithMiddleware(app.TokenMiddleware(context.Background())),
	)

	server.Register("rpc.hello", rpc.HS(Hello))
	server.Register("rpc.records.all", rpc.H(service.AllRecords))
	server.Register("rpc.records.create.one", rpc.H(service.CreateRecord))
	server.Register("rpc.records.lineno", rpc.H(service.GetRecordByNum))
	server.Register("rpc.records.function", rpc.H(service.GetRecordByFunction))
	server.Register("rpc.records.level", rpc.H(service.GetRecordByLevel))
	server.Register("rpc.records.create.many", rpc.H(service.CreateRecords))
	server.Register("rpc.records.duration.15", rpc.H(service.GetLast15MinutesRecords))
	server.Register("rpc.records.duration.x", rpc.H(service.GetLastXMinutesRecords))
	server.Register("rpc.records.delete.one", rpc.H(service.DeleteRecordF))
	server.Register("rpc.records.delete.many", rpc.H(service.DeleteRecordsF))
	server.Register("rpc.users.create", rpc.H(service.CreateUserAccount))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func Hello(ctx context.Context) (string, error) {
	return "world", nil
}
