// Package jobs 提供任务队列相关的功能，包括任务定义、队列操作和任务调度。
package jobs

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// JobStatus 表示任务的当前状态枚举类型。
type JobStatus string

// JobQueued 常量表示任务已加入队列等待执行。
// JobRunning 常量表示任务正在执行中。
// JobSucceeded 常量表示任务执行成功。
// JobFailed 常量表示任务执行失败。
const (
	JobQueued    JobStatus = "queued"
	JobRunning   JobStatus = "running"
	JobSucceeded JobStatus = "succeeded"
	JobFailed    JobStatus = "failed"
)

// Job 表示一个任务实体，包含任务的唯一标识、类型、状态和进度信息。
type Job struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	Status          JobStatus `json:"status"`
	ProgressPercent int       `json:"progressPercent"`
	CurrentStep     string    `json:"currentStep"`
	IdempotencyKey  string    `json:"idempotencyKey,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
}

// Queue 表示基于 Redis 的任务队列，提供任务的入队和出队操作。
type Queue struct {
	redis *redis.Client
	name  string
}

// NewQueue 创建一个新的任务队列实例。
func NewQueue(redis *redis.Client, name string) Queue { return Queue{redis: redis, name: name} }

// Enqueue 将一个任务添加到队列中并返回创建的任务对象。
func (q Queue) Enqueue(ctx context.Context, jobType, idempotencyKey string, payload any) (Job, error) {
	job := Job{ID: uuid.NewString(), Type: jobType, Status: JobQueued, CurrentStep: "queued", IdempotencyKey: idempotencyKey, CreatedAt: time.Now().UTC()}
	body, err := json.Marshal(map[string]any{"job": job, "payload": payload})
	if err != nil {
		return Job{}, err
	}
	if err := q.redis.LPush(ctx, q.name, body).Err(); err != nil {
		return Job{}, err
	}
	return job, nil
}

// Envelope 表示从队列中取出的任务信封，包含任务和负载数据。
type Envelope struct {
	Job     Job             `json:"job"`
	Payload json.RawMessage `json:"payload"`
}

// Pop 从队列中取出一个任务，支持超时设置。
func (q Queue) Pop(ctx context.Context, timeout time.Duration) (Envelope, error) {
	result, err := q.redis.BRPop(ctx, timeout, q.name).Result()
	if err != nil {
		return Envelope{}, err
	}
	var envelope Envelope
	if len(result) < 2 {
		return Envelope{}, nil
	}
	if err := json.Unmarshal([]byte(result[1]), &envelope); err != nil {
		return Envelope{}, err
	}
	return envelope, nil
}
