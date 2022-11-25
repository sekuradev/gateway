package gateway

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var port = flag.Int("p", 8080, "port to bint this server to")

type Server interface {
	Serve()
}

type server struct {
	http.Server
}

func waitForShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

func (s *server) Serve() {
	go func() {
		if err := s.Server.ListenAndServe(); err != nil {
			log.Fatalf("error while serving: %v", err)
		}
	}()
	waitForShutdown()
	s.Shutdown(context.Background())
}

func NewServer(handler http.Handler) Server {
	return &server{
		Server: http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
			Handler: handler,
		},
	}
}
