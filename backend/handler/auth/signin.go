package auth

import (
	authv1 "connect-rpc-sandbox/gen/auth/v1"
	"connect-rpc-sandbox/gen/auth/v1/authv1connect"
	"connectrpc.com/connect"
	"context"
	"github.com/oklog/ulid/v2"
	"log"
)

const TokenHeader = "Auth-Token"

type handler struct{}

func NewHandler() authv1connect.AuthServiceHandler {
	return &handler{}
}

func (h *handler) SignIn(
	ctx context.Context,
	req *connect.Request[authv1.SignInRequest],
) (*connect.Response[authv1.SignInResponse], error) {
	log.Println("Sign in request: ", req.Msg.UserId)

	// TODO: Implement sign in logic
	res := connect.NewResponse(&authv1.SignInResponse{})
	res.Header().Set(TokenHeader, ulid.Make().String())
	return res, nil
}
