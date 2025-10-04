package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bandersnach/sample-app/config"
	"github.com/bandersnach/sample-app/internal/app"
)

func main() {
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	application := app.New(cfg)

	log.Printf("starting http server on %s (env=%s)", cfg.Addr(), cfg.Env)

	if err := application.Run(ctx); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
