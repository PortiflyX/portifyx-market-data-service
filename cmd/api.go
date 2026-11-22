package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
	"github.com/portifyx/portifyx-goutils/pkg/logx"
	"github.com/portifyx/portifyx-market-data-service/config"
	"github.com/portifyx/portifyx-market-data-service/internal/rest/health"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start Market Data HTTP server",
	Run:   runMarketDataServer,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func runMarketDataServer(cmd *cobra.Command, args []string) {
	config.Init()
	cfg := config.GetConfig()

	logger, err := logx.Init(cfg.App.AppName, cfg.App.AppVersion, cfg.App.AppEnv)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	ctx := context.Background()
	logx.Infof(ctx, "[*] starting %s v%s on port :%s (env: %s)", cfg.App.AppName, cfg.App.AppVersion, cfg.App.AppPort, cfg.App.AppEnv)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(echoMw.Recover())
	e.Use(echoMw.CORS())

	hdl := health.NewHealthHandler()
	e.GET("/healthcheck", hdl.HealthCheck)
	e.GET("/healthcheck-readiness", hdl.Readiness)

	go func() {
		addr := fmt.Sprintf(":%s", cfg.App.AppPort)
		if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logx.Fatalf(ctx, err, "[!] fatal error starting echo server: %v", err)
		}
	}()

	gracefulShutdown(e)
}
