package app

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"kcardDesgin/backend/internal/config"
	"kcardDesgin/backend/internal/jobs"
	"kcardDesgin/backend/internal/repository"
	"kcardDesgin/backend/internal/storage"
)

type Container struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Store  storage.Store
	Logger *slog.Logger
}

func New(cfg config.Config) (*Container, error) {
	db, err := repository.OpenPostgres(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	redisClient, err := jobs.NewRedisClient(cfg.RedisURL)
	if err != nil {
		return nil, err
	}
	return &Container{
		Config: cfg,
		DB:     db,
		Redis:  redisClient,
		Store:  storage.NewLocalStore(cfg.StoragePath),
		Logger: slog.Default(),
	}, nil
}

func (c *Container) Close(ctx context.Context) error {
	if c.Redis != nil {
		return c.Redis.Close()
	}
	return nil
}
