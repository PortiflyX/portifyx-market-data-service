package cmd

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/portifyx/portifyx-goutils/pkg/logx"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "market-data",
	Short: "PortifyX Market Data Microservice",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func gracefulShutdown(server *echo.Echo) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	logx.Info(context.Background(), "[*] shutting down Market Data Service gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logx.Errorf(context.Background(), err, "[!] error during server shutdown")
	}

	_ = logx.Sync()
	logx.Info(context.Background(), "[*] market data service cleanly stopped")
}
