package jobs

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type JobStatus string

const (
	JobQueued    JobStatus = "queued"
	JobRunning   JobStatus = "running"
	JobSucceeded JobStatus = "succeeded"
	JobFailed    JobStatus = "failed"
)

type Job struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	Status          JobStatus `json:"status"`
	ProgressPercent int       `json:"progressPercent"`
	CurrentStep     string    `json:"currentStep"`
	IdempotencyKey  string    `json:"idempotencyKey,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
}

type Queue struct {
	redis *redis.Client
	name  string
}

func NewQueue(redis *redis.Client, name string) Queue { return Queue{redis: redis, name: name} }

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
