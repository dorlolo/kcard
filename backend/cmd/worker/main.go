package main

import (
	"context"
	"log/slog"
	"os"

	"kcardDesgin/backend/internal/app"
	"kcardDesgin/backend/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("configuration error", "error", err)
		os.Exit(1)
	}
	container, err := app.New(cfg)
	if err != nil {
		slog.Error("worker bootstrap failed", "error", err)
		os.Exit(1)
	}
	defer container.Close(context.Background())
	slog.Info("worker ready", "queue", cfg.RedisURL)
}
