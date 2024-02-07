package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dannylindquist/gopoints/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer stop()
	server := server.NewServer()
	server.Open()
	<-ctx.Done()
}
