package main

import (
	"chat/internal/server"
	"flag"
)

var (
	addr    = flag.String("addr", ":8080", "http service address")
	network = flag.String("network", "tcp", "type of network connection")
)

func main() {
	flag.Parse()

	srv, err := server.New(*network, *addr)
	if err != nil {
		panic(err)
	}

	srv.Serve()
}
