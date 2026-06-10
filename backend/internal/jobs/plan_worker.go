package jobs

import "context"

type PlanWorker struct{}

func (PlanWorker) Handle(ctx context.Context, payload map[string]any) (map[string]any, error) {
	return map[string]any{"status": "accepted"}, nil
}
