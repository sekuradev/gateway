package main

import (
	"context"
	"flag"

	"github.com/sekuradev/gateway/pkg/gateway"
)

var (
	ctx = context.Background()
)

func main() {
	flag.Parse()
	gateway.NewServer(gateway.AllHandler()).Serve()
}
