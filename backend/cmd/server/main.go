package main

import (
	"connect-rpc-sandbox/gen/greet/v1/greetv1connect"
	"connect-rpc-sandbox/handler"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func main() {
	greetHandler := handler.NewGreetHandler()
	mux := http.NewServeMux()
	mux.Handle(greetv1connect.NewGreetServiceHandler(greetHandler))
	http.ListenAndServe("127.0.0.1:8080", h2c.NewHandler(mux, &http2.Server{}))
}
