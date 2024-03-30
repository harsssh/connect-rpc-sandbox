package main

import (
	greetv1 "connect-rpc-sandbox/gen/greet/v1"
	"connect-rpc-sandbox/gen/greet/v1/greetv1connect"
	"connectrpc.com/connect"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(http.DefaultClient, "http://localhost:8080")
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{
			Name: "Alice",
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res.Msg.Message)
}
