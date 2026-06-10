package ai

import "context"

type CardGenerationWorkflow struct{ Client Client }

func (w CardGenerationWorkflow) Run(ctx context.Context, input WorkflowInput) (WorkflowOutput, error) {
	return Workflows{Client: w.Client}.Run(ctx, WorkflowKind("cardgeneration"), input)
}
