package app

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"

	"github.com/salemzii/swing/service"
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

func TokenMiddleware(ctx context.Context) rpc.Middleware {
	return func(handler rpc.RpcHandler) rpc.RpcHandler {
		return func(ctx context.Context, req *rpc.RpcRequest) *rpc.RpcResponse {
			resp := handler(ctx, req)
			var tokenObj TokenStruct
			//var details TokenDetails
			if err := json.Unmarshal(req.Params, &tokenObj); err != nil {
				return resp
			}
			log.Println("The token is::: " + string(tokenObj.Token))

			details, err := service.VerifyToken(string(tokenObj.Token))
			if err != nil {
				return resp
			}

			if !details.Enabled {
				return resp
			}

			if !details.active {

			}

			return resp
		}
	}

}

func Decode(r io.Reader) (ts *TokenStruct, token string, err error) {

	ts = new(TokenStruct)
	if err = json.NewDecoder(r).Decode(ts); err != nil {
		return &TokenStruct{}, "", err
	}
	var tk string
	if err = json.Unmarshal(ts.Token, &tk); err == nil {
		return ts, tk, nil
	}
	return &TokenStruct{}, "", nil

}
