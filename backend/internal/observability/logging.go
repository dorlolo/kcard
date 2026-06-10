package observability

import (
	"context"
	"log/slog"
)

type contextKey string

const RequestIDKey contextKey = "requestID"
const JobIDKey contextKey = "jobID"

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}
func WithJobID(ctx context.Context, jobID string) context.Context {
	return context.WithValue(ctx, JobIDKey, jobID)
}

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
