package main

import (
	"log/slog"
	"os"

	"kcardDesgin/backend/internal/app"
	"kcardDesgin/backend/internal/config"
	httpapi "kcardDesgin/backend/internal/transport/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("configuration error", "error", err)
		os.Exit(1)
	}
	container, err := app.New(cfg)
	if err != nil {
		slog.Error("application bootstrap failed", "error", err)
		os.Exit(1)
	}
	router := httpapi.NewRouter(container)
	if err := router.Run(cfg.HTTPAddr); err != nil {
		slog.Error("api server stopped", "error", err)
		os.Exit(1)
	}
}
