// Package ai 提供 AI 工作流引擎的核心抽象，包括结构化请求/响应模型和客户端接口。
package ai

import (
	"context"
	"kcardDesgin/backend/internal/ai/model"
)

// CardGenerationWorkflow 封装卡片生成工作流，包含 AI 客户端。
type CardGenerationWorkflow struct{ Client model.Client }

// Run 执行卡片生成工作流，返回工作流输出或错误。
func (w CardGenerationWorkflow) Run(ctx context.Context, input WorkflowInput) (WorkflowOutput, error) {
	return Workflows{Client: w.Client}.Run(ctx, WorkflowCardGeneration, input)
}
