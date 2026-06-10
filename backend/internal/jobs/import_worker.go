package jobs

import "context"

type ImportWorker struct{}

func (ImportWorker) Handle(ctx context.Context, payload map[string]any) (map[string]any, error) {
	return map[string]any{"status": "accepted"}, nil
}
