package server

import (
	"connect-rpc-sandbox/config"
	"connect-rpc-sandbox/gen/auth/v1/authv1connect"
	"connect-rpc-sandbox/gen/greet/v1/greetv1connect"
	"connect-rpc-sandbox/handler/auth"
	"connect-rpc-sandbox/handler/greet"
	"connect-rpc-sandbox/middleware"
	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func Run(conf *config.Config) {
	mux := http.NewServeMux()
	authInterceptor := connect.WithInterceptors(middleware.NewAuthInterceptor())

	greetHandler := greet.NewHandler()
	mux.Handle(greetv1connect.NewGreetServiceHandler(greetHandler, authInterceptor))

	authHandler := auth.NewHandler()
	mux.Handle(authv1connect.NewAuthServiceHandler(authHandler))

	srv := &http.Server{
		Addr:    conf.Host + ":" + conf.Port,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
