package jobs

import "context"

type ExportWorker struct{}

func (ExportWorker) Handle(ctx context.Context, payload map[string]any) (map[string]any, error) {
	return map[string]any{"status": "accepted"}, nil
}
