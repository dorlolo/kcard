// Package ai 提供 AI 工作流引擎的核心抽象，包括结构化请求/响应模型和客户端接口。
package ai

import (
	"context"
	"kcardDesgin/backend/internal/ai/model"
)

// PlanWorkflow 封装复习计划生成工作流，包含 AI 客户端。
type PlanWorkflow struct{ Client model.Client }

// Run 执行复习计划生成工作流，返回工作流输出或错误。
func (w PlanWorkflow) Run(ctx context.Context, input WorkflowInput) (WorkflowOutput, error) {
	return Workflows{Client: w.Client}.Run(ctx, WorkflowPlanGeneration, input)
}
