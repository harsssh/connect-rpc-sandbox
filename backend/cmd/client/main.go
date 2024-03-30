package main

import (
	authv1 "connect-rpc-sandbox/gen/auth/v1"
	"connect-rpc-sandbox/gen/auth/v1/authv1connect"
	greetv1 "connect-rpc-sandbox/gen/greet/v1"
	"connect-rpc-sandbox/gen/greet/v1/greetv1connect"
	"connectrpc.com/connect"
	"context"
	"fmt"
	"net/http"
)

func main() {
	authClient := authv1connect.NewAuthServiceClient(http.DefaultClient, "http://localhost:8080")
	res, _ := authClient.SignIn(
		context.Background(),
		connect.NewRequest(&authv1.SignInRequest{
			UserId:   "alice",
			Password: "password",
		}),
	)
	fmt.Println(res.Header().Get("session_id"))

	greetClient := greetv1connect.NewGreetServiceClient(http.DefaultClient, "http://localhost:8080")
	greetRes, _ := greetClient.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{
			Name: "Alice",
		}),
	)
	fmt.Println(greetRes.Msg.Message)
}
