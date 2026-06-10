package jobs

import "context"

type CardGenerationWorker struct{}

func (CardGenerationWorker) Handle(ctx context.Context, payload map[string]any) (map[string]any, error) {
	return map[string]any{"status": "accepted"}, nil
}
