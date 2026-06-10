package ai

import "context"

type WorkflowKind string

const (
	WorkflowClassification   WorkflowKind = "classification"
	WorkflowCardGeneration   WorkflowKind = "card_generation"
	WorkflowPlanGeneration   WorkflowKind = "plan_generation"
	WorkflowPlanOptimization WorkflowKind = "plan_optimization"
)

type WorkflowInput struct {
	WorkspaceID string
	Prompt      string
	SourceIDs   []string
	Payload     map[string]any
}
type WorkflowOutput struct {
	Drafts   []map[string]any
	Warnings []string
}

type Workflows struct{ Client Client }

func (w Workflows) Run(ctx context.Context, kind WorkflowKind, input WorkflowInput) (WorkflowOutput, error) {
	_, err := w.Client.GenerateStructured(ctx, StructuredRequest{System: string(kind), Messages: []Message{{Role: "user", Content: input.Prompt}}, MaxTokens: 4096})
	if err != nil {
		return WorkflowOutput{}, err
	}
	return WorkflowOutput{Drafts: []map[string]any{}, Warnings: []string{}}, nil
}
