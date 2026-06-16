// Package observability 提供可观测性工具，包括请求追踪与上下文日志。
package observability

import (
	"context"
	"log/slog"
)

type contextKey string

// RequestIDKey 常量用于在 context 中存取请求 ID 的键。
const RequestIDKey contextKey = "requestID"
// JobIDKey 常量用于在 context 中存取任务 ID 的键。
const JobIDKey contextKey = "jobID"

// WithRequestID 将请求 ID 存入 context 并返回新的 context。
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}
// WithJobID 将任务 ID 存入 context 并返回新的 context。
func WithJobID(ctx context.Context, jobID string) context.Context {
	return context.WithValue(ctx, JobIDKey, jobID)
}

// Logger 从 context 中提取请求 ID 和任务 ID，返回带有这些字段的日志器。
func Logger(ctx context.Context) *slog.Logger {
	logger := slog.Default()
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok && requestID != "" {
		logger = logger.With("request_id", requestID)
	}
	if jobID, ok := ctx.Value(JobIDKey).(string); ok && jobID != "" {
		logger = logger.With("job_id", jobID)
	}
	return logger
}
