package app

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/bandersnach/sample-app/config"
	"github.com/bandersnach/sample-app/internal/router"
)

// App owns the HTTP server lifecycle.
type App struct {
	cfg    config.Config
	router *gin.Engine
}

// New wires application dependencies and returns a ready-to-run instance.
func New(cfg config.Config) *App {
	return &App{
		cfg:    cfg,
		router: router.New(cfg),
	}
}

// Run starts the HTTP server and blocks until the context is cancelled.
func (a *App) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:              a.cfg.Addr(),
		Handler:           a.router,
		ReadHeaderTimeout: 10 * time.Second,
	}

	errCh := make(chan error, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return nil
	case err := <-errCh:
		return err
	}
}
