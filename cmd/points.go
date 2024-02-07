package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dannylindquist/gopoints/internal/server"
	"github.com/dannylindquist/gopoints/internal/store"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer stop()
	db := store.NewStore()
	gameServer := store.NewGameService(db)
	server := server.NewServer()
	server.GameService = gameServer
	server.Open()
	<-ctx.Done()
}
