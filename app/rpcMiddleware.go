package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
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

			switch req.Method {
			case "hello":
				return resp
			case "users.login":
				return resp
			case "users.create":
				return resp
			}

			var tokenObj TokenStruct
			var tokenSting string
			if err := json.Unmarshal(req.Params, &tokenObj); err != nil {
				return resp
			}

			tokenByte, err := tokenObj.Token.MarshalJSON()
			if err != nil {
				log.Println(err)
			}
			if err = json.Unmarshal(tokenByte, &tokenSting); err != nil {
				return resp
			}

			valid, userid, username, err := ParseToken(tokenSting)
			if err != nil {
				if !valid {
					return rpc.ErrorResponse(-32000, ErrTokenInValid)
				}
				return resp
			}

			if userid == nil || username == nil {
				return resp
			}

			return resp
		}
	}

}

func Decode(r io.Reader) (ts *TokenStruct, token []byte, err error) {

	ts = new(TokenStruct)
	if err = json.NewDecoder(r).Decode(ts); err != nil {
		log.Println(err)
		return &TokenStruct{}, []byte(""), err
	}
	var tk string
	if err = json.Unmarshal(ts.Token, &tk); err == nil {
		return ts, []byte(tk), nil
	}
	return &TokenStruct{}, []byte(""), nil

}

func ParseToken(tokenString string) (bool, interface{}, interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := []byte(os.Getenv("JwtSecretKey"))
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token.Valid, claims["userid"], claims["user"], nil
	} else {
		return token.Valid, nil, nil, err
	}
}

func ServiceParseToken(tokenString string) (bool, int, interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := []byte(os.Getenv("JwtSecretKey"))
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userid := claims["userid"].(float64)

		return token.Valid, int(userid), claims["user"], nil
	} else {
		return token.Valid, 0, nil, err
	}
}
