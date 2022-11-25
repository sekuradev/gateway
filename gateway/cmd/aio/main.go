package main

import (
	"context"
	"flag"

	"github.com/sekuradev/api/gateway/pkg/gateway"
)

var (
	ctx = context.Background()
)

func main() {
	flag.Parse()
	gateway.NewServer(gateway.AllHandler()).Serve()
}
