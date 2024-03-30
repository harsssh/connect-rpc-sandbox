package greet

import (
	greetv1 "connect-rpc-sandbox/gen/greet/v1"
	"connect-rpc-sandbox/gen/greet/v1/greetv1connect"
	"connectrpc.com/connect"
	"context"
	"fmt"
	"log"
)

type handler struct{}

func NewHandler() greetv1connect.GreetServiceHandler {
	return &handler{}
}

func (h *handler) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
