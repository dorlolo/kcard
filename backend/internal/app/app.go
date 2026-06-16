// Package app 提供应用程序的依赖容器与生命周期管理。
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

// Container 表示应用程序的依赖容器，持有配置、数据库、Redis、存储和日志器等核心依赖。
type Container struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Store  storage.Store
	Logger *slog.Logger
}

// New 根据配置创建并初始化 Container，包括数据库连接、自动迁移、Redis 客户端和本地存储。
func New(cfg config.Config) (*Container, error) {
	db, err := repository.OpenPostgres(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	if err = repository.AutoMigrate(db); err != nil {
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

// Close 关闭容器持有的 Redis 连接，释放相关资源。
func (c *Container) Close(ctx context.Context) error {
	if c.Redis != nil {
		return c.Redis.Close()
	}
	return nil
}
