package api

import (
	"EquityEye/cmd/api/server"
	"EquityEye/internal/logs"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	srv := server.NewServer()
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
		Handler: srv,
	}

	go func() {
		logs.Info("API listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logs.Error(err.Error())
			os.Exit(1)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		// make a new context for the Shutdown (thanks Alessandro Rosetti)
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logs.Error("Error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()

	return nil
}
