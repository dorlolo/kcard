package jobs

import (
	"context"

	"kcardDesgin/backend/internal/ai"
)

type MaterialAnalysisWorker struct{ Workflow ai.ClassificationWorkflow }

type MaterialAnalysisPayload struct {
	MaterialID string `json:"materialId"`
	Text       string `json:"text"`
	Prompt     string `json:"prompt"`
}

func (w MaterialAnalysisWorker) Handle(ctx context.Context, payload MaterialAnalysisPayload) (ai.ClassificationOutput, error) {
	return w.Workflow.Classify(ctx, payload.Text, payload.Prompt)
}
