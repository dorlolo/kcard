// Package jobs 提供任务队列相关的功能，包括任务定义、队列操作和任务调度。
package jobs

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// NewRedisClient 创建一个新的 Redis 客户端实例并验证连接是否成功。
func NewRedisClient(redisURL string) (*redis.Client, error) {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("parse redis url: %w", err)
	}
	client := redis.NewClient(options)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("connect redis: %w", err)
	}
	return client, nil
}
