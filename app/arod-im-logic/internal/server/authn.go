package server

import (
	"context"
	"errors"
	"io/ioutil"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

type authKey struct{}

func JWTAuth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// 获取token
				token := tr.RequestHeader().Get("Authorization")
				// spew.Dump(token)
				// 获取公钥
				publicKeyByte, _ := ioutil.ReadFile("../../internal/conf/public.key")
				publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)

				// 解析token
				tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
					// 检查alg是否正确
					if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
						return nil, errors.New("verification token encryption type error")
					}
					return publicKey, nil
				})
				if err != nil {
					// spew.Dump(err)
					return nil, err
				}

				// 从token中拿取信息
				if tokenClaims != nil {
					if Claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
						// spew.Dump("成功获取Claims")
						ctx = context.WithValue(ctx, authKey{}, Claims)
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
