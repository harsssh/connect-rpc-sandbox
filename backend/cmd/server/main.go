package main

import (
	"connect-rpc-sandbox/config"
	"connect-rpc-sandbox/server"
)

func main() {
	conf := config.GetConfig()
	server.Run(conf)
}
