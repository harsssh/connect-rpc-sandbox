package middleware

import (
	"connect-rpc-sandbox/handler/auth"
	"connectrpc.com/connect"
	"context"
	"errors"
)

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			// TODO: Implement auth logic
			if req.Header().Get(auth.TokenHeader) == "" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("missing auth token"),
				)
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
