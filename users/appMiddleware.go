package users

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"go.neonxp.dev/jsonrpc2/rpc"
)

var (
	ErrTokenNotSet   = errors.New("api token not set")
	ErrTokenInValid  = errors.New("invalid api token")
	ErrTokenExpired  = errors.New("token has expired")
	ErrTokenInActive = errors.New("token not active")
)

type TokenStruct struct {
	Token json.RawMessage `json:"token"`
}

func TokenMiddleware(ctx context.Context) (rpc.Middleware, error) {
	return func(handler rpc.RpcHandler) rpc.RpcHandler {
		return func(ctx context.Context, req *rpc.RpcRequest) *rpc.RpcResponse {
			resp := handler(ctx, req)
			return resp
		}
	}, nil

}

func Decode(r io.Reader) (ts *TokenStruct, err error) {

	ts = new(TokenStruct)
	if err = json.NewDecoder(r).Decode(ts); err != nil {
		return &TokenStruct{}, err
	}
	var tk string
	if err = json.Unmarshal(ts.Token, &tk); err == nil {

	}
	return &TokenStruct{}, nil

}
