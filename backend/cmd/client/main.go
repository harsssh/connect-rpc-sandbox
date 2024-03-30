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
	res, err := authClient.SignIn(
		context.Background(),
		connect.NewRequest(&authv1.SignInRequest{
			UserId:   "alice",
			Password: "password",
		}),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Header().Get("Auth-Token"))

	greetClient := greetv1connect.NewGreetServiceClient(http.DefaultClient, "http://localhost:8080")
	greetReq := connect.NewRequest(&greetv1.GreetRequest{
		Name: "Alice",
	})
	greetReq.Header().Set("Auth-Token", res.Header().Get("Auth-Token"))
	greetRes, err := greetClient.Greet(context.Background(), greetReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(greetRes.Msg.Message)
}
