package ai

import "context"

type PlanWorkflow struct{ Client Client }

func (w PlanWorkflow) Run(ctx context.Context, input WorkflowInput) (WorkflowOutput, error) {
	return Workflows{Client: w.Client}.Run(ctx, WorkflowKind("plan"), input)
}
